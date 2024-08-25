package atapi_client

import (
	"context"

	"github.com/burbru/atapi"
)

type AtapiWrapper struct {
	Token string
	Api   *atapi.APIClient
}

func NewAtapiWrapper(token string) *AtapiWrapper {
	config := atapi.NewConfiguration()
	config.AddDefaultHeader("Authorization", "Bearer "+token)
	api := atapi.NewAPIClient(config)
	return &AtapiWrapper{
		Token: token,
		Api:   api,
	}
}

func (a *AtapiWrapper) GetMyBankAccount() ([]atapi.BankAccountView, error) {
	req := a.Api.BankAccountControllerAPI.GetMyBankAccountUsingGET(context.Background())
	bav, _, err := a.Api.BankAccountControllerAPI.GetMyBankAccountUsingGETExecute(req)
	if err != nil {
		return nil, err
	}
	return bav, nil
}
