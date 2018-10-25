//go:generate gorunpkg github.com/99designs/gqlgen

package resolver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/s-ichikawa/piql/client"
	"github.com/s-ichikawa/piql/middleware"
	"github.com/s-ichikawa/piql/model"
	"io/ioutil"
	"log"
	"net/url"

	"github.com/s-ichikawa/piql"
)

type Resolver struct {
	Host string
}

func (r *Resolver) Mutation() piql.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() piql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

func (r *Resolver) getClient(path string, token *string) client.Client {
	url := url.URL{
		Host: r.Host,
		Path: path,
	}

	return client.Client{
		Url:   &url,
		Token: token,
	}
}
func (r *Resolver) getToken(ctx context.Context) *string {
	v := ctx.Value(middleware.TokenContextKey)
	token, ok := v.(string)
	if !ok {
		return nil
	}
	return &token
}

func (r *Resolver) request(method, path string, token *string, input interface{}) ([]byte, error) {
	fmt.Println(method, path, input)
	client := r.getClient(path, token)

	jsonStr, err := json.Marshal(&input)
	if err != nil {
		log.Fatal("failed to json.Marshal: ", err)
	}

	request, err := client.NewRequest(method, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatalf("failed to NewRequest: %s\n", err)
	}

	res, err := client.HttpClient.Do(request)
	if err != nil {
		log.Fatalf("failed to call api: %s\n", err)
	}

	if res.StatusCode == 404 {
		return nil, fmt.Errorf("404 not found: %s", path)
	}

	return ioutil.ReadAll(res.Body)
}

func (r *Resolver) mutation(ctx context.Context, method, path string, input interface{}) (*model.MutationResponse, error) {
	token := r.getToken(ctx)
	data, err := r.request(method, path, token, input)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %s", err)
	}

	var apiResponse model.MutationResponse
	err = json.Unmarshal(data, &apiResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to json.Unmarshal: %s", err)
	}
	return &apiResponse, nil
}

func (r *Resolver) get(ctx context.Context, path string) ([]byte, error) {
	token := r.getToken(ctx)
	client := r.getClient(path, token)

	req, err := client.NewRequest("GET", nil)
	if err != nil {
		log.Fatalf("failed to NewRequest: %s\n", err)
	}

	res, err := client.HttpClient.Do(req)
	if err != nil {
		log.Fatalf("failed to call api: %s\n", err)
	}

	if res.StatusCode == 404 {
		return nil, fmt.Errorf("not found: %s\n", path)
	}

	return ioutil.ReadAll(res.Body)
}
