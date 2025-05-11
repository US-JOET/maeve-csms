// SPDX-License-Identifier: Apache-2.0

package ocpp201

type NotifyChargingLimitRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Number of the EVSE on which the limit has been applied.
	EvseId int `json:"evseId,omitempty" yaml:"evseId,omitempty" mapstructure:"evseId,omitempty"`
	// Number of the EVSE on which the limit has been applied.
	ChargingLimit ChargingLimitType `json:"chargingLimit,omitempty" yaml:"chargingLimit,omitempty" mapstructure:"chargingLimit,omitempty"`

	// ChargingSchedule corresponds to the JSON schema field "chargingSchedule".
	ChargingSchedule []ChargingScheduleType `json:"chargingSchedule" yaml:"chargingSchedule" mapstructure:"chargingSchedule"`
}

func (*NotifyChargingLimitRequestJson) IsRequest() {}

type ChargingLimitSourceEnumType string

const ChargingLimitSourceEnumEMS ChargingLimitSourceEnumType = "EMS"
const ChargingLimitSourceEnumOther ChargingLimitSourceEnumType = "Other"
const ChargingLimitSourceEnumSO ChargingLimitSourceEnumType = "SO"
const ChargingLimitSourceEnumCSO ChargingLimitSourceEnumType = "CSO"

type ChargingLimitType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

    ChargingLimitSource BootReasonEnumType `json:"chargingLimitSource" yaml:"chargingLimitSource" mapstructure:"chargingLimitSource"`

	IsGridCritical bool `json:"isGridCritical,omitempty" yaml:"isGridCritical,omitempty" mapstructure:"isGridCritical,omitempty"`
}
