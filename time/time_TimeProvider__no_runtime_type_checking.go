//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt time Provider for Terraform CDK (cdktf)
package time

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TimeProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_TimeProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateTimeProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewTimeProviderParameters(scope constructs.Construct, id *string, config *TimeProviderConfig) error {
	return nil
}

