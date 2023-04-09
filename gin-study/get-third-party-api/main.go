package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*func main() {
	routers.go := gin.Default()
	routers.go.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	routers.go.Run(":8086")
}*/

func main() {
	router := gin.Default()
	router.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	router.Run(":8086")
}

const ThirdPartyUrl = "https://raw.githubusercontent.com/gin-gonic/logo/master/color.png"

func getThirdPartyApi(context *gin.Context) {
	response, err := http.Get(ThirdPartyUrl)
	if err != nil && response.StatusCode != http.StatusOK {
		context.Status(http.StatusServiceUnavailable)
		return
	}

	body := response.Body
	length := response.ContentLength
	contentType := response.Header.Get("Content-Type")
	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="gopher.png"`,
	}
	context.DataFromReader(http.StatusOK, length, contentType, body, extraHeaders)
}
