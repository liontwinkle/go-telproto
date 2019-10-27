package otlp

import (
	"math/rand"
	"strconv"
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

func genResource() *Resource {
	return &Resource{
		Labels: []*AttributeKeyValue{
			{Key: "StartTimeUnixnano", Int64Value: 12345678},
			{Key: "Pid", Int64Value: 1234},
			{Key: "HostName", StringValue: "fakehost"},
			{Key: "ServiceName", StringValue: "generator"},
		},
	}
}

func (g *Generator) GenerateSpanBatch(spansPerBatch int, attrsPerSpan int, timedEventsPerSpan int) core.ExportRequest {
	traceID := atomic.AddUint64(&g.tracesSent, 1)

	batch := &TraceExportRequest{ResourceSpans: []*ResourceSpans{{Resource: genResource()}}}
	for i := 0; i < spansPerBatch; i++ {
		startTime := time.Now()

		spanID := atomic.AddUint64(&g.spansSent, 1)

		// Create a span.
		span := &Span{
			TraceId:           core.GenerateTraceID(traceID),
			SpanId:            core.GenerateSpanID(spanID),
			ParentSpanId:      []byte{},
			Name:              "load-generator-span",
			Kind:              Span_CLIENT,
			StartTimeUnixnano: core.TimeToTimestamp(startTime),
			EndTimeUnixnano:   core.TimeToTimestamp(startTime.Add(time.Duration(time.Millisecond))),
		}

		if attrsPerSpan >= 0 {
			// Append attributes.
			span.Attributes = []*AttributeKeyValue{}

			if attrsPerSpan >= 2 {
				span.Attributes = append(span.Attributes,
					&AttributeKeyValue{Key: "load_generator.span_seq_num", Type: AttributeKeyValue_INT64, Int64Value: int64(spanID)})
				span.Attributes = append(span.Attributes,
					&AttributeKeyValue{Key: "load_generator.trace_seq_num", Type: AttributeKeyValue_INT64, Int64Value: int64(traceID)})
			}

			for j := len(span.Attributes); j < attrsPerSpan; j++ {
				attrName := g.genRandByteString(g.random.Intn(50) + 1)
				span.Attributes = append(span.Attributes,
					&AttributeKeyValue{Key: attrName, Type: AttributeKeyValue_STRING, StringValue: g.genRandByteString(g.random.Intn(20) + 1)})
			}
		}

		if timedEventsPerSpan > 0 {
			for i := 0; i < timedEventsPerSpan; i++ {
				span.TimedEvents = append(span.TimedEvents, &Span_TimedEvent{
					TimeUnixnano: core.TimeToTimestamp(startTime),
					Attributes: []*AttributeKeyValue{
						{Key: "te", Type: AttributeKeyValue_INT64, Int64Value: int64(spanID)},
					},
				})
			}
		}

		batch.ResourceSpans[0].Spans = append(batch.ResourceSpans[0].Spans, span)
	}
	return batch
}

func (g *Generator) GenerateMetricBatch(metricsPerBatch int) core.ExportRequest {

	batch := &MetricExportRequest{ResourceMetrics: []*ResourceMetrics{{Resource: genResource()}}}
	for i := 0; i < metricsPerBatch/2; i++ {
		startTime := time.Now()

		labelKeys := []string{
			"label1",
			"label2",
		}

		// Add int64 Gauge
		descr := &MetricDescriptor{
			Name:        "metric" + strconv.Itoa(i),
			Description: "some description: " + strconv.Itoa(i),
			Type:        MetricDescriptor_GAUGE_INT64,
			LabelKeys:   labelKeys,
		}

		var timeseries []*GaugeInt64TimeSeries
		for j := 0; j < 5; j++ {
			var points []*GaugeInt64Value

			prevPointTs := int64(0)
			for k := 0; k < 5; k++ {
				pointTs := core.TimeToTimestamp(startTime.Add(time.Duration(j*k) * time.Millisecond))
				diffTs := pointTs - prevPointTs
				prevPointTs = pointTs

				point := GaugeInt64Value{
					TimestampDeltanano: diffTs,
					Value:              int64(i * j * k),
				}

				//sz := unsafe.Sizeof(SummaryValue{})
				//log.Printf("size=%v", sz)

				points = append(points, &point)
			}

			ts := GaugeInt64TimeSeries{
				LabelValues: []*LabelValue{
					{Value: "val1"},
					{Value: "val2"},
				},
				Points: points,
			}
			timeseries = append(timeseries, &ts)
		}

		metric1 := &Metric{
			MetricDescriptor: descr,
			Data: &Metric_GaugeInt64Data{
				GaugeInt64Data: &GaugeInt64TimeSeriesList{
					List: timeseries,
				},
			},
		}

		batch.ResourceMetrics[0].Metrics = append(batch.ResourceMetrics[0].Metrics, metric1)

		// Add Histogram
		descr = &MetricDescriptor{
			Name:        "metric" + strconv.Itoa(i),
			Description: "some description: " + strconv.Itoa(i),
			Type:        MetricDescriptor_GAUGE_INT64,
			LabelKeys:   labelKeys,
		}

		var timeseries2 []*HistogramTimeSeries
		for j := 0; j < 1; j++ {
			var points []*HistogramValue

			prevPointTs := int64(0)
			for k := 0; k < 5; k++ {
				pointTs := core.TimeToTimestamp(startTime.Add(time.Duration(j*k) * time.Millisecond))
				diffTs := pointTs - prevPointTs
				prevPointTs = pointTs
				val := float64(i * j * k)
				point := HistogramValue{
					TimestampDeltanano: diffTs,
					Count:              1,
					Sum:                val,
					Buckets: []*HistogramValue_Bucket{
						{
							Count: 1,
							Exemplar: &HistogramValue_Bucket_Exemplar{
								Value:             val,
								TimestampUnixnano: pointTs,
							},
						},
					},
				}
				points = append(points, &point)
			}

			ts := HistogramTimeSeries{
				LabelValues: []*LabelValue{
					{Value: "val1"},
					{Value: "val2"},
				},
				BucketOptions: &HistogramTimeSeries_ExplicitBounds_{
					ExplicitBounds: &HistogramTimeSeries_ExplicitBounds{
						Bounds: []float64{0, 1000000},
					},
				},
				Points: points,
			}
			timeseries2 = append(timeseries2, &ts)
		}

		metric2 := &Metric{
			MetricDescriptor: descr,
			Data: &Metric_HistogramData{
				HistogramData: &HistogramTimeSeriesList{
					List: timeseries2,
				},
			},
		}

		batch.ResourceMetrics[0].Metrics = append(batch.ResourceMetrics[0].Metrics, metric2)
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