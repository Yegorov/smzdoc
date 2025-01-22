package main

import "fmt"
import "github.com/Yegorov/smzdoc/cnfg"
import "github.com/Yegorov/smzdoc/action"

func main() {
	fmt.Println(cnfg.WelcomeMsg())
	c := cnfg.NewConfig()
	ud := cnfg.NewUserData()
	ud.Load()
	conf := &cnfg.Conf{Ud: ud, C: c}

	action.GenAcceptCertificate(conf)
}
