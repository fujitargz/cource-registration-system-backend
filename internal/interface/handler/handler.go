package handler

type AppHandler interface {
	UserHandler
}

type appHandler struct {
	UserHandler
}

func NewAppHandler(u UserHandler) AppHandler {
	return appHandler{u}
}
