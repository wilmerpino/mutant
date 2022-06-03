package presenter

import "github.com/kataras/iris/v12/context"

type IContext interface {
	Application() context.Application
	StatusCode(statusCode int)
	JSON(v interface{}, opts ...context.JSON) (n int, err error)
	StopExecution()
	PostValue(name string) string
	Params() *context.RequestParams
	GetHeader(name string) string
	ReadJSON(interface{}) error
}
