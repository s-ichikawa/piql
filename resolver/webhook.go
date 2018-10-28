package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/s-ichikawa/piql/model"
	"log"
)

func (r *mutationResolver) CreateWebhook(ctx context.Context, input model.NewWebhook) (model.NewWebhookResponse, error) {
	path := fmt.Sprintf("/v1/users/%s/webhooks", input.Username)
	token := r.getToken(ctx)

	data, err := r.request("POST", path, token, input)
	if err != nil {
		log.Fatal("failed to read response")
	}

	var res model.NewWebhookResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Fatal("failed to json.Unmarshal:", err)
	}

	return res, nil
}

func (r *mutationResolver) InvokeWebhook(ctx context.Context, input model.CallWebhook) (*model.MutationResponse, error) {
	path := fmt.Sprintf("/v1/users/%s/webhooks/%s", input.Username, input.HashString)
	return r.mutation(ctx, "POST", path, nil)
}

func (r *mutationResolver) DeleteWebhook(ctx context.Context, input model.DeleteWebhook) (*model.MutationResponse, error) {
	path := fmt.Sprintf("/v1/users/%s/webhooks/%s", input.Username, input.HashString)
	return r.mutation(ctx, "DELETE", path, nil)
}

type WebhookSlice struct {
	Webhooks []*model.Webhook
}

func (r *queryResolver) Webhooks(ctx context.Context, input *model.GetWebhooks) ([]*model.Webhook, error) {
	path := fmt.Sprintf("/v1/users/%s/webhooks", input.Username)

	data, err := r.get(ctx, path)
	if err != nil {
		log.Fatalf("failed to get: %s", err)
	}

	var webhooks WebhookSlice
	err = json.Unmarshal(data, &webhooks)
	if err != nil {
		log.Fatalf("failed to json.Unmarshal: %s\n", err)
	}
	return webhooks.Webhooks, nil
}
