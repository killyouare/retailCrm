package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	failMarshal        = "failed to marshal payload: %w"
	failCreateRequest  = "failed to create %s request: %w"
	failExecuteRequest = "failed to execute %s request: %w"
	failReadResponse   = "failed to read response body: %w"
	invalidStatusCode  = "unexpected status code: %d, body: %s"
)

type Client interface {
	Get(ctx context.Context, url string, headers map[string]string) ([]byte, error)
	Post(ctx context.Context, url string, headers map[string]string, payload interface{}) ([]byte, error)
}

type client struct {
	client *http.Client
}

func New(httpClient *http.Client) (Client, error) {
	return &client{
		client: httpClient,
	}, nil
}

func (h *client) Get(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	return h.sendRequest(ctx, http.MethodGet, url, headers, nil)
}

func (h *client) Post(ctx context.Context, url string, headers map[string]string, payload interface{}) ([]byte, error) {
	return h.sendRequest(ctx, http.MethodPost, url, headers, payload)
}

func (h *client) sendRequest(ctx context.Context, method string, url string, headers map[string]string, payload interface{}) ([]byte, error) {
	var bodyBytes []byte
	if payload != nil {
		var err error
		bodyBytes, err = json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf(failMarshal, err)
		}
	}
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf(failCreateRequest, method, err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf(failExecuteRequest, method, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf(failReadResponse, err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf(invalidStatusCode, resp.StatusCode, string(body))
	}

	return body, nil
}
