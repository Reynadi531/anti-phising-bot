package utils

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"mvdan.cc/xurls/v2"
)

var (
	API_URL = "https://phising-checker.reylabs.xyz"
)

type CheckResult struct {
	IsFound   bool      `json:"isFound"`
	IsPhising bool      `json:"isPhising"`
	Domain    string    `json:"domain"`
	Date      time.Time `json:"date"`
}

type ResponseCheck struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    CheckResult `json:"data"`
}

func ExtractURL(message string) [][]byte {
	rxRelaxed := xurls.Relaxed()
	link := rxRelaxed.FindAll([]byte(message), -1)

	return link
}

func CheckPhising(link string) (bool, error) {
	urlparams := url.Values{}
	urlparams.Add("url", link)

	checkurl := API_URL + "/api/v1/check?" + urlparams.Encode()

	resp, err := http.Get(checkurl)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var cResp ResponseCheck

	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return false, err
	}

	if cResp.Data.IsPhising {
		return true, nil
	}
	return false, nil
}
