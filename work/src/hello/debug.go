package main



import (
    "log"
    "fmt"
)

type Stack struct {
    size int64 //栈的容量
    top int64 //栈顶
    data []interface{}
}

func MakeStack( size int64)  Stack{

    s :=Stack{}
    s.size=size
    s.data =make([]interface{},size)

    return s
}

//入栈，空间不足，逐段升高

func (s *Stack) Push(e interface{})  bool{


    if  s.IsFull(){
        log.Printf("栈满，无法入栈")
     return false

    }

    s.data[s.top]=e
    fmt.Println(s.top)
    s.top++

    return true
}
//出栈，栈顶降低
func (s *Stack) Pop()  (r interface{},err error){

    if s.IsEmpty() {
        err =fmt.Errorf("栈已空，无法完成出栈")
        log.Printf("栈已空，无法完成出栈")
        return
    }
    s.top--
    r =s.data[s.top]
    return
}




//判断栈是否满
func (s *Stack) IsFull()  bool{

    return s.top==s.size
}

//判断栈是否为空
func (s *Stack) IsEmpty()  bool{

    return s.top==0
}

func (s *Stack) Traverse(fn func(r interface{}),goorto bool)  {


    //go true遍历进栈   false 遍历出栈
    if goorto {
  var i  int64= 0
        for ;i<s.top;i++ {
            fn(s.data[i])

        }


    }else{
        fmt.Println(s.data)
        for i:=s.top-1;i>=0;i-- {
            fn(s.data[i])
        }
    }

}


//进栈出栈试验
//求将十进制1348转化为八进制

func TestStack()  {


    var fn_c = func(n int) {

        s :=MakeStack(10)

        for   {
            if n==0 {
                break
            }
            s.Push(n%8)
            n =n/8

        }
        s.Traverse(func(r interface{}) {
            fmt.Print(r)
        },false)


    }

    fn_c(1348)

}


func main() {

    TestStack()

}