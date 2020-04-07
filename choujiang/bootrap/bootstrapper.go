package bootrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hongbao/choujiang/comm"
	"html/template"
	"runtime"
	"time"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	Application *gin.Engine
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time
}

var A *Bootstrapper
func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Application:  gin.New(),
	}
	A = b
	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}
func (b *Bootstrapper) Bootstrap()*Bootstrapper  {

	_, file, _, ok := runtime.Caller(1)
	fmt.Println(file,ok)
	b.Application.Static("/public/","./web/public")
	b.Application.SetFuncMap(template.FuncMap{
		"FromUnixtime":comm.FromUnixtime,
		"FromUnixtimeShort":comm.FromUnixtimeShort,
	})
	b.Application.LoadHTMLGlob("./web/views/**/*")
	b.Application.Use(gin.Logger(),gin.Recovery())

	return b
}




















