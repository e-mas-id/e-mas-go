package emas

import (
	"log"
	"os"
	"time"
	"io"
	"net/http"
	"io/ioutil"
)

type Client struct {
	Environment 	string
	AppId 			string
	SecretKey 		string

	Debug 			bool
	Logger 			*log.Logger
}

func NewClient() Client {
	return Client{
		Environment: "",
		Debug: true,
		Logger: log.New(os.Stderr, "", log.LstdFlags),
	}
}

var timeout = 30 * time.Second
var httpClient = &http.Client{Timeout: timeout}

func (c *Client) NewRequest(method string, fullPath string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, fullPath, body)
	if err != nil {
		if c.Debug {
			c.Logger.Println("Request creation failed: ", err)
		}
		return nil, err
	}

	t := time.Now()
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("App-Id", c.AppId)
	req.Header.Add("Timestamp", t.String())
	req.Header.Add("Signature", ShaOneEncrypt(Md5Encrypt(c.AppId + t.String() + c.SecretKey)))

	return req, nil
}

func (c *Client) ExecuteRequest(req *http.Request) ([]byte, error) {
	if c.Debug {
		c.Logger.Println("Request ", req.Method, ": ", req.URL.Host + req.URL.Path)
	}

	start := time.Now()

	res, err := httpClient.Do(req)
	defer res.Body.Close()

	if c.Debug {
		c.Logger.Println("Completed in ", time.Since(start))
	}

	if err != nil {
		if c.Debug {
			c.Logger.Println("Request failed: ", err)
		}
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		if c.Debug {
			c.Logger.Println("Cannot read response body: ", err)
		}
		return nil, err
	}

	if c.Debug {
		c.Logger.Println("API response: ", string(resBody))
	}

	return resBody, nil
}

func (c *Client) Call(method, path string, body io.Reader) ([]byte, error) {
	req, err := c.NewRequest(method, path, body)

	if err != nil {
		return nil, err
	}

	resp, err := c.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
