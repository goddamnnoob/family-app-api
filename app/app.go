package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/goddamnnoob/family-app-api/domain"
	"github.com/goddamnnoob/family-app-api/errs"
	"github.com/goddamnnoob/notReddit/service"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Start() {
	router := httprouter.New()
	dbClient := getDbClient()
	uh := UserHandlers{service.NewUserService(domain.NewUserRepository(dbClient))}
	router.GET("/ping", Ping)
	router.GET("/:userid/getAllFamilyMembers", uh.getAllFamilyMembers)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getDbClient() *mongo.Client {
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://dbUser:lolgowtham@familyapp.zaeth.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		errs.NewUnexpectedError("Unable to connect to DB " + err.Error())
	}
	return client
}
