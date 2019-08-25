package wip

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/amanqi/go-wip/wip"
)

func ShowProfile(format string) error {
	c := wip.NewClient(os.Getenv("WIP_API_KEY"))
	viewer, err := c.GetViewer()
	if err != nil {
		log.Println(err)
		return err
	}
	switch format {
	case "json":
		jsonData, err := json.Marshal(viewer)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println(string(jsonData))

	default:
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
	}

	return nil
}
