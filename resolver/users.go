package resolver

import (
	"context"
	"github.com/s-ichikawa/piql/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.MutationResponse, error) {
	return r.mutation(ctx, "POST", "/v1/users", input)
}

func (r *mutationResolver) UpdateToken(ctx context.Context, input model.NewToken) (*model.MutationResponse, error) {
	return r.mutation(ctx, "PUT", "/v1/users/"+input.Username, input)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input model.DeleteUser) (*model.MutationResponse, error) {
	return r.mutation(ctx, "DELETE", "/v1/users/"+input.Username, nil)
}
