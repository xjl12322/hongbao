package errors

import (
	"github.com/micro/go-micro/errors"
	"hongbao/movies/src/share/config"
)



const (
	errorCodeFilmSuccess = 200

)

var (
	ErrorFilmFailed = errors.New(
		config.ServiceNameUser,"操作异常",errorCodeFilmSuccess,
	)
)
