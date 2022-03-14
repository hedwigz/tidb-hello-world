package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hedwigz/ent-tidb/ent"
)

func main() {
	client, err := ent.Open("mysql", "root@tcp(localhost:4000)/test?parseTime=true", ent.Debug())
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool, with Atlas.
	if err := client.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}

	client.User.Create().
		SetAge(30).
		SetName("hedwigz").
		SaveX(context.Background())
	user := client.User.Query().FirstX(context.Background())
	fmt.Printf("the user: %s is %d years old\n", user.Name, user.Age)
}
