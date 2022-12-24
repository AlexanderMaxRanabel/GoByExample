package main
import "fmt"
func main() {
  
    var x int
    var y int
    var z int

    fmt.Println("Please Enter A Number")
    fmt.Scanf("%v", &x)
    
    fmt.Println("Please Enter A Number")
    fmt.Scanf("%v", &y)
    
    fmt.Println("Please Enter A Number")
    fmt.Scanf("%v", &z)
    
    if x == 0{
        fmt.Println("ERROR DECLARED")
    }
    
    if x == 1{
        fmt.Println(y+z) 
    }
    if x == 2{
        fmt.Println(y-z) 
    }
    if x == 3{
        fmt.Println(y*z) 
    }
    if x == 4{
        fmt.Println(y/z) 
    }
}
