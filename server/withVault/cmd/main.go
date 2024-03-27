package main

import (
	"fmt"
	"withvault/internal/env"
	"withvault/internal/env/vault"
)

func main() {
	vault, err := vault.NewVaultProvider()
	if err != nil {
		fmt.Errorf("internal.NewVaultProvider")
		return
	}

	conf := env.New(vault)

	get := func(v string) string {
		res, err := conf.Get(v)

		if err != nil {
			fmt.Errorf("Couldn't get configuration value for %s: %s", v, err)
		}

		return res
	}

	// XXX: We will revisit this code in future episodes replacing it with another solution
	databaseUser := get("DATABASE_USER")
	databasePassword := get("DATABASE_PASSWORD")

	fmt.Println("env 1 is - ", databaseUser)
	fmt.Println("env 2 is - ", databasePassword)

}
