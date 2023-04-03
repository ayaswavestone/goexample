package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/serviceusage/v1"
)

func main() {
	ctx := context.Background()
	apis_name := ["Compute Engine", "App Engine", "Kubernetes Engine", "Cloud Functions", "Filestore API", "Cloud Storage", "Big Table",
					"Firestore", "Data Migration", "SQL Inspector", "MemoryStore Redis", "MemorySTore Memcached", "Cloud Spanner", "Big Query"
					"Cloud Composer", "Cloud Data Fusion", "Cloud Data Catalog", "DataFlow API", "Cloud DataProc API", "Vertex AI API", "Cloud Translation API"
					"Cloud TPU API", "DialogFlow API", "Cloud Talent Solution API", "Cloud Speech To Text API", "Cloud Vision API",
					"Document AI API", "Cloud Text-To-Speech API", "Notebook API", "Video Intelligence API"]

	apis := ["compute.googleapis.com", "appengine.googleapis.com", "container.googleapis.com", "cloudfunctions.googleapis.com",
	 "file.googleapis.com", "storage.googleapis.com", "bigtableadmin.googleapis.com", "firestore.googleapis.com", "datamigration.googleapis.com"
	 "sqladmin.googleapis.com", "redis.googleapis.com", "memcache.googleapis.com", "spanner.googleapis.com", "bigquery.googleapis.com", "composer.googleapis.com"
	 "datafusion.googleapis.com", "datacatalog.googleapis.com", "dataflow.googleapis.com", "dataproc.googleapis.com", "aiplatform.googleapis.com", "translate.googleapis.com"
	 "tpu.googleapis.com", "dialogflow.googleapis.com", "jobs.googleapis.com", "speech.googleapis.com", "vision.googleapis.com",
	 "documentai.googleapis.com", "texttospeech.googleapis.com", "notebooks.googleapis.com", "videointelligence.googleapis.com"]

	// Replace with your project ID and JSON key file path.
	projectID := "peaceful-nation-305119"
	keyfilePath := "../peaceful-nation-305119-3afbc04ee210.json"

	serviceUsageService, err := serviceusage.NewService(ctx, option.WithCredentialsFile(keyfilePath))
	if err != nil {
		log.Fatalf("Failed to create service usage service: %v", err)
	}

	var enabled [30]bool

	for i:=0; i<len(apis); i++ {
		apiToCheck = apis[i]
		serviceName := fmt.Sprintf("projects/%s/services/%s", projectID, apiToCheck)

		service, err := serviceUsageService.Services.Get(serviceName).Do()
		if err != nil {
			log.Fatalf("Failed to get service information: %v", err)
		}

		if service.State == "ENABLED" {
			enabled[i] = true
			fmt.Printf("Kubernetes API is enabled for project: %s\n", projectID)
		} else {
			enabled[i] = false
			fmt.Printf("Kubernetes API is not enabled for project: %s\n", projectID)
		}
	}

	ind := 0
	for serviceName := range(apis_name){
		fmt.Println("%v enabled : %v", serviceName, enabled[ind])
		ind += 1
	}
}
