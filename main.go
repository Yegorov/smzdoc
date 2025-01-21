package main

import "fmt"

import "github.com/Yegorov/smzdoc/cnfg"

func main() {
	fmt.Println("smzdoc")
	ud := cnfg.NewUserData()
	ud.Load()
	fmt.Println(ud.DateTime())
}
