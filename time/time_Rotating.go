// Prebuilt time Provider for Terraform CDK (cdktf)
package time

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-time-go/time/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-time-go/time/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/time/r/rotating time_rotating}.
type Rotating interface {
	cdktf.TerraformResource
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
	Day() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hour() *float64
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Minute() *float64
	Month() *float64
	// The tree node.
	Node() constructs.Node
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
	Rfc3339() *string
	SetRfc3339(val *string)
	Rfc3339Input() *string
	RotationDays() *float64
	SetRotationDays(val *float64)
	RotationDaysInput() *float64
	RotationHours() *float64
	SetRotationHours(val *float64)
	RotationHoursInput() *float64
	RotationMinutes() *float64
	SetRotationMinutes(val *float64)
	RotationMinutesInput() *float64
	RotationMonths() *float64
	SetRotationMonths(val *float64)
	RotationMonthsInput() *float64
	RotationRfc3339() *string
	SetRotationRfc3339(val *string)
	RotationRfc3339Input() *string
	RotationYears() *float64
	SetRotationYears(val *float64)
	RotationYearsInput() *float64
	Second() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Triggers() *map[string]*string
	SetTriggers(val *map[string]*string)
	TriggersInput() *map[string]*string
	Unix() *float64
	Year() *float64
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRfc3339()
	ResetRotationDays()
	ResetRotationHours()
	ResetRotationMinutes()
	ResetRotationMonths()
	ResetRotationRfc3339()
	ResetRotationYears()
	ResetTriggers()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Rotating
type jsiiProxy_Rotating struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Rotating) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Day() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"day",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Hour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Minute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Month() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"month",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Rfc3339() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rfc3339",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Rfc3339Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rfc3339Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationMonths() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationMonths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationMonthsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationMonthsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationRfc3339() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationRfc3339",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationRfc3339Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationRfc3339Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationYears() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationYears",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationYearsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationYearsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Second() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"second",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Triggers() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) TriggersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Unix() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Year() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"year",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/time/r/rotating time_rotating} Resource.
func NewRotating(scope constructs.Construct, id *string, config *RotatingConfig) Rotating {
	_init_.Initialize()

	if err := validateNewRotatingParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Rotating{}

	_jsii_.Create(
		"@cdktf/provider-time.Rotating",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/time/r/rotating time_rotating} Resource.
func NewRotating_Override(r Rotating, scope constructs.Construct, id *string, config *RotatingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-time.Rotating",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_Rotating)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetRfc3339(val *string) {
	if err := j.validateSetRfc3339Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rfc3339",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetRotationDays(val *float64) {
	if err := j.validateSetRotationDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationDays",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetRotationHours(val *float64) {
	if err := j.validateSetRotationHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationHours",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetRotationMinutes(val *float64) {
	if err := j.validateSetRotationMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationMinutes",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetRotationMonths(val *float64) {
	if err := j.validateSetRotationMonthsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationMonths",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetRotationRfc3339(val *string) {
	if err := j.validateSetRotationRfc3339Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationRfc3339",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetRotationYears(val *float64) {
	if err := j.validateSetRotationYearsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationYears",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetTriggers(val *map[string]*string) {
	if err := j.validateSetTriggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggers",
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
func Rotating_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRotating_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-time.Rotating",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Rotating_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-time.Rotating",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_Rotating) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_Rotating) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_Rotating) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Rotating) ResetRfc3339() {
	_jsii_.InvokeVoid(
		r,
		"resetRfc3339",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Rotating) ResetRotationDays() {
	_jsii_.InvokeVoid(
		r,
		"resetRotationDays",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Rotating) ResetRotationHours() {
	_jsii_.InvokeVoid(
		r,
		"resetRotationHours",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Rotating) ResetRotationMinutes() {
	_jsii_.InvokeVoid(
		r,
		"resetRotationMinutes",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Rotating) ResetRotationMonths() {
	_jsii_.InvokeVoid(
		r,
		"resetRotationMonths",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Rotating) ResetRotationRfc3339() {
	_jsii_.InvokeVoid(
		r,
		"resetRotationRfc3339",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Rotating) ResetRotationYears() {
	_jsii_.InvokeVoid(
		r,
		"resetRotationYears",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Rotating) ResetTriggers() {
	_jsii_.InvokeVoid(
		r,
		"resetTriggers",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Rotating) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

