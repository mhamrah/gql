package main

import (
	"fmt"
	"log"
	"os"

	egql "github.com/mhamrah/gql/encoding"

	"github.com/mhamrah/gql/example/starwars/gql/queryserver"
	"github.com/mhamrah/gql/example/starwars/yarpc/queryserver"
	tgql "github.com/mhamrah/gql/transport"
	"github.com/uber/tchannel-go"
	yarpc "github.com/yarpc/yarpc-go"
	"github.com/yarpc/yarpc-go/encoding/json"
	"github.com/yarpc/yarpc-go/encoding/thrift"
	"github.com/yarpc/yarpc-go/transport"
	yhttp "github.com/yarpc/yarpc-go/transport/http"
	tch "github.com/yarpc/yarpc-go/transport/tchannel"
	"golang.org/x/net/context"
)

func main() {
	channel, err := tchannel.NewChannel("starwars", nil)
	if err != nil {
		log.Fatalln(err)
	}

	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: "starwars",
		Inbounds: []transport.Inbound{
			tch.NewInbound(channel, tch.ListenAddr(":28941")),
			yhttp.NewInbound(":24034"),
			yhttp.NewInbound(":24035", yhttp.Extractor(tgql.GraphQLExtractor)),
		},
		Interceptor: yarpc.Interceptors(requestLogInterceptor{}),
	})

	handler := &handler{humans: getHumans()}

	thrift.Register(dispatcher, queryserver.New(handler))
	json.Register(dispatcher, json.Procedure("human", handler.Human))
	egql.Register(dispatcher, gqlqueryserver.New(handler))

	if err := dispatcher.Start(); err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}

	select {}
}

type requestLogInterceptor struct{}

func (requestLogInterceptor) Handle(
	ctx context.Context, opts transport.Options, req *transport.Request, resw transport.ResponseWriter, handler transport.Handler) error {
	fmt.Printf("received a request to %q from client %q (encoding %q)\n",
		req.Procedure, req.Caller, req.Encoding)
	return handler.Handle(ctx, opts, req, resw)
}
