package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ariefzainuri96/go-api-ecommerce-api-gateway/cmd/api/request"
	"github.com/ariefzainuri96/go-api-ecommerce-api-gateway/cmd/api/response"
	"github.com/ariefzainuri96/go-api-ecommerce-api-gateway/cmd/api/utils"
	authpb "github.com/ariefzainuri96/go-api-ecommerce-api-gateway/proto"
)

// @Summary      Login
// @Description  Perform login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request		body	  request.LoginRequest	true "Login request"
// @Success      200  			{object}  response.LoginResponse
// @Failure      400  			{object}  response.BaseResponse
// @Failure      404  			{object}  response.BaseResponse
// @Router       /auth/login	[post]
func (app *Application) login(w http.ResponseWriter, r *http.Request) {
	var data request.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	defer r.Body.Close()

	err = app.Validator.Struct(data)

	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := app.Service.AuthClient.Client.Login(r.Context(), &authpb.LoginRequest{
		Email:    data.Email,
		Password: data.Password,
	})

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, response.LoginResponse{
		BaseResponse: response.BaseResponse{
			Status:  http.StatusOK,
			Message: "Success",
		},
		Data: response.LoginData{
			ID:    int(user.Id),
			Token: user.Token,
			Name:  user.Name,
			Email: user.Email,
		},
	})
}

// @Summary      Register
// @Description  Perform register
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request		body	  request.RegisterReq	true "Register request"
// @Success      200  			{object}  response.BaseResponse
// @Failure      400  			{object}  response.BaseResponse
// @Failure      404  			{object}  response.BaseResponse
// @Router       /auth/register	[post]
func (app *Application) register(w http.ResponseWriter, r *http.Request) {
	var data request.RegisterReq
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	defer r.Body.Close()

	err = app.Validator.Struct(data)

	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = app.Service.AuthClient.Client.Register(r.Context(), &authpb.RegisterRequest{
		FullName: data.Name,
		Email:    data.Email,
		Password: data.Password,
	})

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, response.BaseResponse{
		Status:  http.StatusOK,
		Message: "Success register account",
	})
}

func (app *Application) AuthRouter() *http.ServeMux {
	authRouter := http.NewServeMux()

	authRouter.HandleFunc("POST /login", app.login)
	authRouter.HandleFunc("POST /register", app.register)

	return authRouter
}
