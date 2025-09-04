package contracts

type NetworkID int

const (
	UOMI_TESTNET NetworkID = iota
)

type DAppID int

const (
	SYTHRA DAppID = iota
)

type ContractID int

const (
	UNISWAPV3_UNIVERSAL_ROUTER ContractID = iota
	UNISWAPV3_FACTORY
	UNISWAPV3_MULTICALL
	UNISWAPV3_QUOTER
	UNISWAPV3_QUOTERV2
	UNISWAPV3_V3MIGRATOR
	UNISWAPV3_NONFUNGIBLEPOSITIONMANAGER
	UNISWAPV3_TICKLENS
	UNISWAPV3_SWAPROUTER
	UNISWAPV3_PERMIT2
)

type TokenID int

const (
	NATIVE TokenID = iota
	WRAPPED
	USDC
	SYNT
)

type UniversalCommand byte

const (
	V3_SWAP_EXACT_IN      UniversalCommand = 0x00
	V3_SWAP_EXACT_OUT     UniversalCommand = 0x01
	PERMIT2_TRANSFER_FROM UniversalCommand = 0x02
	PERMIT2_PERMIT_BATCH  UniversalCommand = 0x03
	SWEEP                 UniversalCommand = 0x04
	TRANSFER              UniversalCommand = 0x05
	PAY_PORTION           UniversalCommand = 0x06
	// 0x07 is unused
	V2_SWAP_EXACT_IN            UniversalCommand = 0x08
	V2_SWAP_EXACT_OUT           UniversalCommand = 0x09
	PERMIT2_PERMIT              UniversalCommand = 0x0a
	WRAP_ETH                    UniversalCommand = 0x0b
	UNWRAP_WETH                 UniversalCommand = 0x0c
	PERMIT2_TRANSFER_FROM_BATCH UniversalCommand = 0x0d
	// 0x0e-0x1f are unused
	EXECUTE_SUB_PLAN UniversalCommand = 0x20
	// 0x21-0x3f are unused
)

var UniversalCommandABIInputs = map[UniversalCommand]string{
	V3_SWAP_EXACT_IN: `[{"name":"swap","type":"function","inputs":[
		{"name":"recipient","type":"address"},
		{"name":"amountIn","type":"uint256"},
		{"name":"amountOutMin","type":"uint256"},
		{"name":"path","type":"bytes"},
		{"name":"payerIsSender","type":"bool"}
	]}]`,
	V3_SWAP_EXACT_OUT: `[{"name":"swap","type":"function","inputs":[
		{"name":"recipient","type":"address"},
		{"name":"amountOut","type":"uint256"},
		{"name":"amountInMax","type":"uint256"},
		{"name":"path","type":"bytes"},
		{"name":"payerIsSender","type":"bool"}
	]}]`,
	WRAP_ETH: `[{"name":"wrap","type":"function","inputs":[
		{"name":"recipient","type":"address"},
		{"name":"amountIn","type":"uint256"}
	]}]`,
	UNWRAP_WETH: `[{"name":"unwrap","type":"function","inputs":[
		{"name":"recipient","type":"address"},
		{"name":"amountMin","type":"uint256"}
	]}]`,
	V2_SWAP_EXACT_IN: `[{"name":"swap","type":"function","inputs":[
		{"name":"recipient","type":"address"},
		{"name":"amountIn","type":"uint256"},
		{"name":"amountOutMin","type":"uint256"},
		{"name":"path","type":"address[]"},
		{"name":"payerIsSender","type":"bool"}]}`,
	V2_SWAP_EXACT_OUT: `[{"name":"swap","type":"function","inputs":[
		{"name":"recipient","type":"address"},
		{"name":"amountOut","type":"uint256"},
		{"name":"amountInMax","type":"uint256"},
		{"name":"path","type":"address[]"},
		{"name":"payerIsSender","type":"bool"}]}`,
	PERMIT2_TRANSFER_FROM: `[{"name":"transfer","type":"function","inputs":[
		{"name":"token","type":"address"},
		{"name":"recipient","type":"address"},
		{"name":"amount","type":"uint160"}]}`,
	SWEEP: `[{"name":"sweep","type":"function","inputs":[
		{"name":"token","type":"address"},
		{"name":"recipient","type":"address"},
		{"name":"amountMin","type":"uint256"}]}`,
	TRANSFER: `[{"name":"transfer","type":"function","inputs":[
		{"name":"token","type":"address"},
		{"name":"recipient","type":"address"},
		{"name":"amount","type":"uint256"}]}`,
	PAY_PORTION: `[{"name":"pay","type":"function","inputs":[
		{"name":"token","type":"address"},
		{"name":"recipient","type":"address"},
		{"name":"bips","type":"uint256"}]}`,
	PERMIT2_PERMIT: `[{"name":"permit","type":"function","inputs":[
		{"name":"permitSingle","type":"tuple","components":[
			{"name":"details","type":"tuple","components":[
				{"name":"token","type":"address"},
				{"name":"amount","type":"uint160"},
				{"name":"expiration","type":"uint48"},
				{"name":"nonce","type":"uint48"}
			]},
			{"name":"spender","type":"address"},
			{"name":"sigDeadline","type":"uint256"}
		]},
		{"name":"signature","type":"bytes"}
	]}]`,
	PERMIT2_PERMIT_BATCH: `[{"name":"permitBatch","type":"function","inputs":[
		{"name":"permitBatch","type":"tuple[]","components":[
			{"name":"details","type":"tuple","components":[
				{"name":"token","type":"address"},
				{"name":"amount","type":"uint160"},
				{"name":"expiration","type":"uint48"},
				{"name":"nonce","type":"uint48"}
			]},
			{"name":"spender","type":"address"},
			{"name":"sigDeadline","type":"uint256"}
		]},
		{"name":"signature","type":"bytes"}
	]}]`,
}
