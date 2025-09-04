package dodoModel

import "encoding/json"

type DodoAPIResponse struct {
	Status int             `json:"status"`
	Data   json.RawMessage `json:"data"`
}

type DodoRouteData struct {
	ResAmount            float64 `json:"resAmount"`
	BaseFeeAmount        float64 `json:"baseFeeAmount"`
	BaseFeeRate          float64 `json:"baseFeeRate"`
	ResPricePerToToken   float64 `json:"resPricePerToToken"`
	ResPricePerFromToken float64 `json:"resPricePerFromToken"`
	PriceImpact          int     `json:"priceImpact"`
	UseSource            string  `json:"useSource"`
	TargetDecimals       int     `json:"targetDecimals"`
	TargetApproveAddr    string  `json:"targetApproveAddr"`
	To                   string  `json:"to"`
	Data                 string  `json:"data"`
	MinReturnAmount      string  `json:"minReturnAmount"`
	GasLimit             string  `json:"gasLimit"`
	RouteInfo            struct {
		SubRouteTotalPart int `json:"subRouteTotalPart"`
		SubRoute          []struct {
			MidPathPart int `json:"midPathPart"`
			MidPath     []struct {
				FromToken         string `json:"fromToken"`
				ToToken           string `json:"toToken"`
				OneSplitTotalPart int    `json:"oneSplitTotalPart"`
				PoolDetails       []struct {
					PoolName      string `json:"poolName"`
					Pool          string `json:"pool"`
					PoolPart      int    `json:"poolPart"`
					PoolInAmount  string `json:"poolInAmount"`
					PoolOutAmount string `json:"poolOutAmount"`
					LpFee         string `json:"lpFee"`
					LpFeeRateBps  string `json:"lpFeeRateBps"`
					UpdatedAt     int    `json:"updatedAt"`
				} `json:"poolDetails"`
				FromAmount string `json:"fromAmount"`
				ToAmount   string `json:"toAmount"`
			} `json:"midPath"`
		} `json:"subRoute"`
	} `json:"routeInfo"`
	RoutePlan []struct {
		Pool        string `json:"pool"`
		Label       string `json:"label"`
		InputToken  string `json:"inputToken"`
		OutputToken string `json:"outputToken"`
		InAmount    string `json:"inAmount"`
		OutAmount   string `json:"outAmount"`
		Direction   int    `json:"direction"`
		MoreInfo    struct {
			Reserve0 string `json:"reserve0"`
			Reserve1 string `json:"reserve1"`
		} `json:"moreInfo"`
		PercentBps int `json:"percentBps"`
		UpdatedAt  int `json:"updatedAt"`
	} `json:"routePlan"`
	Value      string `json:"value"`
	RPCCounter int    `json:"rpcCounter"`
	ID         string `json:"id"`
}
