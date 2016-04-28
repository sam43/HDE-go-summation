package main


import "fmt"


func sumElement(n int, total *int) {
    
if n == 0 {
        
return
    
}
    
    
var element int
    
fmt.Scan(&element)
    
    
if element >= 0 {

        (*total) += (element * element)
    
}
    
    
n--
    
sumElement(n, total)
}


func handleTestcase(n int) {
 
   if n == 0 {

        return
    }
 
   
    var elements int

    fmt.Scan(&elements)
  
  
    var total = 0

    
sumElement(elements, &total)

    fmt.Println(total)

   
 n--
    
handleTestcase(n)

}


func main() {
    
var testcase int
    
fmt.Scan(&testcase)    
    
handleTestcase(testcase)

}
