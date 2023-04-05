package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/api/iam/v1"
	"google.golang.org/api/option"
)

func mainjson() {
	ctx := context.Background()

	// Remplacez par votre ID de projet et le chemin de votre fichier de clé JSON.
	projectID := "your-project-id"
	keyfilePath := "/path/to/your/keyfile.json"

	iamService, err := iam.NewService(ctx, option.WithCredentialsFile(keyfilePath))
	if err != nil {
		log.Fatalf("Failed to create IAM service: %v", err)
	}

	serviceAccountID := "your-service-account-id"
	displayName := "My Service Account"
	resourceName := fmt.Sprintf("projects/%s", projectID)

	// Créez le compte de service.
	serviceAccount, err := iamService.Projects.ServiceAccounts.Create(resourceName, &iam.CreateServiceAccountRequest{
		AccountId: serviceAccountID,
		ServiceAccount: &iam.ServiceAccount{
			DisplayName: displayName,
		},
	}).Do()

	if err != nil {
		log.Fatalf("Failed to create service account: %v", err)
	}

	fmt.Printf("Service account created: %s\n", serviceAccount.Email)

	// Créez une clé pour le compte de service.
	key, err := iamService.Projects.ServiceAccounts.Keys.Create(serviceAccount.Name, &iam.CreateServiceAccountKeyRequest{
		KeyAlgorithm:   "KEY_ALG_RSA_2048",
		PrivateKeyType: "TYPE_GOOGLE_CREDENTIALS_FILE",
	}).Do()

	if err != nil {
		log.Fatalf("Failed to create service account key: %v", err)
	}

	// Téléchargez la clé de compte de service en tant que fichier JSON.
	privateKeyData, err := base64.StdEncoding.DecodeString(key.PrivateKeyData)
	if err != nil {
		log.Fatalf("Failed to decode private key data: %v", err)
	}

	outputFile := "new-service-account-key.json"
	if err := ioutil.WriteFile(outputFile, privateKeyData, 0600); err != nil {
		log.Fatalf("Failed to write private key to file: %v", err)
	}

	fmt.Printf("Service account key created and saved to: %s\n", outputFile)
}
