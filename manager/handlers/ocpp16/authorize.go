// SPDX-License-Identifier: Apache-2.0

package ocpp16

import (
	"context"
	"github.com/thoughtworks/maeve-csms/manager/ocpp"
	types "github.com/thoughtworks/maeve-csms/manager/ocpp/ocpp16"
	"github.com/thoughtworks/maeve-csms/manager/store"
	"log"
)

type AuthorizeHandler struct {
	TokenStore store.TokenStore
}

func (a AuthorizeHandler) HandleCall(ctx context.Context, chargeStationId string, request ocpp.Request) (ocpp.Response, error) {
	req := request.(*types.AuthorizeJson)
	log.Printf("Charge station %s authorize token %s", chargeStationId, req.IdTag)

	status := types.AuthorizeResponseJsonIdTagInfoStatusInvalid
	tok, err := a.TokenStore.LookupToken(ctx, req.IdTag)
	if err != nil {
		return nil, err
	}
	if tok != nil {
		status = types.AuthorizeResponseJsonIdTagInfoStatusAccepted
	}

	return &types.AuthorizeResponseJson{
		IdTagInfo: types.AuthorizeResponseJsonIdTagInfo{
			Status: status,
		},
	}, nil
}
