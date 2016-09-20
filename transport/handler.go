package transport

import (
	"net/http"

	"github.com/yarpc/yarpc-go/transport"
	yhttp "github.com/yarpc/yarpc-go/transport/http"
)

func popHeader(h http.Header, n string) string {
	v := h.Get(n)
	h.Del(n)
	return v
}

// GraphQLExtractor creates a transport request from http headers
func GraphQLExtractor(req *http.Request) (*transport.Request, error) {
	return &transport.Request{
		Caller:    popHeader(req.Header, yhttp.CallerHeader),
		Service:   "graphql",
		Procedure: "graphql",
		Encoding:  "graphql",
		Headers:   transport.Headers{},
		Body:      req.Body,
	}, nil
}
