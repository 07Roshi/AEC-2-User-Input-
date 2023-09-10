package main
import "fmt"

func LPF(n int) int {
    var largest int
    for i:=2;i<=n;i++{
        for n%i==0{
            largest =i
            n/=i
        }
    }
    return largest
}
func main() {
    var cap int
    fmt.Printf("Enter the number: ")
    fmt.Scan(&cap)
    fmt.Println("The largest prime factor of",cap, "is:",LPF(cap))
}


// Output
// Enter a number: 14
// The largest prime factor of 14 is 7