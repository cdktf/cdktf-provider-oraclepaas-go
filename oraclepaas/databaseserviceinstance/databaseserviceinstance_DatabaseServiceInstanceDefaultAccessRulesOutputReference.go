package databaseserviceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/jsii"

	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/databaseserviceinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseServiceInstanceDefaultAccessRulesOutputReference interface {
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
	EnableDbConsole() interface{}
	SetEnableDbConsole(val interface{})
	EnableDbConsoleInput() interface{}
	EnableDbExpress() interface{}
	SetEnableDbExpress(val interface{})
	EnableDbExpressInput() interface{}
	EnableDbListener() interface{}
	SetEnableDbListener(val interface{})
	EnableDbListenerInput() interface{}
	EnableEmConsole() interface{}
	SetEnableEmConsole(val interface{})
	EnableEmConsoleInput() interface{}
	EnableHttp() interface{}
	SetEnableHttp(val interface{})
	EnableHttpInput() interface{}
	EnableHttpSsl() interface{}
	SetEnableHttpSsl(val interface{})
	EnableHttpSslInput() interface{}
	EnableRacDbListener() interface{}
	SetEnableRacDbListener(val interface{})
	EnableRacDbListenerInput() interface{}
	EnableRacOns() interface{}
	SetEnableRacOns(val interface{})
	EnableRacOnsInput() interface{}
	EnableScanListener() interface{}
	SetEnableScanListener(val interface{})
	EnableScanListenerInput() interface{}
	EnableSsh() interface{}
	SetEnableSsh(val interface{})
	EnableSshInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DatabaseServiceInstanceDefaultAccessRules
	SetInternalValue(val *DatabaseServiceInstanceDefaultAccessRules)
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
	ResetEnableDbConsole()
	ResetEnableDbExpress()
	ResetEnableDbListener()
	ResetEnableEmConsole()
	ResetEnableHttp()
	ResetEnableHttpSsl()
	ResetEnableRacDbListener()
	ResetEnableRacOns()
	ResetEnableScanListener()
	ResetEnableSsh()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabaseServiceInstanceDefaultAccessRulesOutputReference
type jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) EnableDbConsole() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDbConsole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) EnableDbConsoleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDbConsoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) EnableDbExpress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDbExpress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) EnableDbExpressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDbExpressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) EnableDbListener() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDbListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) EnableDbListenerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDbListenerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) EnableEmConsole() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEmConsole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) EnableEmConsoleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEmConsoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) EnableHttp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) EnableHttpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) EnableHttpSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttpSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) EnableHttpSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttpSslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) EnableRacDbListener() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRacDbListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) EnableRacDbListenerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRacDbListenerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) EnableRacOns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRacOns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) EnableRacOnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRacOnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) EnableScanListener() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableScanListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) EnableScanListenerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableScanListenerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) EnableSsh() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSsh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) EnableSshInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSshInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) InternalValue() *DatabaseServiceInstanceDefaultAccessRules {
	var returns *DatabaseServiceInstanceDefaultAccessRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatabaseServiceInstanceDefaultAccessRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatabaseServiceInstanceDefaultAccessRulesOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseServiceInstanceDefaultAccessRulesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.databaseServiceInstance.DatabaseServiceInstanceDefaultAccessRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabaseServiceInstanceDefaultAccessRulesOutputReference_Override(d DatabaseServiceInstanceDefaultAccessRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.databaseServiceInstance.DatabaseServiceInstanceDefaultAccessRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference)SetEnableDbConsole(val interface{}) {
	if err := j.validateSetEnableDbConsoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDbConsole",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference)SetEnableDbExpress(val interface{}) {
	if err := j.validateSetEnableDbExpressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDbExpress",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference)SetEnableDbListener(val interface{}) {
	if err := j.validateSetEnableDbListenerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDbListener",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference)SetEnableEmConsole(val interface{}) {
	if err := j.validateSetEnableEmConsoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableEmConsole",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference)SetEnableHttp(val interface{}) {
	if err := j.validateSetEnableHttpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHttp",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference)SetEnableHttpSsl(val interface{}) {
	if err := j.validateSetEnableHttpSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHttpSsl",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference)SetEnableRacDbListener(val interface{}) {
	if err := j.validateSetEnableRacDbListenerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableRacDbListener",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference)SetEnableRacOns(val interface{}) {
	if err := j.validateSetEnableRacOnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableRacOns",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference)SetEnableScanListener(val interface{}) {
	if err := j.validateSetEnableScanListenerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableScanListener",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference)SetEnableSsh(val interface{}) {
	if err := j.validateSetEnableSshParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSsh",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference)SetInternalValue(val *DatabaseServiceInstanceDefaultAccessRules) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) ResetEnableDbConsole() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableDbConsole",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) ResetEnableDbExpress() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableDbExpress",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) ResetEnableDbListener() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableDbListener",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) ResetEnableEmConsole() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableEmConsole",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) ResetEnableHttp() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableHttp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) ResetEnableHttpSsl() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableHttpSsl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) ResetEnableRacDbListener() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableRacDbListener",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) ResetEnableRacOns() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableRacOns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) ResetEnableScanListener() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableScanListener",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) ResetEnableSsh() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableSsh",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabaseServiceInstanceDefaultAccessRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

