package main

import (
	"os"
	"time"

	"github.com/hashicorp/vault-client-go"
)

func NewVault() (*vault.Client, error) {
	client, err := vault.New(
		vault.WithAddress(os.Getenv("VAULT_DEV_LISTEN_ADDRESS")),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}
