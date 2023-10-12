package httpcli

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var Cli *http.Client

const SvrFailedCode = -100

type SvrBody struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
}

func Start() {
	Cli = &http.Client{}
}

func Send(method string, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Set("Cookie", "name=anny")
	response, err := Cli.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()
	return io.ReadAll(response.Body)
}

func SendGamesvrPost(api string, mapinfo map[string]any) (SvrBody, error) {
	mapinfostr, _ := json.Marshal(mapinfo)
	params := url.Values{}
	params.Set("info", string(mapinfostr))
	body := strings.NewReader(params.Encode())
	respbody, err := Send("POST", api, body)
	gamesvrbody := SvrBody{Code: SvrFailedCode}
	if err != nil {
		return gamesvrbody, err
	}
	err = json.Unmarshal(respbody, &gamesvrbody)
	if err != nil {
		return gamesvrbody, err
	}
	return gamesvrbody, nil
}
