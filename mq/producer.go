package mq

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"

	"github.com/airren/echo-bio-backend/config"
	"github.com/airren/echo-bio-backend/model"
)

var mqconn *amqp.Connection

func PublishJob(ctx context.Context, job *model.AnalysisJob) {
	mqconn, err := amqp.Dial(config.Conf.MqURI)
	if err != nil {
		log.Errorf("Failed to connect to RabbitMQ: %s", err)
		return
	}
	defer mqconn.Close()

	ch, err := mqconn.Channel()
	if err != nil {
		log.Error(err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"analysis_job", false, false, false, false, nil)
	if err != nil {
		log.Errorf("declare queue failed: %s", err)
		return
	}

	body, _ := json.Marshal(job)
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/json",
			Body:        []byte(body),
		})

	if err != nil {
		log.Errorf("publish analysis_job failed: %s", err)
	}

}

func ConsumerJob() {
	mqconn, err := amqp.Dial(config.Conf.MqURI)
	if err != nil {
		log.Errorf("Failed to connect to RabbitMQ: %s", err)
		return
	}
	defer mqconn.Close()

	ch, err := mqconn.Channel()
	if err != nil {
		log.Error(err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"analysis_job", false, false, false, false, nil)
	if err != nil {
		log.Errorf("declare queue failed: %s", err)
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Errorf("consumer job failed: %s", err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			job := &model.AnalysisJob{}

			err := json.Unmarshal(d.Body, job)
			if err != nil {
				log.Errorf("consumer unmarshal job failed: %v", err)
				continue
			}
			log.Printf("Received a message: %v", job)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
