package cnfg

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"golang.org/x/text/message"
	"gopkg.in/yaml.v3"
	ntw "moul.io/number-to-words"
)

var userDataPath string = "configs/user_data.yml"

func SetUserDataPath(s string) {
	userDataPath = s
}

func UserDataFullPath() string {
	dir, _ := os.Getwd()
	return path.Join(dir, userDataPath)
}

type UserData struct {
	TitleName           string `yaml:"title_name"`
	FullFIO             string `yaml:"full_fio"`
	ShortFIO            string `yaml:"short_fio"`
	Price               int    `yaml:"price"`
	ContractNo          string `yaml:"contract_no"`
	ShortContragentName string `yaml:"short_contragent_name"`
	Date                string `yaml:"date"`
}

func NewUserData() *UserData {
	return &UserData{}
}

func (ud *UserData) Load() {
	data, err := os.ReadFile(UserDataFullPath())
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(data))
	err = yaml.Unmarshal([]byte(data), ud)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	// fmt.Printf("--- ud:\n%v\n\n", ud)
}

func (ud *UserData) DateTime() time.Time {
	datetime := time.Now().Local()
	if ud.Date != "" {
		var err error
		loc, _ := time.LoadLocation("Local")
		datetime, err = time.ParseInLocation(time.DateOnly, ud.Date, loc)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}
	return datetime
}

func (ud *UserData) FormatDate() string {
	return ud.DateTime().Format("02.01.2006")
}

func (ud *UserData) FormatPrice() string {
	p := message.NewPrinter(message.MatchLanguage("ru"))
	return strings.ReplaceAll(p.Sprintf("%d", ud.Price), ",", " ")
}

func (ud *UserData) FormatPriceWords() string {
	return FirstUpcase(ntw.IntegerToRuRu(ud.Price))
}

func FirstUpcase(s string) string {
	return strings.Title(string([]rune(s)[0:1])) + string([]rune(s)[1:])
}

func (ud *UserData) InitialsFIO() string {
	str := ""
	for _, part := range strings.Split(ud.FullFIO, " ") {
		str += string([]rune(part)[0:1])
	}
	return str
}
