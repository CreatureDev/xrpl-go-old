package main

import (
	"fmt"

	"github.com/CreatureDev/xrpl-go"
	"github.com/CreatureDev/xrpl-go/pkg/api"
)

func main() {
	conn := xrpl.CreateConnection("https://s.altnet.rippletest.net:51234")
	accInfo := api.AccountInfoParams{BaseAccountParams: api.BaseAccountParams{Account: "rGioZYzQmjShBPrtEnktJzzjSBqXZdJvxf"}}
	resp, _ := conn.Submit(&accInfo)
	fmt.Printf("%+v\n", resp)
}
