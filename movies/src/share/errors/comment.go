package errors

import (
	"github.com/micro/go-micro/errors"
	"hongbao/movies/src/share/config"
)

const (
	errorCodeCommentSuccess = 200
)

var (
	ErrorCommentFailed = errors.New(
		config.ServiceNameUser,"操作异常",errorCodeCommentSuccess,
	)
)

