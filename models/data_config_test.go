package models

import (
	"testing"

	"strings"
	"time"
)

func init() {
	dataConfigPath = "../configs/data_config.yml.example"
}

func TestReadFromConfigs(t *testing.T) {
	dc := NewDataConfig()
	dc.ReadFromConfigs()

	wantTitleName := "Артем Егоров"
	if dc.TitleName != wantTitleName {
		t.Errorf("Error got: %s, want: %s.", dc.TitleName, wantTitleName)
	}

	wantPrice := 10000
	if dc.Price != wantPrice {
		t.Errorf("Error got: %v, want: %v.", dc.Price, wantPrice)
	}

	wantDate := "2025-01-01"
	if dc.Date != wantDate {
		t.Errorf("Error got: %v, want: %v.", dc.Date, wantDate)
	}
}

func TestGetDateTime(t *testing.T) {
	dc := NewDataConfig()
	dc.ReadFromConfigs()

	gotDateTime := dc.GetDateTime().String()
	wantDateTime := "2025-01-01 00:00:00"
	if !strings.Contains(gotDateTime, wantDateTime) {
		t.Errorf("Error got: %v, want: %v.", gotDateTime, wantDateTime)
	}

	dc.Date = ""
	gotDateTime = dc.GetDateTime().String()
	wantDateTime = time.Now().Local().String()[0:19]
	if !strings.Contains(gotDateTime, wantDateTime) {
		t.Errorf("Error got: %v, want: %v.", gotDateTime, wantDateTime)
	}
}
