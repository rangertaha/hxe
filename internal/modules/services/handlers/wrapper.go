package handlers

import (
	"github.com/nats-io/nats.go/micro"
	"github.com/rangertaha/hxe/internal/modules/services/models"
	"google.golang.org/protobuf/proto"
)

// HandlerFunc is a generic function type for handlers that return responses
type HandlerFunc[T any, R any] func(req T) (R, error)

// HandlerFuncNoResponse is for handlers that don't return responses
type HandlerFuncNoResponse[T any] func(req T) error

// SimpleHandlerFunc is for handlers that don't need request unmarshaling
type SimpleHandlerFunc[R any] func() (R, error)


// Proto wraps a handler function with automatic marshaling/unmarshaling
func Proto(handler func(*models.Request) *models.Response) micro.HandlerFunc {
	return func(msg micro.Request) {
		var req models.Request

		// Unmarshal request
		if err := proto.Unmarshal(msg.Data(), &req); err != nil {
			msg.Respond([]byte("1"))
			return
		}

		// Call the handler
		res := handler(&req)
		if res == nil {
			msg.Respond([]byte("2"))
			return
		}

		// Marshal and send response
		data, err := proto.Marshal(res)
		if err != nil {
			msg.Respond([]byte("3"))
			return
		}
		msg.Respond(data)
	}
}
