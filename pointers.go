package main

import "fmt"

func main() {
   var a int = 20
   var ip *int

   ip = &a

   fmt.Printf(" address of a variable: %x\n", &a )


   fmt.Printf(" address srored in ip variable: %x\n", ip )


   fmt.Printf(" value of *ip variable: %d\n", *ip )
}