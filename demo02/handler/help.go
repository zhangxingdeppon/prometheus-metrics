package handler

import (
	"github.com/zhangxingdeppon/prometheus-metrics/demo02/stat"
	"github.com/gin-gonic/gin"
)
func jsonError(c *gin.Context, msg interface{}) {
	stat.GaugeVecApiError.WithLabelValues("API").Inc()
	var ms string
	switch v := msg.(type) {
	case string:
		ms = v
	case error:
		ms = v.Error()
	default:
		ms = ""
	}
	c.AbortWithStatusJSON(200, gin.H{"ok": false, "msg": ms})
}