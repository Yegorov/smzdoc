package action

import (
	"os"
	"testing"

	"github.com/Yegorov/smzdoc/cnfg"
)

func TestGenAcceptCertificate(t *testing.T) {
	cnfg.SetUserDataPath("../configs/user_data.yml.example")
	SetAcceptCertificateTemplate("../templates/accept-certificate-example.docx")
	SetAcceptCertificateOutputFolder("../output")
	SetAcceptCertificatePrefix("-dev")

	c := cnfg.NewConfig()
	ud := cnfg.NewUserData()
	ud.Load()
	conf := &cnfg.Conf{Ud: ud, C: c}

	outputFile := "../output/2025-01-01_ЕАА_ИИИ_Акт-dev.docx"
	GenAcceptCertificate(conf)

	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Errorf("Error: file %s must be exists", outputFile)
	} else {
		_ = os.Remove(outputFile)
	}
}
