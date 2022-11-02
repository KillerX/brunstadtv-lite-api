package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	merry "github.com/ansel1/merry/v2"
	"github.com/bcc-code/brunstadtv-lite-api/src/generated"
	"github.com/bcc-code/brunstadtv-lite-api/src/model"
)

// Image is the resolver for the image field.
func (r *episodeResolver) Image(ctx context.Context, obj *model.Episode, style *model.ImageStyle) (*string, error) {
	panic(fmt.Errorf("not implemented: Image - image"))
}

// Streams is the resolver for the streams field.
func (r *episodeResolver) Streams(ctx context.Context, obj *model.Episode) ([]*model.Stream, error) {
	panic(fmt.Errorf("not implemented: Streams - streams"))
}

// Files is the resolver for the files field.
func (r *episodeResolver) Files(ctx context.Context, obj *model.Episode) ([]*model.File, error) {
	panic(fmt.Errorf("not implemented: Files - files"))
}

// Season is the resolver for the season field.
func (r *episodeResolver) Season(ctx context.Context, obj *model.Episode) (*model.Season, error) {
	panic(fmt.Errorf("not implemented: Season - season"))
}

// Progress is the resolver for the progress field.
func (r *episodeResolver) Progress(ctx context.Context, obj *model.Episode) (*int, error) {
	panic(fmt.Errorf("not implemented: Progress - progress"))
}

// Image is the resolver for the image field.
func (r *seasonResolver) Image(ctx context.Context, obj *model.Season, style *model.ImageStyle) (*string, error) {
	panic(fmt.Errorf("not implemented: Image - image"))
}

// Show is the resolver for the show field.
func (r *seasonResolver) Show(ctx context.Context, obj *model.Season) (*model.Show, error) {
	panic(fmt.Errorf("not implemented: Show - show"))
}

// Episodes is the resolver for the episodes field.
func (r *seasonResolver) Episodes(ctx context.Context, obj *model.Season, first *int, offset *int, dir *string) (*model.EpisodePagination, error) {
	panic(fmt.Errorf("not implemented: Episodes - episodes"))
}

// Image is the resolver for the image field.
func (r *showResolver) Image(ctx context.Context, obj *model.Show, style *model.ImageStyle) (*string, error) {
	panic(fmt.Errorf("not implemented: Image - image"))
}

// EpisodeCount is the resolver for the episodeCount field.
func (r *showResolver) EpisodeCount(ctx context.Context, obj *model.Show) (int, error) {
	panic(fmt.Errorf("not implemented: EpisodeCount - episodeCount"))
}

// SeasonCount is the resolver for the seasonCount field.
func (r *showResolver) SeasonCount(ctx context.Context, obj *model.Show) (int, error) {
	panic(fmt.Errorf("not implemented: SeasonCount - seasonCount"))
}

// Seasons is the resolver for the seasons field.
func (r *showResolver) Seasons(ctx context.Context, obj *model.Show, first *int, offset *int, dir *string) (*model.SeasonPagination, error) {
	id, err := strconv.Atoi(obj.ID)
	if err != nil {
		return nil, merry.Wrap(err)
	}

	seasons, err := r.Queries.GetSeasonsForShow(ctx, int64(id))
	if err != nil {
		return nil, merry.Wrap(err)
	}

	gqlSesons := make([]*model.Season, len(seasons))
	for i, s := range seasons {
		gqlSesons[i] = s.AsGQL("no")
	}

	return &model.SeasonPagination{
		Total:  len(gqlSesons),
		First:  len(gqlSesons),
		Offset: 0,
		Items:  gqlSesons,
	}, nil
}

// DefaultEpisode is the resolver for the defaultEpisode field.
func (r *showResolver) DefaultEpisode(ctx context.Context, obj *model.Show) (*model.Episode, error) {
	panic(fmt.Errorf("not implemented: DefaultEpisode - defaultEpisode"))
}

// Episode returns generated.EpisodeResolver implementation.
func (r *Resolver) Episode() generated.EpisodeResolver { return &episodeResolver{r} }

// Season returns generated.SeasonResolver implementation.
func (r *Resolver) Season() generated.SeasonResolver { return &seasonResolver{r} }

// Show returns generated.ShowResolver implementation.
func (r *Resolver) Show() generated.ShowResolver { return &showResolver{r} }

type episodeResolver struct{ *Resolver }
type seasonResolver struct{ *Resolver }
type showResolver struct{ *Resolver }
