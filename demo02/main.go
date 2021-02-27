package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zhangxingdeppon/prometheus-metrics/demo02/handler"
	//"log"
	//"github.com/zhangxingdeppon/prometheus-metrics/demo02/stat"
	//"time"
)
func main()  {
	r := gin.Default()
	r.GET("metrics", gin.WrapH(promhttp.Handler()))
	//r.GET("api", gin.WrapH(handler.MwPrometheusHttp))
	api := r.Group("api").Use(handler.MwPrometheusHttp)
	defaultHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}
	//start := time.Now()
	api.GET("/test",defaultHandler)
	api.GET("/test/wrong",defaultHandler)
	// after request
	
	

	// r.GET("/",func (c *gin.Context)  {
	// 	c.String(200, "Hello, Bryan!")
	// })
	r.Run()

}