package javaserviceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/jsii"

	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/javaserviceinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JavaServiceInstanceWeblogicServerClusterOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	PathPrefixes() *[]*string
	SetPathPrefixes(val *[]*string)
	PathPrefixesInput() *[]*string
	ServerCount() *float64
	SetServerCount(val *float64)
	ServerCountInput() *float64
	ServersPerNode() *float64
	SetServersPerNode(val *float64)
	ServersPerNodeInput() *float64
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
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetPathPrefixes()
	ResetServerCount()
	ResetServersPerNode()
	ResetShape()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JavaServiceInstanceWeblogicServerClusterOutputReference
type jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) PathPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pathPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) PathPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pathPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) ServerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serverCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) ServerCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serverCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) ServersPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serversPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) ServersPerNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serversPerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) ShapeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewJavaServiceInstanceWeblogicServerClusterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) JavaServiceInstanceWeblogicServerClusterOutputReference {
	_init_.Initialize()

	if err := validateNewJavaServiceInstanceWeblogicServerClusterOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.javaServiceInstance.JavaServiceInstanceWeblogicServerClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewJavaServiceInstanceWeblogicServerClusterOutputReference_Override(j JavaServiceInstanceWeblogicServerClusterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.javaServiceInstance.JavaServiceInstanceWeblogicServerClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		j,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference)SetPathPrefixes(val *[]*string) {
	if err := j.validateSetPathPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathPrefixes",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference)SetServerCount(val *float64) {
	if err := j.validateSetServerCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverCount",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference)SetServersPerNode(val *float64) {
	if err := j.validateSetServersPerNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serversPerNode",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference)SetShape(val *string) {
	if err := j.validateSetShapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shape",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) ResetPathPrefixes() {
	_jsii_.InvokeVoid(
		j,
		"resetPathPrefixes",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) ResetServerCount() {
	_jsii_.InvokeVoid(
		j,
		"resetServerCount",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) ResetServersPerNode() {
	_jsii_.InvokeVoid(
		j,
		"resetServersPerNode",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) ResetShape() {
	_jsii_.InvokeVoid(
		j,
		"resetShape",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

