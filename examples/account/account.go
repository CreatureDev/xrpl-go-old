package main

import (
	"fmt"

	"github.com/CreatureDev/xrpl-go"
	"github.com/CreatureDev/xrpl-go/pkg/types"
)

func main() {
	conn := xrpl.CreateConnection("https://s.altnet.rippletest.net:51234")
	accInfo := types.AccountInfoParams{BaseAccountParams: types.BaseAccountParams{Account: "rGioZYzQmjShBPrtEnktJzzjSBqXZdJvxf"}}
	resp, _ := conn.Submit(&accInfo)
	fmt.Printf("%+v\n", resp)
}
