package delivery

import (
	"encoding/json"
	"errors"
	"net/http"

	"golang-course/gateway/internal/domain"
)

type GatewayUseCase interface {
	Execute(owner, repo string) (domain.Repo, error)
}

type HTTPController struct {
	useCase GatewayUseCase
}

func NewHTTPController(uc GatewayUseCase) *HTTPController {
	return &HTTPController{
		useCase: uc,
	}
}

// GetRepository godoc
// @Summary      Get GitHub repository info
// @Description  Get stars, forks and description of a repo
// @Tags         repository
// @Produce      json
// @Param        owner   query      string  true  "Repository Owner"
// @Param        repo    query      string  true  "Repository Name"
// @Success      200     {object}   proto.RepositoryResponse
// @Router       /repo [get]
func (h *HTTPController) GetRepository(w http.ResponseWriter, r *http.Request) {
	owner := r.URL.Query().Get("owner")
	repoName := r.URL.Query().Get("repo")

	if owner == "" || repoName == "" {
		http.Error(w, "parameters 'owner' and 'repo' are required", http.StatusBadRequest)
		return
	}

	resp, err := h.useCase.Execute(owner, repoName)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			http.Error(w, "Repository not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode((resp))
}
