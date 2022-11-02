package sqlc

import (
	"github.com/bcc-code/brunstadtv-lite-api/src/model"
	"strconv"
)

func (show *Show) AsGQL(language string) *model.Show {
	legacyID := ""
	if show.LegacyID.Valid {
		legacyID = strconv.Itoa(int(show.LegacyID.Int64))
	}

	title := model.TranslatedString{}
	_ = title.Scan(show.Title)

	description := model.TranslatedString{}
	_ = description.Scan(show.Description)

	return &model.Show{
		ID:           strconv.Itoa(int(show.ID)),
		LegacyID:     &legacyID,
		Title:        title.Get(language),
		Description:  description.Get(language),
		Image:        &show.Image.String,
		ImageURL:     &show.Image.String, // TODO One of this is wrong
		Images:       nil,
		EpisodeCount: 0,
		SeasonCount:  0,
	}
}
