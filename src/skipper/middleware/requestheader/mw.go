package requestheader

import (
	"skipper/middleware/simpleheader"
	"skipper/skipper"
)

const name = "request-header"

type impl struct {
	simpleheader.Type
}

func Make() skipper.Middleware {
	return &impl{}
}

func (mw *impl) Name() string {
	return name
}

func (mw *impl) MakeFilter(id string, config skipper.MiddlewareConfig) (skipper.Filter, error) {
	f := &impl{}
	err := f.InitFilter(id, config)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (f *impl) Request(ctx skipper.FilterContext) {
	req := ctx.Request()
	if f.Key() == "Host" {
		req.Host = f.Value()
	}

	req.Header.Add(f.Key(), f.Value())
}
