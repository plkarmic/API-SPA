package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func getConfigFromJSON()
{

}

func main() {

	url := "https://svr17.supla.org/oauth/v2/token"

	payload := strings.NewReader("")

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
