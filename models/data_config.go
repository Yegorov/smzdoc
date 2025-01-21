package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type DataConfig struct {
	TitleName  string `yaml:"title_name"`
	FullFIO    string `yaml:"full_fio"`
	ShortFIO   string `yaml:"short_fio"`
	Price      int    `yaml:"price"`
	ContractNo string `yaml:"contract_no"`
	Date       string `yaml:"date"`
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
	// fmt.Println(string(data))
	err = yaml.Unmarshal([]byte(data), dc)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- dc:\n%v\n\n", dc)
}

func (dc *DataConfig) GetDateTime() time.Time {
	datetime := time.Now().Local()
	if dc.Date != "" {
		var err error
		loc, _ := time.LoadLocation("Local")
		datetime, err = time.ParseInLocation(time.DateOnly, dc.Date, loc)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}
	return datetime
}
