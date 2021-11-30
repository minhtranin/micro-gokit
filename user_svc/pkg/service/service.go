package service

import (
	"context"
	"fmt"
	"log"
	"user_svc/pkg/db"
	"user_svc/pkg/model"
)

// UserSvcService describes the service.
type UserSvcService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Create(ctx context.Context, email string, password string) (rs string, err error)
	Get(ctx context.Context) (rs []model.User, err error)
}

type basicUserSvcService struct{}

func (b *basicUserSvcService) Create(ctx context.Context, email string, password string) (rs string, err error) {
	// TODO implement the business logic of Create
	user := model.User{
		Email:    email,
		Password: password,
	}
	u, err := db.GetDb().Collection("user").InsertOne(ctx, user)
	fmt.Println(err, "erro 28")
	if err != nil {
		log.Fatal()
	}
	fmt.Println(u.InsertedID, "----")
	return rs, err
}

// NewBasicUserSvcService returns a naive, stateless implementation of UserSvcService.
func NewBasicUserSvcService() UserSvcService {
	return &basicUserSvcService{}
}

// New returns a UserSvcService with all of the expected middleware wired in.
func New(middleware []Middleware) UserSvcService {
	var svc UserSvcService = NewBasicUserSvcService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicUserSvcService) Get(ctx context.Context) (rs []model.User, err error) {
	// TODO implement the business logic of Get
	return rs, err
}
