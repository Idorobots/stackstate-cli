/*
StackState API

This API documentation page describes the StackState server API. The StackState UI and CLI use the StackState server API to configure and query StackState.  You can use the API for similar purposes.  Each request sent to the StackState server API must be authenticated using one of the available authentication methods.   *Note that the StackState receiver API, used to send topology, telemetry and traces to StackState, is not described on this page and requires a different authentication method.*  For more information on StackState, refer to the [StackState documentation](https://docs.stackstate.com).

API version: 5.2.0
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_api

import (
	"encoding/json"
)

// MetricTelemetryQuery struct for MetricTelemetryQuery
type MetricTelemetryQuery struct {
	Type                string                    `json:"_type"`
	Aggregation         DownsamplingMethod        `json:"aggregation"`
	Baseline            *Baseline                 `json:"baseline,omitempty"`
	Conditions          []TelemetryQueryCondition `json:"conditions"`
	GroupBy             []string                  `json:"groupBy,omitempty"`
	Id                  *int64                    `json:"id,omitempty"`
	LastUpdateTimestamp *int64                    `json:"lastUpdateTimestamp,omitempty"`
	MetricField         *string                   `json:"metricField,omitempty"`
}

// NewMetricTelemetryQuery instantiates a new MetricTelemetryQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricTelemetryQuery(type_ string, aggregation DownsamplingMethod, conditions []TelemetryQueryCondition) *MetricTelemetryQuery {
	this := MetricTelemetryQuery{}
	this.Type = type_
	this.Aggregation = aggregation
	this.Conditions = conditions
	return &this
}

// NewMetricTelemetryQueryWithDefaults instantiates a new MetricTelemetryQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricTelemetryQueryWithDefaults() *MetricTelemetryQuery {
	this := MetricTelemetryQuery{}
	return &this
}

// GetType returns the Type field value
func (o *MetricTelemetryQuery) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MetricTelemetryQuery) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MetricTelemetryQuery) SetType(v string) {
	o.Type = v
}

// GetAggregation returns the Aggregation field value
func (o *MetricTelemetryQuery) GetAggregation() DownsamplingMethod {
	if o == nil {
		var ret DownsamplingMethod
		return ret
	}

	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *MetricTelemetryQuery) GetAggregationOk() (*DownsamplingMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value
func (o *MetricTelemetryQuery) SetAggregation(v DownsamplingMethod) {
	o.Aggregation = v
}

// GetBaseline returns the Baseline field value if set, zero value otherwise.
func (o *MetricTelemetryQuery) GetBaseline() Baseline {
	if o == nil || o.Baseline == nil {
		var ret Baseline
		return ret
	}
	return *o.Baseline
}

// GetBaselineOk returns a tuple with the Baseline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTelemetryQuery) GetBaselineOk() (*Baseline, bool) {
	if o == nil || o.Baseline == nil {
		return nil, false
	}
	return o.Baseline, true
}

// HasBaseline returns a boolean if a field has been set.
func (o *MetricTelemetryQuery) HasBaseline() bool {
	if o != nil && o.Baseline != nil {
		return true
	}

	return false
}

// SetBaseline gets a reference to the given Baseline and assigns it to the Baseline field.
func (o *MetricTelemetryQuery) SetBaseline(v Baseline) {
	o.Baseline = &v
}

// GetConditions returns the Conditions field value
func (o *MetricTelemetryQuery) GetConditions() []TelemetryQueryCondition {
	if o == nil {
		var ret []TelemetryQueryCondition
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *MetricTelemetryQuery) GetConditionsOk() ([]TelemetryQueryCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *MetricTelemetryQuery) SetConditions(v []TelemetryQueryCondition) {
	o.Conditions = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *MetricTelemetryQuery) GetGroupBy() []string {
	if o == nil || o.GroupBy == nil {
		var ret []string
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTelemetryQuery) GetGroupByOk() ([]string, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *MetricTelemetryQuery) HasGroupBy() bool {
	if o != nil && o.GroupBy != nil {
		return true
	}

	return false
}

// SetGroupBy gets a reference to the given []string and assigns it to the GroupBy field.
func (o *MetricTelemetryQuery) SetGroupBy(v []string) {
	o.GroupBy = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetricTelemetryQuery) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTelemetryQuery) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetricTelemetryQuery) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *MetricTelemetryQuery) SetId(v int64) {
	o.Id = &v
}

// GetLastUpdateTimestamp returns the LastUpdateTimestamp field value if set, zero value otherwise.
func (o *MetricTelemetryQuery) GetLastUpdateTimestamp() int64 {
	if o == nil || o.LastUpdateTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdateTimestamp
}

// GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTelemetryQuery) GetLastUpdateTimestampOk() (*int64, bool) {
	if o == nil || o.LastUpdateTimestamp == nil {
		return nil, false
	}
	return o.LastUpdateTimestamp, true
}

// HasLastUpdateTimestamp returns a boolean if a field has been set.
func (o *MetricTelemetryQuery) HasLastUpdateTimestamp() bool {
	if o != nil && o.LastUpdateTimestamp != nil {
		return true
	}

	return false
}

// SetLastUpdateTimestamp gets a reference to the given int64 and assigns it to the LastUpdateTimestamp field.
func (o *MetricTelemetryQuery) SetLastUpdateTimestamp(v int64) {
	o.LastUpdateTimestamp = &v
}

// GetMetricField returns the MetricField field value if set, zero value otherwise.
func (o *MetricTelemetryQuery) GetMetricField() string {
	if o == nil || o.MetricField == nil {
		var ret string
		return ret
	}
	return *o.MetricField
}

// GetMetricFieldOk returns a tuple with the MetricField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTelemetryQuery) GetMetricFieldOk() (*string, bool) {
	if o == nil || o.MetricField == nil {
		return nil, false
	}
	return o.MetricField, true
}

// HasMetricField returns a boolean if a field has been set.
func (o *MetricTelemetryQuery) HasMetricField() bool {
	if o != nil && o.MetricField != nil {
		return true
	}

	return false
}

// SetMetricField gets a reference to the given string and assigns it to the MetricField field.
func (o *MetricTelemetryQuery) SetMetricField(v string) {
	o.MetricField = &v
}

func (o MetricTelemetryQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["aggregation"] = o.Aggregation
	}
	if o.Baseline != nil {
		toSerialize["baseline"] = o.Baseline
	}
	if true {
		toSerialize["conditions"] = o.Conditions
	}
	if o.GroupBy != nil {
		toSerialize["groupBy"] = o.GroupBy
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdateTimestamp != nil {
		toSerialize["lastUpdateTimestamp"] = o.LastUpdateTimestamp
	}
	if o.MetricField != nil {
		toSerialize["metricField"] = o.MetricField
	}
	return json.Marshal(toSerialize)
}

type NullableMetricTelemetryQuery struct {
	value *MetricTelemetryQuery
	isSet bool
}

func (v NullableMetricTelemetryQuery) Get() *MetricTelemetryQuery {
	return v.value
}

func (v *NullableMetricTelemetryQuery) Set(val *MetricTelemetryQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricTelemetryQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricTelemetryQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricTelemetryQuery(val *MetricTelemetryQuery) *NullableMetricTelemetryQuery {
	return &NullableMetricTelemetryQuery{value: val, isSet: true}
}

func (v NullableMetricTelemetryQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricTelemetryQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
