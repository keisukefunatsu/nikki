package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeAPIResponse returns
type HomeAPIResponse struct {
	baseAPIResponse
}

// Homepage root page
func Homepage(c buffalo.Context) error {
	res := HomeAPIResponse{}
	res.Message = "Nothing here"
	return c.Render(http.StatusOK, r.JSON(res))
}
