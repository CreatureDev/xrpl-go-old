package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CreatureDev/xrpl-go/pkg/types"
)

type Client struct {
	Address string
}

type httpRequest struct {
	Method string             `json:"method"`
	Params []types.XRPLParams `json:"params"`
}

func (c *Client) Request(method string, args types.XRPLParams) (types.XRPLResponse, error) {
	req := httpRequest{
		Method: method,
		Params: []types.XRPLParams{args},
	}
	body, _ := json.Marshal(req)
	resp, err := http.Post(c.Address, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return []byte{}, fmt.Errorf("Error executing %s request: %w", method, err)
	}
	defer resp.Body.Close()
	dat := make([]byte, resp.ContentLength)
	_, _ = resp.Body.Read(dat)

	// error check
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(resp.Status)
	}
	var xrpErr types.Error
	decode(dat, xrpErr)
	if xrpErr.Error != "" {
		return nil, fmt.Errorf(xrpErr.Error)
	}

	ret := args.ResponseContainer()
	decode(dat, ret)
	return ret, nil
}

func NewClient(addr string) *Client {
	// TODO validate address is a rippled node
	return &Client{
		Address: addr,
	}
}

type httpResponse struct {
	Result types.XRPLResponse `json:"result"`
}

func decode(d []byte, ret types.XRPLResponse) {
	res := &httpResponse{ret}
	json.Unmarshal(d, res)
}
