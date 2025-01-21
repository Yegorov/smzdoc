package main

import "fmt"

import "github.com/Yegorov/smzdoc/models"

func main() {
	fmt.Println("smzdoc")
	dc := models.NewDataConfig()
	dc.ReadFromConfigs()
	fmt.Println(dc)
}
