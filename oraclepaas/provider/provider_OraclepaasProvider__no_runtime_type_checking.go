//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OraclepaasProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (o *jsiiProxy_OraclepaasProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateOraclepaasProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateOraclepaasProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateOraclepaasProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_OraclepaasProvider) validateSetInsecureParameters(val interface{}) error {
	return nil
}

func validateNewOraclepaasProviderParameters(scope constructs.Construct, id *string, config *OraclepaasProviderConfig) error {
	return nil
}

