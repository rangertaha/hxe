package main

import (
	"fmt"
	"log"

	"github.com/rangertaha/hxe/internal/models"
	"github.com/rangertaha/hxe/pkg/client"
)

func main() {
	// Create a new client
	c := client.NewClient("http://localhost:8080", "admin", "password")

	// Login to get authentication token
	_, err := c.Login()
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}

	// Create a new credential
	credential := &models.Credential{
		Name:        "SSH Key for Production",
		Type:        "ssh_key",
		Username:    "deploy",
		PrivateKey:  "-----BEGIN PRIVATE KEY-----\n...\n-----END PRIVATE KEY-----",
		PublicKey:   "ssh-rsa AAAAB3NzaC1yc2E...",
		Description: "SSH key for production server deployment",
		IsActive:    true,
	}

	// Create the credential
	createdCredential, err := c.Credentials.Create(credential)
	if err != nil {
		log.Fatalf("Failed to create credential: %v", err)
	}
	fmt.Printf("Created credential: %+v\n", createdCredential)

	// Get all credentials
	credentials, err := c.Credentials.List()
	if err != nil {
		log.Fatalf("Failed to list credentials: %v", err)
	}
	fmt.Printf("Found %d credentials\n", len(credentials))

	// Get a specific credential
	credential, err = c.Credentials.Get(createdCredential.ID)
	if err != nil {
		log.Fatalf("Failed to get credential: %v", err)
	}
	fmt.Printf("Retrieved credential: %+v\n", credential)

	// Update the credential
	credential.Description = "Updated description for SSH key"
	updatedCredential, err := c.Credentials.Update(credential)
	if err != nil {
		log.Fatalf("Failed to update credential: %v", err)
	}
	fmt.Printf("Updated credential: %+v\n", updatedCredential)

	// Get credentials by type
	sshCredentials, err := c.Credentials.GetByType("ssh_key")
	if err != nil {
		log.Fatalf("Failed to get SSH credentials: %v", err)
	}
	fmt.Printf("Found %d SSH credentials\n", len(sshCredentials))

	// Get active credentials
	activeCredentials, err := c.Credentials.GetActive()
	if err != nil {
		log.Fatalf("Failed to get active credentials: %v", err)
	}
	fmt.Printf("Found %d active credentials\n", len(activeCredentials))

	// Delete the credential
	err = c.Credentials.Delete(createdCredential.ID)
	if err != nil {
		log.Fatalf("Failed to delete credential: %v", err)
	}
	fmt.Println("Credential deleted successfully")
}
