package client

// import (
// 	"fmt"

// 	"github.com/rangertaha/hxe/internal/api/models"
// )

// // Entity operations

// func (c *Client) ListEntities() ([]models.Entity, error) {
// 	var entities []models.Entity
// 	err := c.Get("/api/entities", &entities)
// 	return entities, err
// }

// func (c *Client) GetEntity(id uint) (*models.Entity, error) {
// 	var entity models.Entity
// 	err := c.Get(fmt.Sprintf("/api/entities/%d", id), &entity)
// 	return &entity, err
// }

// func (c *Client) CreateEntity(entity *models.Entity) (*models.Entity, error) {
// 	var created models.Entity
// 	err := c.Post("/api/entities", entity, &created)
// 	return &created, err
// }

// func (c *Client) UpdateEntity(id uint, entity *models.Entity) (*models.Entity, error) {
// 	var updated models.Entity
// 	err := c.Put(fmt.Sprintf("/api/entities/%d", id), entity, &updated)
// 	return &updated, err
// }

// func (c *Client) DeleteEntity(id uint) error {
// 	return c.Delete(fmt.Sprintf("/api/entities/%d", id))
// }
