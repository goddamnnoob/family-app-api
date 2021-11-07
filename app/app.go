package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/goddamnnoob/family-app-api/domain"
	"github.com/goddamnnoob/family-app-api/errs"
	"github.com/goddamnnoob/family-app-api/logger"
	"github.com/goddamnnoob/family-app-api/service"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Start() {
	router := httprouter.New()
	dbClient := getDbClient()
	uh := UserHandlers{service.NewUserService(domain.NewUserRepository(dbClient))}
	router.GET("/ping", Ping)
	router.GET("/getAllFamilyMembers/:userid", uh.getAllFamilyMembers)
	router.GET("/getUserById/:userid", uh.GetUserByUserId)
	router.POST("/createUser", uh.CreateUser)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getDbClient() *mongo.Client {
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://dbUser:lolgowtham@familyapp.zaeth.mongodb.net/users?retryWrites=true&w=majority") // access restricted by ip
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		errs.NewUnexpectedError("Unable to connect to DB " + err.Error())
		panic(err.Error())
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		errs.NewUnexpectedError("Unable to connect to DB" + err.Error())
	}
	logger.Info("DB connected")
	return client
}
