/**------------------------------------------------------------**
 * @filename quake/client.go
 * @author   jinycoo - admin@jinycoo.com
 * @version  1.0.0
 * @date     2021/8/17 10:52
 * @desc     go-quake - client init
 **------------------------------------------------------------**/
package quake

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinycoo/go-quake/quake/core/query"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	GET = "GET"
	POST = "POST"
	Debug = "debug"

	QToken = "X-QuakeToken"
)

// Client represents Quake HTTP client
type Client struct {
	Mode           string
	ApiKey         string
	Email          string
	BaseURL        string
	Client         *http.Client
}

// NewQuakeClient creates new Quake client using environment variable
func NewQuakeClient(client *http.Client) *Client {
	cfg := NewConfig()

	if client == nil {
		client = http.DefaultClient
	}

	if cfg != nil && cfg.Quake != nil {
		return &Client{
			ApiKey:  cfg.Quake.ApiKey,
			BaseURL: cfg.Quake.BaseUrl,
			Mode: cfg.Mode,
			Client:  client,
		}
	} else {
		log.Fatal("quake config setting err")
	}
	return nil
}

func (c *Client) NewRequest(method string, path string, params interface{}, body io.Reader) (*http.Request, error) {
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}
	return c.newRequest(method, u, params, body)
}

func (c *Client) newRequest(method string, u *url.URL, params interface{}, body io.Reader) (*http.Request, error) {
	qs, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	u.RawQuery = qs.Encode()

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set(QToken, c.ApiKey)
	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	return req, nil
}

func (c *Client) dumpRequest(req *http.Request) {
	var body string

	if req.Body != nil {
		bodyBytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Printf("[DEBUG] failed to read body: %s\n", err)
		}

		req.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))
		body = string(bodyBytes)
	}

	message := fmt.Sprintf("%s %s %s", req.Method, req.URL.String(), body)
	log.Printf("[DEBUG] client request: %s\n", message)
}

func (c *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	if c.Mode == Debug {
		c.dumpRequest(req)
	}
	return c.Client.Do(req)
}

// Do executes common (non-streaming) request.
func (c *Client) Do(ctx context.Context, req *http.Request, destination, page interface{}) error {
	resp, err := c.do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return GetErrorFromResponse(resp)
	}
	if destination == nil {
		return nil
	}
	return c.parseResponse(destination, page, resp.Body)
}

func (c *Client) parseResponse(destination, page interface{}, body io.Reader) (err error) {
	if w, ok := destination.(io.Writer); ok {
		_, err = io.Copy(w, body)
	} else {
		decoder := json.NewDecoder(body)
		var result = &QRes{Data:destination}
		if page != nil {
			switch page.(type) {
			case *Meta:
				result = &QRes{Data:destination, Meta: page.(*Meta)}
			case *Pagination:
				result = &QRes{Data:destination, Meta: &Meta{Page: page.(*Pagination)}}
			}
		}
		if err = decoder.Decode(result); err == nil {
			if code, ok := result.Code.(float64); !ok || code != 0 {
				err = errors.New(result.Message)
			}
		}
	}

	return
}

