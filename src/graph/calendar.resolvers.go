package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bcc-code/brunstadtv-lite-api/src/generated"
	"github.com/bcc-code/brunstadtv-lite-api/src/model"
)

// Period is the resolver for the period field.
func (r *calendarResolver) Period(ctx context.Context, obj *model.Calendar, from string, to string) (*model.CalendarPeriod, error) {
	panic(fmt.Errorf("not implemented: Period - period"))
}

// Day is the resolver for the day field.
func (r *calendarResolver) Day(ctx context.Context, obj *model.Calendar, day string) (*model.CalendarDay, error) {
	panic(fmt.Errorf("not implemented: Day - day"))
}

// Event is the resolver for the event field.
func (r *episodeCalendarEntryResolver) Event(ctx context.Context, obj *model.EpisodeCalendarEntry) (*model.Event, error) {
	panic(fmt.Errorf("not implemented: Event - event"))
}

// Episode is the resolver for the episode field.
func (r *episodeCalendarEntryResolver) Episode(ctx context.Context, obj *model.EpisodeCalendarEntry) (*model.Episode, error) {
	panic(fmt.Errorf("not implemented: Episode - episode"))
}

// Event is the resolver for the event field.
func (r *seasonCalendarEntryResolver) Event(ctx context.Context, obj *model.SeasonCalendarEntry) (*model.Event, error) {
	panic(fmt.Errorf("not implemented: Event - event"))
}

// Season is the resolver for the season field.
func (r *seasonCalendarEntryResolver) Season(ctx context.Context, obj *model.SeasonCalendarEntry) (*model.Season, error) {
	panic(fmt.Errorf("not implemented: Season - season"))
}

// Event is the resolver for the event field.
func (r *showCalendarEntryResolver) Event(ctx context.Context, obj *model.ShowCalendarEntry) (*model.Event, error) {
	panic(fmt.Errorf("not implemented: Event - event"))
}

// Show is the resolver for the show field.
func (r *showCalendarEntryResolver) Show(ctx context.Context, obj *model.ShowCalendarEntry) (*model.Show, error) {
	panic(fmt.Errorf("not implemented: Show - show"))
}

// Event is the resolver for the event field.
func (r *simpleCalendarEntryResolver) Event(ctx context.Context, obj *model.SimpleCalendarEntry) (*model.Event, error) {
	panic(fmt.Errorf("not implemented: Event - event"))
}

// Calendar returns generated.CalendarResolver implementation.
func (r *Resolver) Calendar() generated.CalendarResolver { return &calendarResolver{r} }

// EpisodeCalendarEntry returns generated.EpisodeCalendarEntryResolver implementation.
func (r *Resolver) EpisodeCalendarEntry() generated.EpisodeCalendarEntryResolver {
	return &episodeCalendarEntryResolver{r}
}

// SeasonCalendarEntry returns generated.SeasonCalendarEntryResolver implementation.
func (r *Resolver) SeasonCalendarEntry() generated.SeasonCalendarEntryResolver {
	return &seasonCalendarEntryResolver{r}
}

// ShowCalendarEntry returns generated.ShowCalendarEntryResolver implementation.
func (r *Resolver) ShowCalendarEntry() generated.ShowCalendarEntryResolver {
	return &showCalendarEntryResolver{r}
}

// SimpleCalendarEntry returns generated.SimpleCalendarEntryResolver implementation.
func (r *Resolver) SimpleCalendarEntry() generated.SimpleCalendarEntryResolver {
	return &simpleCalendarEntryResolver{r}
}

type calendarResolver struct{ *Resolver }
type episodeCalendarEntryResolver struct{ *Resolver }
type seasonCalendarEntryResolver struct{ *Resolver }
type showCalendarEntryResolver struct{ *Resolver }
type simpleCalendarEntryResolver struct{ *Resolver }
