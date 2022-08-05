package main

import (
	"fmt"

	"github.com/CreatureDev/xrpl-go"
	"github.com/CreatureDev/xrpl-go/pkg/types"
)

func main() {
	conn := xrpl.CreateConnection("https://s1.ripple.com:51234")
	accInfo := types.AccountInfoParams{BaseAccountParams: types.BaseAccountParams{Account: "rnbAFdgPs2YEmp5PEcDZypUq3839GtmfpJ"}}
	resp := conn.Submit(&accInfo)
	fmt.Printf("%+v\n", resp)
}
