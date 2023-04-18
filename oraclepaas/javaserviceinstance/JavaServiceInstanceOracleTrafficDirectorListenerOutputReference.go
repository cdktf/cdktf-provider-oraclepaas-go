package javaserviceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v4/jsii"

	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v4/javaserviceinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JavaServiceInstanceOracleTrafficDirectorListenerOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *JavaServiceInstanceOracleTrafficDirectorListener
	SetInternalValue(val *JavaServiceInstanceOracleTrafficDirectorListener)
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PrivilegedPort() *float64
	SetPrivilegedPort(val *float64)
	PrivilegedPortInput() *float64
	PrivilegedSecuredPort() *float64
	SetPrivilegedSecuredPort(val *float64)
	PrivilegedSecuredPortInput() *float64
	SecuredPort() *float64
	SetSecuredPort(val *float64)
	SecuredPortInput() *float64
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
	ResetPort()
	ResetPrivilegedPort()
	ResetPrivilegedSecuredPort()
	ResetSecuredPort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JavaServiceInstanceOracleTrafficDirectorListenerOutputReference
type jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) InternalValue() *JavaServiceInstanceOracleTrafficDirectorListener {
	var returns *JavaServiceInstanceOracleTrafficDirectorListener
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) PrivilegedPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privilegedPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) PrivilegedPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privilegedPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) PrivilegedSecuredPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privilegedSecuredPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) PrivilegedSecuredPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privilegedSecuredPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) SecuredPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"securedPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) SecuredPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"securedPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJavaServiceInstanceOracleTrafficDirectorListenerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JavaServiceInstanceOracleTrafficDirectorListenerOutputReference {
	_init_.Initialize()

	if err := validateNewJavaServiceInstanceOracleTrafficDirectorListenerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.javaServiceInstance.JavaServiceInstanceOracleTrafficDirectorListenerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJavaServiceInstanceOracleTrafficDirectorListenerOutputReference_Override(j JavaServiceInstanceOracleTrafficDirectorListenerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.javaServiceInstance.JavaServiceInstanceOracleTrafficDirectorListenerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference)SetInternalValue(val *JavaServiceInstanceOracleTrafficDirectorListener) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference)SetPrivilegedPort(val *float64) {
	if err := j.validateSetPrivilegedPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privilegedPort",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference)SetPrivilegedSecuredPort(val *float64) {
	if err := j.validateSetPrivilegedSecuredPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privilegedSecuredPort",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference)SetSecuredPort(val *float64) {
	if err := j.validateSetSecuredPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securedPort",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		j,
		"resetPort",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) ResetPrivilegedPort() {
	_jsii_.InvokeVoid(
		j,
		"resetPrivilegedPort",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) ResetPrivilegedSecuredPort() {
	_jsii_.InvokeVoid(
		j,
		"resetPrivilegedSecuredPort",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) ResetSecuredPort() {
	_jsii_.InvokeVoid(
		j,
		"resetSecuredPort",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorListenerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

