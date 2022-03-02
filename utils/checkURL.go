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
	IsFound      bool      `json:"isFound"`
	IsPhising    bool      `json:"isPhising"`
	IsSuspicious bool      `json:"isSuspicious"`
	Domain       string    `json:"domain"`
	Date         time.Time `json:"date"`
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

func CheckPhisingAndSuspicious(link string) (bool, bool, error) {
	urlparams := url.Values{}
	urlparams.Add("url", link)

	checkurl := API_URL + "/api/v1/check?" + urlparams.Encode()

	resp, err := http.Get(checkurl)
	if err != nil {
		return false, false, err
	}
	defer resp.Body.Close()

	var cResp ResponseCheck

	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return false, false, err
	}

	if cResp.Data.IsPhising && cResp.Data.IsSuspicious {
		return true, true, nil
	} else if cResp.Data.IsPhising {
		return true, false, nil
	} else if cResp.Data.IsSuspicious {
		return false, true, nil
	}

	return false, false, nil
}
