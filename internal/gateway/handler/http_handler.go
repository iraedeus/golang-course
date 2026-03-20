package handler

import (
	"encoding/json"
	"net/http"

	"golang-course/internal/gateway/client"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HttpHandler struct {
	collectorClient *client.CollectorClient
}

func NewHttpHandler(c *client.CollectorClient) *HttpHandler {
	return &HttpHandler{
		collectorClient: c,
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
func (h *HttpHandler) GetRepository(w http.ResponseWriter, r *http.Request) {
	owner := r.URL.Query().Get("owner")
	repoName := r.URL.Query().Get("repo")

	if owner == "" || repoName == "" {
		http.Error(w, "parameters 'owner' and 'repo' are required", http.StatusBadRequest)
		return
	}

	resp, err := h.collectorClient.GetRepo(owner, repoName)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			http.Error(w, "Repository not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode((resp))
}
