package dogego_server_h2c

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"os"
)

// HTTP/2.0 not over TLS 服务器
func H2CServerProtocol(router *gin.Engine) error {
	log.Printf("HTTP/2.0 not over TLS Server started on %s", os.Getenv("ADDR_H2C"))

	http2Server := &http2.Server{}

	httpServer := &http.Server{
		Addr:    os.Getenv("ADDR_H2C"),
		Handler: h2c.NewHandler(router, http2Server),
	}

	err := httpServer.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}
