/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-31T04:35:53Z.
 *
 * API version: 1.0.9-2110
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// ConnectorHttpRequestAllOf Definition of the list of properties defined in 'connector.HttpRequest', excluding properties defined in parent classes.
type ConnectorHttpRequestAllOf struct {
	// Contents of the request body to send for PUT/PATCH/POST requests.
	Body *string `json:"Body,omitempty"`
	// The timeout for establishing the TCP connection to the target host. If not set the request timeout value is used.
	DialTimeout *int64 `json:"DialTimeout,omitempty"`
	// Collection of key value pairs to set in the request header.
	Header interface{} `json:"Header,omitempty"`
	// The request is for an internal platform API that requires authentication to be inserted by the platform implementation.
	Internal *bool `json:"Internal,omitempty"`
	// Method specifies the HTTP method (GET, POST, PUT, etc.). For client requests an empty string means GET.
	Method *string `json:"Method,omitempty"`
	// The timeout for the HTTP request to complete, from connection establishment to response body read complete. If not set a default timeout of five minutes is used.
	Timeout              *int64        `json:"Timeout,omitempty"`
	Url                  *ConnectorUrl `json:"Url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorHttpRequestAllOf ConnectorHttpRequestAllOf

// NewConnectorHttpRequestAllOf instantiates a new ConnectorHttpRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorHttpRequestAllOf() *ConnectorHttpRequestAllOf {
	this := ConnectorHttpRequestAllOf{}
	return &this
}

// NewConnectorHttpRequestAllOfWithDefaults instantiates a new ConnectorHttpRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorHttpRequestAllOfWithDefaults() *ConnectorHttpRequestAllOf {
	this := ConnectorHttpRequestAllOf{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *ConnectorHttpRequestAllOf) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequestAllOf) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *ConnectorHttpRequestAllOf) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *ConnectorHttpRequestAllOf) SetBody(v string) {
	o.Body = &v
}

// GetDialTimeout returns the DialTimeout field value if set, zero value otherwise.
func (o *ConnectorHttpRequestAllOf) GetDialTimeout() int64 {
	if o == nil || o.DialTimeout == nil {
		var ret int64
		return ret
	}
	return *o.DialTimeout
}

// GetDialTimeoutOk returns a tuple with the DialTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequestAllOf) GetDialTimeoutOk() (*int64, bool) {
	if o == nil || o.DialTimeout == nil {
		return nil, false
	}
	return o.DialTimeout, true
}

// HasDialTimeout returns a boolean if a field has been set.
func (o *ConnectorHttpRequestAllOf) HasDialTimeout() bool {
	if o != nil && o.DialTimeout != nil {
		return true
	}

	return false
}

// SetDialTimeout gets a reference to the given int64 and assigns it to the DialTimeout field.
func (o *ConnectorHttpRequestAllOf) SetDialTimeout(v int64) {
	o.DialTimeout = &v
}

// GetHeader returns the Header field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectorHttpRequestAllOf) GetHeader() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectorHttpRequestAllOf) GetHeaderOk() (*interface{}, bool) {
	if o == nil || o.Header == nil {
		return nil, false
	}
	return &o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *ConnectorHttpRequestAllOf) HasHeader() bool {
	if o != nil && o.Header != nil {
		return true
	}

	return false
}

// SetHeader gets a reference to the given interface{} and assigns it to the Header field.
func (o *ConnectorHttpRequestAllOf) SetHeader(v interface{}) {
	o.Header = v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *ConnectorHttpRequestAllOf) GetInternal() bool {
	if o == nil || o.Internal == nil {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequestAllOf) GetInternalOk() (*bool, bool) {
	if o == nil || o.Internal == nil {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *ConnectorHttpRequestAllOf) HasInternal() bool {
	if o != nil && o.Internal != nil {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *ConnectorHttpRequestAllOf) SetInternal(v bool) {
	o.Internal = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *ConnectorHttpRequestAllOf) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequestAllOf) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *ConnectorHttpRequestAllOf) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *ConnectorHttpRequestAllOf) SetMethod(v string) {
	o.Method = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *ConnectorHttpRequestAllOf) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequestAllOf) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *ConnectorHttpRequestAllOf) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *ConnectorHttpRequestAllOf) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ConnectorHttpRequestAllOf) GetUrl() ConnectorUrl {
	if o == nil || o.Url == nil {
		var ret ConnectorUrl
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequestAllOf) GetUrlOk() (*ConnectorUrl, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ConnectorHttpRequestAllOf) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given ConnectorUrl and assigns it to the Url field.
func (o *ConnectorHttpRequestAllOf) SetUrl(v ConnectorUrl) {
	o.Url = &v
}

func (o ConnectorHttpRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["Body"] = o.Body
	}
	if o.DialTimeout != nil {
		toSerialize["DialTimeout"] = o.DialTimeout
	}
	if o.Header != nil {
		toSerialize["Header"] = o.Header
	}
	if o.Internal != nil {
		toSerialize["Internal"] = o.Internal
	}
	if o.Method != nil {
		toSerialize["Method"] = o.Method
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}
	if o.Url != nil {
		toSerialize["Url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorHttpRequestAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConnectorHttpRequestAllOf := _ConnectorHttpRequestAllOf{}

	if err = json.Unmarshal(bytes, &varConnectorHttpRequestAllOf); err == nil {
		*o = ConnectorHttpRequestAllOf(varConnectorHttpRequestAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Body")
		delete(additionalProperties, "DialTimeout")
		delete(additionalProperties, "Header")
		delete(additionalProperties, "Internal")
		delete(additionalProperties, "Method")
		delete(additionalProperties, "Timeout")
		delete(additionalProperties, "Url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorHttpRequestAllOf struct {
	value *ConnectorHttpRequestAllOf
	isSet bool
}

func (v NullableConnectorHttpRequestAllOf) Get() *ConnectorHttpRequestAllOf {
	return v.value
}

func (v *NullableConnectorHttpRequestAllOf) Set(val *ConnectorHttpRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorHttpRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorHttpRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorHttpRequestAllOf(val *ConnectorHttpRequestAllOf) *NullableConnectorHttpRequestAllOf {
	return &NullableConnectorHttpRequestAllOf{value: val, isSet: true}
}

func (v NullableConnectorHttpRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorHttpRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
