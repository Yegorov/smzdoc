package main

import "fmt"
import "flag"
import _ "time"
import _ "strconv"
import "github.com/Yegorov/smzdoc/cnfg"
import "github.com/Yegorov/smzdoc/action"

func main() {
	fmt.Println(cnfg.WelcomeMsg())
	c := cnfg.NewConfig()
	ud := cnfg.NewUserData()
	ud.Load()
	conf := &cnfg.Conf{Ud: ud, C: c}

	date := conf.Ud.DateTime() // time.Now()
	isDateSet := false
	price := conf.Ud.Price

	flag.Parse()
	acceptCertCmd := "accept-cert"

	switch flag.Arg(0) {
	case acceptCertCmd:
		acceptCert := flag.NewFlagSet(acceptCertCmd, flag.ExitOnError)
		acceptCert.Var(&cnfg.DateArg{Date: &date, IsDateSet: &isDateSet}, "date", "current date")
		acceptCert.IntVar(&price, "price", price, "you price")
		acceptCert.Parse(flag.Args()[1:])
		if isDateSet {
			// fmt.Println("Set date: " + date.String()[0:10])
			conf.Ud.Date = date.String()[0:10]
		}
		if price != conf.Ud.Price {
			// fmt.Println("Set price: " + strconv.Itoa(price))
			conf.Ud.Price = price
		}
		// fmt.Println("Do: " + acceptCertCmd)
		action.GenAcceptCertificate(conf)
	default:
		PrintHowUse()
	}
}

func PrintHowUse() {
	msg := `
use: smzdoc <cmd> [options]

commands:
  accept-cert - generate accept certificate
  options:
    --date - current date (e.g.: 2025-01-01)
    --price - price (e.g.: 10000)
`
	fmt.Println(msg)
}
