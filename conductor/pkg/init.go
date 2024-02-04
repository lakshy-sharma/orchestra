package pkg

import (
	"conductor/pkg/db"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitializeConductor(logger *zap.Logger) {
	setupIdentity(logger)

	// Connect to the database and create users.
	historianConnection, err := db.NewDatabase(
		fmt.Sprintf("%s:%d", viper.GetString("historian.redis_cluster_endpoint"), viper.GetInt("historian.redis_cluster_port")),
		viper.GetString("historian.redis_password"),
		logger,
	)
	if err != nil {
		logger.Fatal("Failed to connect with the database. Shutting Down Conductor.")
	}
	historianConnection.Client.Process(redis.NewSliceCmd("SETUSER", "musician", "on", "deployments", "+get"))
	logger.Info("Conductor Init Complete.")
}

// This function sets up the identity for server.
// Once these keys are ready they become your identity and will be always verified by musicians before they connect back.
func setupIdentity(logger *zap.Logger) {
	// Generate private key.
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		logger.Fatal("Failed to generate private key.")
	}
	err = privateKey.Validate()
	if err != nil {
		logger.Fatal("Failed to validate RSA private key.")
	}

	// Extract public component.
	publicKey := privateKey.Public()

	// Encode private key to PKCS#1 ASN.1 PEM.
	privateKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)

	// Encode public key to PKCS#1 ASN.1 PEM.
	publicKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(publicKey.(*rsa.PublicKey)),
		},
	)

	// Save your identity to files.
	err = writeKeyToFile(privateKeyPEM, path.Join(viper.GetString("security.identity_dir"), "id_rsa.pem"), logger)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = writeKeyToFile([]byte(publicKeyPEM), path.Join(viper.GetString("security.identity_dir"), "id_pub.pem"), logger)
	if err != nil {
		log.Fatal(err.Error())
	}

	logger.Info("Cluster Identity Setup Complete.")
}

// Write key data to a PEM file.
func writeKeyToFile(keyBytes []byte, fileName string, logger *zap.Logger) error {
	err := os.WriteFile(fileName, keyBytes, 0600)
	if err != nil {
		return err
	}
	logger.Sugar().Infof("Key saved to: %s", fileName)
	return nil
}
