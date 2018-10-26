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

func TestMutationResolver_CreateUser(t *testing.T) {
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            w.WriteHeader(405)
            return
        }

        var v model.NewUser
        d := json.NewDecoder(r.Body)
        if err := d.Decode(&v); err != nil {
            t.Fatal(err)
        }

        if got, want := v.Username, "s-ichikawa"; got != want {
            t.Fatalf("got %s, want %s", got, want)
        }

        if got, want := v.Token, "hogehoge"; got != want {
            t.Fatalf("got %s, want %s", got, want)
        }

        if got, want := v.AgreeTermsOfService, "yes"; got != want {
            t.Fatalf("got %s, want %s", got, want)
        }

        if got, want := v.NotMinor, "yes"; got != want {
            t.Fatalf("got %s, want %s", got, want)
        }

        w.WriteHeader(200)
        w.Header().Set("content-type", "applicaton/json")
        switch r.URL.Path {
        case "/v1/users":
            w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
        }
    }))
    defer ts.Close()

    resolvers := &resolver.Resolver{
        Host: ts.URL,
    }

    res, err := resolvers.Mutation().CreateUser(context.TODO(), model.NewUser{
        Username:            "s-ichikawa",
        Token:               "hogehoge",
        AgreeTermsOfService: "yes",
        NotMinor:            "yes",
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

func TestMutationResolver_UpdateToken(t *testing.T) {
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPut {
            w.WriteHeader(405)
            return
        }

        var v model.NewToken
        d := json.NewDecoder(r.Body)
        if err := d.Decode(&v); err != nil {
            t.Fatal(err)
        }

        if got, want := v.Username, "s-ichikawa"; got != want {
            t.Fatalf("got %s, want %s", got, want)
        }

        if got, want := v.NewToken, "fugafuga"; got != want {
            t.Fatalf("got %s, want %s", got, want)
        }

        if got, want := r.Header.Get("X-USER-TOKEN"), "hogehoge"; got != want {
            t.Fatalf("got %s, want %s", got, want)
        }

        w.WriteHeader(200)
        w.Header().Set("content-type", "applicaton/json")
        switch r.URL.Path {
        case "/v1/users/s-ichikawa":
            w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
        }
    }))
    defer ts.Close()

    resolvers := &resolver.Resolver{
        Host: ts.URL,
    }

    ctx := context.WithValue(context.TODO(), middleware.TokenContextKey, "hogehoge")
    res, err := resolvers.Mutation().UpdateToken(ctx, model.NewToken{
        Username: "s-ichikawa",
        NewToken: "fugafuga",
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


func TestMutationResolver_DeleteUser(t *testing.T) {
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodDelete {
            w.WriteHeader(405)
            return
        }

        var v model.DeleteUser
        d := json.NewDecoder(r.Body)
        if err := d.Decode(&v); err != nil {
            t.Fatal(err)
        }

        if got, want := v.Username, ""; got != want {
            t.Fatalf("got %s, want %s", got, want)
        }

        if got, want := r.Header.Get("X-USER-TOKEN"), "hogehoge"; got != want {
            t.Fatalf("got %s, want %s", got, want)
        }

        w.WriteHeader(200)
        w.Header().Set("content-type", "applicaton/json")
        switch r.URL.Path {
        case "/v1/users/s-ichikawa":
            w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
        }
    }))
    defer ts.Close()

    resolvers := &resolver.Resolver{
        Host: ts.URL,
    }

    ctx := context.WithValue(context.TODO(), middleware.TokenContextKey, "hogehoge")
    res, err := resolvers.Mutation().DeleteUser(ctx, model.DeleteUser{
        Username: "s-ichikawa",
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