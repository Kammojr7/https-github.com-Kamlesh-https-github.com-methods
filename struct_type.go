package main
  
import "fmt"
  
type author struct {
    name      string
    branch    string
    particles int
    salary    int
}
  
func (a author) show() {
  
    fmt.Println("Author's Name: ", a.name)
    fmt.Println("Branch Name: ", a.branch)
    fmt.Println("Published articles: ", a.particles)
    fmt.Println("Salary: ", a.salary)
}
  
func main() {
  
     res := author{
        name:      "Kamlesh",
        branch:    "CSE",
        particles: 205,
        salary:    35000,
    }
  
     res.show()
}
