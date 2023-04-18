package databaseserviceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v4/jsii"

	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v4/databaseserviceinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseServiceInstanceHybridDisasterRecoveryOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() *DatabaseServiceInstanceHybridDisasterRecovery
	SetInternalValue(val *DatabaseServiceInstanceHybridDisasterRecovery)
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
	ResetCloudStoragePassword()
	ResetCloudStorageUsername()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabaseServiceInstanceHybridDisasterRecoveryOutputReference
type jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) CloudStorageContainer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) CloudStorageContainerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) CloudStoragePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStoragePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) CloudStoragePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStoragePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) CloudStorageUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) CloudStorageUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) InternalValue() *DatabaseServiceInstanceHybridDisasterRecovery {
	var returns *DatabaseServiceInstanceHybridDisasterRecovery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatabaseServiceInstanceHybridDisasterRecoveryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatabaseServiceInstanceHybridDisasterRecoveryOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseServiceInstanceHybridDisasterRecoveryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.databaseServiceInstance.DatabaseServiceInstanceHybridDisasterRecoveryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabaseServiceInstanceHybridDisasterRecoveryOutputReference_Override(d DatabaseServiceInstanceHybridDisasterRecoveryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.databaseServiceInstance.DatabaseServiceInstanceHybridDisasterRecoveryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference)SetCloudStorageContainer(val *string) {
	if err := j.validateSetCloudStorageContainerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudStorageContainer",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference)SetCloudStoragePassword(val *string) {
	if err := j.validateSetCloudStoragePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudStoragePassword",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference)SetCloudStorageUsername(val *string) {
	if err := j.validateSetCloudStorageUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudStorageUsername",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference)SetInternalValue(val *DatabaseServiceInstanceHybridDisasterRecovery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) ResetCloudStoragePassword() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudStoragePassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) ResetCloudStorageUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudStorageUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceHybridDisasterRecoveryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

