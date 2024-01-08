package main

import (
	context "context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"log"

	"github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func applyMigrations(mongoURI, migrationPath string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}
	if err := client.Connect(context.Background()); err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	driver, err := mongodb.WithInstance(client, &mongodb.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationPath,
		"mongodb", driver,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	log.Println("Migrations applied successfully")
	return nil
}

func main() {

	err := applyMigrations("mongodb+srv://user:qwerty1234@cluster0.wqpyttn.mongodb.net/?retryWrites=true&w=majority", "migrations")
	if err != nil {
		log.Fatal(err)
	}

	clientOptions := options.Client().ApplyURI("mongodb+srv://user:qwerty1234@cluster0.wqpyttn.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	user := User{Name: "John Doe", Email: "john@example.com", Age: 25}
	err = createUser(client, user)
	if err != nil {
		log.Fatal(err)
	}

	createdUser, err := getUserByID(client, user.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created User:", createdUser)

	updatedUser, err := updateUser(client, createdUser.ID, "Doe John")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated User:", updatedUser)

	users, err := getAllUsers(client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All Users:", users)

	err = deleteUser(client, updatedUser.ID)
	if err != nil {
		log.Fatal(err)
	}
}
