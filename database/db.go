package database

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"net/url"
	"sync"
)

var lock = &sync.Mutex{}

type ManagerSQLDB struct {
	DB *sqlx.DB
}

func (m *ManagerSQLDB) OpenConnection() error {

	user := viper.GetString("USER")
	password := viper.GetString("PASSWORD")
	host := viper.GetString("HOSTNAME")
	database := viper.GetString("DATABASE")
	port := viper.GetString("PORT")

	if m.DB == nil {
		lock.Lock()
		defer lock.Unlock()

		if m.DB == nil {
			u := &url.URL{
				Scheme: "sqlserver",
				User:   url.UserPassword(user, password),
				Host:   fmt.Sprintf("%s:%s", host, port),
			}
			fmt.Println("Creating SQL Server connection")
			db, err := sqlx.Open("sqlserver", u.String()+fmt.Sprintf("?database=%s", database))
			if err != nil {
				return err
			}
			m.DB = db
		} else {
			fmt.Println("SQL Server connection already exists, retrieving")
		}
	} else {
		fmt.Println("SQL Server connection already exists, retrieving")
	}

	return nil
}

func (m *ManagerSQLDB) CloseConnection() {
	m.DB.Close()
}
