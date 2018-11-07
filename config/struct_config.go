package config

import "fmt"

type bddDriver interface {
	driverUrl(conf DataBase) string
	driverModule() string
}

type mysql_struct struct{}

func (d mysql_struct) driverUrl(conf DataBase) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.user, conf.password, conf.host, conf.port, conf.name)
}

func (d mysql_struct) driverModule() string {
	return "mysql"
}
