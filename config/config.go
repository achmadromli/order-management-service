package config

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/vault/api"
)

// AppConfig adalah struktur untuk menyimpan konfigurasi aplikasi
type AppConfig struct {
	Database DatabaseConfig
	Server   ServerConfig
	Vault    VaultConfig
}

// DatabaseConfig adalah konfigurasi database
type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

// ServerConfig adalah konfigurasi server HTTP
type ServerConfig struct {
	Port string
}

// VaultConfig adalah konfigurasi HashiCorp Vault
type VaultConfig struct {
	Address     string
	Token       string
	SecretsPath string // Path di Vault tempat menyimpan secrets
}

// LoadConfig membaca konfigurasi dari HashiCorp Vault
func LoadConfig() *AppConfig {
	vaultAddr := os.Getenv("VAULT_ADDR")
	vaultToken := os.Getenv("VAULT_TOKEN")
	secretsPath := os.Getenv("VAULT_SECRETS_PATH")

	if vaultAddr == "" || vaultToken == "" || secretsPath == "" {
		log.Fatal("Missing required environment variables for Vault configuration.")
	}

	client, err := api.NewClient(&api.Config{
		Address: vaultAddr,
	})
	if err != nil {
		log.Fatalf("Failed to create Vault client: %v", err)
	}

	// Baca konfigurasi dari Vault
	secret, err := client.Logical().Read(secretsPath)
	if err != nil {
		log.Fatalf("Failed to read configuration from Vault: %v", err)
	}

	if secret == nil {
		log.Fatal("No configuration found in Vault.")
	}

	var dbConfig DatabaseConfig
	dbConfig.Driver = "postgres" // Ganti dengan driver database yang sesuai (mysql, postgres, dll)
	dbConfig.Host = secret.Data["db_host"].(string)
	dbConfig.Port = secret.Data["db_port"].(string)
	dbConfig.Username = secret.Data["db_username"].(string)
	dbConfig.Password = secret.Data["db_password"].(string)
	dbConfig.Name = secret.Data["db_name"].(string)

	var serverConfig ServerConfig
	serverConfig.Port = os.Getenv("SERVER_PORT")

	return &AppConfig{
		Database: dbConfig,
		Server:   serverConfig,
		Vault:    VaultConfig{Address: vaultAddr, Token: vaultToken, SecretsPath: secretsPath},
	}
}

// GetDatabaseDSN mengembalikan Data Source Name (DSN) untuk koneksi database
func (c *DatabaseConfig) GetDatabaseDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.Username, c.Password, c.Host, c.Port, c.Name)
}
