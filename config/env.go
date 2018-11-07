package config

import (
	"os"
)

var testingMode bool

type DataBase struct {
	driver   bddDriver
	host     string
	port     string
	user     string
	password string
	name     string
}

type Emailer struct {
	apiKey string
	domain string
	mail   string
}

type Environnement struct {
	database map[string]DataBase
	emails   map[string]Emailer
	jwt_key  string
}

var default_bdd = DataBase{
	driver:   mysql_struct{},
	host:     EnvKey("BDD_HOST", "localhost", "localhost"),
	port:     EnvKey("BDD_PORT", "3306", "3306"),
	user:     EnvKey("BDD_USER", "root", "root"),
	password: EnvKey("BDD_PASSWORD", "", ""),
	name:     EnvKey("BDD_NAME", "karp", "karp_test"),
}

var default_email = Emailer{
	apiKey: EnvKey("EMAIL_API_KEY", "key-", "key-"),
	domain: EnvKey("EMAIL_DOMAIN", "p.eu", "p.eu"),
	mail:   EnvKey("EMAIL_MAIL", "master@a.eu", "master@a.eu"),
}

var env = Environnement{
	database: map[string]DataBase{
		"default": default_bdd,
	},
	emails: map[string]Emailer{
		"default": default_email,
	},
	jwt_key: EnvKey("JWT_KEY", "secret", "secret"),
}

func EnvKey(key string, def string, def_test string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		if testingMode {
			return def_test
		} else {
			return def
		}
	} else {
		return val
	}
}
