package dodoModel

type DodoRouteParams struct {
	ChainID          string `url:"chainId"`
	DeadLine         string `url:"deadLine"`
	APIKey           string `url:"apikey"`
	Slippage         string `url:"slippage"`
	Source           string `url:"source"`
	ToTokenAddress   string `url:"toTokenAddress"`
	FromTokenAddress string `url:"fromTokenAddress"`
	UserAddr         string `url:"userAddr"`
	EstimateGas      string `url:"estimateGas"`
	FromAmount       string `url:"fromAmount"`
}
