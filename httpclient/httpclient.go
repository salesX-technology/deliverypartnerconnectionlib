package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"

	"io"
	"net/url"

	netHTTP "net/http"
)

type httpFormPoster[ResponseBody any] struct {
	client *netHTTP.Client
}

func NewHTTPFormPoster[ResponseBody any](client *netHTTP.Client) *httpFormPoster[ResponseBody] {
	return &httpFormPoster[ResponseBody]{
		client: client,
	}
}

func (h *httpFormPoster[ResponseBody]) PostForm(fullurl string, form map[string]string) (ResponseBody, error) {
	fmt.Printf("request: %s\n\n", form)

	var responseBodyModel ResponseBody
	data := url.Values{}
	for k, v := range form {
		data.Set(k, v)
	}

	req, err := netHTTP.NewRequest("POST", fullurl, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return responseBodyModel, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := h.client.Do(req)
	if err != nil {
		return responseBodyModel, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return responseBodyModel, fmt.Errorf("error reading response body: %w", err)
	}

	if err := json.Unmarshal(body, &responseBodyModel); err != nil {
		return responseBodyModel, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	fmt.Printf("body: %s\n\n", string(body))
	return responseBodyModel, nil
}

type httpPosterInstance[Request any, Response any] struct {
	httpClient    *netHTTP.Client
	fullURL       string
	defaultHeader map[string]string
}

func NewHTTPPoster[Request any, Response any](httpClient *netHTTP.Client, fullURL string, defaultHeader map[string]string) httpPosterInstance[Request, Response] {
	return httpPosterInstance[Request, Response]{
		httpClient:    httpClient,
		fullURL:       fullURL,
		defaultHeader: defaultHeader,
	}
}

func (i httpPosterInstance[Request, Response]) Post(headers map[string]string, request Request) (Response, error) {
	var m Response

	reqBody, err := json.Marshal(request)
	if err != nil {
		return m, fmt.Errorf("mashal request error: %s", err)
	}
	fmt.Printf("request: %s\n", string(reqBody))

	req, err := netHTTP.NewRequest(netHTTP.MethodPost, i.fullURL, bytes.NewReader(reqBody))
	if err != nil {
		return m, fmt.Errorf("error creating request: %s", err)
	}

	for hKey, kValue := range i.defaultHeader {
		req.Header.Set(hKey, kValue)
	}

	for hKey, hValue := range headers {
		req.Header.Set(hKey, hValue)
	}

	httpResponse, err := i.httpClient.Do(req)
	if err != nil {
		return m, fmt.Errorf("error sending request to %s: %v", i.fullURL, err)
	}

	defer httpResponse.Body.Close()
	responseBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return m, fmt.Errorf("error reading response body: %s", err)
	}

	if httpResponse.StatusCode < 200 || httpResponse.StatusCode >= 300 {
		return m, fmt.Errorf("request not successful with status code %d", httpResponse.StatusCode)
	}

	if err := json.Unmarshal(responseBody, &m); err != nil {
		return m, fmt.Errorf("error occurred when unmarshalling response body: %s", err)
	}

	fmt.Printf("response: %s\n", string(responseBody))
	return m, nil
}
