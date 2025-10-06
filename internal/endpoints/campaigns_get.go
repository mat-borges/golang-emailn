package endpoints

import (
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

func (h Handler) CampaignGet(w http.ResponseWriter, r *http.Request) {
	campaigns, err := h.CampaignService.Repo.GetAll()
	if err != nil {
		if errors.Is(err, internalerrors.ErrInternal) {
			render.Status(r, http.StatusInternalServerError)
		} else {
			render.Status(r, http.StatusBadRequest)
		}
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, campaigns)
}