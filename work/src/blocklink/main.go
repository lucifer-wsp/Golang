package main

import (
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "io"
    "log"
    "net/http"
    "os"
    "time"

    "blocklink/github.com/davecgh/go-spew/spew"
    "blocklink/github.com/gorilla/mux"
    "blocklink/github.com/joho/godotenv"
)

type Block struct{
    Index int
    Timestamp string
    BPM int
    Hash string
    PrevHash string
}

type Message struct {
    BPM int
}

func calculateHash(b Block) string{
    record := string(b.Index) + b.Timestamp + string(b.BPM) + b.PrevHash
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}

func generateBlock(ob Block, BPM int) (Block, error) {
    var nb Block

    t := time.Now()
    nb.Index = ob.Index + 1
    nb.Timestamp = t.String()
    nb.BPM = BPM
    nb.PrevHash = ob.Hash
    nb.Hash = calculateHash(nb)

    return nb, nil
}


func isBlockValid(nb, ob Block) bool {
    if ob.Index + 1 != nb.Index {
        return false
    }
    if ob.Hash != nb.PrevHash {
        return false
    }
    if calculateHash(nb) != nb.Hash {
        return false
    }
    return true
}

func replaceChain(nb []Block) {
    if len(nb) > len(blockchain) {
        blockchain = nb
    }
}

func run() error {
    mux := makeMuxRouter()
    httpAddr := os.Getenv("ADDR")
    log.Println("Listen on ", os.Getenv("ADDR"))
    s := &http.Server{
        Addr: ":" + httpAddr,
        Handler: mux,
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    if err := s.ListenAndServe(); err != nil {
        return err
    }
    return nil
}

func makeMuxRouter() http.Handler {
    muxRouter := mux.NewRouter()
    muxRouter.HandleFunc("/", handleGetBlockChain).Methods("GET")
    muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")

    return muxRouter
}


func handleGetBlockChain(w http.ResponseWriter, r *http.Request) {
    bytes, err := json.MarshalIndent(blockchain, "", " ")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    io.WriteString(w, string(bytes))
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
    var m Message

    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&m); err != nil {
        respondWithJSON(w, r, http.StatusBadRequest, r.Body)
        return
    }
    defer r.Body.Close()

    nb , err := generateBlock(blockchain[len(blockchain) - 1], m.BPM) 
    if err != nil {
        respondWithJSON(w, r, http.StatusInternalServerError, m)
        return
    }
    if isBlockValid(nb, blockchain[len(blockchain) - 1]) {
        nbChain := append(blockchain, nb)
        replaceChain(nbChain)
        spew.Dump(blockchain)
    }
    respondWithJSON(w, r, http.StatusCreated, nb)
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}){
    response, err := json.MarshalIndent(payload, "", " ")
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("HTTP 500: Internal Server Error"))
        return
    }
    w.WriteHeader(code)
    w.Write(response)
}

var blockchain []Block
func main(){
    err := godotenv.Load()
    if err != nil {
        log.Fatal(err)
    }

    go func() {
        t := time.Now()
        genesisBlock := Block{0, t.String(), 0, "", ""}
        spew.Dump(genesisBlock)
        blockchain = append(blockchain, genesisBlock)
    }()
    log.Fatal(run())
}