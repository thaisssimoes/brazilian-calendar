package model

import (
	"time"
)

type Calendario struct {
	Calendar []Date
}

type Date struct {
	Data    string `db:"data"`
	Dia     int    `db:"dia"`
	Mes     int    `db:"mes"`
	Ano     int    `db:"ano"`
	Feriado int8   `db:"feriado"`
	DiaUtil int8   `db:"dia_util"`
}

func NewCalendar(firstDate string, numYears int) (Calendario, error) {
	startDate, err := time.Parse("2006-01-02", firstDate)
	if err != nil {
	}
	startYear, _, _ := startDate.Date()

	var d Date
	var h Holiday
	var c Calendario
	mapDate, err := d.getDate(firstDate, numYears)
	if err != nil {
		return Calendario{}, err
	}
	holidaysList, err := h.getHolidays(startYear, numYears-1)
	if err != nil {
		return Calendario{}, err
	}
	mapDate = d.getHoliday(mapDate, holidaysList)

	for _, date := range mapDate {
		c.Calendar = append(c.Calendar, date)
	}

	return c, nil
}

func (c *Date) getDate(startDate string, numberYears int) (map[string]Date, error) {
	var mapDates = map[string]Date{}
	var mes time.Month

	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, err
	}
	end := start.AddDate(numberYears, 0, -1)

	for d := start; d.After(end) == false; d = d.AddDate(0, 0, 1) {
		c.Ano, mes, c.Dia = d.Date()
		c.Mes = int(mes)
		c.Data = d.Format("2006-01-02")
		c.DiaUtil = 1
		if int(d.Weekday()) == 0 || int(d.Weekday()) == 6 {
			c.DiaUtil = 0
		}
		mapDates[c.Data] = *c
	}

	return mapDates, nil
}

func (c *Date) getHoliday(calendar map[string]Date, listHoliday [][]Holiday) map[string]Date {
	var d Date
	for _, holidaysByYear := range listHoliday {
		for _, holiday := range holidaysByYear {
			d = calendar[holiday.Date]
			d.DiaUtil = 0
			d.Feriado = 1
			checkCarnaval(calendar, holiday)
			calendar[holiday.Date] = d
		}
	}
	return calendar
}

func checkCarnaval(calendar map[string]Date, holiday Holiday) error {
	var d Date
	carnavalDate, err := time.Parse("2006-01-02", holiday.Date)
	if err != nil {
		return err
	}

	if holiday.Name == "Carnaval" && int(carnavalDate.Weekday()) == 2 {
		date, err := time.Parse("2006-01-02", holiday.Date)
		if err != nil {
			return err
		}
		date = date.AddDate(0, 0, -1)
		previousDate := date.Format("2006-01-02")
		d = calendar[previousDate]
		d.DiaUtil = 0
		d.Feriado = 1
		calendar[previousDate] = d
	}
	return nil
}
