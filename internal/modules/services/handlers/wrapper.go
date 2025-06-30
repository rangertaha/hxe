package handlers

import (
	"fmt"

	"github.com/nats-io/nats.go/micro"
	"google.golang.org/protobuf/proto"
)

// HandlerFunc is a generic function type for handlers that return responses
type HandlerFunc[T proto.Message, R proto.Message] func(req T) (R, error)

// HandlerFuncNoResponse is for handlers that don't return responses
type HandlerFuncNoResponse[T proto.Message] func(req T) error

// SimpleHandlerFunc is for handlers that don't need request unmarshaling
type SimpleHandlerFunc[R proto.Message] func() (R, error)

// Wrap wraps a handler function with automatic marshaling/unmarshaling
func Wrap[T proto.Message, R proto.Message](handler HandlerFunc[T, R]) micro.HandlerFunc {
	return func(msg micro.Request) {
		var req T

		// Unmarshal request
		if err := proto.Unmarshal(msg.Data(), &req); err != nil {
			msg.Error("500", "Failed to unmarshal request", nil)
			return
		}

		// Call the handler
		response, err := handler(req)
		if err != nil {
			msg.Error("500", err.Error(), nil)
			return
		}

		// Marshal and send response
		data, err := proto.Marshal(response)
		if err != nil {
			msg.Error("500", "Failed to marshal response", nil)
			return
		}
		msg.Respond(data)
	}
}

// WrapNoResponse wraps a handler that doesn't return a response
func WrapNoResponse[T proto.Message](handler HandlerFuncNoResponse[T]) micro.HandlerFunc {
	return func(msg micro.Request) {
		var req T

		// Unmarshal request
		if err := proto.Unmarshal(msg.Data(), &req); err != nil {
			msg.Error("500", "Failed to unmarshal request", nil)
			return
		}

		// Call the handler
		if err := handler(req); err != nil {
			msg.Error("500", err.Error(), nil)
			return
		}

		// Send empty response
		msg.Respond([]byte{})
	}
}

// WrapSimple wraps a handler that doesn't need request unmarshaling
func WrapSimple[R proto.Message](handler SimpleHandlerFunc[R]) micro.HandlerFunc {
	return func(msg micro.Request) {
		// Call the handler
		response, err := handler()
		if err != nil {
			msg.Error("500", err.Error(), nil)
			return
		}

		// Marshal and send response
		data, err := proto.Marshal(response)
		if err != nil {
			msg.Error("500", "Failed to marshal response", nil)
			return
		}
		msg.Respond(data)
	}
}

// WrapWithErrorHandler allows custom error handling
func WrapWithErrorHandler[T proto.Message, R proto.Message](
	handler HandlerFunc[T, R],
	errorHandler func(error) (string, string, map[string]string),
) micro.HandlerFunc {
	return func(msg micro.Request) {
		var req T

		// Unmarshal request
		if err := proto.Unmarshal(msg.Data(), &req); err != nil {
			code, message, metadata := errorHandler(fmt.Errorf("failed to unmarshal request: %w", err))
			msg.Error(code, message, metadata)
			return
		}

		// Call the handler
		response, err := handler(req)
		if err != nil {
			code, message, metadata := errorHandler(err)
			msg.Error(code, message, metadata)
			return
		}

		// Marshal and send response
		data, err := proto.Marshal(response)
		if err != nil {
			code, message, metadata := errorHandler(fmt.Errorf("failed to marshal response: %w", err))
			msg.Error(code, message, metadata)
			return
		}
		msg.Respond(data)
	}
}
