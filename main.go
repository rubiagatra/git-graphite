package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	// Login form route
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
            <html>
                <body>
                    <form action="/login" method="post">
                        <input type="text" name="username" placeholder="Username" required>
                        <input type="password" name="password" placeholder="Password" required>
                        <input type="submit" value="Login">
                    </form>
                </body>
            </html>
        `))
	})

	http.ListenAndServe(":3000", r)
}
