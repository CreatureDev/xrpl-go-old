package xrpl

import (
	"fmt"
	"os"

	"github.com/CreatureDev/xrpl-go/internal/http"
	"github.com/CreatureDev/xrpl-go/pkg/types"
)

type Connection struct {
	c Client
}

func CreateConnection(address string) *Connection {
	return &Connection{
		c: http.NewClient(address),
	}
}

// Submit is a generic function to pass requests to connected XRPL Node
// Response will be in the format associated with request parameters
func (c *Connection) Submit(p types.XRPLParams) types.XRPLResponse {
	if !p.Valid() {
		fmt.Fprintln(os.Stderr, "Request improperly formatted")
		return nil
	}
	resp, err := c.c.Request(p.MethodString(), p)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}
	return resp
}