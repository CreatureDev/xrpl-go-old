package api

type Error struct {
	Error   string      `json:"error"`
	Status  string      `json:"status"`
	Request interface{} `json:"request"`
}
