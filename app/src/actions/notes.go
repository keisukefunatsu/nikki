package actions

import (
	"net/http"
	"nikki/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/pkg/errors"
)

// NoteAPIResponse returns notes
type NoteAPIResponse struct {
	baseAPIResponse
	Note  *models.Note  `json:"note,omitempty"`
	Notes *models.Notes `json:"notes,omitempty"`
}

// NoteList returns notes model
func NoteList(c buffalo.Context) error {
	res := NoteAPIResponse{}
	notes := &models.Notes{}
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	q := tx.PaginateFromParams(c.Params())
	if err := q.All(notes); err != nil {
		// TODO Implement here
	}
	res.Notes = notes
	c.Response().Header().Set("X-Pagination", q.Paginator.String())

	return c.Render(http.StatusOK, r.JSON(res))
}

// NoteShow returns note model
func NoteShow(c buffalo.Context) error {
	res := NoteAPIResponse{}
	note := &models.Note{}
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	if err := tx.Find(note, c.Param("note_id")); err != nil {
		// TODO Implement here
	}
	res.Note = note

	return c.Render(http.StatusOK, r.JSON(res))
}

// NoteCreate creates note record
func NoteCreate(c buffalo.Context) error {
	res := NoteAPIResponse{}
	note := &models.Note{}
	if err := c.Bind(note); err != nil {
		return errors.WithStack(err)
	}
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}
	verrs, err := tx.ValidateAndCreate(note)

	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
	}
	res.Note = note

	return c.Render(http.StatusCreated, r.JSON(res))
}
