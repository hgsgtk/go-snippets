package controller

import (
	"net/http"

	"github.com/hgsgtk/go-snippets/layerd-port-arch/service"
)

type UserViewController struct {
	Service service.UserViewService
}

type UserViewRequest struct {
	ID int `json:"id"`
}

type UserViewResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
}

func (c *UserViewController) Handler(w http.ResponseWriter, r *http.Request) {
	rb := UserViewRequest{}
	if err := decodeRequestBodyJSON(r, &rb); err != nil {
		respondError(w, err)
		return
	}
	output, err := c.Service.Run(service.UserViewInput{ID: rb.ID})
	if err != nil {
		// respond error
	}
	rs := UserViewResponse{
		ID:       output.User.ID,
		FullName: output.User.GetFullName(),
	}
	respond(w, rs, http.StatusOK)
}

func decodeRequestBodyJSON(r *http.Request, v interface{}) *ErrorResponse {
	// json decoding
	return nil
}

func respond(w http.ResponseWriter, v interface{}, status int) {
	// respond
}

func respondError(w http.ResponseWriter, err *ErrorResponse) {
	// respond error
	return
}

// error response format
type ErrorResponse struct{}
