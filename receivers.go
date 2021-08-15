package main
 
import (
    "fmt"
)
type Animal struct{
    name string
}
 
func (a Animal)Run(){
    fmt.Println(a.name, "is running...")
}
 
func (a *Animal)RunFaster(){
    fmt.Println(a.name, "is running...")
}
 
func main() {
    a := Animal{"Tiger"}
    a.Run()                // Tiger is running...
    a.RunFaster()          // Tiger is running...
}
