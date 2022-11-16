package databaseserviceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v2/jsii"

	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v2/databaseserviceinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseServiceInstanceBackupsOutputReference interface {
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
	CreateIfMissing() interface{}
	SetCreateIfMissing(val interface{})
	CreateIfMissingInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DatabaseServiceInstanceBackups
	SetInternalValue(val *DatabaseServiceInstanceBackups)
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
	ResetCreateIfMissing()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabaseServiceInstanceBackupsOutputReference
type jsiiProxy_DatabaseServiceInstanceBackupsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) CloudStorageContainer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) CloudStorageContainerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) CloudStoragePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStoragePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) CloudStoragePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStoragePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) CloudStorageUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) CloudStorageUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) CreateIfMissing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createIfMissing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) CreateIfMissingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createIfMissingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) InternalValue() *DatabaseServiceInstanceBackups {
	var returns *DatabaseServiceInstanceBackups
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatabaseServiceInstanceBackupsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatabaseServiceInstanceBackupsOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseServiceInstanceBackupsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseServiceInstanceBackupsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.databaseServiceInstance.DatabaseServiceInstanceBackupsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabaseServiceInstanceBackupsOutputReference_Override(d DatabaseServiceInstanceBackupsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.databaseServiceInstance.DatabaseServiceInstanceBackupsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference)SetCloudStorageContainer(val *string) {
	if err := j.validateSetCloudStorageContainerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudStorageContainer",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference)SetCloudStoragePassword(val *string) {
	if err := j.validateSetCloudStoragePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudStoragePassword",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference)SetCloudStorageUsername(val *string) {
	if err := j.validateSetCloudStorageUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudStorageUsername",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference)SetCreateIfMissing(val interface{}) {
	if err := j.validateSetCreateIfMissingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createIfMissing",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference)SetInternalValue(val *DatabaseServiceInstanceBackups) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) ResetCloudStoragePassword() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudStoragePassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) ResetCloudStorageUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudStorageUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) ResetCreateIfMissing() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateIfMissing",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabaseServiceInstanceBackupsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

