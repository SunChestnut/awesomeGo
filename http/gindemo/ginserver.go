package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"time"
)

/**
==> gin 框架：Go 语言实现的 Web 框架
		🐱https://github.com/gin-gonic/gin
==> zap 框架：日志框架
		🐱https://github.com/uber-go/zap
*/

const keyRequestId = "requestId"

func main() {

	r := gin.Default()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// Use() 函数可以将全局中间件附加到路由器上
	r.Use(

		// 使用 middleware 来实现当访问不同的 url 时，该访问都能被记录到日志中。 原理为：不论请求是访问到了 /ping 还是 /hello，都需要先执行 middleware 中的内容
		// 使用 gin 创建 middleware，使用 zap 记录日志
		func(context *gin.Context) {
			start := time.Now()

			// 将控制权交给下一个 middleware，待下一个 middleware 执行完毕后，再跳回来执行当前 middle 中 context.Next() 剩下的部分
			context.Next()

			// 记录 path, log latency, response code
			logger.Info("🎃incoming request",
				zap.String("path", context.Request.URL.Path),
				zap.Int("status", context.Writer.Status()),
				zap.Duration("elapsed", time.Now().Sub(start)),
				zap.Int(keyRequestId, context.GetInt(keyRequestId)),
			)
		},

		// 为每一个请求生成 requestId，每个 request 在整个生命周期中都会携带 requestId
		func(context *gin.Context) {
			context.Set(keyRequestId, rand.Int())

			// 当前的中间件为最后一个，因此会继续执行 context.Next() 后续的内容
			context.Next()

			logger.Info("This is placeholder~~")
		},
	)

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if rid, exists := c.Get(keyRequestId); exists {
			h[keyRequestId] = rid
		}
		c.JSON(http.StatusOK, h)
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello you~~")
	})

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run()

}
