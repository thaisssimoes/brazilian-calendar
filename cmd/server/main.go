package main

import (
	"calendar-dataset/database"
	"calendar-dataset/internal/Utils"
	"calendar-dataset/internal/model"
	"calendar-dataset/internal/repository"
	"fmt"
)

func main() {
	Utils.ViperConfig()

	calendar, err := model.NewCalendar("2022-01-01", 10)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	var db database.ManagerSQLDB
	defer db.CloseConnection()

	err = db.OpenConnection()
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	err = repository.InsertCalendar(db.DB, calendar)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

}
