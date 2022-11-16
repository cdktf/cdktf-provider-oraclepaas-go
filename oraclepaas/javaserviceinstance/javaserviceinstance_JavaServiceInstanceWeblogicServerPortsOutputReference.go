package javaserviceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v2/jsii"

	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v2/javaserviceinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JavaServiceInstanceWeblogicServerPortsOutputReference interface {
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
	ContentPort() *float64
	SetContentPort(val *float64)
	ContentPortInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeploymentChannelPort() *float64
	SetDeploymentChannelPort(val *float64)
	DeploymentChannelPortInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *JavaServiceInstanceWeblogicServerPorts
	SetInternalValue(val *JavaServiceInstanceWeblogicServerPorts)
	PrivilegedContentPort() *float64
	SetPrivilegedContentPort(val *float64)
	PrivilegedContentPortInput() *float64
	PrivilegedSecuredContentPort() *float64
	SetPrivilegedSecuredContentPort(val *float64)
	PrivilegedSecuredContentPortInput() *float64
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
	ResetContentPort()
	ResetDeploymentChannelPort()
	ResetPrivilegedContentPort()
	ResetPrivilegedSecuredContentPort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JavaServiceInstanceWeblogicServerPortsOutputReference
type jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) ContentPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"contentPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) ContentPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"contentPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) DeploymentChannelPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentChannelPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) DeploymentChannelPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentChannelPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) InternalValue() *JavaServiceInstanceWeblogicServerPorts {
	var returns *JavaServiceInstanceWeblogicServerPorts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) PrivilegedContentPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privilegedContentPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) PrivilegedContentPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privilegedContentPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) PrivilegedSecuredContentPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privilegedSecuredContentPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) PrivilegedSecuredContentPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privilegedSecuredContentPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJavaServiceInstanceWeblogicServerPortsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JavaServiceInstanceWeblogicServerPortsOutputReference {
	_init_.Initialize()

	if err := validateNewJavaServiceInstanceWeblogicServerPortsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.javaServiceInstance.JavaServiceInstanceWeblogicServerPortsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJavaServiceInstanceWeblogicServerPortsOutputReference_Override(j JavaServiceInstanceWeblogicServerPortsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.javaServiceInstance.JavaServiceInstanceWeblogicServerPortsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference)SetContentPort(val *float64) {
	if err := j.validateSetContentPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentPort",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference)SetDeploymentChannelPort(val *float64) {
	if err := j.validateSetDeploymentChannelPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentChannelPort",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference)SetInternalValue(val *JavaServiceInstanceWeblogicServerPorts) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference)SetPrivilegedContentPort(val *float64) {
	if err := j.validateSetPrivilegedContentPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privilegedContentPort",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference)SetPrivilegedSecuredContentPort(val *float64) {
	if err := j.validateSetPrivilegedSecuredContentPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privilegedSecuredContentPort",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) ResetContentPort() {
	_jsii_.InvokeVoid(
		j,
		"resetContentPort",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) ResetDeploymentChannelPort() {
	_jsii_.InvokeVoid(
		j,
		"resetDeploymentChannelPort",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) ResetPrivilegedContentPort() {
	_jsii_.InvokeVoid(
		j,
		"resetPrivilegedContentPort",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) ResetPrivilegedSecuredContentPort() {
	_jsii_.InvokeVoid(
		j,
		"resetPrivilegedSecuredContentPort",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := j.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerPortsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

