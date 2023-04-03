package repository

import (
	"calendar-dataset/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func InsertCalendar(db *sqlx.DB, calendar model.Calendario) error {
	var i int
	fmt.Println("Starting insert into database")
	for _, item := range calendar.Calendar {
		query := fmt.Sprintf(`INSERT INTO Calendario (data,dia,mes,ano,feriado,dia_util) VALUES ('%v', %v, %v, %v, %v, %v)`, item.Data, item.Dia, item.Mes, item.Ano, item.Feriado, item.DiaUtil)
		_, err := db.Exec(query)
		if err != nil {
			fmt.Printf("%v", err)
			return err
		}
		i++
	}

	return nil
}
