package main

import "fmt"
import "time"
import "strings"

var name string = "Dezza Rizqi"
var hobby = [] string {"Music", "Swimming", "Coding"}
var dob string = "12 October"

func main() {
  fmt.Printf("Name:\t\t%v\n", name)
  time.Sleep(time.Second * 5)
  fmt.Printf("Hobbies:\t%v\n", strings.Join(hobby, ","))
  time.Sleep(time.Second * 5)
  fmt.Printf("Date of Birth:\t%v\n", dob)

  // for i:=0 ; i<=2; i++ {
  //   if i>0 {
  //     time.Sleep(time.Second * 5)
  //     switch i {
  //     case 1:
  //       fmt.Printf("Hobbies:\t%v\n", strings.Join(hobby, ","))
  //     case 2:
  //       fmt.Printf("Date of Birth:\t%v\n", dob)
  //     }
  //   } else {
  //     fmt.Printf("Name:\t\t%v\n", name)
  //   }
  // }
}
