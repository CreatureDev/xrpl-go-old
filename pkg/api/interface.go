package api

import "encoding/json"

type XRPLParams interface {
	// MethodString returns the method/command string
	// associated with these parameters
	MethodString() string
	// ResponseContainer returns a new struct to store the response
	DecodeResponse(json.RawMessage) XRPLResponse
	// Valid confirms that all fields are in appropriate format
	Validate() error
}

type XRPLResponse interface {
}
