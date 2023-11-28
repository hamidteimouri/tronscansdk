package tronscansdk

type TronBalanceOfAccount struct {
	Total int         `json:"total"`
	Data  []TokenList `json:"data"`
}

type TokenList struct {
	Amount           interface{} `json:"amount,omitempty"`
	Quantity         interface{} `json:"quantity"`
	TokenId          string      `json:"tokenId"`
	Level            string      `json:"level"`
	TokenPriceInUsd  float64     `json:"tokenPriceInUsd,omitempty"`
	TokenName        string      `json:"tokenName"`
	TokenAbbr        string      `json:"tokenAbbr"`
	TokenCanShow     int         `json:"tokenCanShow"`
	TokenLogo        string      `json:"tokenLogo"`
	TokenPriceInTrx  float64     `json:"tokenPriceInTrx,omitempty"`
	AmountInUsd      float64     `json:"amountInUsd,omitempty"`
	Balance          string      `json:"balance"`
	TokenDecimal     int         `json:"tokenDecimal"`
	TokenType        string      `json:"tokenType"`
	Vip              bool        `json:"vip"`
	NrOfTokenHolders int         `json:"nrOfTokenHolders,omitempty"`
	TransferCount    int         `json:"transferCount,omitempty"`
	Project          string      `json:"project,omitempty"`
}
