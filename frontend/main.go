package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Food struct {
	ID       int
	Name     string
	Calories int
	Proteine int
}

type ToDo struct {
	ID          int
	TaskName    string
	Deadline    string
	Urgency     string
	Description string
}

type Weight struct {
	ID   int
	Name string
	Reps int
	Sets int
}

type Cardio struct {
	ID       int
	Name     string
	Distance int
	Duration int
}

func main() {
	// Connect to PostgreSQL database
	db, err := sql.Open("postgres", "postgresql://postgres:123@localhost/hestia?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Load the HTML template for the navigation
	navTmpl, err := template.ParseFiles("templates/nav.html")
	if err != nil {
		log.Fatal("Error loading navigation template:", err)
	}

	// Serve the HTML content for the navigation
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := navTmpl.Execute(w, nil); err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	// Load the HTML template for the food table
	foodTmpl, err := template.ParseFiles("templates/food.html")
	if err != nil {
		log.Fatal("Error loading food template:", err)
	}

	// Serve the HTML content for the Food table
	http.HandleFunc("/food", func(w http.ResponseWriter, r *http.Request) {

		// Query the database for existing food entries
		rows, err := db.Query("SELECT * FROM Food")
		if err != nil {
			log.Fatal("Error querying the database:", err)
		}

		defer rows.Close()

		// Retrieve data from rows
		var foods []Food
		for rows.Next() {
			var food Food
			if err := rows.Scan(&food.ID, &food.Name, &food.Calories, &food.Proteine); err != nil {
				log.Fatal("Error scanning row:", err)
			}
			foods = append(foods, food)
		}

		// Execute the HTML template to display the Food table
		if err := foodTmpl.Execute(w, foods); err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

	})

	// Handle form submission to add a new food entry
	http.HandleFunc("/addfood", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		name := r.FormValue("name")
		calories := r.FormValue("calories")
		proteine := r.FormValue("proteine")

		// Insert the new food entry into the database
		_, err := db.Exec("INSERT INTO Food (Name, Calories, Proteine) VALUES ($1, $2, $3)", name, calories, proteine)
		if err != nil {
			log.Println("Error inserting new food entry:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Redirect back to the Food table page
		http.Redirect(w, r, "/food", http.StatusSeeOther)
	})

	// Load the HTML template for the food table
	weightTmpl, err := template.ParseFiles("templates/weight.html")
	if err != nil {
		log.Fatal("Error loading food template:", err)
	}

	// Serve the HTML content for the Weight table
	http.HandleFunc("/weight", func(w http.ResponseWriter, r *http.Request) {
		// Query the database for existing food entries
		rows, err := db.Query("SELECT * FROM Weight")
		if err != nil {
			log.Fatal("Error querying the database:", err)
		}
		defer rows.Close()

		// Retrieve data from rows
		var weights []Weight
		for rows.Next() {
			var weight Weight
			if err := rows.Scan(&weight.ID, &weight.Name, &weight.Reps, &weight.Sets); err != nil {
				log.Fatal("Error scanning row:", err)
			}
			weights = append(weights, weight)
		}

		// Execute the HTML template to display the weight table
		if err := weightTmpl.Execute(w, weights); err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	// Handle form submission to add a new weight entry
	http.HandleFunc("/addweight", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		name := r.FormValue("name")
		Reps := r.FormValue("Reps")
		Sets := r.FormValue("Sets")

		// Insert the new weight entry into the database
		_, err := db.Exec("INSERT INTO Weight (Name, Reps, Sets) VALUES ($1, $2, $3)", name, Reps, Sets)
		if err != nil {
			log.Println("Error inserting new food entry:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Redirect back to the weight table page
		http.Redirect(w, r, "/weight", http.StatusSeeOther)
	})

	// Load the HTML template for the cardio table
	cardioTmpl, err := template.ParseFiles("templates/cardio.html")
	if err != nil {
		log.Fatal("Error loading navigation template:", err)
	}

	// Serve the HTML content for the Cardio table
	http.HandleFunc("/cardio", func(w http.ResponseWriter, r *http.Request) {
		// Query the database for existing cardio entries
		rows, err := db.Query("SELECT * FROM Cardio")
		if err != nil {
			log.Fatal("Error querying the database:", err)
		}
		defer rows.Close()

		// Retrieve data from rows
		var cardios []Cardio
		for rows.Next() {
			var cardio Cardio
			if err := rows.Scan(&cardio.ID, &cardio.Name, &cardio.Distance, &cardio.Duration); err != nil {
				log.Fatal("Error scanning row:", err)
			}
			cardios = append(cardios, cardio)
		}

		// Execute the HTML template to display the Cardio table
		if err := cardioTmpl.Execute(w, cardios); err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	// Handle form submission to add a new cardio entry
	http.HandleFunc("/addcardio", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		name := r.FormValue("name")
		Distance := r.FormValue("Distance")
		Duration := r.FormValue("Duration")

		// Insert the new cardio entry into the database
		_, err := db.Exec("INSERT INTO Cardio (Name, Distance, Duration) VALUES ($1, $2, $3)", name, Distance, Duration)
		if err != nil {
			log.Println("Error inserting new food entry:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Redirect back to the cardio table page
		http.Redirect(w, r, "/cardio", http.StatusSeeOther)
	})

	// Load the HTML template for the To-Do list
	todoTmpl, err := template.ParseFiles("templates/todo.html")
	if err != nil {
		log.Fatal("Error loading navigation template:", err)
	}

	// Serve the HTML content for the To-Do list
	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		// Query the database for existing tasks
		rows, err := db.Query("SELECT * FROM ToDo")
		if err != nil {
			log.Fatal("Error querying the database:", err)
		}
		defer rows.Close()

		// Retrieve data from rows
		var todos []ToDo
		for rows.Next() {
			var todo ToDo
			if err := rows.Scan(&todo.ID, &todo.TaskName, &todo.Deadline, &todo.Urgency, &todo.Description); err != nil {
				log.Fatal("Error scanning row:", err)
			}
			todos = append(todos, todo)
		}

		// Execute the HTML template to display the To-Do list
		if err := todoTmpl.Execute(w, todos); err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	// Handle form submission to add a new To-Do task
	http.HandleFunc("/addtodo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		taskName := r.FormValue("taskName")
		deadline := r.FormValue("deadline")
		urgency := r.FormValue("urgency")
		description := r.FormValue("description")

		// Insert the new task into the database
		_, err := db.Exec("INSERT INTO ToDo (TaskName, Deadline, Urgency, Description) VALUES ($1, $2, $3, $4)", taskName, deadline, urgency, description)
		if err != nil {
			log.Println("Error inserting new task:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Redirect back to the To-Do list page
		http.Redirect(w, r, "/todo", http.StatusSeeOther)
	})

	// Load the HTML template for the day page
	dayTmpl, err := template.ParseFiles("templates/day.html")
	if err != nil {
		log.Fatal("Error loading day template:", err)
	}

	// Serve the HTML content for the Day page
	http.HandleFunc("/day", func(w http.ResponseWriter, r *http.Request) {
		// Query the database for existing food entries
		rows, err := db.Query("SELECT * FROM Food")
		if err != nil {
			log.Fatal("Error querying the database:", err)
		}
		defer rows.Close()

		// Retrieve data from rows
		var foods []Food
		for rows.Next() {
			var food Food
			if err := rows.Scan(&food.ID, &food.Name, &food.Calories, &food.Proteine); err != nil {
				log.Fatal("Error scanning row:", err)
			}
			foods = append(foods, food)
		}

		// Query the database for existing weight entries
		rows, err = db.Query("SELECT * FROM Weight")
		if err != nil {
			log.Fatal("Error querying the database:", err)
		}
		defer rows.Close()

		// Retrieve data from rows
		var weights []Weight
		for rows.Next() {
			var weight Weight
			if err := rows.Scan(&weight.ID, &weight.Name, &weight.Reps, &weight.Sets); err != nil {
				log.Fatal("Error scanning row:", err)
			}
			weights = append(weights, weight)
		}

		// Query the database for existing cardio entries
		rows, err = db.Query("SELECT * FROM Cardio")
		if err != nil {
			log.Fatal("Error querying the database:", err)
		}
		defer rows.Close()

		// Retrieve data from rows
		var cardios []Cardio
		for rows.Next() {
			var cardio Cardio
			if err := rows.Scan(&cardio.ID, &cardio.Name, &cardio.Distance, &cardio.Duration); err != nil {
				log.Fatal("Error scanning row:", err)
			}
			cardios = append(cardios, cardio)
		}

		// Execute the HTML template to display the Day page
		if err := dayTmpl.Execute(w, struct {
			Foods   []Food
			Weights []Weight
			Cardios []Cardio
		}{
			Foods:   foods,
			Weights: weights,
			Cardios: cardios,
		}); err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	// Handle form submission to add a new food entry for the day
	http.HandleFunc("/adddayfood", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		foodID := r.FormValue("food")

		// Insert the new food entry for the day
		_, err := db.Exec("INSERT INTO DayFood (Food_ID) VALUES ($1)", foodID)
		if err != nil {
			log.Println("Error inserting new day food entry:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Redirect back to the Day page
		http.Redirect(w, r, "/day", http.StatusSeeOther)
	})

	// Handle form submission to add a new weight entry for the day
	http.HandleFunc("/adddayweight", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		weightID := r.FormValue("weight")

		// Insert the new weight entry for the day
		_, err := db.Exec("INSERT INTO DayWeight (Weight_ID) VALUES ($1)", weightID)
		if err != nil {
			log.Println("Error inserting new day weight entry:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Redirect back to the Day page
		http.Redirect(w, r, "/day", http.StatusSeeOther)
	})

	// Handle formsubmission to add a new cardio entry for the day
	http.HandleFunc("/adddaycardio", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		cardioID := r.FormValue("cardio")

		// Insert the new cardio entry for the day
		_, err := db.Exec("INSERT INTO DayCardio (Cardio_ID) VALUES ($1)", cardioID)
		if err != nil {
			log.Println("Error inserting new day cardio entry:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Redirect back to the Day page
		http.Redirect(w, r, "/day", http.StatusSeeOther)
	})

	// Start the HTTP server
	fmt.Println("Server listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
