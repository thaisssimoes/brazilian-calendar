package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Holiday struct {
	Date string `json:"date"`
	Name string `json:"name"`
}

var apiURL = "https://brasilapi.com.br/api/feriados/v1"

func (h *Holiday) getHolidays(startYear, numYears int) ([][]Holiday, error) {
	var holidayList [][]Holiday
	endYear := startYear + numYears

	for i := startYear; i <= endYear; i++ {
		requestURL := fmt.Sprintf("%v/%v", apiURL, i)
		annualHolidayList, err := request(requestURL)
		if err != nil {
			return nil, err
		}
		holidayList = append(holidayList, annualHolidayList)
	}
	return holidayList, nil
}
func request(requestURL string) ([]Holiday, error) {
	var jsonHolidayList []Holiday
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return nil, err
	}

	json.Unmarshal(resBody, &jsonHolidayList)

	return jsonHolidayList, nil
}
