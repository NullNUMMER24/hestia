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

func main() {
	// Connect to PostgreSQL database
	db, err := sql.Open("postgres", "postgresql://postgres:123@localhost/hestia?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

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
				form {
					margin-top: 20px;
				}
				input[type="text"], input[type="submit"] {
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
            <a href="/food">Food</a> | <a href="/todo">To-Do</a> | <a href="/weight">Weight</a>
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
				form {
					margin-top: 20px;
				}
				input[type="text"], input[type="submit"] {
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
			<a href="/nav">Navigation</a> | <a href="/todo">To-Do</a> | <a href="/weight">Weight</a>
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
            <h2>Add New Food Entry</h2>
            <form action="/addfood" method="post">
                <label for="name">Name:</label>
                <input type="text" id="name" name="name">
                <label for="calories">Calories:</label>
                <input type="text" id="calories" name="calories">
                <label for="proteine">Proteine:</label>
                <input type="text" id="proteine" name="proteine">
                <input type="submit" value="Add Food">
            </form>
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
				form {
					margin-top: 20px;
				}
				input[type="text"], input[type="submit"] {
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
			<a href="/nav">Navigation</a> | <a href="/food">Food</a> | <a href="/todo">To-Do</a>
            <h1>Food Table</h1>
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
            <h2>Add New Weight Entry</h2>
            <form action="/addweight" method="post">
                <label for="name">Name:</label>
                <input type="text" id="name" name="name">
                <label for="Reps">Reps:</label>
                <input type="text" id="Reps" name="Reps">
                <label for="Sets">Sets:</label>
                <input type="text" id="Sets" name="Sets">
                <input type="submit" value="Add Weight">
            </form>
        </body>
        </html>
    `))

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

		// Execute the HTML template to display the Food table
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

	// Define the HTML template for the To-Do list
	var todoTmpl = template.Must(template.New("todo").Parse(`
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>HESTIA - To-Do List</title>
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
				form {
					margin-top: 20px;
				}
				input[type="text"], input[type="submit"] {
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
			<a href="/nav">Navigation</a> | <a href="/food">Food</a> | <a href="/weight">Weight</a>
            <h1>To-Do List</h1>
            <ul>
                {{range .}}
                <li>{{.TaskName}} - Deadline: {{.Deadline}}, Urgency: {{.Urgency}}, Description: {{.Description}}</li>
                {{end}}
            </ul>
            <h2>Add New Task</h2>
            <form action="/addtodo" method="post">
                <label for="taskName">Task Name:</label>
                <input type="text" id="taskName" name="taskName"><br>
                <label for="deadline">Deadline:</label>
                <input type="date" id="deadline" name="deadline"><br>
                <label for="urgency">Urgency:</label>
                <input type="text" id="urgency" name="urgency"><br>
                <label for="description">Description:</label><br>
                <textarea id="description" name="description"></textarea><br>
                <input type="submit" value="Add Task">
            </form>
        </body>
        </html>
    `))

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

	// Start the HTTP server
	fmt.Println("Server listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
