package config

import "github.com/astaxie/beego/config"

type DBConfig struct {
	DBUser string
	DBPass string
	DBName string
}

func (dc *DBConfig) Read() error {
	fullConfigIni, err := config.NewConfig("ini", "config.ini")
	if err != nil {
		fullConfigIni, err = config.NewConfig("ini", "../config.ini")
		if err != nil {
			return err
		}
	}

	configIni, err := fullConfigIni.GetSection("default")

	if err != nil {
		println(2)
		return err
	}

	dc.DBUser = configIni["user"] // postgres
	dc.DBPass = configIni["pass"] // 1
	dc.DBName = configIni["name"] // users
	return nil
}
