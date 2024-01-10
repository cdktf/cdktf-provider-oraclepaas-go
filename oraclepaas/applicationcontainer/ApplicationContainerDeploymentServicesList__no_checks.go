// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package applicationcontainer

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationContainerDeploymentServicesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationContainerDeploymentServicesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApplicationContainerDeploymentServicesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApplicationContainerDeploymentServicesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationContainerDeploymentServicesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationContainerDeploymentServicesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApplicationContainerDeploymentServicesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApplicationContainerDeploymentServicesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

