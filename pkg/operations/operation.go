package operations

import "github.com/markgoddard/reductionist/pkg/request"

type Operation interface {
	Execute([]byte, request.Data) ([]byte, error)
}
