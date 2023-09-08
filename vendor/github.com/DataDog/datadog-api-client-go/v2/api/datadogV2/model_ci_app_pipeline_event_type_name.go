// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"
)

// CIAppPipelineEventTypeName Type of the event.
type CIAppPipelineEventTypeName string

// List of CIAppPipelineEventTypeName.
const (
	CIAPPPIPELINEEVENTTYPENAME_CIPIPELINE CIAppPipelineEventTypeName = "cipipeline"
)

var allowedCIAppPipelineEventTypeNameEnumValues = []CIAppPipelineEventTypeName{
	CIAPPPIPELINEEVENTTYPENAME_CIPIPELINE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppPipelineEventTypeName) GetAllowedValues() []CIAppPipelineEventTypeName {
	return allowedCIAppPipelineEventTypeNameEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppPipelineEventTypeName) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppPipelineEventTypeName(value)
	return nil
}

// NewCIAppPipelineEventTypeNameFromValue returns a pointer to a valid CIAppPipelineEventTypeName
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppPipelineEventTypeNameFromValue(v string) (*CIAppPipelineEventTypeName, error) {
	ev := CIAppPipelineEventTypeName(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppPipelineEventTypeName: valid values are %v", v, allowedCIAppPipelineEventTypeNameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppPipelineEventTypeName) IsValid() bool {
	for _, existing := range allowedCIAppPipelineEventTypeNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppPipelineEventTypeName value.
func (v CIAppPipelineEventTypeName) Ptr() *CIAppPipelineEventTypeName {
	return &v
}