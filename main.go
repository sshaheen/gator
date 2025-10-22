package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/sshaheen/gator/internal/app"
	"github.com/sshaheen/gator/internal/commands"
	"github.com/sshaheen/gator/internal/config"
	"github.com/sshaheen/gator/internal/database"
)

func main() {
	config_data, err := config.Read()
	if err != nil {
		fmt.Println("Failed to read config")
		os.Exit(1)
	}

	app_state := app.State{
		CFG: &config_data,
	}

	db, err := sql.Open("postgres", app_state.CFG.DbURL)
	if err != nil {
		fmt.Println("Error connecting to DB")
		os.Exit(1)
	}

	dbQueries := database.New(db)
	app_state.DB = dbQueries

	cmds := commands.Commands{
		Commands: make(map[string]func(*app.State, commands.Command) error),
	}

	cmds.Register("login", commands.LoginHandler)
	cmds.Register("register", commands.RegisterHandler)
	cmds.Register("reset", commands.ResetHandler)
	cmds.Register("users", commands.GetUsersHandler)

	in_args := os.Args
	if len(in_args) < 2 {
		fmt.Println("You must add a command")
		os.Exit(1)
	}

	cmd := commands.Command{
		Name: in_args[1],
		Args: in_args[2:],
	}

	err = cmds.Run(&app_state, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
