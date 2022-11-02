package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bcc-code/brunstadtv-lite-api/src/generated"
	"github.com/bcc-code/brunstadtv-lite-api/src/model"
)

// Show is the resolver for the show field.
func (r *episodeSearchItemResolver) Show(ctx context.Context, obj *model.EpisodeSearchItem) (*model.Show, error) {
	panic(fmt.Errorf("not implemented: Show - show"))
}

// Season is the resolver for the season field.
func (r *episodeSearchItemResolver) Season(ctx context.Context, obj *model.EpisodeSearchItem) (*model.Season, error) {
	panic(fmt.Errorf("not implemented: Season - season"))
}

// Show is the resolver for the show field.
func (r *seasonSearchItemResolver) Show(ctx context.Context, obj *model.SeasonSearchItem) (*model.Show, error) {
	panic(fmt.Errorf("not implemented: Show - show"))
}

// EpisodeSearchItem returns generated.EpisodeSearchItemResolver implementation.
func (r *Resolver) EpisodeSearchItem() generated.EpisodeSearchItemResolver {
	return &episodeSearchItemResolver{r}
}

// SeasonSearchItem returns generated.SeasonSearchItemResolver implementation.
func (r *Resolver) SeasonSearchItem() generated.SeasonSearchItemResolver {
	return &seasonSearchItemResolver{r}
}

type episodeSearchItemResolver struct{ *Resolver }
type seasonSearchItemResolver struct{ *Resolver }
