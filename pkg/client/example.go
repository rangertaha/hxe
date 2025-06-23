package client

// import (
// 	"fmt"
// 	"log"
// )

// // Example usage of the HXE client
// func Example() {
// 	// Create a new client
// 	// client := NewClient("http://localhost:8080")

// 	// // Get the program client
// 	// programClient := client.Program()

// 	// // List all programs
// 	// programs, err := client.program.List()
// 	// if err != nil {
// 	// 	log.Fatal("Failed to list programs:", err)
// 	// }
	
// 	// fmt.Printf("Found %d programs\n", len(programs))

// 	// // Create a new program
// 	// newProgram := &models.Program{
// 	// 	Name:        "Example Program",
// 	// 	Description: "A test program",
// 	// 	Command:     "/usr/bin/echo",
// 	// 	Args:        "Hello World",
// 	// 	Directory:   "/tmp",
// 	// 	User:        "nobody",
// 	// 	Group:       "nobody",
// 	// 	Status:      "stopped",
// 	// 	Autostart:   false,
// 	// 	Enabled:     true,
// 	// 	Retries:     0,
// 	// 	MaxRetries:  3,
// 	// }

// 	// created, err := programClient.CreateProgram(newProgram)
// 	// if err != nil {
// 	// 	log.Fatal("Failed to create program:", err)
// 	// }
// 	// fmt.Printf("Created program with ID: %d\n", created.ID)

// 	// // Get a specific program
// 	// program, err := programClient.GetProgram(created.ID)
// 	// if err != nil {
// 	// 	log.Fatal("Failed to get program:", err)
// 	// }
// 	// fmt.Printf("Program: %s - Status: %s\n", program.Name, program.Status)

// 	// // Start the program
// 	// err = programClient.StartProgram(created.ID)
// 	// if err != nil {
// 	// 	log.Fatal("Failed to start program:", err)
// 	// }
// 	// fmt.Println("Program started")

// 	// // Update the program
// 	// program.Description = "Updated description"
// 	// updated, err := programClient.UpdateProgram(created.ID, program)
// 	// if err != nil {
// 	// 	log.Fatal("Failed to update program:", err)
// 	// }
// 	// fmt.Printf("Updated program: %s\n", updated.Description)

// 	// // Enable autostart
// 	// err = programClient.EnableProgram(created.ID)
// 	// if err != nil {
// 	// 	log.Fatal("Failed to enable program:", err)
// 	// }
// 	// fmt.Println("Program enabled for autostart")

// 	// // Restart the program
// 	// err = programClient.RestartProgram(created.ID)
// 	// if err != nil {
// 	// 	log.Fatal("Failed to restart program:", err)
// 	// }
// 	// fmt.Println("Program restarted")

// 	// // Stop the program
// 	// err = programClient.StopProgram(created.ID)
// 	// if err != nil {
// 	// 	log.Fatal("Failed to stop program:", err)
// 	// }
// 	// fmt.Println("Program stopped")

// 	// // Delete the program
// 	// err = programClient.DeleteProgram(created.ID)
// 	// if err != nil {
// 	// 	log.Fatal("Failed to delete program:", err)
// 	// }
// 	// fmt.Println("Program deleted")
// }
