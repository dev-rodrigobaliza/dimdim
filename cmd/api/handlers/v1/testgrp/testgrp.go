package testgrp

import (
	"context"
	"errors"
	"math/rand"
	"net/http"

	"github.com/dev-rodrigobaliza/dimdim/pkg/web"
	v1 "github.com/dev-rodrigobaliza/dimdim/internal/web/v1"
)

// Test is our example route.
func Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		return v1.NewRequestError(errors.New("TRUSTED ERROR"), http.StatusBadRequest)
	}

	// Validate the data
	// Call into the business layer

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
