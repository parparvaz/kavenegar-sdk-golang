package kavenegar

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// Endpoints
const (
	baseAPIMainURL = "https://api.kavenegar.com"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		BaseURL:    baseAPIMainURL,
		UserAgent:  "Kavenegar/golang",
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Kavenegar-golang ", log.LstdFlags),
	}
}

func NewProxyClient(apiKey, proxyUrl string) *Client {
	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		log.Println(err)

		return nil
	}

	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &Client{
		APIKey:    apiKey,
		BaseURL:   baseAPIMainURL,
		UserAgent: "Kavenegar/golang",
		HTTPClient: &http.Client{
			Transport: tr,
		},
		Logger: log.New(os.Stderr, "Kavenegar-golang ", log.LstdFlags),
	}
}

type doFunc func(req *http.Request) (*http.Response, error)

type Client struct {
	APIKey     string
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Debug      bool
	Logger     *log.Logger
	TimeOffset int64
	do         doFunc
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

func (c *Client) parseRequest(r *request, opts ...RequestOption) (err error) {
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)
	fullURL = fmt.Sprintf(fullURL, c.APIKey)

	queryString := r.query.Encode()
	body := &bytes.Buffer{}
	bodyString := r.form.Encode()
	header := http.Header{}
	if r.header != nil {
		header = r.header.Clone()
	}
	if bodyString != "" {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		body = bytes.NewBufferString(bodyString)
	}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}

	c.debug("full url: %s, body: %s", fullURL, bodyString)

	r.fullURL = fullURL
	r.header = header
	r.body = body
	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("request: %#v", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		if err == nil && cerr != nil {
			err = cerr
		}
	}()

	c.debug("response: %#v", res)
	c.debug("response body: %s", string(data))
	c.debug("response status code: %d", res.StatusCode)

	if res.StatusCode >= http.StatusBadRequest {
		apiErr := new(APIError)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s", e)
		}
		return nil, apiErr
	}
	return data, nil
}

func (c *Client) SetApiEndpoint(url string) *Client {
	c.BaseURL = url
	return c
}

func (c *Client) NewSendService() *SendService {
	return &SendService{c: c}
}
func (c *Client) NewSendArrayService() *SendArrayService {
	return &SendArrayService{c: c}
}
func (c *Client) NewStatusService() *StatusService {
	return &StatusService{c: c}
}
func (c *Client) NewStatusByLocalIDService() *StatusByLocalIDService {
	return &StatusByLocalIDService{c: c}
}
func (c *Client) NewSelectService() *SelectService {
	return &SelectService{c: c}
}
func (c *Client) NewSelectOutboxService() *SelectOutboxService {
	return &SelectOutboxService{c: c}
}
func (c *Client) NewLatestOutBoxService() *LatestOutBoxService {
	return &LatestOutBoxService{c: c}
}
func (c *Client) NewCountOutboxService() *CountOutboxService {
	return &CountOutboxService{c: c}
}
func (c *Client) NewCancelService() *CancelService {
	return &CancelService{c: c}
}
func (c *Client) NewReceiveService() *ReceiveService {
	return &ReceiveService{c: c}
}
func (c *Client) NewCountInboxService() *CountInboxService {
	return &CountInboxService{c: c}
}
func (c *Client) NewInfoService() *InfoService {
	return &InfoService{c: c}
}
func (c *Client) NewLookupService() *LookupService {
	return &LookupService{c: c}
}

type (
	sendType int8
)

const (
	SendTypeNews        sendType = iota
	SendTypeRegular     sendType = iota
	SendTypeSaveStorage sendType = iota
	SendTypeApp         sendType = iota

	StatusQueue           int = 1
	StatusScheduled       int = 2
	StatusTelecom         int = 4
	StatusTelecomSimilar4 int = 5
	StatusFailed          int = 6
	StatusDelivered       int = 10
	StatusUndelivered     int = 11
	StatusRejectByUser    int = 13
	StatusBlock           int = 14
	StatusInvalid         int = 100
)

func (s sendType) int8() int8 {
	return int8(s)
}
