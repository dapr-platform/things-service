package eventpub

import (
	"context"
	"github.com/dapr-platform/common"
	"github.com/pkg/errors"
)

func PublishInternalMessage(ctx context.Context, pubsub, topic string, msg any) error {
	err := common.GetDaprClient().PublishEvent(ctx, pubsub, topic, msg)
	if err != nil {
		err = errors.Wrap(err, "PublishInternalMessage")
	}
	return err
}
