package xrpl

import (
	"github.com/CreatureDev/xrpl-go/internal/http"
	"github.com/CreatureDev/xrpl-go/pkg/types"
)

type Connection struct {
	c                 Client
	disableValidation bool
}

// CreateConnection will create a new synchronous http connection to the specified XRPL server address.
func CreateConnection(address string) *Connection {
	return &Connection{
		c: http.NewClient(address),
	}
}

// DisableValidation will disable parameter validation for all subsequent requests.
func (c *Connection) DisableValidation() {
	c.disableValidation = true
}

// EnableValidation will enable parameter validation for all subsequent requests. Enabled by default.
func (c *Connection) EnableValidation() {
	c.disableValidation = false
}

// Submit is a generic function to pass requests to connected XRPL Node
// Response will be in the format associated with request parameters
func (c *Connection) Submit(p types.XRPLParams) (types.XRPLResponse, error) {
	if !c.disableValidation {
		if err := p.Validate(); err != nil {
			return nil, err
		}
	}
	resp, err := c.c.Request(p.MethodString(), p)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
