package main

import (
    "reflect"
    "fmt"
)
type Foo struct {
    A int `tag1:"Tag1" tag2:"Second Tag"`
    B string
}
func main(){
    f := Foo{A: 10, B: "Salutations"}
    fPtr := &f
    m := map[string]int{"A": 1 , "B":2}
    ch := make(chan int)
    sl:= []int{1,32,34}
    str := "string var"
    strPtr := &str

    tMap := examiner(reflect.TypeOf(f), 0)
    tMapPtr := examiner(reflect.TypeOf(fPtr), 0)
    tMapM := examiner(reflect.TypeOf(m), 0)
    tMapCh := examiner(reflect.TypeOf(ch), 0)
    tMapSl := examiner(reflect.TypeOf(sl), 0)
    tMapStr := examiner(reflect.TypeOf(str), 0)
    tMapStrPtr := examiner(reflect.TypeOf(strPtr), 0)

    fmt.Println("tMap :", tMap)
    fmt.Println("tMapPtr: ",tMapPtr)
    fmt.Println("tMapM: ",tMapM)
    fmt.Println("tMapCh: ",tMapCh)
    fmt.Println("tMapSl: ",tMapSl)
    fmt.Println("tMapStr: ",tMapStr)
    fmt.Println("tMapStrPtr: ",tMapStrPtr)
}

func examiner(t reflect.Type, depth int) map[int]map[string]string{
    outType := make(map[int]map[string]string)
    outType = map[int] map[string]string{depth:{"Name":t.Name(), "Kind":t.Kind().String()}}
  
    // 如果需要继续检测元素的类型
    switch t.Kind() {
    case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
        outType = make(map[int]map[string]string)
        tMap := examiner(t.Elem(), depth)
        for k, v := range tMap{
            outType[k] = v
        }

    case reflect.Struct:
        outType = make(map[int]map[string]string)
        for i := 0; i < t.NumField(); i++ {
            f := t.Field(i)
            outType[i] = map[string]string{
                "Name":f.Name,
                "Kind":f.Type.String(),
            }
        }
    }

    return outType
}