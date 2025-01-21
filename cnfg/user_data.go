package cnfg

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"gopkg.in/yaml.v3"
)

var userDataPath string = "configs/user_data.yml"

func UserDataFullPath() string {
	dir, _ := os.Getwd()
	return path.Join(dir, userDataPath)
}

type UserData struct {
	TitleName  string `yaml:"title_name"`
	FullFIO    string `yaml:"full_fio"`
	ShortFIO   string `yaml:"short_fio"`
	Price      int    `yaml:"price"`
	ContractNo string `yaml:"contract_no"`
	Date       string `yaml:"date"`
}

func NewUserData() *UserData {
	return &UserData{}
}

func (dc *UserData) Load() {
	data, err := os.ReadFile(UserDataFullPath())
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(data))
	err = yaml.Unmarshal([]byte(data), dc)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	// fmt.Printf("--- dc:\n%v\n\n", dc)
}

func (dc *UserData) DateTime() time.Time {
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
