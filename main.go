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
	client, err := ent.Open("mysql", "root@tcp(localhost:4000)/test?parseTime=true")
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
		SetName("hedwigz").SaveX(context.Background())

	users := client.User.Query().AllX(context.Background())

	for i := 0; i < len(users); i++ {
		u := users[i]
		fmt.Printf("the user: %s is %d years old\n", u.Name, u.Age)
	}
}
