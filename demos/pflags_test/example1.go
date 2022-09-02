package main

import (
    flag "github.com/spf13/pflag"  //替换原生的flag，并兼容
    "fmt"
)

var flagvar1 int
var flagvar2 bool

func init() {
    flag.IntVar(&flagvar1, "varname1", 1, "help message for flagname")
   flag.BoolVarP(&flagvar2, "boolname1", "b", true, "help message")
}


func main() {
   var ip1 *int = flag.Int("flagname1", 1, "help message for flagname")
   
   var ip2 = flag.IntP("flagname2", "f", 2, "help message") 
   
   

   flag.Parse()
   
   fmt.Println("ip1 has value ", *ip1)
   fmt.Println("ip2 has value ", *ip2)
   fmt.Println("flagvar1 has value ", flagvar1) 
   fmt.Println("flagvar2 has value ", flagvar2) 
}
