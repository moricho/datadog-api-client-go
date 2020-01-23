# MetricsQueryResponseSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggr** | Pointer to **string** | Aggregation type | [optional] [readonly] 
**DisplayName** | Pointer to **string** | Display name of the metric | [optional] [readonly] 
**End** | Pointer to **int64** | End of the time window, milliseconds since Unix epoch | [optional] [readonly] 
**Expression** | Pointer to **string** | Metric expression | [optional] [readonly] 
**Interval** | Pointer to **int64** | Number of seconds between data samples | [optional] [readonly] 
**Length** | Pointer to **int64** | Number of data samples | [optional] [readonly] 
**Metric** | Pointer to **string** | Metric name | [optional] [readonly] 
**Pointlist** | Pointer to [**[][]float64**](array.md) | List of points of the time series | [optional] [readonly] 
**Scope** | Pointer to **string** | Metric scope, comma separated list of tags | [optional] [readonly] 
**Start** | Pointer to **int64** | Start of the time window, milliseconds since Unix epoch | [optional] [readonly] 
**Unit** | Pointer to [**[]MetricsQueryResponseUnit**](MetricsQueryResponse_unit.md) | Detailed information about the metric unit. First element describes the \&quot;primary unit\&quot; (e.g. &#x60;bytes&#x60; in &#x60;bytes per second&#x60;), second describes the \&quot;per unit\&quot; (e.g. &#x60;second&#x60; in &#x60;bytes per second&#x60;) | [optional] [readonly] 

## Methods

### GetAggr

`func (o *MetricsQueryResponseSeries) GetAggr() string`

GetAggr returns the Aggr field if non-nil, zero value otherwise.

### GetAggrOk

`func (o *MetricsQueryResponseSeries) GetAggrOk() (string, bool)`

GetAggrOk returns a tuple with the Aggr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAggr

`func (o *MetricsQueryResponseSeries) HasAggr() bool`

HasAggr returns a boolean if a field has been set.

### SetAggr

`func (o *MetricsQueryResponseSeries) SetAggr(v string)`

SetAggr gets a reference to the given string and assigns it to the Aggr field.

### GetDisplayName

`func (o *MetricsQueryResponseSeries) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MetricsQueryResponseSeries) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MetricsQueryResponseSeries) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MetricsQueryResponseSeries) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetEnd

`func (o *MetricsQueryResponseSeries) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *MetricsQueryResponseSeries) GetEndOk() (int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnd

`func (o *MetricsQueryResponseSeries) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEnd

`func (o *MetricsQueryResponseSeries) SetEnd(v int64)`

SetEnd gets a reference to the given int64 and assigns it to the End field.

### GetExpression

`func (o *MetricsQueryResponseSeries) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *MetricsQueryResponseSeries) GetExpressionOk() (string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpression

`func (o *MetricsQueryResponseSeries) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### SetExpression

`func (o *MetricsQueryResponseSeries) SetExpression(v string)`

SetExpression gets a reference to the given string and assigns it to the Expression field.

### GetInterval

`func (o *MetricsQueryResponseSeries) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *MetricsQueryResponseSeries) GetIntervalOk() (int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInterval

`func (o *MetricsQueryResponseSeries) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetInterval

`func (o *MetricsQueryResponseSeries) SetInterval(v int64)`

SetInterval gets a reference to the given int64 and assigns it to the Interval field.

### GetLength

`func (o *MetricsQueryResponseSeries) GetLength() int64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *MetricsQueryResponseSeries) GetLengthOk() (int64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLength

`func (o *MetricsQueryResponseSeries) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLength

`func (o *MetricsQueryResponseSeries) SetLength(v int64)`

SetLength gets a reference to the given int64 and assigns it to the Length field.

### GetMetric

`func (o *MetricsQueryResponseSeries) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *MetricsQueryResponseSeries) GetMetricOk() (string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMetric

`func (o *MetricsQueryResponseSeries) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### SetMetric

`func (o *MetricsQueryResponseSeries) SetMetric(v string)`

SetMetric gets a reference to the given string and assigns it to the Metric field.

### GetPointlist

`func (o *MetricsQueryResponseSeries) GetPointlist() [][]float64`

GetPointlist returns the Pointlist field if non-nil, zero value otherwise.

### GetPointlistOk

`func (o *MetricsQueryResponseSeries) GetPointlistOk() ([][]float64, bool)`

GetPointlistOk returns a tuple with the Pointlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPointlist

`func (o *MetricsQueryResponseSeries) HasPointlist() bool`

HasPointlist returns a boolean if a field has been set.

### SetPointlist

`func (o *MetricsQueryResponseSeries) SetPointlist(v [][]float64)`

SetPointlist gets a reference to the given [][]float64 and assigns it to the Pointlist field.

### GetScope

`func (o *MetricsQueryResponseSeries) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *MetricsQueryResponseSeries) GetScopeOk() (string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScope

`func (o *MetricsQueryResponseSeries) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScope

`func (o *MetricsQueryResponseSeries) SetScope(v string)`

SetScope gets a reference to the given string and assigns it to the Scope field.

### GetStart

`func (o *MetricsQueryResponseSeries) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *MetricsQueryResponseSeries) GetStartOk() (int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStart

`func (o *MetricsQueryResponseSeries) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStart

`func (o *MetricsQueryResponseSeries) SetStart(v int64)`

SetStart gets a reference to the given int64 and assigns it to the Start field.

### GetUnit

`func (o *MetricsQueryResponseSeries) GetUnit() []MetricsQueryResponseUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricsQueryResponseSeries) GetUnitOk() ([]MetricsQueryResponseUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnit

`func (o *MetricsQueryResponseSeries) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnit

`func (o *MetricsQueryResponseSeries) SetUnit(v []MetricsQueryResponseUnit)`

SetUnit gets a reference to the given []MetricsQueryResponseUnit and assigns it to the Unit field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

