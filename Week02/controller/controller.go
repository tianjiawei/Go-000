package controller

type Handler struct {
	UserController *UserController
}

func NewHandler(userController *UserController) *Handler {
	return &Handler{
		UserController: userController,
	}
}
