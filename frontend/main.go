package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type Food struct {
	ID       int
	Name     string
	Calories int
	Proteine int
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

type PageData struct {
	Foods   []Food
	Weights []Weight
	Cardios []Cardio
}

const (
	DBHost     = "localhost"
	DBPort     = 5432
	DBUser     = "postgres"
	DBPassword = "123"
	DBName     = "hestia"
)

Gvar db *sql.DB

func main() {
	// Connect to PostgreSQL database
	db, err := sql.Open("postgres", "postgresql://postgres:123@localhost/hestia?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	http.HandleFunc("/add-food", addFoodHandler)
	http.HandleFunc("/add-exercise", addExerciseHandler)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


	// Define the HTML template for the Day page with a form to add meals and exercises using dropdown menus
	var dayTmpl = template.Must(template.New("day").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>HESTIA - Day View</title>
	<style>
		body {
			background-color: #222;
			color: #fff;
		}
		form {
			margin-top: 20px;
		}
		select, input[type="submit"] {
			padding: 5px;
			font-size: 16px;
			border: none;
			border-radius: 4px;
			background-color: #444;
			color: #fff;
			margin-right: 10px;
		}
	</style>
</head>
<body>
	<a href="/">Navigation</a> | <a href="/food">Food</a> | <a href="/weight">Weight</a> | <a href="/cardio">Cardio</a>
	<h1>Day View</h1>
	<h2>Today's Activities</h2>
	<ul>
		{{range .Activities}}
		<li>{{.}}</li>
		{{end}}
	</ul>
	<h2>Add Meals and Exercises</h2>
	<form action="/adddayactivity" method="post">
		<label for="meal">Meal:</label>
		<select id="meal" name="meal">
			<option value="">Select a meal...</option>
			{{range .Foods}}
			<option value="{{.}}">{{.}}</option>
			{{end}}
		</select>
		<label for="exercise">Exercise:</label>
		<select id="exercise" name="exercise">
			<option value="">Select an exercise...</option>
			{{range .Exercises}}
			<option value="{{.}}">{{.}}</option>
			{{end}}
		</select>
		<input type="submit" value="Add Activity">
	</form>
</body>
</html>
`))

	// Serve the HTML content for the Day page
	http.HandleFunc("/day", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve today's date
		today := time.Now().Format("2006-01-02") // Format: YYYY-MM-DD

		// Retrieve today's activities from the database
		rows, err := db.Query("SELECT meal, exercise FROM day WHERE date = $1", today)
		if err != nil {
			log.Println("Error querying the database:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var activities []string
		for rows.Next() {
			var meal, exercise string
			if err := rows.Scan(&meal, &exercise); err != nil {
				log.Println("Error scanning row:", err)
				continue
			}
			activities = append(activities, fmt.Sprintf("Meal: %s, Exercise: %s", meal, exercise))
		}
		if err := rows.Err(); err != nil {
			log.Println("Error iterating rows:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Retrieve today's meals and exercises from the database
		foods, err := GetFoods(db)
		if err != nil {
			log.Println("Error retrieving foods:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		exercises, err := GetExercises(db)
		if err != nil {
			log.Println("Error retrieving exercises:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Execute the HTML template to display the Day view
		if err := dayTmpl.Execute(w, struct {
			Activities []string
			Foods      []string
			Exercises  []string
		}{
			Activities: activities,
			Foods:      foods,
			Exercises:  exercises,
		}); err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	// Handle form submission to add a new meal or exercise for the day
	http.HandleFunc("/adddayactivity", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		meal := r.FormValue("meal")
		exercise := r.FormValue("exercise")

		// Retrieve today's date
		today := time.Now().Format("2006-01-02") // Format: YYYY-MM-DD

		// Insert the new meal and exercise into the database
		_, err := db.Exec("INSERT INTO day (date, meal, exercise) VALUES ($1, $2, $3)", today, meal, exercise)
		if err != nil {
			log.Println("Error inserting new activity:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Redirect back to the Day page
		http.Redirect(w, r, "/day", http.StatusSeeOther)
	})

	// Define the HTML template for the navigation
	var navTmpl = template.Must(template.New("nav").Parse(`
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>HESTIA - Navigation</title>
			<style>
				body {
					background-color: #222;
					color: #fff;
				}
				a {
					color: #fff;
					text-decoration: none;
				}
			</style>
        </head>
        <body>
            <a href="/food">Food</a> | <a href="/weight">Weight</a> | <a href="/cardio">Cardio</a> | <a href="/day">Day</a>
        </body>
        </html>
    `))

	// Serve the HTML content for the navigation
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := navTmpl.Execute(w, nil); err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	// Define the HTML template for the Food table
	var foodTmpl = template.Must(template.New("food").Parse(`
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>HESTIA - Food Table</title>
			<style>
				body {
					background-color: #222;
					color: #fff;
				}
				table {
					border-collapse: collapse;
					width: 100%;
				}
				th, td {
					padding: 8px;
					text-align: left;
					border-bottom: 1px solid #444;
				}
				th {
					background-color: #333;
				}
				tr:nth-child(even) {
					background-color: #333;
				}
			</style>
        </head>
        <body>
			<a href="/">Navigation</a> | <a href="/weight">Weight</a> | <a href="/cardio">Cardio</a> | <a href="/day">Day</a>
            <h1>Food Table</h1>
            <table>
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Calories</th>
                    <th>Proteine</th>
                </tr>
                {{range .}}
                <tr>
                    <td>{{.ID}}</td>
                    <td>{{.Name}}</td>
                    <td>{{.Calories}}</td>
                    <td>{{.Proteine}}</td>
                </tr>
                {{end}}
            </table>
        </body>
        </html>
    `))

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

	// Define the HTML template for the Weight table
	var weightTmpl = template.Must(template.New("weight").Parse(`
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>HESTIA - Weight Table</title>
			<style>
				body {
					background-color: #222;
					color: #fff;
				}
				table {
					border-collapse: collapse;
					width: 100%;
				}
				th, td {
					padding: 8px;
					text-align: left;
					border-bottom: 1px solid #444;
				}
				th {
					background-color: #333;
				}
				tr:nth-child(even) {
					background-color: #333;
				}
			</style>
        </head>
        <body>
			<a href="/">Navigation</a> | <a href="/food">Food</a> | <a href="/cardio">Cardio</a> | <a href="/day">Day</a>
            <h1>Weight Table</h1>
            <table>
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Reps</th>
                    <th>Sets</th>
                </tr>
                {{range .}}
                <tr>
                    <td>{{.ID}}</td>
                    <td>{{.Name}}</td>
                    <td>{{.Reps}}</td>
                    <td>{{.Sets}}</td>
                </tr>
                {{end}}
            </table>
        </body>
        </html>
    `))

	// Serve the HTML content for the Weight table
	http.HandleFunc("/weight", func(w http.ResponseWriter, r *http.Request) {
		// Query the database for existing weight entries
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

		// Execute the HTML template to display the Weight table
		if err := weightTmpl.Execute(w, weights); err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	// Define the HTML template for the Cardio table
	var cardioTmpl = template.Must(template.New("cardio").Parse(`
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>HESTIA - Cardio Table</title>
			<style>
				body {
					background-color: #222;
					color: #fff;
				}
				table {
					border-collapse: collapse;
					width: 100%;
				}
				th, td {
					padding: 8px;
					text-align: left;
					border-bottom: 1px solid #444;
				}
				th {
					background-color: #333;
				}
				tr:nth-child(even) {
					background-color: #333;
				}
			</style>
        </head>
        <body>
			<a href="/">Navigation</a> | <a href="/food">Food</a> | <a href="/weight">Weight</a> | <a href="/day">Day</a>
            <h1>Cardio Table</h1>
            <table>
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Distance</th>
                    <th>Duration</th>
                </tr>
                {{range .}}
                <tr>
                    <td>{{.ID}}</td>
                    <td>{{.Name}}</td>
                    <td>{{.Distance}}</td>
                    <td>{{.Duration}}</td>
                </tr>
                {{end}}
            </table>
        </body>
        </html>
    `))

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

	// Start the HTTP server
	fmt.Println("Server listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

// Function to fetch data from the Food table
func getFoodData(db *sql.DB) []Food {
	rows, err := db.Query("SELECT * FROM Food")
	if err != nil {
		log.Fatal("Error querying the database:", err)
	}
	defer rows.Close()

	var foods []Food
	for rows.Next() {
		var food Food
		if err := rows.Scan(&food.ID, &food.Name, &food.Calories, &food.Proteine); err != nil {
			log.Fatal("Error scanning row:", err)
		}
		foods = append(foods, food)
	}

	return foods
}

func addFoodHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	foodName := r.FormValue("food_name")
	// Retrieve other form values

	// Insert data into the Food table
	_, err = db.Exec("INSERT INTO Food (Name, Calories, Proteine) VALUES ($1, $2, $3)",
		foodName, /* Add Calories and Proteine values here */)
	if err != nil {
		http.Error(w, "Error inserting data into Food table", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Food added successfully!")
}

func addExerciseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	exerciseType := r.FormValue("exercise_type") // Should be "cardio" or "weightlifting"
	exerciseName := r.FormValue("exercise_name")
	distance := r.FormValue("distance")
	duration := r.FormValue("duration")
	reps := r.FormValue("reps")
	sets := r.FormValue("sets")

	// Validate form data
	if exerciseType != "cardio" && exerciseType != "weightlifting" {
		http.Error(w, "Invalid exercise type", http.StatusBadRequest)
		return
	}

	// Insert data into the appropriate table based on exercise type
	var result sql.Result
	var insertErr error
	switch exerciseType {
	case "cardio":
		result, insertErr = db.Exec("INSERT INTO Cardio (Name, Distance, Duration) VALUES ($1, $2, $3)",
			exerciseName, distance, duration)
	case "weightlifting":
		result, insertErr = db.Exec("INSERT INTO Weight (Name, Reps, Sets) VALUES ($1, $2, $3)",
			exerciseName, reps, sets)
	}

	if insertErr != nil {
		http.Error(w, "Error inserting data into database", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Error getting rows affected", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "No rows affected", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Exercise added successfully!")
}