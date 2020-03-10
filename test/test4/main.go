package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type helper struct {



}
func (h *Helper) HelloHandler() (r *ginHelper.Router) {
	return &ginHelper.Router{
		Path:   "/HelloHandler",
		Method: "GET",
		Handlers: []gin.HandlerFunc{
			helloHandler,
		}}
}
func helloHandler(c *gin.Context) {

	c.String(http.StatusOK, "Hello world!")
}

func main() {


}