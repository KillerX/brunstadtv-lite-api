package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/bcc-code/brunstadtv-lite-api/src/generated"
	"github.com/bcc-code/brunstadtv-lite-api/src/model"
)

// SetDevicePushToken is the resolver for the setDevicePushToken field.
func (r *mutationRootResolver) SetDevicePushToken(ctx context.Context, token string) (*model.Device, error) {
	panic(fmt.Errorf("not implemented: SetDevicePushToken - setDevicePushToken"))
}

// SetEpisodeProgress is the resolver for the setEpisodeProgress field.
func (r *mutationRootResolver) SetEpisodeProgress(ctx context.Context, id string, progress *int, duration *int) (*model.Episode, error) {
	panic(fmt.Errorf("not implemented: SetEpisodeProgress - setEpisodeProgress"))
}

// Application is the resolver for the application field.
func (r *queryRootResolver) Application(ctx context.Context) (*model.Application, error) {
	panic(fmt.Errorf("not implemented: Application - application"))
}

// Export is the resolver for the export field.
func (r *queryRootResolver) Export(ctx context.Context, groups []string) (*model.Export, error) {
	panic(fmt.Errorf("not implemented: Export - export"))
}

// Page is the resolver for the page field.
func (r *queryRootResolver) Page(ctx context.Context, id *string, code *string) (*model.Page, error) {
	panic(fmt.Errorf("not implemented: Page - page"))
}

// Section is the resolver for the section field.
func (r *queryRootResolver) Section(ctx context.Context, id string, timestamp *string) (model.Section, error) {
	panic(fmt.Errorf("not implemented: Section - section"))
}

// Show is the resolver for the show field.
func (r *queryRootResolver) Show(ctx context.Context, id string) (*model.Show, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	show, err := r.Queries.GetShowByID(ctx, int64(idInt))
	if err != nil {
		return nil, err
	}

	return show.AsGQL("no"), err
}

// Season is the resolver for the season field.
func (r *queryRootResolver) Season(ctx context.Context, id string) (*model.Season, error) {
	panic(fmt.Errorf("not implemented: Season - season"))
}

// Episode is the resolver for the episode field.
func (r *queryRootResolver) Episode(ctx context.Context, id string) (*model.Episode, error) {
	panic(fmt.Errorf("not implemented: Episode - episode"))
}

// Collection is the resolver for the collection field.
func (r *queryRootResolver) Collection(ctx context.Context, id string) (*model.Collection, error) {
	panic(fmt.Errorf("not implemented: Collection - collection"))
}

// Search is the resolver for the search field.
func (r *queryRootResolver) Search(ctx context.Context, queryString string, first *int, offset *int, typeArg *string, minScore *int) (*model.SearchResult, error) {
	panic(fmt.Errorf("not implemented: Search - search"))
}

// Calendar is the resolver for the calendar field.
func (r *queryRootResolver) Calendar(ctx context.Context) (*model.Calendar, error) {
	panic(fmt.Errorf("not implemented: Calendar - calendar"))
}

// Event is the resolver for the event field.
func (r *queryRootResolver) Event(ctx context.Context, id string) (*model.Event, error) {
	panic(fmt.Errorf("not implemented: Event - event"))
}

// Faq is the resolver for the faq field.
func (r *queryRootResolver) Faq(ctx context.Context) (*model.Faq, error) {
	panic(fmt.Errorf("not implemented: Faq - faq"))
}

// Me is the resolver for the me field.
func (r *queryRootResolver) Me(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Me - me"))
}

// Config is the resolver for the config field.
func (r *queryRootResolver) Config(ctx context.Context) (*model.Config, error) {
	panic(fmt.Errorf("not implemented: Config - config"))
}

// Profiles is the resolver for the profiles field.
func (r *queryRootResolver) Profiles(ctx context.Context) ([]*model.Profile, error) {
	panic(fmt.Errorf("not implemented: Profiles - profiles"))
}

// Profile is the resolver for the profile field.
func (r *queryRootResolver) Profile(ctx context.Context) (*model.Profile, error) {
	panic(fmt.Errorf("not implemented: Profile - profile"))
}

// MutationRoot returns generated.MutationRootResolver implementation.
func (r *Resolver) MutationRoot() generated.MutationRootResolver { return &mutationRootResolver{r} }

// QueryRoot returns generated.QueryRootResolver implementation.
func (r *Resolver) QueryRoot() generated.QueryRootResolver { return &queryRootResolver{r} }

type mutationRootResolver struct{ *Resolver }
type queryRootResolver struct{ *Resolver }
