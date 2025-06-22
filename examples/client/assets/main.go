package main

import (
	"log"

	"github.com/rangertaha/hxe/internal/db"
	"github.com/rangertaha/hxe/pkg/client"
)

func main() {
	client := client.New("http://localhost:8080")

	// Create an asset
	newAsset := &db.Asset{
		Name:        "Bitcoin",
		Symbol:      "BTC",
		Description: "Digital gold",
	}
	created, err := client.CreateAsset(newAsset)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created asset: %v", created)

	// List assets
	assets, err := client.ListAssets()
	if err != nil {
		log.Fatal(err)
	}
	for _, asset := range assets {
		log.Printf("Asset: %v", asset.Name)
	}

	// Get an asset
	asset, err := client.GetAsset(created.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Asset: %v", asset)
	// Update an asset
	asset.Description = "Updated description"
	updated, err := client.UpdateAsset(asset.ID, asset)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Updated asset: %v", updated)
	// Delete an asset
	err = client.DeleteAsset(asset.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Deleted asset: %v", asset.ID)
}
