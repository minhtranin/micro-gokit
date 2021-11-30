package endpoint

import (
	"context"
	"user_svc/pkg/model"
	service "user_svc/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateRequest collects the request parameters for the Create method.
// type CreateRequest struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// CreateResponse collects the response parameters for the Create method.
type CreateResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeCreateEndpoint returns an endpoint that invokes Create on the service.
func MakeCreateEndpoint(s service.UserSvcService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.User)
		rs, err := s.Create(ctx, req.Email, req.Password)
		return CreateResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Create implements Service. Primarily useful in a client.
func (e Endpoints) Create(ctx context.Context, email string, password string) (rs string, err error) {
	request := model.User{
		Email:    email,
		Password: password,
	}
	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateResponse).Rs, response.(CreateResponse).Err
}

// CreateRequest collects the request parameters for the Create method.
type CreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetRequest collects the request parameters for the Get method.
type GetRequest struct{}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	Rs  []model.User `json:"rs"`
	Err error        `json:"err"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.UserSvcService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		rs, err := s.Get(ctx)
		return GetResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.Err
}

// Get implements Service. Primarily useful in a client.
func (e Endpoints) Get(ctx context.Context) (rs []model.User, err error) {
	request := GetRequest{}
	response, err := e.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).Rs, response.(GetResponse).Err
}
