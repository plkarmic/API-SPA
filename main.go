package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type expirationTime int32

// tokenJSON is the struct representing the HTTP response from OAuth2
// providers returning a token in JSON form.
type tokenJSON struct {
	AccessToken  string         `json:"access_token"`
	TokenType    string         `json:"token_type"`
	RefreshToken string         `json:"refresh_token"`
	ExpiresIn    expirationTime `json:"expires_in"` // at least PayPal returns string, while most return number
	Expires      expirationTime `json:"expires"`    // broken Facebook spelling of expires_in
}

type config struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

func getConfigFromJSON(file string) string {
	var config config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	payloadSTR := "------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"grant_type\"\r\n\r\n" + config.GrantType + "\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"client_id\"\r\n\r\n" + config.ClientID + "\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"client_secret\"\r\n\r\n" + config.ClientSecret + "\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"username\"\r\n\r\n" + config.Username + "\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"password\"\r\n\r\n" + config.Password + "\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--"

	fmt.Println(payloadSTR)

	return payloadSTR
}

func main() {

	url := "https://svr17.supla.org/oauth/v2/token"

	payload := strings.NewReader(getConfigFromJSON("./config.json"))
	//fmt.Println(payload)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "37172fb7-c7ef-dc00-2f0d-599649c07135")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))
	var tj tokenJSON
	json.Unmarshal(body, &tj)
	fmt.Println(tj.AccessToken)

}
