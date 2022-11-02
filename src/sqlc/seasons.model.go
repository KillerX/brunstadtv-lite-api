package sqlc

import (
	"github.com/bcc-code/brunstadtv-lite-api/src/model"
	"strconv"
)

func (s Season) AsGQL(language string) *model.Season {
	legacyID := ""
	if s.LegacyID.Valid {
		legacyID = strconv.Itoa(int(s.LegacyID.Int64))
	}

	title := model.TranslatedString{}
	_ = title.Scan(s.Title)

	description := model.TranslatedString{}
	_ = description.Scan(s.Description)

	return &model.Season{
		ID:          string(s.ID),
		LegacyID:    &legacyID,
		AgeRating:   s.AgeRating,
		Title:       title.Get(language),
		Description: description.Get(language),
		Image:       &s.Image.String,
		ImageURL:    nil,
		Images:      nil,
		Number:      int(s.Number),
	}

}
