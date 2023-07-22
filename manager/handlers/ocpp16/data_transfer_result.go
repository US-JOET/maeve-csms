// SPDX-License-Identifier: Apache-2.0

package ocpp16

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"

	"github.com/thoughtworks/maeve-csms/manager/handlers"
	"github.com/thoughtworks/maeve-csms/manager/ocpp"
	types "github.com/thoughtworks/maeve-csms/manager/ocpp/ocpp16"
	"github.com/thoughtworks/maeve-csms/manager/schemas"
	"golang.org/x/exp/slog"
)

type DataTransferResultHandler struct {
	SchemaFS         fs.FS
	CallResultRoutes map[string]map[string]handlers.CallResultRoute
}

func (d DataTransferResultHandler) HandleCallResult(ctx context.Context, chargeStationId string, request ocpp.Request, response ocpp.Response, state any) error {
	req := request.(*types.DataTransferJson)
	resp := response.(*types.DataTransferResponseJson)

	messageId := ""
	if req.MessageId != nil {
		messageId = *req.MessageId
	}
	slog.Info("data transfer result",
		slog.String("vendorId", req.VendorId), slog.String("messageId", messageId))

	vendorMap, ok := d.CallResultRoutes[req.VendorId]
	if !ok {
		return fmt.Errorf("unknown data transfer result vendor: %s", req.VendorId)
	}
	route, ok := vendorMap[messageId]
	if !ok {
		return fmt.Errorf("unknown data transfer result message id: %s", messageId)
	}

	var dataTransferRequest ocpp.Request
	if req.Data != nil {
		data := []byte(*req.Data)
		err := schemas.Validate(data, d.SchemaFS, route.RequestSchema)
		if err != nil {
			return fmt.Errorf("validating %s:%s data transfer result request data: %w", req.VendorId, messageId, err)
		}
		dataTransferRequest = route.NewRequest()
		err = json.Unmarshal(data, &dataTransferRequest)
		if err != nil {
			return fmt.Errorf("unmarshalling %s:%s data transfer request data: %w", req.VendorId, messageId, err)
		}
	}

	var dataTransferResponse ocpp.Response
	if resp.Data != nil {
		data := []byte(*resp.Data)
		err := schemas.Validate(data, d.SchemaFS, route.ResponseSchema)
		if err != nil {
			return fmt.Errorf("validating %s:%s data transfer result response data: %w", req.VendorId, messageId, err)
		}
		dataTransferResponse = route.NewResponse()
		err = json.Unmarshal(data, &dataTransferResponse)
		if err != nil {
			return fmt.Errorf("unmarshalling %s:%s data transfer response data: %w", req.VendorId, messageId, err)
		}
	}

	return route.Handler.HandleCallResult(ctx, chargeStationId, dataTransferRequest, dataTransferResponse, state)
}
