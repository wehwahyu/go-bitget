package account

/**
 * @Author: bitget-sdk-team
 * @Date: 2022-09-30 10:46
 * @DES: Coin withdrawals request
 */
type WithdrawReq struct {
	Coin         string `json:"coin"`
	TransferType string `json:"transferType"`
	Address      string `json:"address"`
	Chain        string `json:"chain"`
	Size         string `json:"size"`
}
