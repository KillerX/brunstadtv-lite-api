package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bcc-code/brunstadtv-lite-api/src/generated"
	"github.com/bcc-code/brunstadtv-lite-api/src/model"
)

// Global is the resolver for the global field.
func (r *configResolver) Global(ctx context.Context, obj *model.Config, timestamp *string) (*model.GlobalConfig, error) {
	panic(fmt.Errorf("not implemented: Global - global"))
}

// Config returns generated.ConfigResolver implementation.
func (r *Resolver) Config() generated.ConfigResolver { return &configResolver{r} }

type configResolver struct{ *Resolver }
