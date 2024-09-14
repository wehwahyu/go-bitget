package v2

import (
	"github.com/wehwahyu/go-bitget/constants"
	"github.com/wehwahyu/go-bitget/internal"
	"github.com/wehwahyu/go-bitget/internal/common"
)

type UserClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *UserClient) Init() *UserClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

func (p *UserClient) ApiKeyList(subAccountUid string) (string, error) {

	// "/api/v2/user/virtual-subaccount-apikey-list?subAccountUid=1"

	params := internal.NewParams()
	params["subAccountUid"] = subAccountUid

	uri := constants.V2_User + "/virtual-subaccount-apikey-list"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err
}
