package client

import (
	"fmt"

	"github.com/rangertaha/hxe/internal/models"
)

// CredentialClient handles credential-related API requests
type CredentialClient struct {
	client *Client
}

// List returns all credentials
func (c *CredentialClient) List() ([]models.Credential, error) {
	var credentials []models.Credential
	err := c.client.Get("/api/credentials", &credentials)
	return credentials, err
}

// Get returns a single credential by ID
func (c *CredentialClient) Get(id uint) (*models.Credential, error) {
	var credential models.Credential
	err := c.client.Get(fmt.Sprintf("/api/credentials/%d", id), &credential)
	return &credential, err
}

// Create creates a new credential
func (c *CredentialClient) Create(credential *models.Credential) (*models.Credential, error) {
	err := c.client.Post("/api/credentials", credential, &credential)
	return credential, err
}

// Update updates an existing credential
func (c *CredentialClient) Update(credential *models.Credential) (*models.Credential, error) {
	err := c.client.Put(fmt.Sprintf("/api/credentials/%d", credential.ID), credential, &credential)
	return credential, err
}

// Delete deletes a credential
func (c *CredentialClient) Delete(id uint) error {
	var response interface{}
	return c.client.Delete(fmt.Sprintf("/api/credentials/%d", id), &response)
}

// GetByType returns credentials filtered by type
func (c *CredentialClient) GetByType(credentialType string) ([]models.Credential, error) {
	var credentials []models.Credential
	err := c.client.Get(fmt.Sprintf("/api/credentials?type=%s", credentialType), &credentials)
	return credentials, err
}

// GetByGroup returns credentials filtered by group ID
func (c *CredentialClient) GetByGroup(groupID uint) ([]models.Credential, error) {
	var credentials []models.Credential
	err := c.client.Get(fmt.Sprintf("/api/credentials?group_id=%d", groupID), &credentials)
	return credentials, err
}

// GetActive returns only active credentials
func (c *CredentialClient) GetActive() ([]models.Credential, error) {
	var credentials []models.Credential
	err := c.client.Get("/api/credentials?is_active=true", &credentials)
	return credentials, err
}
