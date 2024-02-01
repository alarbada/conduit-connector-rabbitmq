package rabbitmq

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/alarbada/conduit-connector-rabbitmq/test"
	sdk "github.com/conduitio/conduit-connector-sdk"
	"github.com/matryer/is"
	"github.com/rabbitmq/amqp091-go"
)

func TestTeardownSource_NoOpen(t *testing.T) {
	is := is.New(t)
	con := NewSource()
	err := con.Teardown(context.Background())
	is.NoErr(err)
}

func newSourceCfg(queueName string) map[string]string {
	return map[string]string{
		"url":       test.URL,
		"queueName": queueName,
	}
}

func TestSource_Integration_RestartFull(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	is := is.New(t)

	queueName := test.SetupQueueName(t, is)
	cfgMap := newSourceCfg(queueName)

	recs1 := generateRabbitmqMsgs(1, 3)
	go produceRabbitmqMsgs(ctx, is, queueName, recs1)
	lastPosition := testSourceIntegrationRead(ctx, is, cfgMap, nil, recs1, false)

	recs2 := generateRabbitmqMsgs(4, 6)
	go produceRabbitmqMsgs(ctx, is, queueName, recs2)

	testSourceIntegrationRead(ctx, is, cfgMap, lastPosition, recs2, false)
}

func TestSource_Integration_RestartPartial(t *testing.T) {
	t.Parallel()

	is := is.New(t)
	ctx := context.Background()
	queueName := test.SetupQueueName(t, is)

	cfgMap := newSourceCfg(queueName)

	recs1 := generateRabbitmqMsgs(1, 3)
	go produceRabbitmqMsgs(ctx, is, queueName, recs1)

	lastPosition := testSourceIntegrationRead(ctx, is, cfgMap, nil, recs1, true)

	// only first record was acked, produce more records and expect to resume
	// from last acked record
	recs2 := generateRabbitmqMsgs(4, 6)
	go produceRabbitmqMsgs(ctx, is, queueName, recs2)

	var wantRecs []amqp091.Publishing
	wantRecs = append(wantRecs, recs1[1:]...)
	wantRecs = append(wantRecs, recs2...)

	testSourceIntegrationRead(ctx, is, cfgMap, lastPosition, wantRecs, false)
}

func generateRabbitmqMsgs(from, to int) []amqp091.Publishing {
	var msgs []amqp091.Publishing

	for i := from; i <= to; i++ {
		msg := amqp091.Publishing{
			Body: []byte(fmt.Sprintf("test-payload-%d", i)),
		}

		msgs = append(msgs, msg)
	}

	return msgs
}

func produceRabbitmqMsgs(ctx context.Context, is *is.I, queueName string, msgs []amqp091.Publishing) {
	conn, err := amqp091.Dial(test.URL)
	is.NoErr(err)

	defer conn.Close()

	ch, err := conn.Channel()
	is.NoErr(err)

	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	is.NoErr(err)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	for _, msg := range msgs {
		err = ch.PublishWithContext(ctx, "", q.Name, false, false, msg)
		is.NoErr(err)
	}
}

// testSourceIntegrationRead reads and acks messages in range [from,to].
// If ackFirst is true, only the first message will be acknowledged.
// Returns the position of the last message read.
func testSourceIntegrationRead(
	ctx context.Context,
	is *is.I,
	cfgMap map[string]string,
	startFrom sdk.Position,
	wantRecords []amqp091.Publishing,
	ackFirstOnly bool,
) sdk.Position {
	underTest := NewSource()
	defer func() {
		err := underTest.Teardown(ctx)
		is.NoErr(err)
	}()

	err := underTest.Configure(ctx, cfgMap)
	is.NoErr(err)
	err = underTest.Open(ctx, startFrom)
	is.NoErr(err)

	var positions []sdk.Position
	for _, wantRecord := range wantRecords {
		rec, err := underTest.Read(ctx)
		is.NoErr(err)

		recPayload := string(rec.Payload.After.Bytes())
		wantPayload := string(wantRecord.Body)
		is.Equal(wantPayload, recPayload)

		positions = append(positions, rec.Position)
	}

	for i, p := range positions {
		if i > 0 && ackFirstOnly {
			break
		}
		err = underTest.Ack(ctx, p)
		is.NoErr(err)
	}

	return positions[len(positions)-1]
}
