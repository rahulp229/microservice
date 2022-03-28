package executor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

type TestExecutor interface {
	FetchData(fsyms, tsyms string) (*Response, error)
}

type testExecutor struct {
}

func NewTestExecutor() TestExecutor {
	return testExecutor{}
}

type Response struct {
	Raw     *RawCrypto     `json:"RAW,omitempty"`
	Display *DisplayCrypto `json:"DISPLAY,omitempty"`
}

type RawCrypto struct {
	Bitcoin *RawCurrency `json:"BTC,omitempty"`
	Riple   *RawCurrency `json:"XRP,omitempty"`
}

type DisplayCrypto struct {
	Bitcoin *DisplayCurrency `json:"BTC,omitempty"`
	Riple   *DisplayCurrency `json:"XRP,omitempty"`
}

type RawCurrency struct {
	USD *RawData `json:"USD,omitempty"`
	EUR *RawData `json:"EUR,omitempty"`
}

type DisplayCurrency struct {
	USD *DisplayData `json:"USD,omitempty"`
	EUR *DisplayData `json:"EUR,omitempty"`
}
type RawData struct {
	CHANGE24HOUR    float32 `json:"CHANGE24HOUR,omitempty"`
	CHANGEPCT24HOUR float32 `json:"CHANGEPCT24HOUR,omitempty"`
	OPEN24HOUR      float32 `json:"OPEN24HOUR,omitempty"`
	VOLUME24HOUR    float32 `json:"VOLUME24HOUR,omitempty"`
	VOLUME24HOURTO  float32 `json:"VOLUME24HOURTO,omitempty"`
	LOW24HOUR       float32 `json:"LOW24HOUR,omitempty"`
	HIGH24HOUR      float32 `json:"HIGH24HOUR,omitempty"`
	PRICE           float32 `json:"PRICE,omitempty"`
	SUPPLY          float32 `json:"SUPPLY,omitempty"`
	MKTCAP          float32 `json:"MKTCAP,omitempty"`
}

type DisplayData struct {
	CHANGE24HOUR    string `json:"CHANGE24HOUR,omitempty"`
	CHANGEPCT24HOUR string `json:"CHANGEPCT24HOUR,omitempty"`
	OPEN24HOUR      string `json:"OPEN24HOUR,omitempty"`
	VOLUME24HOUR    string `json:"VOLUME24HOUR,omitempty"`
	VOLUME24HOURTO  string `json:"VOLUME24HOURTO,omitempty"`
	LOW24HOUR       string `json:"LOW24HOUR,omitempty"`
	HIGH24HOUR      string `json:"HIGH24HOUR,omitempty"`
	PRICE           string `json:"PRICE,omitempty"`
	SUPPLY          string `json:"SUPPLY,omitempty"`
	MKTCAP          string `json:"MKTCAP,omitempty"`
}

func (te testExecutor) FetchData(fsyms, tsyms string) (*Response, error) {
	queryString := "fsyms=" + fsyms + "&tsyms=" + tsyms
	fmt.Println("base url : ", viper.GetString("app.thirdPartyUrl"))
	url := viper.GetString("app.thirdPartyUrl") + queryString
	resp1, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	r, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(r, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
