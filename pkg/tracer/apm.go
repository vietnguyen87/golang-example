package tracer

import (
	"context"
	"example-service/pkg/config"
	"example-service/pkg/logger"
	"example-service/pkg/utils/apiwrapper"
	"github.com/gin-gonic/gin"
	"gitlab.marathon.edu.vn/pkg/go/xerrors"
	"go.elastic.co/apm/module/apmhttp/v2"
	"go.elastic.co/apm/v2"
	"net/http"
	"runtime/debug"
)

type middleware struct {
	engine         *gin.Engine
	tracer         *apm.Tracer
	requestIgnorer apmhttp.RequestIgnorerFunc
}

/*func NewTracer(engine *gin.Engine, o ...Option) Tracer {
	fmt.Println(os.Getenv("ELASTIC_APM_SERVICE_NAME"))
	fmt.Println(os.Getenv("ELASTIC_APM_SERVER_URL"))
	serviceName := os.Getenv("ELASTIC_APM_SERVICE_NAME")
	serviceVersion := os.Getenv("ELASTIC_APM_SERVICE_VERSION")
	t, _ := apm.NewTracer(serviceName, serviceVersion)
	newTracer := &tracer{
		engine: engine,
		tracer: t,
	}
	for _, o := range o {
		o(newTracer)
	}
	if newTracer.requestIgnorer == nil {
		newTracer.requestIgnorer = apmhttp.NewDynamicServerRequestIgnorer(newTracer.tracer)
	}
	return newTracer
}*/

// Option sets options for tracing.
type Option func(*middleware)

// Middleware returns a new Gin middleware handler for tracing
// requests and reporting errors.
//
// This middleware will recover and report panics, so it can
// be used instead of the standard gin.Recovery middleware.
//
// By default, the middleware will use apm.DefaultTracer.
// Use WithTracer to specify an alternative tracer.
func Middleware(engine *gin.Engine, o ...Option) gin.HandlerFunc {
	serviceName := config.GetElasticApmConfig().ServiceName
	serviceVersion := config.GetElasticApmConfig().ServiceVersion
	t, _ := apm.NewTracer(serviceName, serviceVersion)
	m := &middleware{
		engine: engine,
		tracer: t,
	}
	for _, o := range o {
		o(m)
	}
	if m.requestIgnorer == nil {
		m.requestIgnorer = apmhttp.NewDynamicServerRequestIgnorer(m.tracer)
	}
	return m.handle
}

// Middleware returns a new Gin middleware handler for tracing
// requests and reporting errors.
//
// This middleware will recover and report panics, so it can
// be used instead of the standard gin.Recovery middleware.
//
// By default, the middleware will use apm.DefaultTracer.
// Use WithTracer to specify an alternative tracer.
func (t *middleware) handle(c *gin.Context) {
	if !t.tracer.Recording() || t.requestIgnorer(c.Request) {
		c.Next()
		return
	}

	requestName := t.getRequestName(c)
	tx, body, req := StartTransactionWithBody(t.tracer, requestName, c.Request)
	defer tx.End()
	c.Request = req

	defer func() {
		log := logger.CToL(c.Request.Context(), "gin-recover")
		if v := recover(); v != nil {
			e := t.tracer.Recovered(v)
			e.SetTransaction(tx) // or e.SetSpan(span)
			setContext(&e.Context, c, body)
			e.Send()
			log.Errorf("[Recovery] panic method %v, path %v, err %v, trace %v",
				c.Request.Method,
				c.Request.URL.EscapedPath(),
				v,
				string(debug.Stack()))
			apiwrapper.Abort(c, &apiwrapper.Response{Error: xerrors.InternalServerError.New()})
		}
		c.Writer.WriteHeaderNow()
		tx.Result = apmhttp.StatusCodeResult(c.Writer.Status())

		if tx.Sampled() {
			setContext(&tx.Context, c, body)
		}

		for _, err := range c.Errors {
			e := t.tracer.NewError(err.Err)
			e.SetTransaction(tx)
			setContext(&e.Context, c, body)
			e.Handled = true
			e.Send()
		}
		body.Discard()
	}()
	c.Next()
}

func setContext(ctx *apm.Context, c *gin.Context, body *apm.BodyCapturer) {
	ctx.SetFramework("gin", gin.Version)
	ctx.SetHTTPRequest(c.Request)
	ctx.SetHTTPRequestBody(body)
	ctx.SetHTTPStatusCode(c.Writer.Status())
	ctx.SetHTTPResponseHeaders(c.Writer.Header())
}

func (t *middleware) getRequestName(c *gin.Context) string {
	if fullPath := c.FullPath(); fullPath != "" {
		return c.Request.Method + " " + fullPath
	}
	return apmhttp.UnknownRouteRequestName(c.Request)
}

// StartTransactionWithBody returns a new Transaction with name,
// created with tracer, and taking trace context from req.
//
// If the transaction is not ignored, the request and the request body
// capturer will be returned with the transaction added to its context.
func StartTransactionWithBody(tracer *apm.Tracer, name string, req *http.Request) (*apm.Transaction, *apm.BodyCapturer, *http.Request) {
	tx, req := StartTransaction(tracer, name, req)
	bc := tracer.CaptureHTTPRequestBody(req)
	if bc != nil {
		req = RequestWithContext(apm.ContextWithBodyCapturer(req.Context(), bc), req)
	}
	return tx, bc, req
}

// StartTransaction returns a new Transaction with name,
// created with tracer, and taking trace context from req.
//
// If the transaction is not ignored, the request will be
// returned with the transaction added to its context.
//
// DEPRECATED. Use StartTransactionWithBody instead.
func StartTransaction(tracer *apm.Tracer, name string, req *http.Request) (*apm.Transaction, *http.Request) {
	traceContext, ok := getRequestTraceparent(req, apmhttp.W3CTraceparentHeader)
	if !ok {
		traceContext, ok = getRequestTraceparent(req, apmhttp.ElasticTraceparentHeader)
	}
	if ok {
		traceContext.State, _ = apmhttp.ParseTracestateHeader(req.Header[apmhttp.TracestateHeader]...)
	}
	tx := tracer.StartTransactionOptions(name, "request", apm.TransactionOptions{TraceContext: traceContext})
	ctx := apm.ContextWithTransaction(req.Context(), tx)
	req = RequestWithContext(ctx, req)
	return tx, req
}

func getRequestTraceparent(req *http.Request, header string) (apm.TraceContext, bool) {
	if values := req.Header[header]; len(values) == 1 && values[0] != "" {
		if c, err := apmhttp.ParseTraceparentHeader(values[0]); err == nil {
			return c, true
		}
	}
	return apm.TraceContext{}, false
}

// RequestWithContext is equivalent to req.WithContext, except that the URL
// pointer is copied, rather than the contents.
func RequestWithContext(ctx context.Context, req *http.Request) *http.Request {
	url := req.URL
	req.URL = nil
	reqCopy := req.WithContext(ctx)
	reqCopy.URL = url
	req.URL = url
	return reqCopy
}
