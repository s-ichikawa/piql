package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/s-ichikawa/piql/model"
	"github.com/vektah/gqlparser/gqlerror"
	"log"
)

func pixelId(userName, id, date string) string {
	return fmt.Sprintf("%s_%s_%s", userName, id, date)
}

func (r *mutationResolver) CreatePixel(ctx context.Context, input model.NewPixel) (*model.MutationResponse, error) {
	path := fmt.Sprintf("/v1/users/%s/graphs/%s/", input.Username, input.ID)
	return r.mutation(ctx, "POST", path, input)
}

func (r *mutationResolver) UpdatePixel(ctx context.Context, input model.NewPixel) (*model.MutationResponse, error) {
	path := fmt.Sprintf("/v1/users/%s/graphs/%s/", input.Username, input.ID)
	return r.mutation(ctx, "PUT", path, input)
}

func (r *mutationResolver) IncrementPixel(ctx context.Context, input model.IncrementPixel) (*model.MutationResponse, error) {
	path := fmt.Sprintf("/v1/users/%s/graphs/%s/increment", input.Username, input.ID)
	return r.mutation(ctx, "PUT", path, input)
}

func (r *mutationResolver) DecrementPixel(ctx context.Context, input model.DecrementPixel) (*model.MutationResponse, error) {
	path := fmt.Sprintf("/v1/users/%s/graphs/%s/decrement", input.Username, input.ID)
	return r.mutation(ctx, "PUT", path, input)
}

func (r *mutationResolver) DeletePixel(ctx context.Context, input model.DeletePixel) (*model.MutationResponse, error) {
	path := fmt.Sprintf("/v1/users/%s/graphs/%s/%s", input.Username, input.ID, input.Date)
	return r.mutation(ctx, "DELETE", path, input)
}

func (r *queryResolver) Pixel(ctx context.Context, input *model.GetPixel) (*model.Pixel, error) {
	path := fmt.Sprintf("/v1/users/%s/graphs/%s/%s", input.Username, input.ID, input.Date)
	data, err := r.get(ctx, path)
	if err != nil {
		log.Fatalf("failed to get: %s", err)
	}

	var apiResponse model.MutationResponse
	err = json.Unmarshal(data, &apiResponse)
	if err == nil && len(apiResponse.Message) > 0 {
		return &model.Pixel{}, gqlerror.Errorf("error: %s", apiResponse.Message)
	}

	var pixel model.Pixel
	err = json.Unmarshal(data, &pixel)
	if err != nil {
		return &model.Pixel{}, gqlerror.Errorf("failed to json.Unmarshal: %s", err)
	}

	pixel.Id = pixelId(input.Username, input.ID, input.Date)
	return &pixel, nil
}
