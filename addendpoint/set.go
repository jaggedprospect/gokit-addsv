package addendpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"

	"github.com/jaggedprospect/gokit-addsvc/addservice"
)

// Set collects all of the endpoints that compose an add service.
// It's meant to be used as a helper struct to collect all of the
// endpoints into a single parameter.
type Set struct {
	SumEndpoint    endpoint.Endpoint
	ConcatEndpoint endpoint.Endpoint
}

// New returns a Set that wraps the provided server and wires in
// all of the expected endpoint middlewares via the various
// parameters.
func New(svc addservice.Service, logger log.Logger, duration metrics.Histogram, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer) Set {
	panic("Not implemented")
}
