//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package staticresource

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StaticResource) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_StaticResource) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StaticResource) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StaticResource) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StaticResource) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StaticResource) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StaticResource) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StaticResource) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StaticResource) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StaticResource) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StaticResource) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StaticResource) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateStaticResource_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_StaticResource) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StaticResource) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_StaticResource) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_StaticResource) validateSetRfc3339Parameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StaticResource) validateSetTriggersParameters(val *map[string]*string) error {
	return nil
}

func validateNewStaticResourceParameters(scope constructs.Construct, id *string, config *StaticResourceConfig) error {
	return nil
}

