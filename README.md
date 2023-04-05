# brazilian-calendar

This is a Brazilian Calendar generator with a pre-programmed insertion to a SQL Server database created.
It uses the <a href="https://brasilapi.com.br/api/feriados/v1">Brasil API</a> to get the Nationals holidays and was verified using Anbimas's bank calendars.

To generate the calendar you must call the constructor *model.NewCalendar()*, passing the *start date* and *how many years* you would like to generate ahead.
Ex.:

```
calendar, err := model.NewCalendar("2022-01-01", 10)
```
The calendar (an object Calendar) will be generated automatically and assigned to the *calendar* variable, as you see in the example above. 

Calendar is basically a list of Dates that can be exported to a csv or a database of your choice (only the SQL Server connection is implemented).

In this Calendar, Dates have the following infos:

|Column   |  Type  |
|---------|--------|
| Data    | string | 
| Dia     | int    | 
| Mes     | int    | 
| Ano     | int    | 
| Feriado | int8   | 
| DiaUtil | int8   |

In which, Feriado and DiaUtil works as flags 0 or 1 (The int8 was choose because of database issues)

The main.go (found <a href="https://github.com/thaisssimoes/brazilian-calendar/blob/main/cmd/server/main.go">here</a>) is already built to create and insert a calendar into your SQL Server database.
You can also use the migration file to create and drop the calendar table.

