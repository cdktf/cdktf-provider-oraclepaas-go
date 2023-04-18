package javaserviceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v4/jsii"

	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v4/javaserviceinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JavaServiceInstanceBackupsOutputReference interface {
	cdktf.ComplexObject
	AutoGenerate() interface{}
	SetAutoGenerate(val interface{})
	AutoGenerateInput() interface{}
	CloudStorageContainer() *string
	SetCloudStorageContainer(val *string)
	CloudStorageContainerInput() *string
	CloudStoragePassword() *string
	SetCloudStoragePassword(val *string)
	CloudStoragePasswordInput() *string
	CloudStorageUsername() *string
	SetCloudStorageUsername(val *string)
	CloudStorageUsernameInput() *string
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
	InternalValue() *JavaServiceInstanceBackups
	SetInternalValue(val *JavaServiceInstanceBackups)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseOauthForStorage() interface{}
	SetUseOauthForStorage(val interface{})
	UseOauthForStorageInput() interface{}
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
	ResetAutoGenerate()
	ResetCloudStoragePassword()
	ResetCloudStorageUsername()
	ResetUseOauthForStorage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JavaServiceInstanceBackupsOutputReference
type jsiiProxy_JavaServiceInstanceBackupsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) AutoGenerate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoGenerate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) AutoGenerateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoGenerateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) CloudStorageContainer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) CloudStorageContainerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) CloudStoragePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStoragePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) CloudStoragePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStoragePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) CloudStorageUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) CloudStorageUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) InternalValue() *JavaServiceInstanceBackups {
	var returns *JavaServiceInstanceBackups
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) UseOauthForStorage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOauthForStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) UseOauthForStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOauthForStorageInput",
		&returns,
	)
	return returns
}


func NewJavaServiceInstanceBackupsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JavaServiceInstanceBackupsOutputReference {
	_init_.Initialize()

	if err := validateNewJavaServiceInstanceBackupsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JavaServiceInstanceBackupsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.javaServiceInstance.JavaServiceInstanceBackupsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJavaServiceInstanceBackupsOutputReference_Override(j JavaServiceInstanceBackupsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.javaServiceInstance.JavaServiceInstanceBackupsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference)SetAutoGenerate(val interface{}) {
	if err := j.validateSetAutoGenerateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoGenerate",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference)SetCloudStorageContainer(val *string) {
	if err := j.validateSetCloudStorageContainerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudStorageContainer",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference)SetCloudStoragePassword(val *string) {
	if err := j.validateSetCloudStoragePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudStoragePassword",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference)SetCloudStorageUsername(val *string) {
	if err := j.validateSetCloudStorageUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudStorageUsername",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference)SetInternalValue(val *JavaServiceInstanceBackups) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference)SetUseOauthForStorage(val interface{}) {
	if err := j.validateSetUseOauthForStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useOauthForStorage",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) ResetAutoGenerate() {
	_jsii_.InvokeVoid(
		j,
		"resetAutoGenerate",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) ResetCloudStoragePassword() {
	_jsii_.InvokeVoid(
		j,
		"resetCloudStoragePassword",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) ResetCloudStorageUsername() {
	_jsii_.InvokeVoid(
		j,
		"resetCloudStorageUsername",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) ResetUseOauthForStorage() {
	_jsii_.InvokeVoid(
		j,
		"resetUseOauthForStorage",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JavaServiceInstanceBackupsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

