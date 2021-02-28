package handler

import (
	"github.com/zhangxingdeppon/prometheus-metrics/demo02/stat"
	//导入第一步中初始换好的stat(prometheus client_golang 初始化包中的监控项)
	"github.com/gin-gonic/gin"
	"time"
	"log"
	//"net/http"
)

func MwPrometheusHttp(c *gin.Context) {
	start := time.Now()
	method := c.Request.Method
	fullPath := c.FullPath()
	stat.GaugeVecApiMethod.WithLabelValues(method).Inc()
	log.Println("The full path: ", fullPath)
	if fullPath == "/api/test/wrong" {
		stat.GaugeVecApiError.WithLabelValues("API").Inc()
	}
	c.Next()
    // after request
	end := time.Now()
	d := end.Sub(start) // time.Millisecond
	log.Println("the request duration time : ",d)
	stat.GaugeVecApiDuration.WithLabelValues(method).Set(float64(d))

	
	
}