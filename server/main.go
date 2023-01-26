package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()

	cfg := struct {
		Type                    string `json:"type"`
		ProjectId               string `json:"project_id"`
		PrivateKeyId            string `json:"private_key_id"`
		PrivateKey              string `json:"private_key"`
		ClientEmail             string `json:"client_email"`
		ClientId                string `json:"client_id"`
		AuthURI                 string `json:"auth_uri"`
		TokenURI                string `json:"token_uri"`
		AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
		ClientX509CertUrl       string `json:"client_x509_cert_url"`
	}{
		Type:                    os.Getenv("FIREBASE_TYPE"),
		ProjectId:               os.Getenv("FIREBASE_PROJECT_ID"),
		PrivateKeyId:            os.Getenv("FIREBASE_PRIVATE_KEY_ID"),
		PrivateKey:              os.Getenv("FIREBASE_PRIVATE_KEY"),
		ClientEmail:             os.Getenv("FIREBASE_CLIENT_EMAIL"),
		ClientId:                os.Getenv("FIREBASE_CLIENT_ID"),
		AuthURI:                 os.Getenv("FIREBASE_AUTH_URI"),
		TokenURI:                os.Getenv("FIREBASE_TOKEN_URI"),
		AuthProviderX509CertURL: os.Getenv("FIREBASE_PROVIDER_X509_CERT_URL"),
		ClientX509CertUrl:       os.Getenv("FIREBASE_CLIENT_X509_CERT_URL"),
	}

	configJson, _ := json.Marshal(cfg)

	opt := option.WithCredentialsJSON(configJson)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		panic(err)
	}

	registrationToken := "dP2texqfPk1LACLeDS-3et:APA91bEAaC9-URDzc4I2TuHbs3fnpRwdNeuVeawHNB3KmLLwwofuP7JDJq840YboehlWCPcRfAcVwJaTwEFBZyBL0lzUzWWXfx9r8lIiHrrM8wINo1BW6NhC_8v3BwfGKauV55Sui4AM"

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Notification Test",
			Body:  "Hello React!!",
		},
		Token: registrationToken,
		Data: map[string]string{
			"nama": "Adhiana Mastur",
			"umur": "25",
		},
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully sent message:", response)

	messages := []*messaging.Message{}
	for i := 0; i < 100; i++ {
		messages = append(messages, &messaging.Message{
			Notification: &messaging.Notification{
				Title: "Notification Test",
				Body:  "Hello React!!",
			},
			Token: registrationToken,
			Data: map[string]string{
				"nama": "Adhiana Mastur",
				"umur": "25",
			},
		})
	}

	responses, err := client.SendAll(ctx, messages)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(responses.FailureCount)
	fmt.Println(responses.SuccessCount)
	fmt.Println(responses.Responses)

	// for _, r := range responses.Responses {
	// 	//
	// }

	fmt.Println("Sucessfully sent messages:", responses)
}
