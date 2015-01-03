package bls

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func SingleSeries(series string) string {
	url := "http://api.bls.gov/publicAPI/v2/timeseries/data/"
	req, err := http.NewRequest("POST", url+series, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return string(body)
}

// "LAUCN040010000000005", "LAUCN040010000000006"
func MultipleSeries(series ...string) {
	s := "\"" + strings.Join(series, "\",\"") + "\""
	payload := `{"seriesid":[` + s + `]}`
	url := "http://api.bls.gov/publicAPI/v2/timeseries/data/"

	jsonStr := []byte(payload)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func check(err error) {
	if err != nil {
		panic(err)
	}

}
