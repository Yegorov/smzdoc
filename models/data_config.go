package models

import (
	"fmt"
	"os"
)

type DataConfig struct {
	TitleName   string
	FullFIO     string
	ShortFIO    string
	Price       int
	ContractNo  string
	CurrentDate string
}

func NewDataConfig() *DataConfig {
	return &DataConfig{}
}

func (dc *DataConfig) ReadFromConfigs() {
	data, err := os.ReadFile("configs/data_config.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
