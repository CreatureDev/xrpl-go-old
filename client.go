package xrpl

import "github.com/CreatureDev/xrpl-go/pkg/api"

type Client interface {
	Request(string, api.XRPLParams) (api.XRPLResponse, error)
}
