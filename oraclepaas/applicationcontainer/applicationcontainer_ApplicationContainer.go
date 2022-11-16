package applicationcontainer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v2/applicationcontainer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/oraclepaas/r/application_container oraclepaas_application_container}.
type ApplicationContainer interface {
	cdktf.TerraformResource
	AppUrl() *string
	ArchiveUrl() *string
	SetArchiveUrl(val *string)
	ArchiveUrlInput() *string
	AuthType() *string
	SetAuthType(val *string)
	AuthTypeInput() *string
	AvailabilityDomain() *[]*string
	SetAvailabilityDomain(val *[]*string)
	AvailabilityDomainInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Deployment() ApplicationContainerDeploymentOutputReference
	DeploymentFile() *string
	SetDeploymentFile(val *string)
	DeploymentFileInput() *string
	DeploymentInput() *ApplicationContainerDeployment
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GitPassword() *string
	SetGitPassword(val *string)
	GitPasswordInput() *string
	GitRepository() *string
	SetGitRepository(val *string)
	GitRepositoryInput() *string
	GitUsername() *string
	SetGitUsername(val *string)
	GitUsernameInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancerSubnets() *[]*string
	SetLoadBalancerSubnets(val *[]*string)
	LoadBalancerSubnetsInput() *[]*string
	Manifest() ApplicationContainerManifestOutputReference
	ManifestFile() *string
	SetManifestFile(val *string)
	ManifestFileInput() *string
	ManifestInput() *ApplicationContainerManifest
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Notes() *string
	SetNotes(val *string)
	NotesInput() *string
	NotificationEmail() *string
	SetNotificationEmail(val *string)
	NotificationEmailInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	Runtime() *string
	SetRuntime(val *string)
	RuntimeInput() *string
	SubscriptionType() *string
	SetSubscriptionType(val *string)
	SubscriptionTypeInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ApplicationContainerTimeoutsOutputReference
	TimeoutsInput() interface{}
	WebUrl() *string
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutDeployment(value *ApplicationContainerDeployment)
	PutManifest(value *ApplicationContainerManifest)
	PutTimeouts(value *ApplicationContainerTimeouts)
	ResetArchiveUrl()
	ResetAuthType()
	ResetAvailabilityDomain()
	ResetDeployment()
	ResetDeploymentFile()
	ResetGitPassword()
	ResetGitRepository()
	ResetGitUsername()
	ResetId()
	ResetLoadBalancerSubnets()
	ResetManifest()
	ResetManifestFile()
	ResetNotes()
	ResetNotificationEmail()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetRuntime()
	ResetSubscriptionType()
	ResetTags()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ApplicationContainer
type jsiiProxy_ApplicationContainer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApplicationContainer) AppUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) ArchiveUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archiveUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) ArchiveUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archiveUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) AvailabilityDomain() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) AvailabilityDomainInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) Deployment() ApplicationContainerDeploymentOutputReference {
	var returns ApplicationContainerDeploymentOutputReference
	_jsii_.Get(
		j,
		"deployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) DeploymentFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) DeploymentFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) DeploymentInput() *ApplicationContainerDeployment {
	var returns *ApplicationContainerDeployment
	_jsii_.Get(
		j,
		"deploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) GitPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) GitPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) GitRepository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) GitRepositoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) GitUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) GitUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) LoadBalancerSubnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) LoadBalancerSubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerSubnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) Manifest() ApplicationContainerManifestOutputReference {
	var returns ApplicationContainerManifestOutputReference
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) ManifestFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) ManifestFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) ManifestInput() *ApplicationContainerManifest {
	var returns *ApplicationContainerManifest
	_jsii_.Get(
		j,
		"manifestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) Notes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) NotesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) NotificationEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) NotificationEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) SubscriptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) SubscriptionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) Timeouts() ApplicationContainerTimeoutsOutputReference {
	var returns ApplicationContainerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationContainer) WebUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webUrl",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/oraclepaas/r/application_container oraclepaas_application_container} Resource.
func NewApplicationContainer(scope constructs.Construct, id *string, config *ApplicationContainerConfig) ApplicationContainer {
	_init_.Initialize()

	if err := validateNewApplicationContainerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationContainer{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.applicationContainer.ApplicationContainer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/oraclepaas/r/application_container oraclepaas_application_container} Resource.
func NewApplicationContainer_Override(a ApplicationContainer, scope constructs.Construct, id *string, config *ApplicationContainerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.applicationContainer.ApplicationContainer",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetArchiveUrl(val *string) {
	if err := j.validateSetArchiveUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archiveUrl",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetAuthType(val *string) {
	if err := j.validateSetAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetAvailabilityDomain(val *[]*string) {
	if err := j.validateSetAvailabilityDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityDomain",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetDeploymentFile(val *string) {
	if err := j.validateSetDeploymentFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentFile",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetGitPassword(val *string) {
	if err := j.validateSetGitPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitPassword",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetGitRepository(val *string) {
	if err := j.validateSetGitRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitRepository",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetGitUsername(val *string) {
	if err := j.validateSetGitUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitUsername",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetLoadBalancerSubnets(val *[]*string) {
	if err := j.validateSetLoadBalancerSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerSubnets",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetManifestFile(val *string) {
	if err := j.validateSetManifestFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestFile",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetNotes(val *string) {
	if err := j.validateSetNotesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notes",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetNotificationEmail(val *string) {
	if err := j.validateSetNotificationEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationEmail",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetSubscriptionType(val *string) {
	if err := j.validateSetSubscriptionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscriptionType",
		val,
	)
}

func (j *jsiiProxy_ApplicationContainer)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ApplicationContainer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationContainer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-oraclepaas.applicationContainer.ApplicationContainer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApplicationContainer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationContainer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-oraclepaas.applicationContainer.ApplicationContainer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApplicationContainer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationContainer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-oraclepaas.applicationContainer.ApplicationContainer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApplicationContainer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-oraclepaas.applicationContainer.ApplicationContainer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApplicationContainer) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApplicationContainer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationContainer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationContainer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationContainer) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationContainer) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationContainer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationContainer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationContainer) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationContainer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationContainer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationContainer) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApplicationContainer) PutDeployment(value *ApplicationContainerDeployment) {
	if err := a.validatePutDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeployment",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationContainer) PutManifest(value *ApplicationContainerManifest) {
	if err := a.validatePutManifestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManifest",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationContainer) PutTimeouts(value *ApplicationContainerTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationContainer) ResetArchiveUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetArchiveUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainer) ResetAuthType() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainer) ResetAvailabilityDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailabilityDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainer) ResetDeployment() {
	_jsii_.InvokeVoid(
		a,
		"resetDeployment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainer) ResetDeploymentFile() {
	_jsii_.InvokeVoid(
		a,
		"resetDeploymentFile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainer) ResetGitPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetGitPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainer) ResetGitRepository() {
	_jsii_.InvokeVoid(
		a,
		"resetGitRepository",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainer) ResetGitUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetGitUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainer) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainer) ResetLoadBalancerSubnets() {
	_jsii_.InvokeVoid(
		a,
		"resetLoadBalancerSubnets",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainer) ResetManifest() {
	_jsii_.InvokeVoid(
		a,
		"resetManifest",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainer) ResetManifestFile() {
	_jsii_.InvokeVoid(
		a,
		"resetManifestFile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainer) ResetNotes() {
	_jsii_.InvokeVoid(
		a,
		"resetNotes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainer) ResetNotificationEmail() {
	_jsii_.InvokeVoid(
		a,
		"resetNotificationEmail",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainer) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainer) ResetRuntime() {
	_jsii_.InvokeVoid(
		a,
		"resetRuntime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainer) ResetSubscriptionType() {
	_jsii_.InvokeVoid(
		a,
		"resetSubscriptionType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainer) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainer) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationContainer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationContainer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationContainer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationContainer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

