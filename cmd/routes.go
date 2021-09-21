package main

import (
	"github.com/RamiroCuenca/go-jwt-notes/notes/controllers"
	usersControllers "github.com/RamiroCuenca/go-jwt-notes/users/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes() *chi.Mux {
	// Create a new multiplexer
	r := chi.NewMux()

	// Use logger middleware from chi
	r.Use(middleware.Logger)

	// Notes Handlers
	r.Post("/api/v1/notes/create", AuthenticationMiddleware(controllers.NotesCreate))
	r.Get("/api/v1/notes/readbyid", controllers.NotesGetById)
	r.Get("/api/v1/notes/readall", controllers.NotesGetAll)
	r.Put("/api/v1/notes/update", controllers.NotesUpdateById)
	r.Delete("/api/v1/notes/delete", controllers.NotesDeleteById)

	// Users Handlers
	r.Post("/api/v1/users/create", usersControllers.UsersLogin)

	return r
}
