// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rotating

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-time.rotating.Rotating",
		reflect.TypeOf((*Rotating)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "day", GoGetter: "Day"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "hour", GoGetter: "Hour"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "minute", GoGetter: "Minute"},
			_jsii_.MemberProperty{JsiiProperty: "month", GoGetter: "Month"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRfc3339", GoMethod: "ResetRfc3339"},
			_jsii_.MemberMethod{JsiiMethod: "resetRotationDays", GoMethod: "ResetRotationDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetRotationHours", GoMethod: "ResetRotationHours"},
			_jsii_.MemberMethod{JsiiMethod: "resetRotationMinutes", GoMethod: "ResetRotationMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetRotationMonths", GoMethod: "ResetRotationMonths"},
			_jsii_.MemberMethod{JsiiMethod: "resetRotationRfc3339", GoMethod: "ResetRotationRfc3339"},
			_jsii_.MemberMethod{JsiiMethod: "resetRotationYears", GoMethod: "ResetRotationYears"},
			_jsii_.MemberMethod{JsiiMethod: "resetTriggers", GoMethod: "ResetTriggers"},
			_jsii_.MemberProperty{JsiiProperty: "rfc3339", GoGetter: "Rfc3339"},
			_jsii_.MemberProperty{JsiiProperty: "rfc3339Input", GoGetter: "Rfc3339Input"},
			_jsii_.MemberProperty{JsiiProperty: "rotationDays", GoGetter: "RotationDays"},
			_jsii_.MemberProperty{JsiiProperty: "rotationDaysInput", GoGetter: "RotationDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "rotationHours", GoGetter: "RotationHours"},
			_jsii_.MemberProperty{JsiiProperty: "rotationHoursInput", GoGetter: "RotationHoursInput"},
			_jsii_.MemberProperty{JsiiProperty: "rotationMinutes", GoGetter: "RotationMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "rotationMinutesInput", GoGetter: "RotationMinutesInput"},
			_jsii_.MemberProperty{JsiiProperty: "rotationMonths", GoGetter: "RotationMonths"},
			_jsii_.MemberProperty{JsiiProperty: "rotationMonthsInput", GoGetter: "RotationMonthsInput"},
			_jsii_.MemberProperty{JsiiProperty: "rotationRfc3339", GoGetter: "RotationRfc3339"},
			_jsii_.MemberProperty{JsiiProperty: "rotationRfc3339Input", GoGetter: "RotationRfc3339Input"},
			_jsii_.MemberProperty{JsiiProperty: "rotationYears", GoGetter: "RotationYears"},
			_jsii_.MemberProperty{JsiiProperty: "rotationYearsInput", GoGetter: "RotationYearsInput"},
			_jsii_.MemberProperty{JsiiProperty: "second", GoGetter: "Second"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "triggers", GoGetter: "Triggers"},
			_jsii_.MemberProperty{JsiiProperty: "triggersInput", GoGetter: "TriggersInput"},
			_jsii_.MemberProperty{JsiiProperty: "unix", GoGetter: "Unix"},
			_jsii_.MemberProperty{JsiiProperty: "year", GoGetter: "Year"},
		},
		func() interface{} {
			j := jsiiProxy_Rotating{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-time.rotating.RotatingConfig",
		reflect.TypeOf((*RotatingConfig)(nil)).Elem(),
	)
}
