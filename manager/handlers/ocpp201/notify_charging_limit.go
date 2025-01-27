// SPDX-License-Identifier: Apache-2.0

package ocpp201

import (
	"context"
	"github.com/thoughtworks/maeve-csms/manager/ocpp"
	"github.com/thoughtworks/maeve-csms/manager/ocpp/ocpp201"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type NotifyChargingLimitHandler struct{}

func (h NotifyChargingLimitHandler) HandleCall(ctx context.Context, chargeStationId string, request ocpp.Request) (response ocpp.Response, err error) {
	req := request.(*ocpp201.NotifyChargingLimitRequestJson)

	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("notify_charging_limit.charging_limit_source", string(req.ChargingLimit.ChargingLimitSource)),
		attribute.Int("notify_charging_limit.evse_id", req.EvseId))

	return &ocpp201.NotifyChargingLimitResponseJson{}, nil
}
