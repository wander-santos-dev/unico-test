package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/teste-unico/db"
	"github.com/teste-unico/models"
)

var feiraIdKey = "feiraId"

func feiras(r chi.Router) {
	r.Get("/", getAllFeiras)
	r.Post("/", createFeira)
	r.Route("/{feiraId}", func(r chi.Router) {
		r.Use(FeiraContext)
		r.Get("/", getFeira)
		r.Put("/", updateFeira)
		r.Delete("/", deleteFeira)
	})
}

func FeiraContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		feiraId := chi.URLParam(r, "feiraId")
		if feiraId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("Feira ID is required")))
			return
		}
		id, err := strconv.Atoi(feiraId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("Invalid feira ID")))
		}
		ctx := context.WithValue(r.Context(), feiraIdKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func createFeira(w http.ResponseWriter, r *http.Request) {
	feira := &models.Feira{}

	if err := render.Bind(r, feira); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	if err := dbInstance.AddFeira(*feira); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, feira); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func getAllFeiras(w http.ResponseWriter, r *http.Request) {
	feiras, err := dbInstance.GetAllFeiras()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, feiras); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func getFeira(w http.ResponseWriter, r *http.Request) {
	FeiraID := r.Context().Value(feiraIdKey).(int)
	feira, err := dbInstance.GetFeiraById(FeiraID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &feira); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteFeira(w http.ResponseWriter, r *http.Request) {
	feiraId := r.Context().Value(feiraIdKey).(int)
	err := dbInstance.DeleteFeira(feiraId)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}

func updateFeira(w http.ResponseWriter, r *http.Request) {
	feiraId := r.Context().Value(feiraIdKey).(int)
	feiraData := models.Feira{}
	if err := render.Bind(r, &feiraData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	feira, err := dbInstance.UpdateFeira(feiraId, feiraData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &feira); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
