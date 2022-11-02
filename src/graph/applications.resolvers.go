package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bcc-code/brunstadtv-lite-api/src/generated"
	"github.com/bcc-code/brunstadtv-lite-api/src/model"
)

// Page is the resolver for the page field.
func (r *applicationResolver) Page(ctx context.Context, obj *model.Application) (*model.Page, error) {
	panic(fmt.Errorf("not implemented: Page - page"))
}

// Application returns generated.ApplicationResolver implementation.
func (r *Resolver) Application() generated.ApplicationResolver { return &applicationResolver{r} }

type applicationResolver struct{ *Resolver }
