/*
 * HXE - Host-based Process Execution Engine
 * Copyright (C) 2025 rangertaha@gmail.com
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
	"log"

	"github.com/rangertaha/hxe/pkg/client"
)

func main() {
	ListPrograms()
	PrintListPrograms()
	PrintGetProgram()
}

func ListPrograms() {
	client := client.NewClient("http://localhost:8080", "admin", "password")
	var err error

	// Login to get JWT token
	if client, err = client.Login(); err != nil {
		log.Fatal("Failed to login:", err)
	} else {
		fmt.Println("Successfully logged in")
	}

	programs, err := client.Program.List()
	if err != nil {
		log.Fatal("Failed to list programs:", err)
	}
	fmt.Printf("Found %d programs\n", len(programs))
}

func PrintListPrograms() {
	client := client.NewClient("http://localhost:8080", "admin", "password")
	var err error

	// Login to get JWT token
	if client, err = client.Login(); err != nil {
		log.Fatal("Failed to login:", err)
	} else {
		fmt.Println("Successfully logged in")
	}

	programs, err := client.Program.List()
	if err != nil {
		log.Fatal("Failed to list programs:", err)
	}
	fmt.Printf("Found %d programs\n", len(programs))
	client.Program.PrintList(programs)
}

func PrintGetProgram() {
	client := client.NewClient("http://localhost:8080", "admin", "password")
	var err error

	// Login to get JWT token
	if client, err = client.Login(); err != nil {
		log.Fatal("Failed to login:", err)
	} else {
		fmt.Println("Successfully logged in")
	}

	program, err := client.Program.Get("1")
	if err != nil {
		log.Fatal("Failed to list programs:", err)
	}
	fmt.Printf("Program: %s\n", program.Name)
	client.Program.PrintDetail(program)
}

// // Example usage of the HXE client with JWT authentication
// func Example() {
// 	// Create an authenticated client
// 	client := client.NewAuthenticatedClient("http://localhost:8080", "admin", "password")

// 	// Login to get JWT token
// 	err := client.Login()
// 	if err != nil {
// 		log.Fatal("Failed to login:", err)
// 	}
// 	fmt.Println("Successfully logged in")

// 	// Get the program client
// 	programClient := client.Program()

// 	// List all programs
// 	programs, err := programClient.ListPrograms()
// 	if err != nil {
// 		log.Fatal("Failed to list programs:", err)
// 	}
// 	fmt.Printf("Found %d programs\n", len(programs))

// 	// Create a new program
// 	newProgram := &models.Program{
// 		Name:        "Example Program",
// 		Description: "A test program",
// 		Command:     "/usr/bin/echo",
// 		Args:        "Hello World",
// 		Directory:   "/tmp",
// 		User:        "nobody",
// 		Group:       "nobody",
// 		Status:      "stopped",
// 		Autostart:   false,
// 		Enabled:     true,
// 		Retries:     0,
// 		MaxRetries:  3,
// 	}

// 	created, err := programClient.CreateProgram(newProgram)
// 	if err != nil {
// 		log.Fatal("Failed to create program:", err)
// 	}
// 	fmt.Printf("Created program with ID: %d\n", created.ID)

// 	// Get a specific program
// 	program, err := programClient.GetProgram(created.ID)
// 	if err != nil {
// 		log.Fatal("Failed to get program:", err)
// 	}
// 	fmt.Printf("Program: %s - Status: %s\n", program.Name, program.Status)

// 	// Start the program
// 	err = programClient.StartProgram(created.ID)
// 	if err != nil {
// 		log.Fatal("Failed to start program:", err)
// 	}
// 	fmt.Println("Program started")

// 	// Update the program
// 	program.Description = "Updated description"
// 	updated, err := programClient.UpdateProgram(created.ID, program)
// 	if err != nil {
// 		log.Fatal("Failed to update program:", err)
// 	}
// 	fmt.Printf("Updated program: %s\n", updated.Description)

// 	// Enable autostart
// 	err = programClient.EnableProgram(created.ID)
// 	if err != nil {
// 		log.Fatal("Failed to enable program:", err)
// 	}
// 	fmt.Println("Program enabled for autostart")

// 	// Restart the program
// 	err = programClient.RestartProgram(created.ID)
// 	if err != nil {
// 		log.Fatal("Failed to restart program:", err)
// 	}
// 	fmt.Println("Program restarted")

// 	// Stop the program
// 	err = programClient.StopProgram(created.ID)
// 	if err != nil {
// 		log.Fatal("Failed to stop program:", err)
// 	}
// 	fmt.Println("Program stopped")

// 	// Delete the program
// 	err = programClient.DeleteProgram(created.ID)
// 	if err != nil {
// 		log.Fatal("Failed to delete program:", err)
// 	}
// 	fmt.Println("Program deleted")

// 	// Logout
// 	client.Logout()
// 	fmt.Println("Logged out")
// }

// // Example with manual token management
// func ExampleWithManualToken() {
// 	// Create a basic client
// 	client := client.NewClient("http://localhost:8080")

// 	// Set a token manually (if you already have one)
// 	client.SetToken("your-jwt-token-here")

// 	// Get the program client
// 	programClient := client.Program()

// 	// Use the client as normal
// 	programs, err := programClient.ListPrograms()
// 	if err != nil {
// 		log.Fatal("Failed to list programs:", err)
// 	}
// 	fmt.Printf("Found %d programs\n", len(programs))
// }
