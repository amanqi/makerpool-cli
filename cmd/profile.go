/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/amanqi/go-wip/wip"
	"github.com/spf13/cobra"
)

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Manage user profile",
	Run: func(cmd *cobra.Command, args []string) {
		showProfile()
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
}

func showProfile() error {
	c := wip.NewClient(os.Getenv("WIP_API_KEY"))
	viewer, err := c.GetViewer()
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println()
	fmt.Println("Platform: Work In Progress")
	fmt.Println("ID:", viewer.ID)
	fmt.Println("Username:", viewer.Username)
	fmt.Println("Firstname:", viewer.FirstName)
	fmt.Println("Lastname:", viewer.LastName)
	fmt.Println("URL:", viewer.URL)
	productsCount := len(viewer.Products)
	productsLabel := fmt.Sprintf("%d products:", productsCount)
	if productsCount == 0 {
		productsLabel = "No product yet"
	}
	fmt.Println(productsLabel)
	for _, product := range viewer.Products {
		fmt.Printf(" - %s\n", product.Name)
	}
	fmt.Println("Completed todos:", viewer.CompletedTodosCount)
	if viewer.Streaking {
		fmt.Println("Streak:", viewer.Streak)
		fmt.Println("Best streak:", viewer.BestStreak)
	}
	todosCount := len(viewer.Todos)
	todosLabel := fmt.Sprintf("%d pending todos:", todosCount)
	if productsCount == 0 {
		todosLabel = "No pending todo"
	}
	fmt.Println(todosLabel)
	for _, todo := range viewer.Todos {
		fmt.Printf(" - %s\n", todo.Body)
	}

	return nil
}
