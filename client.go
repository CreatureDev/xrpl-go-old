package xrpl

import "github.com/CreatureDev/xrpl-go/pkg/types"

type Client interface {
	Request(string, types.XRPLParams) (types.XRPLResponse, error)
}
