package testgrp

import (
	"context"
	"errors"
	"math/rand"
	"net/http"

	"github.com/sudonite/service/foundation/web"
)

// Test is an example route.
func Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		return errors.New("UNTRUSTED ERROR")
	}

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
