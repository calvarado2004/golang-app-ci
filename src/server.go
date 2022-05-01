package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	_ "github.com/lib/pq"
)

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var todos []string

	//Idempotent way to create a Schema and a Table
	_, err := db.Exec("CREATE SCHEMA IF NOT EXISTS todos AUTHORIZATION postgres; CREATE TABLE IF NOT EXISTS todos.todos (item text);")
	if err != nil {
		log.Fatalf("An error occured while trying to create schema and table")
	}

	rows, err := db.Query("SELECT * FROM todos.todos")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}

	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}
	return c.Render("index", fiber.Map{
		"Todos": todos,
	})
}

type todo struct {
	Item string
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	newTodo := todo{}
	if err := c.BodyParser(&newTodo); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}
	fmt.Printf("%v", newTodo)
	if newTodo.Item != "" {
		_, err := db.Exec("INSERT into todos.todos VALUES ($1)", newTodo.Item)
		if err != nil {
			log.Fatalf("An error occured while executing query: %v", err)
		}
	}

	return c.Redirect("/app-golang")
}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
	olditem := c.Query("olditem")
	newitem := c.Query("newitem")
	db.Exec("UPDATE todos.todos SET item=$1 WHERE item=$2", newitem, olditem)
	return c.Redirect("/app-golang")
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	todoToDelete := c.Query("item")
	db.Exec("DELETE from todos.todos WHERE item=$1", todoToDelete)
	return c.SendString("deleted")
}

func main() {

	DB_SERVER := os.Getenv("DB_SERVER")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable", DB_USER, DB_PASSWORD, DB_SERVER, DB_PORT)

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

        app.Static(
                "/app-golang/static/",
                "./static")

	app.Get("/app-golang", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	app.Post("/app-golang", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})

	app.Put("/app-golang/update", func(c *fiber.Ctx) error {
		return putHandler(c, db)
	})

	app.Delete("/app-golang/delete", func(c *fiber.Ctx) error {
		return deleteHandler(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
