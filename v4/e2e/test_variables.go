package e2e

import "github.com/andy-a-o/rox-go/v4/core/model"

var (
	TestVarsIsComputedBooleanPropCalled bool
	TestVarsIsComputedStringPropCalled  bool
	TestVarsIsComputedIntPropCalled     bool
	TestVarsIsComputedFloatPropCalled   bool
	TestVarsIsComputedSemverPropCalled  bool

	TestVarsTargetGroup1 bool
	TestVarsTargetGroup2 bool

	TestVarsIsImpressionRaised     bool
	TestVarsImpressionReturnedArgs *model.ImpressionArgs

	TestVarsIsPropForTargetGroupForDependency bool

	TestVarsConfigurationFetchedCount int
)
