package cnfg

import (
	"testing"

	"strings"
	"time"
)

func init() {
	userDataPath = "../configs/user_data.yml.example"
}

func TestReadFromConfigs(t *testing.T) {
	user_cnfg := NewUserData()
	user_cnfg.Load()

	wantTitleName := "Артем Егоров"
	if user_cnfg.TitleName != wantTitleName {
		t.Errorf("Error got: %s, want: %s.", user_cnfg.TitleName, wantTitleName)
	}

	wantPrice := 10000
	if user_cnfg.Price != wantPrice {
		t.Errorf("Error got: %v, want: %v.", user_cnfg.Price, wantPrice)
	}

	wantDate := "2025-01-01"
	if user_cnfg.Date != wantDate {
		t.Errorf("Error got: %v, want: %v.", user_cnfg.Date, wantDate)
	}
}

func TestGetDateTime(t *testing.T) {
	user_cnfg := NewUserData()
	user_cnfg.Load()

	gotDateTime := user_cnfg.DateTime().String()
	wantDateTime := "2025-01-01 00:00:00"
	if !strings.Contains(gotDateTime, wantDateTime) {
		t.Errorf("Error got: %v, want: %v.", gotDateTime, wantDateTime)
	}

	user_cnfg.Date = ""
	gotDateTime = user_cnfg.DateTime().String()
	wantDateTime = time.Now().Local().String()[0:19]
	if !strings.Contains(gotDateTime, wantDateTime) {
		t.Errorf("Error got: %v, want: %v.", gotDateTime, wantDateTime)
	}
}

func TestInitialsFIO(t *testing.T) {
	user_cnfg := NewUserData()
	user_cnfg.Load()

	want := "ЕАА"
	got := user_cnfg.InitialsFIO()
	if got != want {
		t.Errorf("Error got: %v, want: %v.", got, want)
	}
}
