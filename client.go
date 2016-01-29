package vindalu

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/vindalu/vindalu/config"
)

const CREDS_CONF = ".vindalu/credentials"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string
}

type Client struct {
	Url    string
	Config *config.ClientConfig
	creds  Credentials
}

func NewClient(url string) (c *Client, err error) {
	c = &Client{Url: url}
	err = c.loadUserConfig()
	return
}

func (c *Client) SetCredentials(user string, pass string) {
	c.creds = Credentials{user, pass, ""}
}

func (c *Client) loadUserConfig() (err error) {
	if c.Config, err = c.getConfig(); err == nil {

		var b []byte
		b, err = ioutil.ReadFile(os.Getenv("HOME") + "/" + CREDS_CONF)
		if err == nil {
			var v struct {
				Auth Credentials `json:"auth"`
			}

			if err = json.Unmarshal(b, &v); err == nil {
				c.creds = v.Auth
			}
		}
	}
	return
}

func (c *Client) doRequest(method, urlpath string, body []byte) (resp *http.Response, b []byte, err error) {
	var req *http.Request

	if body != nil {
		req, err = http.NewRequest(method, c.Url+urlpath, bytes.NewBuffer(body))
	} else {
		req, err = http.NewRequest(method, c.Url+urlpath, nil)
	}

	if err == nil {
		switch method {
		case "PUT", "POST", "DELETE":
			if len(c.creds.Token) > 0 {
				req.Header.Set("Authorization", "BEARER "+c.creds.Token)
			} else {
				req.SetBasicAuth(c.creds.Username, c.creds.Password)
			}
			break
		default:
			break
		}

		if resp, err = http.DefaultClient.Do(req); err == nil {
			defer resp.Body.Close()

			b, err = ioutil.ReadAll(resp.Body)
		}
	}
	return
}

func (c *Client) getConfig() (cfg *config.ClientConfig, err error) {
	var b []byte
	if _, b, err = c.doRequest("GET", "/config", nil); err == nil {
		cfg = &config.ClientConfig{}
		err = json.Unmarshal(b, cfg)
	}
	return
}

func (c *Client) getOpaque(items ...string) string {
	out := c.Config.ApiPrefix
	for _, v := range items {
		out += "/" + v
	}
	return out
}
