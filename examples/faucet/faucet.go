package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CreatureDev/xrpl-go"
	"github.com/CreatureDev/xrpl-go/pkg/api"
	"github.com/CreatureDev/xrpl-go/pkg/wallet"
)

const (
	faucet_url string = "https://faucet.altnet.rippletest.net/accounts"
)

type RequestFunding struct {
	Account api.Account `json:"destination"`
}

func main() {
	w, err := wallet.GenerateWallet([]byte{})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Generated wallet: %s\n", w.Account)

	conn := xrpl.CreateConnection("https://s.altnet.rippletest.net:51234")

	res, err := conn.Submit(&api.AccountInfoParams{
		BaseAccountParams: api.BaseAccountParams{Account: w.Account},
	})
	if err != nil && err.Error() != "actNotFound" {
		fmt.Println(err.Error())
	}

	var info *api.AccountInfoResponse

	info, _ = res.(*api.AccountInfoResponse)

	fmt.Printf("%+v\n", info)

	// request funding
	req, _ := json.Marshal(RequestFunding{Account: w.Account})
	resp, err := http.Post(faucet_url, "application/json", bytes.NewBuffer(req))
	if err != nil {
		fmt.Println("Failed to request funding: " + err.Error())
	}
	defer resp.Body.Close()
	dat := make([]byte, resp.ContentLength)
	_, _ = resp.Body.Read(dat)

	fmt.Println("RESP\n" + string(dat))

	res, err = conn.Submit(&api.AccountInfoParams{
		BaseAccountParams: api.BaseAccountParams{Account: w.Account},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	info, _ = res.(*api.AccountInfoResponse)

	fmt.Printf("%+v\n", info)

}
