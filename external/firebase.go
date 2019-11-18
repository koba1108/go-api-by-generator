package external

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var app *firebase.App
var messagingClient *messaging.Client

func initFirebase() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("serviceAccount/nl-smart-match-dev-firebase-adminsdk.json")
	if app, err = firebase.NewApp(ctx, nil, opt); err != nil {
		panic(err.Error())
	}
	if messagingClient, err = app.Messaging(ctx); err != nil {
		panic(err.Error())
	}
}

func GetFirebaseApp() *firebase.App {
	return app
}

func SendFCM(token string, data map[string]string) (string, error) {
	message := &messaging.Message{
		Data: data,
		Token: token,
	}
	return messagingClient.Send(context.Background(), message)
}
