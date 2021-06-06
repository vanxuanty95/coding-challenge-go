package main

import (
	"coding-challenge-go/cmd/api/config"
	"coding-challenge-go/pkg/api"
	"coding-challenge-go/pkg/logger"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

var gfgLogger logger.Logger

func main() {
	state := flag.String("state", "local", "state of service")
	gfgLogger = logger.WithPrefix("main")
	cfg := initConfig(*state)

	db, err := initDB(cfg)
	if err != nil {
		gfgLogger.Errorf("Fail to create server: %v", err)
		return
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			gfgLogger.Errorf("Fail to close db connection: %v", err)
		}
	}(db)

	initRestfulAPI(cfg, db)
}

func initConfig(state string) *config.Config {
	cfg := getConfig(state)

	domain := os.Getenv("LISTEN")
	if domain != "" {
		DomainElement := strings.Split(domain, ":")
		cfg.RestfulAPI.Host = DomainElement[0]
		cfg.RestfulAPI.Port = DomainElement[1]
	}

	return cfg
}

func getConfig(state string) *config.Config {
	cfgPath := fmt.Sprintf("config/config.%v.yml", state)
	f, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		gfgLogger.Panicf("Fail to open configurations file: %v", err)
	}

	var cfg config.Config
	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		gfgLogger.Panicf("Fail to decode configurations file: %v", err)
	}
	cfg.State = state
	return &cfg
}

func initDB(cfg *config.Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Database)
	return sql.Open("mysql", connectionString)
}

func initRestfulAPI(cfg *config.Config, db *sql.DB) {
	engine, err := api.CreateAPIEngine(cfg, db)
	if err != nil {
		gfgLogger.Errorf("Fail to create server: %v", err)
		return
	}

	gfgLogger.Info("Start server")

	apiDomainString := fmt.Sprintf("%v:%v", cfg.RestfulAPI.Host, cfg.RestfulAPI.Port)
	err = engine.Run(apiDomainString)
	if err != nil {
		gfgLogger.Panicf("Fail to listen and server: %v", err)
	}
}
