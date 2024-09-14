package spot

import (
	"github.com/wehwahyu/go-bitget/internal"
	"github.com/wehwahyu/go-bitget/internal/common"
)

type SpotWalletClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *SpotWalletClient) Init() *SpotWalletClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

func (p *SpotWalletClient) Transfer(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/wallet/transfer", postBody)
	return resp, err
}

func (p *SpotWalletClient) DepositAddress(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/wallet/deposit-address", params)
	return resp, err
}

func (p *SpotWalletClient) Withdrawal(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/wallet/withdrawal", postBody)
	return resp, err
}

func (p *SpotWalletClient) WithdrawalRecords(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/wallet/withdrawal-records", params)
	return resp, err
}

func (p *SpotWalletClient) DepositRecords(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/wallet/deposit-records", params)
	return resp, err
}
