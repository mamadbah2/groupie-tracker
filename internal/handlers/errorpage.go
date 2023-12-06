package handlers

import (
	"bim/internal/models"
	"bim/internal/pkg"
	"net/http"
)

func errorResponse(w http.ResponseWriter, code int) {
	err := pkg.Error(code)
	w.WriteHeader(code)
	renderTemplates(w, "error", &models.TemplateData{AppInfos: models.App{AppName: appName, PageTitle: err["code"] + " " + err["msg"]}, Error: models.ErrorData{Code: err["code"], Message: err["msg"]}})
	return
}
