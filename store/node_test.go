package store

import (
	"database/sql"
	"github.com/4whomtbts/timefabric/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"testing"
)

type TestSuite struct {
	suite.Suite
	db *sqlx.DB
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) SetupTest() {
	db, err := sqlx.Connect("mysql", "root:1234@(localhost:3307)/db?multiStatements=true&parseTime=true")
	if err != nil {
		panic(err)
	}
	suite.db = db

	initSqlFile, err := ioutil.ReadFile("./init-dev.sql")
	if err != nil {
		panic(err)
	}
	_, err = suite.db.Exec(string(initSqlFile))
	if err != nil {
		panic(err)
	}
}

func (suite *TestSuite) TestSaveNode() {
	err := saveNode(suite.db, &model.TimeFabricNode{
		NodeId:          0,
		StorageGroupId:  0,
		HostName:        "ubuntu~",
		NodeName:        "1",
		HostIP:          "192.168.1.15",
		Port:            "8081",
		Excluded:        false,
		Exclusive:       false,
		RegisteredAt:    sql.NullTime{},
		ExcludedAt:      sql.NullTime{},
		LastAllocatedAt: sql.NullTime{},
		LastHeartbeatAt: sql.NullTime{},
	})
	if err != nil {
		log.Fatal("failed test!")
	}
}
