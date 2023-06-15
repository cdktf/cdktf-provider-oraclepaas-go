package databaseserviceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v5/jsii"

	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v5/databaseserviceinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseServiceInstanceInstantiateFromBackupOutputReference interface {
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
	DatabaseId() *string
	SetDatabaseId(val *string)
	DatabaseIdInput() *string
	DecryptionKey() *string
	SetDecryptionKey(val *string)
	DecryptionKeyInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DatabaseServiceInstanceInstantiateFromBackup
	SetInternalValue(val *DatabaseServiceInstanceInstantiateFromBackup)
	OnPremise() interface{}
	SetOnPremise(val interface{})
	OnPremiseInput() interface{}
	ServiceId() *string
	SetServiceId(val *string)
	ServiceIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WalletFileContent() *string
	SetWalletFileContent(val *string)
	WalletFileContentInput() *string
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
	ResetDecryptionKey()
	ResetOnPremise()
	ResetServiceId()
	ResetWalletFileContent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabaseServiceInstanceInstantiateFromBackupOutputReference
type jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) CloudStorageContainer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) CloudStorageContainerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) CloudStoragePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStoragePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) CloudStoragePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStoragePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) CloudStorageUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) CloudStorageUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudStorageUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) DatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) DatabaseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) DecryptionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"decryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) DecryptionKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"decryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) InternalValue() *DatabaseServiceInstanceInstantiateFromBackup {
	var returns *DatabaseServiceInstanceInstantiateFromBackup
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) OnPremise() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onPremise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) OnPremiseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onPremiseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) ServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) ServiceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) WalletFileContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"walletFileContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) WalletFileContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"walletFileContentInput",
		&returns,
	)
	return returns
}


func NewDatabaseServiceInstanceInstantiateFromBackupOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatabaseServiceInstanceInstantiateFromBackupOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseServiceInstanceInstantiateFromBackupOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.databaseServiceInstance.DatabaseServiceInstanceInstantiateFromBackupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabaseServiceInstanceInstantiateFromBackupOutputReference_Override(d DatabaseServiceInstanceInstantiateFromBackupOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.databaseServiceInstance.DatabaseServiceInstanceInstantiateFromBackupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference)SetCloudStorageContainer(val *string) {
	if err := j.validateSetCloudStorageContainerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudStorageContainer",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference)SetCloudStoragePassword(val *string) {
	if err := j.validateSetCloudStoragePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudStoragePassword",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference)SetCloudStorageUsername(val *string) {
	if err := j.validateSetCloudStorageUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudStorageUsername",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference)SetDatabaseId(val *string) {
	if err := j.validateSetDatabaseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseId",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference)SetDecryptionKey(val *string) {
	if err := j.validateSetDecryptionKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"decryptionKey",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference)SetInternalValue(val *DatabaseServiceInstanceInstantiateFromBackup) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference)SetOnPremise(val interface{}) {
	if err := j.validateSetOnPremiseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onPremise",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference)SetServiceId(val *string) {
	if err := j.validateSetServiceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceId",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference)SetWalletFileContent(val *string) {
	if err := j.validateSetWalletFileContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"walletFileContent",
		val,
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) ResetCloudStoragePassword() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudStoragePassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) ResetCloudStorageUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudStorageUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) ResetDecryptionKey() {
	_jsii_.InvokeVoid(
		d,
		"resetDecryptionKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) ResetOnPremise() {
	_jsii_.InvokeVoid(
		d,
		"resetOnPremise",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) ResetServiceId() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) ResetWalletFileContent() {
	_jsii_.InvokeVoid(
		d,
		"resetWalletFileContent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabaseServiceInstanceInstantiateFromBackupOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

