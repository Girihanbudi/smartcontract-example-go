package rest

func (h Handler) RegisterApi() {
	filesystem := h.Router.Group("/filesystem")
	{
		filesystem.POST("", h.Upload)
		filesystem.GET("", h.Download)
	}
}
