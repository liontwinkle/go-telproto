package otlptimewrapped

import (
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/tigrannajaryan/exp-otelproto/core"
)

// Generator allows to generate a ExportRequest.
type Generator struct {
	random     *rand.Rand
	tracesSent uint64
	spansSent  uint64
}

func NewGenerator() *Generator {
	return &Generator{
		random: rand.New(rand.NewSource(99)),
	}
}

func (g *Generator) genRandByteString(len int) string {
	b := make([]byte, len)
	for i := range b {
		b[i] = byte(g.random.Intn(10) + 33)
	}
	return string(b)
}

func (g *Generator) GenerateSpanBatch(spansPerBatch int, attrsPerSpan int, timedEventsPerSpan int) core.ExportRequest {
	traceID := atomic.AddUint64(&g.tracesSent, 1)

	resource := Resource{
		Process: &Process{
			StartTimeUnixnano: timeToTimestamp(time.Now()),
			Pid:               1234,
			HostName:          "fakehost",
		},
		ServiceName: "generator",
	}

	batch := &TraceExportRequest{ResourceSpans: []*ResourceSpans{{Resource: &resource}}}
	for i := 0; i < spansPerBatch; i++ {
		startTime := time.Date(2019, 10, 31, 10, 11, 12, 13, time.UTC)

		spanID := atomic.AddUint64(&g.spansSent, 1)

		// Create a span.
		span := &Span{
			TraceId:           core.GenerateTraceID(traceID),
			SpanId:            core.GenerateSpanID(spanID),
			Name:              "load-generator-span",
			Kind:              Span_CLIENT,
			StartTimeUnixnano: timeToTimestamp(startTime),
			EndTimeUnixnano:   timeToTimestamp(startTime.Add(time.Duration(i) * time.Millisecond)),
		}

		if attrsPerSpan >= 0 {
			// Append attributes.
			span.Attributes = []*AttributeKeyValue{}

			if attrsPerSpan >= 2 {
				span.Attributes = append(span.Attributes,
					&AttributeKeyValue{Key: "load_generator.span_seq_num", Type: AttributeKeyValue_INT64, IntValue: int64(spanID)})
				span.Attributes = append(span.Attributes,
					&AttributeKeyValue{Key: "load_generator.trace_seq_num", Type: AttributeKeyValue_INT64, IntValue: int64(traceID)})
			}

			for j := len(span.Attributes); j < attrsPerSpan; j++ {
				attrName := g.genRandByteString(g.random.Intn(20) + 1)
				span.Attributes = append(span.Attributes,
					&AttributeKeyValue{Key: attrName, Type: AttributeKeyValue_STRING, StringValue: g.genRandByteString(g.random.Intn(20) + 1)})
			}
		}

		if timedEventsPerSpan > 0 {
			span.TimeEvents = &Span_TimeEvents{}
			for i := 0; i < timedEventsPerSpan; i++ {
				span.TimeEvents.TimeEvent = append(span.TimeEvents.TimeEvent, &Span_TimeEvent{
					TimeUnixnano: timeToTimestamp(startTime),
					Attributes: []*AttributeKeyValue{
						{Key: "te", Type: AttributeKeyValue_INT64, IntValue: int64(spanID)},
					},
				})
			}
		}

		batch.ResourceSpans[0].Spans = append(batch.ResourceSpans[0].Spans, span)
	}
	return batch
}

func timeToTimestamp(t time.Time) *timestamp.Timestamp {
	nanoTime := t.UnixNano()
	return &timestamp.Timestamp{
		Seconds: nanoTime / 1e9,
		Nanos:   int32(nanoTime % 1e9),
	}
}

func (g *Generator) GenerateMetricBatch(metricsPerBatch int) core.ExportRequest {
	return nil
}
