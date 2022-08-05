package types

type XRPLParams interface {
	// MethodString returns the method/command string
	// associated with these parameters
	MethodString() string
	// ResponseContainer returns a new struct to store the response
	ResponseContainer() XRPLResponse
	// Valid confirms that all fields are in appropriate format
	Valid() error
}

type XRPLResponse interface {
}
