package config

import "github.com/casdoor/casdoor-go-sdk/casdoorsdk"

func AuthInit() {
	casdoorsdk.InitConfig(Conf.CasdoorConf.Endpoint, Conf.CasdoorConf.ClientId,
		Conf.CasdoorConf.ClientSecret, Conf.CasdoorConf.JwtSecret, Conf.CasdoorConf.CasdoorOrganization,
		Conf.CasdoorConf.CasdoorApplication)
}
