package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

const (
	baseURL    = "https://opendata.cwb.gov.tw/api/v1" // 中央氣象局開放資料平臺
	path       = "/rest/datastore/F-D0047-063"        // 臺北市未來1週天氣預報
	queryField = "locationName"
)

var (
	queryParams = map[string]string{
		"Authorization": "CWB-B565A7A8-D4E7-4CBC-8DA9-3CD30B123027",
		"elementName":   "WeatherDescription",
	}
)

func requestExternalAPI(conn net.Conn, keyword string) error {
	queryParams[queryField] = keyword

	req, err := http.NewRequest("GET", baseURL+path, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	q := req.URL.Query()
	for k, v := range queryParams {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	fmt.Println("Request URL:", req.URL.String())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	respBody := &respBody{}
	if err := json.NewDecoder(resp.Body).Decode(respBody); err != nil {
		fmt.Println(err)
		return err
	}

	if respBody.Success == "true" {
		conn.Write([]byte("Query succeeded.\n"))
		if locs := respBody.Records.Locations; len(locs) > 0 {
			if loc := locs[0].Locations; len(loc) > 0 {
				if we := loc[0].WeatherElements; len(we) > 0 {
					if td := we[0].TimeData; len(td) > 0 {
						if ev := td[0].ElementValues; len(ev) > 0 {
							conn.Write([]byte("Weather in " + keyword + " in next 12 hours: " + ev[0].Value + "\n"))
						}
					}
				}
			}
		}
	}

	return nil
}
