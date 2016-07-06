package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"platform_points/conf"
)

type AccountResponse struct {
	Id         int    `json:"id"`
	Nickname   string `json:"nickname"`
	Status     string `json:"status"`
	LoginToken string `json:"login_token"`
}

func Authenticate(loginToken string, appCode string) *AccountResponse {
	// Preparing POST Data
	values := url.Values{}
	values.Add("login_token", loginToken)
	values.Add("app_code", appCode)
	values.Add("without_session", "1")

	// Set login url
	url := conf.ACCOUNT_DOMAIN + "login/token"

	// 通信
	resp, _ := http.PostForm(url, values)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	// Create data from response
	var ar AccountResponse
	json.Unmarshal(body, &ar)
	//logrus.Debug(ar)
	//logrus.Debug(ar.Nickname)
	return &ar
}
