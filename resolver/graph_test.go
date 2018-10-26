package resolver_test

import (
    "context"
    "encoding/json"
    "github.com/s-ichikawa/piql/middleware"
    "github.com/s-ichikawa/piql/model"
    "github.com/s-ichikawa/piql/resolver"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestMutationResolver_CreateGraph(t *testing.T) {
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            w.WriteHeader(405)
            return
        }

        var v model.NewGraph
        d := json.NewDecoder(r.Body)
        if err := d.Decode(&v); err != nil {
            t.Fatal(err)
        }

        if got, want := v.Username, "s-ichikawa"; got != want {
            t.Fatalf("got %s, want %s", got, want)
        }
        if got, want := v.ID, "G12345"; got != want {
            t.Fatalf("got %s, want %s", got, want)
        }
        if got, want := v.Name, "TestGraph"; got != want {
            t.Fatalf("got %s, want %s", got, want)
        }
        if got, want := v.Unit, "commit"; got != want {
            t.Fatalf("got %s, want %s", got, want)
        }
        if got, want := v.Type, model.GraphTypeInt; got != want {
            t.Fatalf("got %s, want %s", got, want)
        }
        if got, want := v.Color, model.GraphColorShibafu; got != want {
            t.Fatalf("got %s, want %s", got, want)
        }


        if got, want := r.Header.Get("X-USER-TOKEN"), "hogehoge"; got != want {
            t.Fatalf("got %s, want %s", got, want)
        }

        w.WriteHeader(200)
        w.Header().Set("content-type", "applicaton/json")
        switch r.URL.Path {
        case "/v1/users/s-ichikawa/graphs":
            w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
        }
    }))
    defer ts.Close()

    resolvers := &resolver.Resolver{
        Host: ts.URL,
    }

    ctx := context.WithValue(context.TODO(), middleware.TokenContextKey, "hogehoge")
    res, err := resolvers.Mutation().CreateGraph(ctx, model.NewGraph{
        Username: "s-ichikawa",
        ID: "G12345",
        Name: "TestGraph",
        Unit: "commit",
        Type: "int",
        Color: "shibafu",
    })
    if err != nil {
        t.Fatalf("err: %s", err)
    }
    if res.IsSuccess != true {
        t.Fatalf("got %v, want true", res.IsSuccess)
    }
    if res.Message != "Success." {
        t.Fatalf("got %s, want Success", res.Message)
    }
}
