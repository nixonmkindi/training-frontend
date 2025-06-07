// Package rest allows for quick and easy access any REST or REST-like API.
package rest

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"net/url"
	"training-frontend/package/log"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type Client struct {
	user   string
	key    string
	secret string
	url    string //base url
}

type Response struct {
	StatusCode int
	Body       []byte
}

// New function creates a restful api client.
// baseUrl should be in the format: https://localhost:4321
func New(baseUrl string) *Client {
	//thread safe singletone initialised
	c := &Client{
		//user:   user,
		//key:    key,
		//secret: secret,
		url: baseUrl,
		//wsUrl:  WsUrl,
	}

	return c
}

// GetMethod call the end point and returns body in binary form
func (c *Client) GetMethod(ctx echo.Context, u string) *Response {

	response := &Response{}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		log.Errorf("error initializing new get request: %v", err)
		return response
	}

	accessToken, err := ctx.Cookie("access-token")
	if err == nil {
		//bearer := "Bearer " + accessToken.Value
		req.Header.Add("authorization", accessToken.Value)
		//pp.Printf("token: %v\n", accessToken.Value)
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Errorf("error obtaining post response: %v", err)
		return response
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Errorf("error reading post response body: %v", err)
		return response
	}
	response.StatusCode = res.StatusCode
	response.Body = data
	return response
}

func (c *Client) PostMethod(ctx echo.Context, u string, v url.Values) *Response {

	response := &Response{}
	req, err := http.NewRequest("POST", u, strings.NewReader(v.Encode()))
	if err != nil {
		log.Errorf("error initializing new request: %v", err)
		return response
	}
	accessToken, err := ctx.Cookie("access-token")
	if err == nil {
		bearer := accessToken.Value
		req.Header.Add("authorization", bearer)
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Errorf("error obtaining post response: %v", err)
		return response
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Errorf("error reading post response body: %v", err)
		return response
	}
	response.StatusCode = res.StatusCode
	response.Body = data
	return response
}

// Send consumes the api
// endpoint shoud be in this format '/user/list
// params should be filled by form values for post requests
// isPost should be true for post request and false for get request
func (c *Client) Send(ctx echo.Context, endpoint string, params map[string]string, isPost bool) *Response {

	u := c.url + endpoint
	var res *Response
	if isPost {
		urlValues := url.Values{}
		for k, v := range params {
			urlValues.Add(k, v)
		}
		urlValues.Encode()
		res = c.PostMethod(ctx, u, urlValues)
	} else {
		// Get method for public method
		res = c.GetMethod(ctx, u)
	}
	return res
}

//after adding signature

func (c *Client) Nonce() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func (c *Client) ToHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func (c *Client) Signature() (string, string) {
	nonce := c.Nonce()
	message := nonce + c.user + c.key
	signature := c.ToHmac256(message, c.secret)
	return signature, nonce
}
