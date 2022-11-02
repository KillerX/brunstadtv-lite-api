package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bcc-code/brunstadtv-lite-api/src/generated"
	"github.com/bcc-code/brunstadtv-lite-api/src/model"
)

// Categories is the resolver for the categories field.
func (r *fAQResolver) Categories(ctx context.Context, obj *model.Faq, first *int, offset *int) (*model.FAQCategoryPagination, error) {
	panic(fmt.Errorf("not implemented: Categories - categories"))
}

// Category is the resolver for the category field.
func (r *fAQResolver) Category(ctx context.Context, obj *model.Faq, id string) (*model.FAQCategory, error) {
	panic(fmt.Errorf("not implemented: Category - category"))
}

// Question is the resolver for the question field.
func (r *fAQResolver) Question(ctx context.Context, obj *model.Faq, id string) (*model.Question, error) {
	panic(fmt.Errorf("not implemented: Question - question"))
}

// Questions is the resolver for the questions field.
func (r *fAQCategoryResolver) Questions(ctx context.Context, obj *model.FAQCategory, first *int, offset *int) (*model.QuestionPagination, error) {
	panic(fmt.Errorf("not implemented: Questions - questions"))
}

// Category is the resolver for the category field.
func (r *questionResolver) Category(ctx context.Context, obj *model.Question) (*model.FAQCategory, error) {
	panic(fmt.Errorf("not implemented: Category - category"))
}

// FAQ returns generated.FAQResolver implementation.
func (r *Resolver) FAQ() generated.FAQResolver { return &fAQResolver{r} }

// FAQCategory returns generated.FAQCategoryResolver implementation.
func (r *Resolver) FAQCategory() generated.FAQCategoryResolver { return &fAQCategoryResolver{r} }

// Question returns generated.QuestionResolver implementation.
func (r *Resolver) Question() generated.QuestionResolver { return &questionResolver{r} }

type fAQResolver struct{ *Resolver }
type fAQCategoryResolver struct{ *Resolver }
type questionResolver struct{ *Resolver }
