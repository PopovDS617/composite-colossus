package client

import (
	"bytes"
	"dist_calc/types"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	Endpoint string
}

func NewClient(endpoint string) *Client {
	return &Client{
		Endpoint: endpoint,
	}
}

func (c *Client) AggregateInvoice(distance types.Distance) error {
	httpc := http.DefaultClient
	b, err := json.Marshal(distance)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewReader(b))

	if err != nil {
		return err
	}

	res, err := httpc.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("the service responded with non 200 code %d", res.StatusCode)
	}
	return nil
}
