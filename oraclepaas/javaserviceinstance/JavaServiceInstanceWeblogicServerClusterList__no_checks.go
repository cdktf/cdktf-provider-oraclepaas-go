// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package javaserviceinstance

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterList) validateGetParameters(index *float64) error {
	return nil
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_JavaServiceInstanceWeblogicServerClusterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewJavaServiceInstanceWeblogicServerClusterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

