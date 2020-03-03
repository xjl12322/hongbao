package errors

import (
	"github.com/micro/go-micro/errors"
	"hongbao/movies/src/share/config"
)


const (
	errorCodeOrderSuccess = 200

)

var (
	ErrorOrderFailed = errors.New(
		config.ServiceNameOrder,"操作异常",errorCodeOrderSuccess,
	)
	ErrorOrderAlreadyScore= errors.New(
		config.ServiceNameOrder,"已经评分了！",errorCodeOrderSuccess,
	)
)
