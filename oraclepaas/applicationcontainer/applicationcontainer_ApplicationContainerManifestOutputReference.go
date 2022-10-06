package applicationcontainer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/jsii"

	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/applicationcontainer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationContainerManifestOutputReference interface {
	cdktf.ComplexObject
	Clustered() interface{}
	SetClustered(val interface{})
	ClusteredInput() interface{}
	Command() *string
	SetCommand(val *string)
	CommandInput() *string
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
	HealthCheckEndpoint() *string
	SetHealthCheckEndpoint(val *string)
	HealthCheckEndpointInput() *string
	Home() *string
	SetHome(val *string)
	HomeInput() *string
	InternalValue() *ApplicationContainerManifest
	SetInternalValue(val *ApplicationContainerManifest)
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	Notes() *string
	SetNotes(val *string)
	NotesInput() *string
	Release() ApplicationContainerManifestReleaseOutputReference
	ReleaseInput() *ApplicationContainerManifestRelease
	Runtime() ApplicationContainerManifestRuntimeOutputReference
	RuntimeInput() *ApplicationContainerManifestRuntime
	ShutdownTime() *float64
	SetShutdownTime(val *float64)
	ShutdownTimeInput() *float64
	StartupTime() *float64
	SetStartupTime(val *float64)
	StartupTimeInput() *float64
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
	PutRelease(value *ApplicationContainerManifestRelease)
	PutRuntime(value *ApplicationContainerManifestRuntime)
	ResetClustered()
	ResetCommand()
	ResetHealthCheckEndpoint()
	ResetHome()
	ResetMode()
	ResetNotes()
	ResetRelease()
	ResetRuntime()
	ResetShutdownTime()
	ResetStartupTime()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationContainerManifestOutputReference
type jsiiProxy_ApplicationContainerManifestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) Clustered() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clustered",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) ClusteredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusteredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) Command() *string {
	var returns *string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) CommandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) HealthCheckEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) HealthCheckEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) Home() *string {
	var returns *string
	_jsii_.Get(
		j,
		"home",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) HomeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) InternalValue() *ApplicationContainerManifest {
	var returns *ApplicationContainerManifest
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) Notes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) NotesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) Release() ApplicationContainerManifestReleaseOutputReference {
	var returns ApplicationContainerManifestReleaseOutputReference
	_jsii_.Get(
		j,
		"release",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) ReleaseInput() *ApplicationContainerManifestRelease {
	var returns *ApplicationContainerManifestRelease
	_jsii_.Get(
		j,
		"releaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) Runtime() ApplicationContainerManifestRuntimeOutputReference {
	var returns ApplicationContainerManifestRuntimeOutputReference
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) RuntimeInput() *ApplicationContainerManifestRuntime {
	var returns *ApplicationContainerManifestRuntime
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) ShutdownTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shutdownTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) ShutdownTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shutdownTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) StartupTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startupTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) StartupTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startupTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewApplicationContainerManifestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApplicationContainerManifestOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationContainerManifestOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationContainerManifestOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.applicationContainer.ApplicationContainerManifestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApplicationContainerManifestOutputReference_Override(a ApplicationContainerManifestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.applicationContainer.ApplicationContainerManifestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference)SetClustered(val interface{}) {
	if err := j.validateSetClusteredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clustered",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference)SetCommand(val *string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference)SetHealthCheckEndpoint(val *string) {
	if err := j.validateSetHealthCheckEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckEndpoint",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference)SetHome(val *string) {
	if err := j.validateSetHomeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"home",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference)SetInternalValue(val *ApplicationContainerManifest) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference)SetNotes(val *string) {
	if err := j.validateSetNotesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notes",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference)SetShutdownTime(val *float64) {
	if err := j.validateSetShutdownTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shutdownTime",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference)SetStartupTime(val *float64) {
	if err := j.validateSetStartupTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startupTime",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainerManifestOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) PutRelease(value *ApplicationContainerManifestRelease) {
	if err := a.validatePutReleaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRelease",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) PutRuntime(value *ApplicationContainerManifestRuntime) {
	if err := a.validatePutRuntimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRuntime",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) ResetClustered() {
	_jsii_.InvokeVoid(
		a,
		"resetClustered",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		a,
		"resetCommand",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) ResetHealthCheckEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthCheckEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) ResetHome() {
	_jsii_.InvokeVoid(
		a,
		"resetHome",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		a,
		"resetMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) ResetNotes() {
	_jsii_.InvokeVoid(
		a,
		"resetNotes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) ResetRelease() {
	_jsii_.InvokeVoid(
		a,
		"resetRelease",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) ResetRuntime() {
	_jsii_.InvokeVoid(
		a,
		"resetRuntime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) ResetShutdownTime() {
	_jsii_.InvokeVoid(
		a,
		"resetShutdownTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) ResetStartupTime() {
	_jsii_.InvokeVoid(
		a,
		"resetStartupTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		a,
		"resetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationContainerManifestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

