package mocks

import (
	"encoding/json"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"

	"github.com/stretchr/testify/mock"
)

type MockContext struct {
	mock.Mock
	dataReadJSON string
}

func (ctx *MockContext) ReadParams(ptr interface{}) error {
	return nil
}

func (ctx *MockContext) Application() context.Application {
	ctx.Called()
	app := iris.Default()
	return app
}

func (ctx *MockContext) JSON(v interface{}, options ...context.JSON) (int, error) {
	args := ctx.Called(v)
	return args.Get(0).(int), args.Error(1)
}

func (ctx *MockContext) StatusCode(statusCode int) {
	ctx.Called(statusCode)
}

func (ctx *MockContext) StopExecution() {
	ctx.Called()
}

func (ctx *MockContext) PostValue(name string) string {
	args := ctx.Called(name)
	return args.String(0)
}

func (ctx *MockContext) Params() *context.RequestParams {
	args := ctx.Called()
	return args.Get(0).(*context.RequestParams)
}

func (ctx *MockContext) GetHeader(name string) string {
	args := ctx.Called(name)
	return args.String(0)
}

func (ctx *MockContext) ReadJSON(data interface{}) error {
	args := ctx.Called(data)
	json.Unmarshal([]byte(ctx.dataReadJSON), &data) // nolint:errcheck
	return args.Error(0)
}

func (ctx *MockContext) SetDNAInputJSON(data string) {
	ctx.dataReadJSON = data
}
