package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	pq "github.com/lib/pq"
)

type Storage struct {
	Host     string `json:"host"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	User     string `json:"user"`
}
type Compute struct {
	ComputeURL string `json:"url"`
}

type ResourceInformation struct {
	ServingURL string `json:"serving_url"`
}

type Metadata struct {
	Compute Compute `json:"compute"`
	Storage Storage `json:"storage"`
}

type InfrastructureMetadata struct {
	Metadata            Metadata            `json:"metadata"`
	ResourceInformation ResourceInformation `json:"resource_information"`
}

func main() {
	fmt.Println("initiation started...")

	var metadataFlg string
	flag.StringVar(&metadataFlg, "metadata", "", "infrastructure metadata to be processed")
	flag.Parse()

	if metadataFlg == "" {
		fmt.Println("empty metadata")
		return
	}

	pq.Array("")

	fmt.Println(metadataFlg)

	var infraMetadata InfrastructureMetadata

	err := json.Unmarshal([]byte(metadataFlg), &infraMetadata)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(infraMetadata.Metadata)

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			infraMetadata.Metadata.Storage.Host,
			infraMetadata.Metadata.Storage.Port,
			infraMetadata.Metadata.Storage.User,
			infraMetadata.Metadata.Storage.Password,
			infraMetadata.Metadata.Storage.Name,
		),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.ReadFile("../../../../migrations/2_todos.sql")
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := db.Exec(string(f))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res, "initiation success")
}
