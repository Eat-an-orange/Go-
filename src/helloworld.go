package main


import  (
    "fmt"
   
)


func add(a int,b int) int{
    var sum int
    sum = a+b  
    return sum
}

func main() {
   var c int
   c = add(100,200)
   for i:=0; i<100; i++{
    go test_goroute(i,300)
   }
   //go语言中go关键字即可开启线程调用函数

   fmt.Println("add结果",c,"啊哈哈")
}
   