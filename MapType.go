package main

import "fmt"


func MapType (){
 
 fmt.Println("Map Types ::")
 m := map[string] int {"foo":42}
 fmt.Println(m)

 m["foo"] = 56
 fmt.Println(m["foo"])

 delete(m,"foo")

 fmt.Println(m)
}