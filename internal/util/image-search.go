package util

import (
	"ChattyDiaryBot/internal/config"
	param "ChattyDiaryBot/internal/util/params"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
)

type BingEndpointStruct struct {
	SearchAPIEndpoint      string
	ImageSearchAPIEndpoint string
}

var BingEndpoint BingEndpointStruct
var BingHost string

func init() {
	BingHost = "api.bing.microsoft.com"
	BingEndpoint.SearchAPIEndpoint = "api.bing.microsoft.com/"
	BingEndpoint.ImageSearchAPIEndpoint = "/v7.0/images/search"
}

//
//Return a url for the querying picture
//
func BingImageSearch(word string) (string, error) {
	cli := new(http.Client)
	req := new(http.Request)
	req.Header = http.Header{"Ocp-Apim-Subscription-Key": {config.Config.Bot.BingAPIKey}}
	//req.Header.Add()
	req.URL = &url.URL{
		Scheme:   "https",
		Host:     BingHost,
		Path:     BingEndpoint.ImageSearchAPIEndpoint,
		RawQuery: fmt.Sprintf("q=%s", word),
	}
	cli.Timeout = time.Duration(10) * time.Second
	resp, err := cli.Do(req)
	if err != nil {
		logrus.Error("Error when posting,msg:", err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	//n, err := resp.Body.Read(body)
	if err != nil {
		logrus.Error(err.Error())
	}
	//logrus.Info("Got response from server,msg=", string(body))
	//re, _ := regexp.Compile(`[.*?]`)
	//marshaled := re.FindStringSubmatch(string(body))
	bodyUnmarshaled := new(param.RespBingImageAPISuccess)
	err = json.Unmarshal(body, &bodyUnmarshaled)
	if err != nil {
		logrus.Error("Error at unmarshaling,msg:" + err.Error())
	}
	rand.Seed(time.Now().Unix())
	ind := rand.Int() % len(bodyUnmarshaled.Value)
	return bodyUnmarshaled.Value[ind].ContentUrl, err
}
