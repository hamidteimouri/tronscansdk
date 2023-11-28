package tronscansdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	httpClientGoLib "github.com/trustwallet/go-libs/client"
	"net/url"
)

func GetBalanceOfUsdt(ctx context.Context, addressWallet string) (string, error) {
	path := fmt.Sprintf("%s", AccountTokenList)

	req := httpClientGoLib.InitClient(BaseUrl, nil)

	nb := httpClientGoLib.NewReqBuilder()
	nb.Query(url.Values{
		"address": []string{addressWallet},
	})
	nb.Method("GET").PathStatic(path)
	body, err := req.Execute(ctx, nb.Build())
	if err != nil {
		return "", err
	}

	b := TronBalanceOfAccount{}
	err = json.Unmarshal(body, &b)
	if err != nil {
		return "", err
	}

	if b.Data == nil {
		return "", errors.New("failed to parse response of request")
	}

	for _, tokenList := range b.Data {
		if tokenList.TokenId == ContractAddressUsdt {
			return fmt.Sprint(tokenList.Quantity), nil
		}
	}
	return "0", nil
}
