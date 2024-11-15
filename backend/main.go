package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Project struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl"`
	Link        string `json:"link"`
}

type Skill struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Level    int    `json:"level"`
}

type About struct {
	Name        string `json:"name"`
	Role        string `json:"role"`
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl"`
}

type Contact struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func main() {
	r := mux.NewRouter()

	// Initialize routes
	initializeRoutes(r)

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(r)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func initializeRoutes(r *mux.Router) {
	r.HandleFunc("/api/projects", getProjects).Methods("GET")
	r.HandleFunc("/api/skills", getSkills).Methods("GET")
	r.HandleFunc("/api/about", getAbout).Methods("GET")
	r.HandleFunc("/api/contact", submitContact).Methods("POST")
}

func getProjects(w http.ResponseWriter, r *http.Request) {
	projects := []Project{
		{
			ID:          "1",
			Title:       "Groupie-Tracker Project",
			Description: "A full-stack web application",
			ImageURL:    "/images/project1.jpg",
			Link:        "https://github.com/stkisengese/groupie-tracker",
		},
		// Add more projects
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func getSkills(w http.ResponseWriter, r *http.Request) {
	skills := []Skill{
		{
			ID:       "1",
			Name:     "React",
			Category: "Frontend",
			Level:    70,
		},
		// Add more skills
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(skills)
}

func getAbout(w http.ResponseWriter, r *http.Request) {
	about := About{
		Name:        "Your Name",
		Role:        "Full Stack Developer",
		Description: "Passionate developer with experience in Go and React",
		ImageURL:    "/images/profile.jpg",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(about)
}

func submitContact(w http.ResponseWriter, r *http.Request) {
	var contact Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Here you would typically send an email or store the contact info
	// For now, we'll just return a success message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Thank you for your message. I'll get back to you soon!",
	})
}

// package main

// import (
//     "log"
//     "net/http"
//     "github.com/gorilla/mux"
//     "github.com/rs/cors"
// )

// func main() {
//     r := mux.NewRouter()

//     // Initialize routes
//     initializeRoutes(r)

//     // Setup CORS
//     c := cors.New(cors.Options{
//         AllowedOrigins: []string{"http://localhost:3000"},
//         AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
//         AllowedHeaders: []string{"Content-Type", "Authorization"},
//     })

//     handler := c.Handler(r)

//     log.Println("Server starting on port 8080...")
//     log.Fatal(http.ListenAndServe(":8080", handler))
// }

// func initializeRoutes(r *mux.Router) {
//     r.HandleFunc("/api/projects", getProjects).Methods("GET")
//     r.HandleFunc("/api/skills", getSkills).Methods("GET")
//     r.HandleFunc("/api/about", getAbout).Methods("GET")
//     r.HandleFunc("/api/contact", submitContact).Methods("POST")
// }
