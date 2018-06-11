package linode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Domain struct {
	ID     int    `json:"id,omitempty"`
	Domain string `json:"domain,omitempty"`
}

type Domains struct {
	Data   []Domain `json:"data,omitempty"`
	Page   int      `json:"page,omitempty"`
	Pages  int      `json:"pages,omitempty"`
	Result int      `json:"result,omitempty"`
}

type DomainRecord struct {
	ID       *int    `json:"id,omitempty"`
	Weight   *int    `json:"weight,omitempty"`
	Name     *string `json:"name,omitempty"`
	Target   *string `json:"target,omitempty"`
	Priority *int    `json:"priority,omitempty"`
	Type     string  `json:"type,omitempty"`
	Port     *int    `json:"port,omitempty"`
	Service  *string `json:"service"`
	Protocol *string `json:"protocol"`
	TTLSec   *int    `json:"ttl_sec,omitempty"`
	Tag      *string `json:"tag,omitempty"`
}

type LinodeClient interface {
	Request(method string, snippet string, body interface{}, res interface{}) error
}

type LinodeClientImpl struct {
	token string
}

func NewLinodeClient(token string) LinodeClient {
	return LinodeClientImpl{
		token: token,
	}
}

func (c LinodeClientImpl) Request(method string, snippet string, body interface{}, res interface{}) error {
	var err error

	endpoint := "https://api.linode.com/v4/"

	client := &http.Client{}
	url := endpoint + snippet

	var req *http.Request

	if body == nil {
		log.Printf("req = %s %s", method, url)
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return err
		}

	} else {
		content, err := json.Marshal(body)
		if err != nil {
			return err
		}
		log.Printf("req = %s %s %s", method, url, content)
		req, err = http.NewRequest(method, url, bytes.NewReader(content))
		if err != nil {
			return err
		}
		req.Header.Add("Content-Type", "application/json")

	}
	req.Header.Add("Authorization", "Bearer "+c.token)

	rsp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(rsp.Body)
	bufStr := buf.String()

	if rsp.StatusCode != 200 {
		return fmt.Errorf("status %s, body = %s", bufStr)
	}

	err = json.Unmarshal([]byte(bufStr), res)
	if err != nil {
		return err
	}

	log.Printf("statusCode = %d, rsp = %s", rsp.StatusCode, bufStr)

	return nil
}
