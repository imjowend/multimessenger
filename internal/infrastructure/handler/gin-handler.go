package handler

import (
	ports "github.com/imjowend/multimessenger/internal/core/services/ports"
)

type HTTPGinHandler struct {
	multimessengerService ports.MultimessengerService
}

func NewHTTPGinHandler(mms ports.MultimessengerService) *HTTPGinHandler {
	return &HTTPGinHandler{
		multimessengerService: mms,
	}
}
