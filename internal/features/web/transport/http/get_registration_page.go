package web_transport_http

import (
	"net/http"

	core_logger "github.com/nkondratev/golang-todoapp/internal/core/logger"
	core_http_response "github.com/nkondratev/golang-todoapp/internal/core/transport/http/response"
)

func (h *WebHTTPHandler) GetRegistrationPage(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	html, err := h.webService.GetRegistrationPage()
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get registration.html from main page",
		)
		return
	}

	responseHandler.HTMLResponse(html)
}
