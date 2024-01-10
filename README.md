# tronscansdk
A small sdk for TronScan Api

## How to install
```shell
go get -u github.com/hamidteimouri/tronscansdk
```

### Get Balance Of Usdt
There is a method named `GetBalanceOfUsdt` that returns `USDT (TRC20)` balance of a wallet.
```go
package main

import (
	"context"
	"fmt"
	"github.com/hamidteimouri/tronscansdk"
)

func main() {
	balance, err := tronscansdk.GetBalanceOfUsdt(context.Background(), "your-wallet-address")
	if err != nil {
		panic(err)
	}
	fmt.Println("balance : " + balance)
}
```