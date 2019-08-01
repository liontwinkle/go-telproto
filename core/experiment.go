package core

import (
	"log"
	"net"
	"os"
	"sync"
	"time"

	"github.com/shirou/gopsutil/process"
)

func onBatchReceive(batch ExportRequest) {
	log.Printf("Server received a batch")
}

func RunTest(clnt Client, srv Server, gen Generator) {

	// Listen locally for Agent's forwarded data
	go srv.Listen("0.0.0.0:3465", onBatchReceive)

	// Connect to Agent
	clnt.Connect("localhost:3465")

	// Generate and send a batch
	for i := 0; i < 2; i++ {
		batch := gen.GenerateBatch(100, 2)
		clnt.Export(batch)
	}
}

type Options struct {
	Batches       int
	SpansPerBatch int
	AttrPerSpan   int
}

func BenchmarkLocalDelivery(
	clientFactory func() Client,
	serverFactory func() Server,
	generatorFactory func() Generator,
	options Options,
) (cpuSecs float64, wallSecs float64) {
	// Create client, server and generator from factories
	clnt := clientFactory()
	srv := serverFactory()
	gen := generatorFactory()

	// Find a local address for delivery.
	endpoint := GetAvailableLocalAddress()

	// Create a WaitGroup to count sent/received Batches.
	wg := sync.WaitGroup{}

	// Server listen locally.
	go srv.Listen(endpoint, func(batch ExportRequest) {
		// Count delivered batch.
		wg.Done()
	})

	// Clietn connect to the server.
	clnt.Connect(endpoint)

	// Begin measuring CPU time.
	proc, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		log.Fatal(err)
	}

	startCPUTimes, err := proc.Times()
	if err != nil {
		log.Fatal(err)
	}

	// Begin measuring wall time.
	startWallTime := time.Now()

	// Generate and send Batches.
	for i := 0; i < options.Batches; i++ {
		// Count sent batch.
		wg.Add(1)
		batch := gen.GenerateBatch(options.SpansPerBatch, options.AttrPerSpan)
		clnt.Export(batch)
	}

	// Wait until all Batches are delivered.
	wg.Wait()

	// Measure used CPU time.
	endCPUTimes, err := proc.Times()
	if err != nil {
		log.Fatal(err)
	}
	cpuSecs = endCPUTimes.Total() - startCPUTimes.Total()

	// Measure used wall time.
	endWallTime := time.Now()
	wallSecs = endWallTime.Sub(startWallTime).Seconds()

	// Stop the server.
	srv.Stop()

	return
}

func RunAgent(clnt Client, srv Server, listenAddress, destination string) {

	log.Printf("Agent: listening on %s", listenAddress)
	log.Printf("Agent: forwarding to %s", destination)

	err := clnt.Connect(destination)
	if err != nil {
		log.Fatalf("Cannot connection to %v: %v", destination, err)
	}

	srv.Listen(listenAddress, func(batch ExportRequest) {
		log.Printf("Agent: forwarding span batch")
		clnt.Export(batch)
	})
}

// GetAvailableLocalAddress finds an available local port and returns an endpoint
// describing it. The port is available for opening when this function returns
// provided that there is no race by some other code to grab the same port
// immediately.
func GetAvailableLocalAddress() string {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("failed to get a free local port: %v", err)
	}
	// There is a possible race if something else takes this same port before
	// the test uses it, however, that is unlikely in practice.
	defer ln.Close()
	return ln.Addr().String()
}
