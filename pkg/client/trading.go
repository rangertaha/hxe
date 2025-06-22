package client

import (
	"fmt"

	"github.com/rangertaha/hxe/internal/db"
)

// Asset operations
func (c *Client) ListAssets() ([]db.Asset, error) {
	var assets []db.Asset
	err := c.Get("/api/assets", &assets)
	return assets, err
}

func (c *Client) GetAsset(id uint) (*db.Asset, error) {
	var asset db.Asset
	err := c.Get(fmt.Sprintf("/api/assets/%d", id), &asset)
	return &asset, err
}

func (c *Client) CreateAsset(asset *db.Asset) (*db.Asset, error) {
	var created db.Asset
	err := c.Post("/api/assets", asset, &created)
	return &created, err
}

func (c *Client) UpdateAsset(id uint, asset *db.Asset) (*db.Asset, error) {
	var updated db.Asset
	err := c.Put(fmt.Sprintf("/api/assets/%d", id), asset, &updated)
	return &updated, err
}

func (c *Client) DeleteAsset(id uint) error {
	return c.Delete(fmt.Sprintf("/api/assets/%d", id))
}
