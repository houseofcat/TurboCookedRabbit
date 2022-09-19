package main_test

import (
	"fmt"
	"testing"

	"github.com/houseofcat/turbocookedrabbit/v2/pkg/tcr"
	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestCreateConsumer(t *testing.T) {
	defer goleak.VerifyNone(t)

	consumer1 := tcr.NewConsumerFromConfig(AckableConsumerConfig, ConnectionPool)
	assert.NotNil(t, consumer1)

	consumer2 := tcr.NewConsumerFromConfig(ConsumerConfig, ConnectionPool)
	assert.NotNil(t, consumer2)

	TestCleanup(t)
}

func TestStartStopConsumer(t *testing.T) {
	defer goleak.VerifyNone(t)

	consumer := tcr.NewConsumerFromConfig(ConsumerConfig, ConnectionPool)
	assert.NotNil(t, consumer)

	consumer.StartConsuming()
	err := consumer.StopConsuming(false, false)
	assert.NoError(t, err)

	TestCleanup(t)
}

func TestStartWithActionStopConsumer(t *testing.T) {
	defer goleak.VerifyNone(t)

	consumer := tcr.NewConsumerFromConfig(ConsumerConfig, ConnectionPool)
	assert.NotNil(t, consumer)

	consumer.StartConsumingWithAction(
		func(msg *tcr.ReceivedMessage) {
			if err := msg.Acknowledge(); err != nil {
				fmt.Printf("Error acking message: %v\r\n", msg.Delivery.Body)
			}
		})
	err := consumer.StopConsuming(false, false)
	assert.NoError(t, err)

	TestCleanup(t)
}

func TestConsumerGet(t *testing.T) {
	defer goleak.VerifyNone(t)

	consumer := tcr.NewConsumerFromConfig(ConsumerConfig, ConnectionPool)
	assert.NotNil(t, consumer)

	delivery, err := consumer.Get("TcrTestQueue")
	assert.Nil(t, delivery) // empty queue should be nil
	assert.NoError(t, err)

	TestCleanup(t)
}
