package javaserviceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/jsii"

	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/javaserviceinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JavaServiceInstanceOracleTrafficDirectorOutputReference interface {
	cdktf.ComplexObject
	Admin() JavaServiceInstanceOracleTrafficDirectorAdminOutputReference
	AdminInput() *JavaServiceInstanceOracleTrafficDirectorAdmin
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
	HighAvailability() interface{}
	SetHighAvailability(val interface{})
	HighAvailabilityInput() interface{}
	InternalValue() *JavaServiceInstanceOracleTrafficDirector
	SetInternalValue(val *JavaServiceInstanceOracleTrafficDirector)
	IpReservations() *[]*string
	SetIpReservations(val *[]*string)
	IpReservationsInput() *[]*string
	Listener() JavaServiceInstanceOracleTrafficDirectorListenerOutputReference
	ListenerInput() *JavaServiceInstanceOracleTrafficDirectorListener
	LoadBalancingPolicy() *string
	SetLoadBalancingPolicy(val *string)
	LoadBalancingPolicyInput() *string
	RootUrl() *string
	Shape() *string
	SetShape(val *string)
	ShapeInput() *string
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
	PutAdmin(value *JavaServiceInstanceOracleTrafficDirectorAdmin)
	PutListener(value *JavaServiceInstanceOracleTrafficDirectorListener)
	ResetHighAvailability()
	ResetIpReservations()
	ResetListener()
	ResetLoadBalancingPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JavaServiceInstanceOracleTrafficDirectorOutputReference
type jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) Admin() JavaServiceInstanceOracleTrafficDirectorAdminOutputReference {
	var returns JavaServiceInstanceOracleTrafficDirectorAdminOutputReference
	_jsii_.Get(
		j,
		"admin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) AdminInput() *JavaServiceInstanceOracleTrafficDirectorAdmin {
	var returns *JavaServiceInstanceOracleTrafficDirectorAdmin
	_jsii_.Get(
		j,
		"adminInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) HighAvailability() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"highAvailability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) HighAvailabilityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"highAvailabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) InternalValue() *JavaServiceInstanceOracleTrafficDirector {
	var returns *JavaServiceInstanceOracleTrafficDirector
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) IpReservations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipReservations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) IpReservationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipReservationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) Listener() JavaServiceInstanceOracleTrafficDirectorListenerOutputReference {
	var returns JavaServiceInstanceOracleTrafficDirectorListenerOutputReference
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) ListenerInput() *JavaServiceInstanceOracleTrafficDirectorListener {
	var returns *JavaServiceInstanceOracleTrafficDirectorListener
	_jsii_.Get(
		j,
		"listenerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) LoadBalancingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) LoadBalancingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) RootUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) ShapeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJavaServiceInstanceOracleTrafficDirectorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JavaServiceInstanceOracleTrafficDirectorOutputReference {
	_init_.Initialize()

	if err := validateNewJavaServiceInstanceOracleTrafficDirectorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.javaServiceInstance.JavaServiceInstanceOracleTrafficDirectorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJavaServiceInstanceOracleTrafficDirectorOutputReference_Override(j JavaServiceInstanceOracleTrafficDirectorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.javaServiceInstance.JavaServiceInstanceOracleTrafficDirectorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference)SetHighAvailability(val interface{}) {
	if err := j.validateSetHighAvailabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"highAvailability",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference)SetInternalValue(val *JavaServiceInstanceOracleTrafficDirector) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference)SetIpReservations(val *[]*string) {
	if err := j.validateSetIpReservationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipReservations",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference)SetLoadBalancingPolicy(val *string) {
	if err := j.validateSetLoadBalancingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancingPolicy",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference)SetShape(val *string) {
	if err := j.validateSetShapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shape",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) PutAdmin(value *JavaServiceInstanceOracleTrafficDirectorAdmin) {
	if err := j.validatePutAdminParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAdmin",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) PutListener(value *JavaServiceInstanceOracleTrafficDirectorListener) {
	if err := j.validatePutListenerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putListener",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) ResetHighAvailability() {
	_jsii_.InvokeVoid(
		j,
		"resetHighAvailability",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) ResetIpReservations() {
	_jsii_.InvokeVoid(
		j,
		"resetIpReservations",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) ResetListener() {
	_jsii_.InvokeVoid(
		j,
		"resetListener",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) ResetLoadBalancingPolicy() {
	_jsii_.InvokeVoid(
		j,
		"resetLoadBalancingPolicy",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JavaServiceInstanceOracleTrafficDirectorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

