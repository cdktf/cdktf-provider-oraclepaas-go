// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package javaserviceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v8/jsii"

	"github.com/cdktf/cdktf-provider-oraclepaas-go/oraclepaas/v8/javaserviceinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JavaServiceInstanceWeblogicServerOutputReference interface {
	cdktf.ComplexObject
	Admin() JavaServiceInstanceWeblogicServerAdminOutputReference
	AdminInput() *JavaServiceInstanceWeblogicServerAdmin
	ApplicationDatabase() JavaServiceInstanceWeblogicServerApplicationDatabaseList
	ApplicationDatabaseInput() interface{}
	BackupVolumeSize() *string
	SetBackupVolumeSize(val *string)
	BackupVolumeSizeInput() *string
	Cluster() JavaServiceInstanceWeblogicServerClusterList
	ClusterInput() interface{}
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
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
	ConnectString() *string
	SetConnectString(val *string)
	ConnectStringInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Database() JavaServiceInstanceWeblogicServerDatabaseOutputReference
	DatabaseInput() *JavaServiceInstanceWeblogicServerDatabase
	Domain() JavaServiceInstanceWeblogicServerDomainOutputReference
	DomainInput() *JavaServiceInstanceWeblogicServerDomain
	// Experimental.
	Fqn() *string
	InternalValue() *JavaServiceInstanceWeblogicServer
	SetInternalValue(val *JavaServiceInstanceWeblogicServer)
	IpReservations() *[]*string
	SetIpReservations(val *[]*string)
	IpReservationsInput() *[]*string
	ManagedServers() JavaServiceInstanceWeblogicServerManagedServersOutputReference
	ManagedServersInput() *JavaServiceInstanceWeblogicServerManagedServers
	MiddlewareVolumeSize() *string
	SetMiddlewareVolumeSize(val *string)
	MiddlewareVolumeSizeInput() *string
	NodeManager() JavaServiceInstanceWeblogicServerNodeManagerOutputReference
	NodeManagerInput() *JavaServiceInstanceWeblogicServerNodeManager
	Ports() JavaServiceInstanceWeblogicServerPortsOutputReference
	PortsInput() *JavaServiceInstanceWeblogicServerPorts
	RootUrl() *string
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
	UpperStackProductName() *string
	SetUpperStackProductName(val *string)
	UpperStackProductNameInput() *string
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
	PutAdmin(value *JavaServiceInstanceWeblogicServerAdmin)
	PutApplicationDatabase(value interface{})
	PutCluster(value interface{})
	PutDatabase(value *JavaServiceInstanceWeblogicServerDatabase)
	PutDomain(value *JavaServiceInstanceWeblogicServerDomain)
	PutManagedServers(value *JavaServiceInstanceWeblogicServerManagedServers)
	PutNodeManager(value *JavaServiceInstanceWeblogicServerNodeManager)
	PutPorts(value *JavaServiceInstanceWeblogicServerPorts)
	ResetApplicationDatabase()
	ResetBackupVolumeSize()
	ResetCluster()
	ResetClusterName()
	ResetConnectString()
	ResetDomain()
	ResetIpReservations()
	ResetManagedServers()
	ResetMiddlewareVolumeSize()
	ResetNodeManager()
	ResetPorts()
	ResetUpperStackProductName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JavaServiceInstanceWeblogicServerOutputReference
type jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) Admin() JavaServiceInstanceWeblogicServerAdminOutputReference {
	var returns JavaServiceInstanceWeblogicServerAdminOutputReference
	_jsii_.Get(
		j,
		"admin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) AdminInput() *JavaServiceInstanceWeblogicServerAdmin {
	var returns *JavaServiceInstanceWeblogicServerAdmin
	_jsii_.Get(
		j,
		"adminInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ApplicationDatabase() JavaServiceInstanceWeblogicServerApplicationDatabaseList {
	var returns JavaServiceInstanceWeblogicServerApplicationDatabaseList
	_jsii_.Get(
		j,
		"applicationDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ApplicationDatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) BackupVolumeSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) BackupVolumeSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVolumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) Cluster() JavaServiceInstanceWeblogicServerClusterList {
	var returns JavaServiceInstanceWeblogicServerClusterList
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ClusterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ConnectString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ConnectStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) Database() JavaServiceInstanceWeblogicServerDatabaseOutputReference {
	var returns JavaServiceInstanceWeblogicServerDatabaseOutputReference
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) DatabaseInput() *JavaServiceInstanceWeblogicServerDatabase {
	var returns *JavaServiceInstanceWeblogicServerDatabase
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) Domain() JavaServiceInstanceWeblogicServerDomainOutputReference {
	var returns JavaServiceInstanceWeblogicServerDomainOutputReference
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) DomainInput() *JavaServiceInstanceWeblogicServerDomain {
	var returns *JavaServiceInstanceWeblogicServerDomain
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) InternalValue() *JavaServiceInstanceWeblogicServer {
	var returns *JavaServiceInstanceWeblogicServer
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) IpReservations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipReservations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) IpReservationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipReservationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ManagedServers() JavaServiceInstanceWeblogicServerManagedServersOutputReference {
	var returns JavaServiceInstanceWeblogicServerManagedServersOutputReference
	_jsii_.Get(
		j,
		"managedServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ManagedServersInput() *JavaServiceInstanceWeblogicServerManagedServers {
	var returns *JavaServiceInstanceWeblogicServerManagedServers
	_jsii_.Get(
		j,
		"managedServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) MiddlewareVolumeSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"middlewareVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) MiddlewareVolumeSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"middlewareVolumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) NodeManager() JavaServiceInstanceWeblogicServerNodeManagerOutputReference {
	var returns JavaServiceInstanceWeblogicServerNodeManagerOutputReference
	_jsii_.Get(
		j,
		"nodeManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) NodeManagerInput() *JavaServiceInstanceWeblogicServerNodeManager {
	var returns *JavaServiceInstanceWeblogicServerNodeManager
	_jsii_.Get(
		j,
		"nodeManagerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) Ports() JavaServiceInstanceWeblogicServerPortsOutputReference {
	var returns JavaServiceInstanceWeblogicServerPortsOutputReference
	_jsii_.Get(
		j,
		"ports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) PortsInput() *JavaServiceInstanceWeblogicServerPorts {
	var returns *JavaServiceInstanceWeblogicServerPorts
	_jsii_.Get(
		j,
		"portsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) RootUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ShapeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) UpperStackProductName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upperStackProductName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) UpperStackProductNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upperStackProductNameInput",
		&returns,
	)
	return returns
}


func NewJavaServiceInstanceWeblogicServerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JavaServiceInstanceWeblogicServerOutputReference {
	_init_.Initialize()

	if err := validateNewJavaServiceInstanceWeblogicServerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.javaServiceInstance.JavaServiceInstanceWeblogicServerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJavaServiceInstanceWeblogicServerOutputReference_Override(j JavaServiceInstanceWeblogicServerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-oraclepaas.javaServiceInstance.JavaServiceInstanceWeblogicServerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference)SetBackupVolumeSize(val *string) {
	if err := j.validateSetBackupVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupVolumeSize",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference)SetConnectString(val *string) {
	if err := j.validateSetConnectStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectString",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference)SetInternalValue(val *JavaServiceInstanceWeblogicServer) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference)SetIpReservations(val *[]*string) {
	if err := j.validateSetIpReservationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipReservations",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference)SetMiddlewareVolumeSize(val *string) {
	if err := j.validateSetMiddlewareVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"middlewareVolumeSize",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference)SetShape(val *string) {
	if err := j.validateSetShapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shape",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference)SetUpperStackProductName(val *string) {
	if err := j.validateSetUpperStackProductNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upperStackProductName",
		val,
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) PutAdmin(value *JavaServiceInstanceWeblogicServerAdmin) {
	if err := j.validatePutAdminParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAdmin",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) PutApplicationDatabase(value interface{}) {
	if err := j.validatePutApplicationDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putApplicationDatabase",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) PutCluster(value interface{}) {
	if err := j.validatePutClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putCluster",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) PutDatabase(value *JavaServiceInstanceWeblogicServerDatabase) {
	if err := j.validatePutDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDatabase",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) PutDomain(value *JavaServiceInstanceWeblogicServerDomain) {
	if err := j.validatePutDomainParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDomain",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) PutManagedServers(value *JavaServiceInstanceWeblogicServerManagedServers) {
	if err := j.validatePutManagedServersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putManagedServers",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) PutNodeManager(value *JavaServiceInstanceWeblogicServerNodeManager) {
	if err := j.validatePutNodeManagerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putNodeManager",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) PutPorts(value *JavaServiceInstanceWeblogicServerPorts) {
	if err := j.validatePutPortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPorts",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ResetApplicationDatabase() {
	_jsii_.InvokeVoid(
		j,
		"resetApplicationDatabase",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ResetBackupVolumeSize() {
	_jsii_.InvokeVoid(
		j,
		"resetBackupVolumeSize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ResetCluster() {
	_jsii_.InvokeVoid(
		j,
		"resetCluster",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ResetClusterName() {
	_jsii_.InvokeVoid(
		j,
		"resetClusterName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ResetConnectString() {
	_jsii_.InvokeVoid(
		j,
		"resetConnectString",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ResetDomain() {
	_jsii_.InvokeVoid(
		j,
		"resetDomain",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ResetIpReservations() {
	_jsii_.InvokeVoid(
		j,
		"resetIpReservations",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ResetManagedServers() {
	_jsii_.InvokeVoid(
		j,
		"resetManagedServers",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ResetMiddlewareVolumeSize() {
	_jsii_.InvokeVoid(
		j,
		"resetMiddlewareVolumeSize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ResetNodeManager() {
	_jsii_.InvokeVoid(
		j,
		"resetNodeManager",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ResetPorts() {
	_jsii_.InvokeVoid(
		j,
		"resetPorts",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ResetUpperStackProductName() {
	_jsii_.InvokeVoid(
		j,
		"resetUpperStackProductName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

