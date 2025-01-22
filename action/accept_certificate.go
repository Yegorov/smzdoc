package action

import (
	"fmt"
	"github.com/Yegorov/smzdoc/cnfg"

	"github.com/lukasjarosch/go-docx"
)

func GenAcceptCertificate(c *cnfg.Conf) {
	replaceMap := docx.PlaceholderMap{
		"contract_no": c.Ud.ContractNo,
		"date":        c.Ud.FormatDate(),
		"full_fio":    c.Ud.FullFIO,
		"short_fio":   c.Ud.ShortFIO,
		"price":       c.Ud.FormatPrice(),
		"price_words": c.Ud.FormatPriceWords(),
	}
	doc, err := docx.Open(AcceptCertificateTemplatePath(c))
	if err != nil {
		panic(err)
	}
	err = doc.ReplaceAll(replaceMap)
	if err != nil {
		panic(err)
	}
	err = doc.WriteToFile(AcceptCertificateOutputPath(c))
	if err != nil {
		panic(err)
	}
}

var (
	acceptCertificateTemplate     string = "templates/accept-certificate.docx"
	acceptCertificateOutputFolder string = "output"
	acceptCertificatePrefix       string = ""
)

func SetAcceptCertificateTemplate(s string) {
	acceptCertificateTemplate = s
}

func SetAcceptCertificateOutputFolder(s string) {
	acceptCertificateOutputFolder = s
}

func SetAcceptCertificatePrefix(s string) {
	acceptCertificatePrefix = s
}

func AcceptCertificateTemplatePath(_ *cnfg.Conf) string {
	return acceptCertificateTemplate
}

func AcceptCertificateOutputName(c *cnfg.Conf, suffix string) string {
	name := fmt.Sprintf("%s_%s_%s_%s", c.Ud.Date, c.Ud.InitialsFIO(), c.Ud.ShortContragentName, "Акт")
	name += suffix
	name += ".docx"
	return name
}

func AcceptCertificateOutputFolderPath(_ *cnfg.Conf) string {
	return acceptCertificateOutputFolder
}

func AcceptCertificateOutputPath(c *cnfg.Conf) string {
	return AcceptCertificateOutputFolderPath(c) + "/" +
		AcceptCertificateOutputName(c, acceptCertificatePrefix)
}
