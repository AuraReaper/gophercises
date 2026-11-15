package internal

import (
	"net/http"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) RootHandler(w http.ResponseWriter, r *http.Request) {

	intro := h.svc.Data["intro"]
	err := h.svc.Template.Execute(w, intro)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) LoopHandler(w http.ResponseWriter, r *http.Request) {
	arcKey := r.PathValue("arc")
	storyArc := h.svc.Data[arcKey]

	err := h.svc.Template.Execute(w, storyArc)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
