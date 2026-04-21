package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"photo-upload-service/internal/models"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (p *Publisher) PublishPhoto(ctx context.Context, id uuid.UUID, ext string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	message := models.EvaluationMessage{
		PhotoID: id,
		Method:  "Mock method",
		Ext:     ext,
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	var publishCtx context.Context
	var cancel context.CancelFunc

	if _, ok := ctx.Deadline(); !ok {
		publishCtx, cancel = context.WithTimeout(ctx, p.publishTimeout)
		defer cancel()
	} else {
		publishCtx = ctx
	}

	err = p.channel.PublishWithContext(
		publishCtx,
		"",
		p.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         jsonData,
			DeliveryMode: amqp.Persistent,
			MessageId:    id.String(),
			Timestamp:    time.Now(),
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish photo as base64: %w", err)
	}

	return nil
}
