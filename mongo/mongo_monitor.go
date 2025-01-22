package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/event"
)

func (d *documentDB) setupCommandMonitor() *event.CommandMonitor {
	return &event.CommandMonitor{
		Started: func(ctx context.Context, evt *event.CommandStartedEvent) {
			d.log.Info(ctx, fmt.Sprintf("MongoDB query started: %s %v", evt.CommandName, evt.Command))
		},
		Succeeded: func(ctx context.Context, evt *event.CommandSucceededEvent) {
			d.log.Info(ctx, fmt.Sprintf("MongoDB query succeeded: %s", evt.CommandName))
		},
		Failed: func(ctx context.Context, evt *event.CommandFailedEvent) {
			d.log.Error(ctx, fmt.Sprintf("MongoDB query failed: %s %v", evt.CommandName, evt.Failure))
		},
	}
}
