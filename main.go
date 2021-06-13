package main

import (
	"fmt"
	"github.com/4whomtbts/timefabric/config"
	"github.com/4whomtbts/timefabric/model"
	"github.com/4whomtbts/timefabric/store"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func main() {

	configFile, _ := filepath.Abs("example.yaml")
	yamlFile, err := ioutil.ReadFile(configFile)
	var config config.TimeFabricConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	fmt.Print(config)
	store.Db, err = sqlx.Connect("mysql", "root:1234@(localhost:3307)/db?multiStatements=true&parseTime=true")
	if err != nil {
		panic(err)
	}
	initSqlFile, err := ioutil.ReadFile("./init-dev.sql")
	if err != nil {
		panic(err)
	}
	_, err = store.Db.Exec(string(initSqlFile))
	if err != nil {
		panic(err)
	}
	var node []model.TimeFabricNode
	err = store.Db.Select(&node, `SELECT * FROM NODE`)
	if err != nil {
		panic(err)
	}
	fmt.Print("노드 = ", node)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/" , hello)
	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}