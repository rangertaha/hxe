package utils

import (
	"github.com/nats-io/nats.go/micro"
	"google.golang.org/protobuf/proto"
)

// HandlerFunc is a generic function type for handlers that return responses
type HandlerFunc[T proto.Message, R proto.Message] func(req T) (R, error)

// Wrap wraps a handler function with automatic marshaling/unmarshaling
func Proto[T proto.Message, R proto.Message](handler HandlerFunc[T, R]) micro.HandlerFunc {
	return func(msg micro.Request) {
		var req T

		// Unmarshal request
		if err := proto.Unmarshal(msg.Data(), req); err != nil {
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

// type (
// 	// Handler is used to respond to service requests.
// 	Handler interface {
// 		Handle(Request)
// 	}

// 	// HandlerFunc is a function implementing [Handler].
// 	// It allows using a function as a request handler, without having to implement Handle
// 	// on a separate type.
// 	HandlerFunc func(Request)

// 	// Request represents service request available in the service handler.
// 	// It exposes methods to respond to the request, as well as
// 	// getting the request data and headers.
// 	Request interface {
// 		// Respond sends the response for the request.
// 		// Additional headers can be passed using [WithHeaders] option.
// 		Respond([]byte, ...RespondOpt) error

// 		// Proto marshals the given response value and responds to the request.
// 		// Additional headers can be passed using [WithHeaders] option.
// 		ProtoResponse(any, ...RespondOpt) error

// 		// RespondJSON marshals the given response value and responds to the request.
// 		// Additional headers can be passed using [WithHeaders] option.
// 		RespondJSON(any, ...RespondOpt) error

// 		// Error prepares and publishes error response from a handler.
// 		// A response error should be set containing an error code and description.
// 		// Optionally, data can be set as response payload.
// 		Error(code, description string, data []byte, opts ...RespondOpt) error

// 		// Data returns request data.
// 		Data() []byte

// 		// Headers returns request headers.
// 		Headers() Headers

// 		// Subject returns underlying NATS message subject.
// 		Subject() string

// 		// Reply returns underlying NATS message reply subject.
// 		Reply() string
// 	}

// 	// Headers is a wrapper around [*nats.Header]
// 	Headers nats.Header

// 	// RespondOpt is a function used to configure [Request.Respond] and [Request.RespondJSON] methods.
// 	RespondOpt func(*nats.Msg)

// 	// request is a default implementation of Request interface
// 	request struct {
// 		msg          *nats.Msg
// 		respondError error
// 	}

// 	serviceError struct {
// 		Code        string `json:"code"`
// 		Description string `json:"description"`
// 	}
// )

// var (
// 	ErrRespond         = errors.New("NATS error when sending response")
// 	ErrMarshalResponse = errors.New("marshaling response")
// 	ErrArgRequired     = errors.New("argument required")
// )

// func (fn HandlerFunc) Handle(req Request) {
// 	fn(req)
// }

// // ContextHandler is a helper function used to utilize [context.Context]
// // in request handlers.
// func ContextHandler(ctx context.Context, handler func(context.Context, Request)) Handler {
// 	return HandlerFunc(func(req Request) {
// 		handler(ctx, req)
// 	})
// }

// // Respond sends the response for the request.
// // Additional headers can be passed using [WithHeaders] option.
// func (r *request) Respond(response []byte, opts ...RespondOpt) error {
// 	respMsg := &nats.Msg{
// 		Data: response,
// 	}
// 	for _, opt := range opts {
// 		opt(respMsg)
// 	}

// 	if err := r.msg.RespondMsg(respMsg); err != nil {
// 		r.respondError = fmt.Errorf("%w: %s", ErrRespond, err)
// 		return r.respondError
// 	}

// 	return nil
// }

// // RespondJSON marshals the given response value and responds to the request.
// // Additional headers can be passed using [WithHeaders] option.
// func (r *request) RespondJSON(response any, opts ...RespondOpt) error {
// 	resp, err := json.Marshal(response)
// 	if err != nil {
// 		return ErrMarshalResponse
// 	}
// 	return r.Respond(resp, opts...)
// }

// func (r *request) ProtoResponse(response proto.Message, opts ...RespondOpt) error {
// 	resp, err := proto.Marshal(response)
// 	if err != nil {
// 		return ErrMarshalResponse
// 	}
// 	return r.Respond(resp, opts...)
// }

// // Error prepares and publishes error response from a handler.
// // A response error should be set containing an error code and description.
// // Optionally, data can be set as response payload.
// func (r *request) Error(code, description string, data []byte, opts ...RespondOpt) error {
// 	if code == "" {
// 		return fmt.Errorf("%w: error code", ErrArgRequired)
// 	}
// 	if description == "" {
// 		return fmt.Errorf("%w: description", ErrArgRequired)
// 	}
// 	response := &nats.Msg{
// 		Header: nats.Header{
// 			ErrorHeader:     []string{description},
// 			ErrorCodeHeader: []string{code},
// 		},
// 	}
// 	for _, opt := range opts {
// 		opt(response)
// 	}

// 	response.Data = data
// 	if err := r.msg.RespondMsg(response); err != nil {
// 		r.respondError = err
// 		return err
// 	}
// 	r.respondError = &serviceError{
// 		Code:        code,
// 		Description: description,
// 	}

// 	return nil
// }

// // WithHeaders can be used to configure response with custom headers.
// func WithHeaders(headers Headers) RespondOpt {
// 	return func(m *nats.Msg) {
// 		if m.Header == nil {
// 			m.Header = nats.Header(headers)
// 			return
// 		}

// 		for k, v := range headers {
// 			m.Header[k] = v
// 		}
// 	}
// }

// // Data returns request data.
// func (r *request) Data() []byte {
// 	return r.msg.Data
// }

// // Headers returns request headers.
// func (r *request) Headers() Headers {
// 	return Headers(r.msg.Header)
// }

// // Subject returns underlying NATS message subject.
// func (r *request) Subject() string {
// 	return r.msg.Subject
// }

// // Reply returns underlying NATS message reply subject.
// func (r *request) Reply() string {
// 	return r.msg.Reply
// }

// // Get gets the first value associated with the given key.
// // It is case-sensitive.
// func (h Headers) Get(key string) string {
// 	return nats.Header(h).Get(key)
// }

// // Values returns all values associated with the given key.
// // It is case-sensitive.
// func (h Headers) Values(key string) []string {
// 	return nats.Header(h).Values(key)
// }

// func (e *serviceError) Error() string {
// 	return fmt.Sprintf("%s:%s", e.Code, e.Description)
// }
