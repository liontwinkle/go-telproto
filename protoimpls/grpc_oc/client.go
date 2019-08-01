package grpc_oc

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/octraceprotobuf"
)

// Client can connect to a server and send a batch of spans.
type Client struct {
	client     octraceprotobuf.OCStreamTracerClient
	stream     octraceprotobuf.OCStreamTracer_SendBatchClient
	WaitForAck bool
}

func (c *Client) Connect(server string) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c.client = octraceprotobuf.NewOCStreamTracerClient(conn)

	// Establish stream to server.
	c.stream, err = c.client.SendBatch(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Export(batch core.SpanBatch) {
	// Send the batch via stream.
	c.stream.Send(batch.(*octraceprotobuf.SpanBatch))

	if c.WaitForAck {
		// Wait for response from server. This is full synchronous operation,
		// we do not send batches concurrently.
		_, err := c.stream.Recv()

		if err != nil {
			log.Fatal("Error from server when expecting batch response")
		}
	}
}
