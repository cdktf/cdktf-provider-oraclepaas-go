package mysqlserviceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v3/jsii"

	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v3/mysqlserviceinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MysqlServiceInstanceBackupsOutputReference interface {
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
	InternalValue() *MysqlServiceInstanceBackups
	SetInternalValue(val *MysqlServiceInstanceBackups)
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

// The jsii proxy struct for MysqlServiceInstanceBackupsOutputReference
type jsiiProxy_MysqlServiceInstanceBackupsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) CloudStorageContainer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) CloudStorageContainerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) CloudStoragePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStoragePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) CloudStoragePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStoragePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) CloudStorageUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) CloudStorageUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) CreateIfMissing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createIfMissing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) CreateIfMissingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createIfMissingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) InternalValue() *MysqlServiceInstanceBackups {
	var returns *MysqlServiceInstanceBackups
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMysqlServiceInstanceBackupsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MysqlServiceInstanceBackupsOutputReference {
	_init_.Initialize()

	if err := validateNewMysqlServiceInstanceBackupsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MysqlServiceInstanceBackupsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.mysqlServiceInstance.MysqlServiceInstanceBackupsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMysqlServiceInstanceBackupsOutputReference_Override(m MysqlServiceInstanceBackupsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.mysqlServiceInstance.MysqlServiceInstanceBackupsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference)SetCloudStorageContainer(val *string) {
	if err := j.validateSetCloudStorageContainerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudStorageContainer",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference)SetCloudStoragePassword(val *string) {
	if err := j.validateSetCloudStoragePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudStoragePassword",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference)SetCloudStorageUsername(val *string) {
	if err := j.validateSetCloudStorageUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudStorageUsername",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference)SetCreateIfMissing(val interface{}) {
	if err := j.validateSetCreateIfMissingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createIfMissing",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference)SetInternalValue(val *MysqlServiceInstanceBackups) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MysqlServiceInstanceBackupsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) ResetCloudStoragePassword() {
	_jsii_.InvokeVoid(
		m,
		"resetCloudStoragePassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) ResetCloudStorageUsername() {
	_jsii_.InvokeVoid(
		m,
		"resetCloudStorageUsername",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) ResetCreateIfMissing() {
	_jsii_.InvokeVoid(
		m,
		"resetCreateIfMissing",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MysqlServiceInstanceBackupsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

