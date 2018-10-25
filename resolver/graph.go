package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/s-ichikawa/piql/model"
	"github.com/vektah/gqlparser/gqlerror"
	"log"
)

func (r *mutationResolver) CreateGraph(ctx context.Context, input model.NewGraph) (*model.MutationResponse, error) {
	path := fmt.Sprintf("/v1/users/%s/graphs", input.Username)
	return r.mutation(ctx, "POST", path, input)
}

func (r *mutationResolver) UpdateGraph(ctx context.Context, input model.UpdateGraph) (*model.MutationResponse, error) {
	path := fmt.Sprintf("/v1/users/%s/graphs/%s", input.Username, input.ID)
	return r.mutation(ctx, "PUT", path, input)
}

func (r *mutationResolver) DeleteGraph(ctx context.Context, input model.DeleteGraph) (*model.MutationResponse, error) {
	path := fmt.Sprintf("/v1/users/%s/graphs/%s", input.Username, input.ID)
	return r.mutation(ctx, "DELETE", path, nil)
}

type GraphSlice struct {
	Graphs []model.GraphInfo
}

func (r *queryResolver) Graphs(ctx context.Context, input *model.GetGraphs) ([]model.GraphInfo, error) {
	path := fmt.Sprintf("/v1/users/%s/graphs", input.Username)
	data, err := r.get(ctx, path)
	if err != nil {
		log.Fatalf("failed to read response: %s\n", err)
	}

	var graphs GraphSlice
	err = json.Unmarshal(data, &graphs)
	if err != nil {
		log.Fatalf("failed to json.Unmarshal: %s\n", err)
	}
	return graphs.Graphs, nil
}

func (r *queryResolver) Graph(ctx context.Context, input *model.GetGraph) (*model.Graph, error) {
	path := fmt.Sprintf("/v1/users/%s/graphs/%s", input.Username, input.ID)
	data, err := r.get(ctx, path)
	if err != nil {
		return nil, gqlerror.Errorf("get error: %s", err)
	}

	graph := model.Graph{
		Id:  input.ID,
		Svg: string(data),
	}
	return &graph, nil
}
