package rabbitmq

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func (p *Publisher) PublishPhotoAsBase64(ctx context.Context, bytes []byte, id uuid.UUID) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if len(bytes) == 0 {
		return fmt.Errorf("photo bytes is empty")
	}

	encodedData := base64.StdEncoding.EncodeToString(bytes)

	message := map[string]interface{}{
		"photo_id":  id.String(),
		"size":      len(bytes),
		"data":      encodedData,
		"timestamp": time.Now().Unix(),
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	var publishCtx context.Context
	var cancel context.CancelFunc

	if _, ok := ctx.Deadline(); !ok {
		publishCtx, cancel = context.WithTimeout(ctx, 30*time.Second)
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
			Headers: amqp.Table{
				"photo_id":   id.String(),
				"photo_size": len(bytes),
				"encoding":   "base64",
			},
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish photo as base64: %w", err)
	}

	return nil
}
