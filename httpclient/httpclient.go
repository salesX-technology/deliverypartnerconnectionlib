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
	baseURL string
	client  *netHTTP.Client
}

func NewHTTPFormPoster[ResponseBody any](baseURL string, client *netHTTP.Client) *httpFormPoster[ResponseBody] {
	return &httpFormPoster[ResponseBody]{
		client:  client,
		baseURL: baseURL,
	}
}

func (h *httpFormPoster[ResponseBody]) PostForm(endpoint string, form map[string]string) (ResponseBody, error) {
	fmt.Printf("making request to %s\n", h.baseURL+endpoint)
	fmt.Printf("request: \n%s\n\n", form)

	var responseBodyModel ResponseBody
	data := url.Values{}
	for k, v := range form {
		data.Set(k, v)
	}

	req, err := netHTTP.NewRequest("POST", h.baseURL+endpoint, bytes.NewBufferString(data.Encode()))
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

	fmt.Printf("body: \n%s\n\n", string(body))
	return responseBodyModel, nil
}

type httpPosterInstance[Request any, Response any] struct {
	httpClient    *netHTTP.Client
	baseURL       string
	defaultHeader map[string]string
}

func NewHTTPPoster[Request any, Response any](httpClient *netHTTP.Client, baseUrl string, defaultHeader map[string]string) httpPosterInstance[Request, Response] {
	return httpPosterInstance[Request, Response]{
		httpClient:    httpClient,
		baseURL:       baseUrl,
		defaultHeader: defaultHeader,
	}
}

func (i httpPosterInstance[Request, Response]) Post(url string, headers map[string]string, request Request) (Response, error) {
	var m Response

	fmt.Printf("making request to %s\n", i.baseURL+url)
	reqBody, err := json.Marshal(request)
	if err != nil {
		return m, fmt.Errorf("mashal request error: %s", err)
	}
	fmt.Printf("request: \n%s\n", string(reqBody))

	req, err := netHTTP.NewRequest(netHTTP.MethodPost, i.baseURL+url, bytes.NewReader(reqBody))
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
		return m, fmt.Errorf("error sending request to %s: %v", i.baseURL+url, err)
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

	fmt.Printf("response: \n%s\n", string(responseBody))
	return m, nil
}

func (i httpPosterInstance[Request, Response]) PostHook(url string, headers map[string]string, request Request) (Response, error) {
	var m Response

	fmt.Printf("making request to %s\n", i.baseURL+url)
	reqBody, err := json.Marshal(request)
	if err != nil {
		return m, fmt.Errorf("mashal request error: %s", err)
	}
	fmt.Printf("request: \n%s\n", string(reqBody))

	req, err := netHTTP.NewRequest(netHTTP.MethodPost, i.baseURL+url, bytes.NewReader(reqBody))
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
		return m, fmt.Errorf("error sending request to %s: %v", i.baseURL+url, err)
	}

	defer httpResponse.Body.Close()
	responseBody, err := io.ReadAll(httpResponse.Body)

	// fmt.Printf("response: \n%s\n", string(responseBody))
	if err != nil {
		return m, fmt.Errorf("error reading response body: %s", err)
	}

	if httpResponse.StatusCode < 200 || httpResponse.StatusCode >= 300 {
		return m, fmt.Errorf("request not successful with status code %d", httpResponse.StatusCode)
	}

	if err := json.Unmarshal(responseBody, &m); err != nil {
		return m, fmt.Errorf("error occurred when unmarshalling response body: %s", err)
	} else {
		fmt.Printf("response: \n%s\n", string(responseBody))
	}

	return m, nil
}

func (i httpPosterInstance[Request, Response]) PostHookSPX(url string, headers map[string]string, request Request) (Response, []byte, error) {
	var m Response

	// Construct full URL
	fullURL := i.baseURL + url
	fmt.Printf("making request to %s\n", fullURL)

	// Marshal the request body
	reqBody, err := json.Marshal(request)
	if err != nil {
		return m, nil, fmt.Errorf("marshal request error: %v", err)
	}
	fmt.Printf("request: \n%s\n", string(reqBody))

	// Create a new HTTP POST request
	req, err := netHTTP.NewRequest(netHTTP.MethodPost, fullURL, bytes.NewReader(reqBody))
	if err != nil {
		return m, nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set default headers
	for hKey, kValue := range i.defaultHeader {
		req.Header.Set(hKey, kValue)
	}

	// Set additional headers
	for hKey, hValue := range headers {
		req.Header.Set(hKey, hValue)
	}

	// Send the HTTP request
	httpResponse, err := i.httpClient.Do(req)
	if err != nil {
		return m, nil, fmt.Errorf("error sending request to %s: %v", fullURL, err)
	}
	defer httpResponse.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return m, nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Check for HTTP status codes indicating failure
	if httpResponse.StatusCode < 200 || httpResponse.StatusCode >= 300 {
		return m, responseBody, fmt.Errorf("request not successful with status code %d", httpResponse.StatusCode)
	}

	// Try to unmarshal the response body into the response object
	if err := json.Unmarshal(responseBody, &m); err != nil {
		return m, responseBody, fmt.Errorf("error unmarshalling response body: %v", err)
	}

	// Print the response for debugging purposes
	fmt.Printf("response: \n%s\n", responseBody)

	return m, responseBody, nil
}

func NewHTTPGetter[Request any, Response any](httpClient *netHTTP.Client, baseURL string, defaultHeader map[string]string) httpPosterInstance[Request, Response] {
	return httpPosterInstance[Request, Response]{
		httpClient:    httpClient,
		baseURL:       baseURL,
		defaultHeader: defaultHeader,
	}
}

func (i httpPosterInstance[Request, Response]) Get(endpoint string, headers map[string]string, request Request) (Response, error) {
	var m Response

	fmt.Printf("making request to %s\n", i.baseURL+endpoint)
	req, err := netHTTP.NewRequest(netHTTP.MethodGet, i.baseURL+endpoint, nil)
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
		return m, fmt.Errorf("error sending request to %s: %v", i.baseURL+endpoint, err)
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

	fmt.Printf("response: \n%s\n", string(responseBody))
	return m, nil
}
