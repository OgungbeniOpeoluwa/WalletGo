package httpRequest

import (
	"WalletService/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func MakePostRequest(body map[string]any, requestUrl, key string) (io.Reader, error) {
	bodys, err := marshalRequestBody(body)
	if err != nil {
		logger.ErrorLogger(err)
		return nil, err
	}
	uri, err := url.Parse(requestUrl)
	if err != nil {
		logger.ErrorLogger(err)
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, uri.String(), bytes.NewReader(bodys))
	if err != nil {
		logger.ErrorLogger(err)
		return nil, err
	}
	request.Header.Set("Authorization", key)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resposne, err := client.Do(request)
	if err != nil {
		logger.ErrorLogger(err)
		return nil, err
	}
	if resposne.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed,status code:%v", resposne.StatusCode)
	}
	return resposne.Body, err
}

func DecodeRequestBody[T any](responseBody io.Reader, types T) (T, error) {
	err := json.NewDecoder(responseBody).Decode(&types)
	if err != nil {
		return types, err
	}
	return types, nil
}

func marshalRequestBody(body map[string]any) ([]byte, error) {
	bodys, err := json.Marshal(body)
	if err != nil {
		return bodys, err
	}
	return bodys, err
}
