package main

import "fmt"

import "github.com/Yegorov/smzdoc/models"

func main() {
	fmt.Println("smzdoc")
	dc := models.DataConfig{
		TitleName: "Артем Егоров",
	}
	fmt.Println(dc)
}
