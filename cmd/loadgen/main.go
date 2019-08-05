package main

import (
	"flag"
	"log"
	"time"

	"github.com/tigrannajaryan/exp-otelproto/protoimpls/ws_stream_async"

	"github.com/tigrannajaryan/exp-otelproto/protoimpls/ws_stream_sync"

	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_stream_lb_async"

	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_stream_lb"

	"github.com/tigrannajaryan/exp-otelproto/encodings/traceprotobuf"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_stream"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_unary"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/octraceprotobuf"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_oc"
)

func main() {
	log.Println("Load Generator.")

	var destination string
	flag.StringVar(&destination, "dest", "localhost:3465", "destination endpoint to forward to")

	//var listenAddress string
	//flag.StringVar(&listenAddress, "listen", "0.0.0.0:3465", "local address to listen on")

	var protocol string
	flag.StringVar(&protocol, "protocol", "",
		"protocol to benchmark (opencensus,ocack,unary,streamsync,streamlbtimedsync,streamlbalwayssync,streamlbasync,wsstreamsync,wsstreamasync,wsstreamasynczlib)")

	var spansPerSecond int
	flag.IntVar(&spansPerSecond, "spanspersec", 100, "spans per second")

	var rebalancePeriodStr = flag.String("rebalance", "30s", "rebalance period (Valid time units are ns, us, ms, s, m, h)")

	flag.Parse()

	rebalancePeriod, err := time.ParseDuration(*rebalancePeriodStr)
	if err != nil {
		log.Fatal(err)
	}

	switch protocol {
	case "opencensus":
		core.LoadGenerator(
			func() core.Client { return &grpc_oc.Client{} },
			func() core.Generator { return octraceprotobuf.NewGenerator() },
			destination,
			spansPerSecond,
		)

	case "ocack":
		core.LoadGenerator(
			func() core.Client { return &grpc_oc.Client{WaitForAck: true} },
			func() core.Generator { return octraceprotobuf.NewGenerator() },
			destination,
			spansPerSecond,
		)

	case "unary":
		core.LoadGenerator(
			func() core.Client { return &grpc_unary.Client{} },
			func() core.Generator { return traceprotobuf.NewGenerator() },
			destination,
			spansPerSecond,
		)

	case "streamsync":
		core.LoadGenerator(
			func() core.Client { return &grpc_stream.Client{} },
			func() core.Generator { return traceprotobuf.NewGenerator() },
			destination,
			spansPerSecond,
		)

	case "streamlbtimedsync":
		core.LoadGenerator(
			func() core.Client { return &grpc_stream_lb.Client{StreamReopenPeriod: rebalancePeriod} },
			func() core.Generator { return traceprotobuf.NewGenerator() },
			destination,
			spansPerSecond,
		)

	case "streamlbalwayssync":
		core.LoadGenerator(
			func() core.Client { return &grpc_stream_lb.Client{ReopenAfterEveryRequest: true} },
			func() core.Generator { return traceprotobuf.NewGenerator() },
			destination,
			spansPerSecond,
		)

	case "streamlbasync":
		core.LoadGenerator(
			func() core.Client { return &grpc_stream_lb_async.Client{StreamReopenPeriod: rebalancePeriod} },
			func() core.Generator { return traceprotobuf.NewGenerator() },
			destination,
			spansPerSecond,
		)

	case "wsstreamsync":
		core.LoadGenerator(
			func() core.Client { return &ws_stream_sync.Client{} },
			func() core.Generator { return traceprotobuf.NewGenerator() },
			destination,
			spansPerSecond,
		)

	case "wsstreamasync":
		core.LoadGenerator(
			func() core.Client { return &ws_stream_async.Client{Compression: traceprotobuf.CompressionMethod_NONE} },
			func() core.Generator { return traceprotobuf.NewGenerator() },
			destination,
			spansPerSecond,
		)

	case "wsstreamasynczlib":
		core.LoadGenerator(
			func() core.Client { return &ws_stream_async.Client{Compression: traceprotobuf.CompressionMethod_ZLIB} },
			func() core.Generator { return traceprotobuf.NewGenerator() },
			destination,
			spansPerSecond,
		)

	default:
		flag.Usage()
	}

}
