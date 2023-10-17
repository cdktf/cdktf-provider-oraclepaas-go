// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mysqlserviceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v7/jsii"

	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v7/mysqlserviceinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EmAgentPassword() *string
	SetEmAgentPassword(val *string)
	EmAgentPasswordInput() *string
	EmAgentUsername() *string
	SetEmAgentUsername(val *string)
	EmAgentUsernameInput() *string
	EmPassword() *string
	SetEmPassword(val *string)
	EmPasswordInput() *string
	EmPort() *float64
	SetEmPort(val *float64)
	EmPortInput() *float64
	EmUsername() *string
	SetEmUsername(val *string)
	EmUsernameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfiguration
	SetInternalValue(val *MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfiguration)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetEmAgentPassword()
	ResetEmAgentUsername()
	ResetEmPassword()
	ResetEmPort()
	ResetEmUsername()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference
type jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) EmAgentPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emAgentPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) EmAgentPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emAgentPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) EmAgentUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emAgentUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) EmAgentUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emAgentUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) EmPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) EmPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) EmPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"emPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) EmPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"emPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) EmUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) EmUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) InternalValue() *MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfiguration {
	var returns *MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewMysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.mysqlServiceInstance.MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference_Override(m MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.mysqlServiceInstance.MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference)SetEmAgentPassword(val *string) {
	if err := j.validateSetEmAgentPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emAgentPassword",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference)SetEmAgentUsername(val *string) {
	if err := j.validateSetEmAgentUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emAgentUsername",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference)SetEmPassword(val *string) {
	if err := j.validateSetEmPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emPassword",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference)SetEmPort(val *float64) {
	if err := j.validateSetEmPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emPort",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference)SetEmUsername(val *string) {
	if err := j.validateSetEmUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emUsername",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference)SetInternalValue(val *MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) ResetEmAgentPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetEmAgentPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) ResetEmAgentUsername() {
	_jsii_.InvokeVoid(
		m,
		"resetEmAgentUsername",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) ResetEmPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetEmPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) ResetEmPort() {
	_jsii_.InvokeVoid(
		m,
		"resetEmPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) ResetEmUsername() {
	_jsii_.InvokeVoid(
		m,
		"resetEmUsername",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstanceMysqlConfigurationEnterpriseMonitorConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

