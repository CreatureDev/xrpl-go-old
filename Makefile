.PHONY: \
	examples \
	account \
	clean

examples: account

account: 
	go build -o bin/account examples/account/account.go

clean:
	rm -rf bin
