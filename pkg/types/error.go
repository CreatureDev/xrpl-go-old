package types

type Error struct {
	Error   string     `json:"error"`
	Status  string     `json:"error"`
	Request XRPLParams `json:"request"`
}
