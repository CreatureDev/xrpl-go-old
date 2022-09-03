.PHONY: \
	examples \
	account \
	clean

examples: account faucet

account: 
	go build -o bin/account examples/account/account.go

faucet:
	go build -o bin/faucet examples/faucet/faucet.go

clean:
	rm -rf bin
