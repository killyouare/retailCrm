package retail

import (
	"context"
	"crmtest/internal/pkg/http"
	logger "crmtest/internal/pkg/util"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
)

var (
	unsuccessfulError = fmt.Errorf("unsuccessful retail")
)

type Client interface {
	GetOrders(ctx context.Context, parameters *OrdersRequest) (OrdersResponse, error)
}

type client struct {
	httpClient http.Client
	logger     logger.Logger
	url        string
	token      string
}

func New(httpClient http.Client, logger logger.Logger, url, token string) (Client, error) {
	return &client{
		httpClient: httpClient,
		logger:     logger,
		url:        url,
		token:      token,
	}, nil
}

func (c *client) GetOrders(ctx context.Context, parameters *OrdersRequest) (OrdersResponse, error) {
	var resp OrdersResponse

	params, _ := query.Values(parameters)

	data, err := c.GetRequest(ctx, fmt.Sprintf("/orders?%s", params.Encode()))
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(data, &resp)

	if err != nil {
		return resp, err
	}

	if !resp.Success {
		return resp, unsuccessfulError
	}

	return resp, err
}

func (c *client) GetRequest(ctx context.Context, urlWithParameters string) ([]byte, error) {
	var prefix = "/api/v5"

	headers := map[string]string{
		"X-API-KEY": c.token,
	}

	return c.httpClient.Get(ctx, fmt.Sprintf("%s%s%s", c.url, prefix, urlWithParameters), headers)
}
