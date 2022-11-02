package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bcc-code/brunstadtv-lite-api/src/generated"
	"github.com/bcc-code/brunstadtv-lite-api/src/model"
)

// Items is the resolver for the items field.
func (r *collectionResolver) Items(ctx context.Context, obj *model.Collection, first *int, offset *int) (*model.CollectionItemPagination, error) {
	panic(fmt.Errorf("not implemented: Items - items"))
}

// Items is the resolver for the items field.
func (r *defaultGridSectionResolver) Items(ctx context.Context, obj *model.DefaultGridSection, first *int, offset *int) (*model.SectionItemPagination, error) {
	panic(fmt.Errorf("not implemented: Items - items"))
}

// Items is the resolver for the items field.
func (r *defaultSectionResolver) Items(ctx context.Context, obj *model.DefaultSection, first *int, offset *int) (*model.SectionItemPagination, error) {
	panic(fmt.Errorf("not implemented: Items - items"))
}

// Items is the resolver for the items field.
func (r *featuredSectionResolver) Items(ctx context.Context, obj *model.FeaturedSection, first *int, offset *int) (*model.SectionItemPagination, error) {
	panic(fmt.Errorf("not implemented: Items - items"))
}

// Items is the resolver for the items field.
func (r *iconSectionResolver) Items(ctx context.Context, obj *model.IconSection, first *int, offset *int) (*model.SectionItemPagination, error) {
	panic(fmt.Errorf("not implemented: Items - items"))
}

// Items is the resolver for the items field.
func (r *labelSectionResolver) Items(ctx context.Context, obj *model.LabelSection, first *int, offset *int) (*model.SectionItemPagination, error) {
	panic(fmt.Errorf("not implemented: Items - items"))
}

// Messages is the resolver for the messages field.
func (r *messageSectionResolver) Messages(ctx context.Context, obj *model.MessageSection) ([]*model.Message, error) {
	panic(fmt.Errorf("not implemented: Messages - messages"))
}

// Image is the resolver for the image field.
func (r *pageResolver) Image(ctx context.Context, obj *model.Page, style *model.ImageStyle) (*string, error) {
	panic(fmt.Errorf("not implemented: Image - image"))
}

// Sections is the resolver for the sections field.
func (r *pageResolver) Sections(ctx context.Context, obj *model.Page, first *int, offset *int) (*model.SectionPagination, error) {
	panic(fmt.Errorf("not implemented: Sections - sections"))
}

// Items is the resolver for the items field.
func (r *posterGridSectionResolver) Items(ctx context.Context, obj *model.PosterGridSection, first *int, offset *int) (*model.SectionItemPagination, error) {
	panic(fmt.Errorf("not implemented: Items - items"))
}

// Items is the resolver for the items field.
func (r *posterSectionResolver) Items(ctx context.Context, obj *model.PosterSection, first *int, offset *int) (*model.SectionItemPagination, error) {
	panic(fmt.Errorf("not implemented: Items - items"))
}

// Collection returns generated.CollectionResolver implementation.
func (r *Resolver) Collection() generated.CollectionResolver { return &collectionResolver{r} }

// DefaultGridSection returns generated.DefaultGridSectionResolver implementation.
func (r *Resolver) DefaultGridSection() generated.DefaultGridSectionResolver {
	return &defaultGridSectionResolver{r}
}

// DefaultSection returns generated.DefaultSectionResolver implementation.
func (r *Resolver) DefaultSection() generated.DefaultSectionResolver {
	return &defaultSectionResolver{r}
}

// FeaturedSection returns generated.FeaturedSectionResolver implementation.
func (r *Resolver) FeaturedSection() generated.FeaturedSectionResolver {
	return &featuredSectionResolver{r}
}

// IconSection returns generated.IconSectionResolver implementation.
func (r *Resolver) IconSection() generated.IconSectionResolver { return &iconSectionResolver{r} }

// LabelSection returns generated.LabelSectionResolver implementation.
func (r *Resolver) LabelSection() generated.LabelSectionResolver { return &labelSectionResolver{r} }

// MessageSection returns generated.MessageSectionResolver implementation.
func (r *Resolver) MessageSection() generated.MessageSectionResolver {
	return &messageSectionResolver{r}
}

// Page returns generated.PageResolver implementation.
func (r *Resolver) Page() generated.PageResolver { return &pageResolver{r} }

// PosterGridSection returns generated.PosterGridSectionResolver implementation.
func (r *Resolver) PosterGridSection() generated.PosterGridSectionResolver {
	return &posterGridSectionResolver{r}
}

// PosterSection returns generated.PosterSectionResolver implementation.
func (r *Resolver) PosterSection() generated.PosterSectionResolver { return &posterSectionResolver{r} }

type collectionResolver struct{ *Resolver }
type defaultGridSectionResolver struct{ *Resolver }
type defaultSectionResolver struct{ *Resolver }
type featuredSectionResolver struct{ *Resolver }
type iconSectionResolver struct{ *Resolver }
type labelSectionResolver struct{ *Resolver }
type messageSectionResolver struct{ *Resolver }
type pageResolver struct{ *Resolver }
type posterGridSectionResolver struct{ *Resolver }
type posterSectionResolver struct{ *Resolver }
