package goC8

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/marvin-hansen/goC8/types"
	"github.com/marvin-hansen/goC8/utils"
	"github.com/valyala/fasthttp"
	"net/url"
	"time"
)

// used for JSON unmarshaling
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (c *Client) requestJsonResponse(req types.Requester, results types.JsonResponder) error {
	res, reqErr := c.do(req)
	if reqErr != nil {
		//println("Request error")
		return reqErr
	}
	// println(string(res.Body()))
	results.SetJsonMessage(res.Body())
	decErr := decode(res.Body(), results)
	if decErr != nil {
		// println("Decode error")
		return decErr
	}
	return nil
}

func (c *Client) Request(req types.Requester, results types.Responder) error {
	res, reqErr := c.do(req)
	if reqErr != nil {
		utils.DbgPrint("Request error", debug)
		utils.DbgPrint(reqErr.Error(), debug)
		return reqErr
	}

	utils.DbgPrint(string(res.Body()), debug)
	decErr := decode(res.Body(), results)
	if decErr != nil {
		// println("Decode error")
		return decErr
	}

	return nil
}

// do build & executes the actual Request from the rquester
// requester - implementation
// targetStatusCode the expected http status code i.e. 200
func (c *Client) do(r types.Requester) (*fasthttp.Response, *APIError) {
	req := c.newRequest(r)

	utils.DbgPrint("URI:"+req.URI().String(), debug)
	utils.DbgPrint("Payload: "+string(r.Payload()), debug)

	// fasthttp for http2.0
	res := fasthttp.AcquireResponse()
	err := c.HTTPC.DoTimeout(req, res, c.HTTPTimeout)
	if err != nil {
		utils.DbgPrint("http req error", debug)
		apiErr := getApiError(res)
		return nil, apiErr
	}

	if !checkStatusCode(res.StatusCode(), r.ResponseCode()) {
		utils.DbgPrint("Status: Not ok", debug)
		var resp = new(Response)
		if jsonErr := json.Unmarshal(res.Body(), resp); jsonErr != nil {
			apiErr := getApiError(res)
			return nil, apiErr
		}

		if !resp.Success {
			utils.DbgPrint("Response not success", debug)
			apiErr := getApiError(res)
			return nil, apiErr
		}
	}

	return res, nil
}

// checkStatusCode determines if an OK response came through.
func checkStatusCode(responseStatusCode, expectedStatusCode int) bool {
	switch responseStatusCode {
	case 200:
		return true
	case 201:
		return true
	case 202:
		return true
	case expectedStatusCode:
		return true
	default:
		return false
	}
}

func getApiError(res *fasthttp.Response) *APIError {
	apiErr := &APIError{}
	_ = decode(res.Body(), apiErr)
	return apiErr
}

func getUri(endpoint string, r types.Requester) *url.URL {
	var u = new(url.URL)

	if r.HasQueryParameter() {
		utils.DbgPrint("Has Parameters: ", debug)
		utils.DbgPrint("Parameters: "+r.GetQueryParameter(), debug)
		u, _ = url.ParseRequestURI(endpoint + r.GetQueryParameter())
	} else {
		u, _ = url.ParseRequestURI(endpoint)
	}

	u.Path = r.Path()

	return u
}

func (c *Client) newRequest(r types.Requester) *fasthttp.Request {
	u := getUri(c.Endpoint, r)
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(r.Method())
	req.SetRequestURI(u.String())
	body := r.Payload()
	req.SetBody(body)

	nonce := fmt.Sprintf("%d", int64(time.Now().UTC().UnixNano()/int64(time.Millisecond)))
	payload := nonce + r.Method() + u.Path

	if u.RawQuery != "" {
		u.RawQuery = r.Query()
		payload += "?" + u.RawQuery
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "apikey "+c.getApiKey())

	return req
}

func decode(inputJson []byte, outputObject interface{}) error {
	if len(inputJson) > 2 { // check for empty "{}" JSON first
		err := json.Unmarshal(inputJson, outputObject)
		if err != nil {
			return fmt.Errorf("JSON decode error")
		} else {
			return nil
		}
	}
	return nil // return nil in case of empty json
}
