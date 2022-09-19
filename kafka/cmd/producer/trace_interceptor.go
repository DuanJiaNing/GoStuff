package main

import (
	"context"
	"strings"

	"github.com/Shopify/sarama"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type OTelInterceptor struct {
	tracer     trace.Tracer
	fixedAttrs []attribute.KeyValue
}

// NewOTelInterceptor processes span for intercepted messages and add some
// headers with the span data.
func NewOTelInterceptor(brokers []string) *OTelInterceptor {
	oi := OTelInterceptor{}
	oi.tracer = otel.GetTracerProvider().Tracer("shopify.com/sarama/examples/interceptors")

	// These are based on the spec, which was reachable as of 2020-05-15
	// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/semantic_conventions/messaging.md
	oi.fixedAttrs = []attribute.KeyValue{
		attribute.String("messaging.destination_kind", "topic"),
		attribute.String("span.otel.kind", "PRODUCER"),
		attribute.String("messaging.system", "kafka"),
		attribute.String("net.transport", "IP.TCP"),
		attribute.String("messaging.url", strings.Join(brokers, ",")),
	}
	return &oi
}

func (oi *OTelInterceptor) OnSend(msg *sarama.ProducerMessage) {
	_, span := oi.tracer.Start(context.TODO(), msg.Topic)
	defer span.End()

	attWithTopic := append(oi.fixedAttrs,
		attribute.String("messaging.destination", msg.Topic),
	)
	span.SetAttributes(attWithTopic...)

	msg.Headers = append(msg.Headers,
		sarama.RecordHeader{Key: []byte("messaging.extra_info"), Value: []byte("asdf")},
	)
}
