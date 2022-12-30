package ginutils

import (
	"github.com/vietnguyen87/pkg-golang/xcontext"

	"github.com/gin-gonic/gin"
)

const (
	headerRequestID = "X-Request-ID"
	formatDate      = "0601021504" // YYMMDDHHMM
)

/*func InjectTraceID(c *gin.Context) {
	id := xid.New()
	// Format: YYMMDDHHMM_ID => 210914_0050_c4voumo6n88nq9t8dk20
	traceID := fmt.Sprintf("%s_%s", id.Time().Format(formatDate), id.String())
	rid := c.GetHeader(headerRequestID)
	if rid == "" {
		rid = traceID
	}
	ctx := context.WithValue(c.Request.Context(), xcontext.KeyContextID.String(), rid)
	c.Request = c.Request.WithContext(ctx)
	c.Set(xcontext.KeyContextID.String(), rid)
	c.Next()
}*/

func GetTraceIDFromCtx(c *gin.Context) string {
	if result, ok := c.Get(xcontext.KeyContextID.String()); ok {
		return result.(string)
	}
	return ""
}
