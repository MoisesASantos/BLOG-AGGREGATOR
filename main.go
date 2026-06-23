package main

import (
	"fmt"

	"github.com/MoisesASantos/BLOG-AGGREGATOR/internal/config"
)

func main() {

	cfg := config.Read()
	cfg.SetUser("mosantos")
	cfg = config.Read()
	fmt.Printf(
		"user_logged: %s\ndb_credentials: %s\n",
		cfg.Current_user_name,
		cfg.Db_url,
	)
}
