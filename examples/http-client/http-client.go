package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type RequestClient struct {
	Headers map[string]string
	Body    interface{}
}
type Response struct {
	StatusCode int
	Body       []byte
}

func (request RequestClient) setHeaders(req *http.Request) {
	for key, value := range request.Headers {
		req.Header.Add(key, value)
	}
}

func (request RequestClient) Get(url string) (Response, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Response{}, nil
	}

	request.setHeaders(req)

	var response Response

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, nil
	}

	response.StatusCode = res.StatusCode
	response.Body = body

	return response, nil
}

func (request RequestClient) Post(url string) (Response, error) {

	if request.Body == nil {
		return Response{}, errors.New("body empty")
	}

	fmt.Println(request.Body)
	jsonBody, err := json.Marshal(request.Body)
	if err != nil {
		return Response{}, nil
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return Response{}, nil
	}

	request.setHeaders(req)

	var response Response

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, err
	}

	response.StatusCode = res.StatusCode
	response.Body = body

	return response, nil
}
