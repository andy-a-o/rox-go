package impression_test

import (
	"github.com/andy-a-o/rox-go/v5/core/context"
	"github.com/andy-a-o/rox-go/v5/core/impression"
	"github.com/andy-a-o/rox-go/v5/core/mocks"
	"github.com/andy-a-o/rox-go/v5/core/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImpressionInvokerEmptyInvokeNotThrowingException(t *testing.T) {
	internalFlags := &mocks.InternalFlags{}
	impressionInvoker := impression.NewImpressionInvoker(internalFlags, nil, nil, false)
	impressionInvoker.Invoke(nil, nil)
}

func TestImpressionInvokerInvokeAndParameters(t *testing.T) {
	internalFlags := &mocks.InternalFlags{}
	impressionInvoker := impression.NewImpressionInvoker(internalFlags, nil, nil, false)

	ctx := context.NewContext(map[string]interface{}{"obj1": 1})
	reportingValue := model.NewReportingValue("name", "value", true)

	isImpressionRaised := false
	impressionInvoker.RegisterImpressionHandler(func(e model.ImpressionArgs) {
		assert.Equal(t, reportingValue, e.ReportingValue)
		assert.Equal(t, ctx, e.Context)

		isImpressionRaised = true
	})

	impressionInvoker.Invoke(reportingValue, ctx)

	assert.True(t, isImpressionRaised)
}

func TestImpressionInvokerInvokeHandleUserCodePanic(t *testing.T) {
	internalFlags := &mocks.InternalFlags{}
	impressionInvoker := impression.NewImpressionInvoker(internalFlags, nil, nil, false)

	ctx := context.NewContext(map[string]interface{}{"obj1": 1})
	reportingValue := model.NewReportingValue("name", "value", true)

	isImpressionRaised := false
	impressionInvoker.RegisterImpressionHandler(func(e model.ImpressionArgs) {
		panic("mwahahahahaEvilUser")
	})

	impressionInvoker.Invoke(reportingValue, ctx)

	assert.False(t, isImpressionRaised)
}

func TestImpressionInvokerWillNotInvokeAnalyticsWhenFlagIsOff(t *testing.T) {
	// TODO
}

func TestImpressionInvokerWillNotInvokeAnalyticsWhenIsRoxy(t *testing.T) {
	// TODO
}

func TestImpressionInvokerWillInvokeAnalytics(t *testing.T) {
	// TODO
}

func TestImpressionInvokerWillInvokeAnalyticsWithBadDistinctID(t *testing.T) {
	// TODO
}
