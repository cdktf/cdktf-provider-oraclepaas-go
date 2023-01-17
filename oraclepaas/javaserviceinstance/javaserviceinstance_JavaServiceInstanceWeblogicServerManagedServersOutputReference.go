package javaserviceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v3/jsii"

	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v3/javaserviceinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JavaServiceInstanceWeblogicServerManagedServersOutputReference interface {
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
	InitialHeapSize() *float64
	SetInitialHeapSize(val *float64)
	InitialHeapSizeInput() *float64
	InitialPermanentGeneration() *float64
	SetInitialPermanentGeneration(val *float64)
	InitialPermanentGenerationInput() *float64
	InternalValue() *JavaServiceInstanceWeblogicServerManagedServers
	SetInternalValue(val *JavaServiceInstanceWeblogicServerManagedServers)
	JvmArgs() *string
	SetJvmArgs(val *string)
	JvmArgsInput() *string
	MaxHeapSize() *float64
	SetMaxHeapSize(val *float64)
	MaxHeapSizeInput() *float64
	MaxPermanentGeneration() *float64
	SetMaxPermanentGeneration(val *float64)
	MaxPermanentGenerationInput() *float64
	OverwriteJvmArgs() interface{}
	SetOverwriteJvmArgs(val interface{})
	OverwriteJvmArgsInput() interface{}
	ServerCount() *float64
	SetServerCount(val *float64)
	ServerCountInput() *float64
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
	ResetInitialHeapSize()
	ResetInitialPermanentGeneration()
	ResetJvmArgs()
	ResetMaxHeapSize()
	ResetMaxPermanentGeneration()
	ResetOverwriteJvmArgs()
	ResetServerCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JavaServiceInstanceWeblogicServerManagedServersOutputReference
type jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) InitialHeapSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialHeapSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) InitialHeapSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialHeapSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) InitialPermanentGeneration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialPermanentGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) InitialPermanentGenerationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialPermanentGenerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) InternalValue() *JavaServiceInstanceWeblogicServerManagedServers {
	var returns *JavaServiceInstanceWeblogicServerManagedServers
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) JvmArgs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jvmArgs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) JvmArgsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jvmArgsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) MaxHeapSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxHeapSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) MaxHeapSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxHeapSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) MaxPermanentGeneration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPermanentGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) MaxPermanentGenerationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPermanentGenerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) OverwriteJvmArgs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overwriteJvmArgs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) OverwriteJvmArgsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overwriteJvmArgsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) ServerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serverCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) ServerCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serverCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJavaServiceInstanceWeblogicServerManagedServersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JavaServiceInstanceWeblogicServerManagedServersOutputReference {
	_init_.Initialize()

	if err := validateNewJavaServiceInstanceWeblogicServerManagedServersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.javaServiceInstance.JavaServiceInstanceWeblogicServerManagedServersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJavaServiceInstanceWeblogicServerManagedServersOutputReference_Override(j JavaServiceInstanceWeblogicServerManagedServersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.javaServiceInstance.JavaServiceInstanceWeblogicServerManagedServersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference)SetInitialHeapSize(val *float64) {
	if err := j.validateSetInitialHeapSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialHeapSize",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference)SetInitialPermanentGeneration(val *float64) {
	if err := j.validateSetInitialPermanentGenerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialPermanentGeneration",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference)SetInternalValue(val *JavaServiceInstanceWeblogicServerManagedServers) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference)SetJvmArgs(val *string) {
	if err := j.validateSetJvmArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jvmArgs",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference)SetMaxHeapSize(val *float64) {
	if err := j.validateSetMaxHeapSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxHeapSize",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference)SetMaxPermanentGeneration(val *float64) {
	if err := j.validateSetMaxPermanentGenerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPermanentGeneration",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference)SetOverwriteJvmArgs(val interface{}) {
	if err := j.validateSetOverwriteJvmArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overwriteJvmArgs",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference)SetServerCount(val *float64) {
	if err := j.validateSetServerCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverCount",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) ResetInitialHeapSize() {
	_jsii_.InvokeVoid(
		j,
		"resetInitialHeapSize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) ResetInitialPermanentGeneration() {
	_jsii_.InvokeVoid(
		j,
		"resetInitialPermanentGeneration",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) ResetJvmArgs() {
	_jsii_.InvokeVoid(
		j,
		"resetJvmArgs",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) ResetMaxHeapSize() {
	_jsii_.InvokeVoid(
		j,
		"resetMaxHeapSize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) ResetMaxPermanentGeneration() {
	_jsii_.InvokeVoid(
		j,
		"resetMaxPermanentGeneration",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) ResetOverwriteJvmArgs() {
	_jsii_.InvokeVoid(
		j,
		"resetOverwriteJvmArgs",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) ResetServerCount() {
	_jsii_.InvokeVoid(
		j,
		"resetServerCount",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerManagedServersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

