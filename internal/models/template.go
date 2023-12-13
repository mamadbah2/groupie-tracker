package models

type TemplateData struct {
	AllArtist []Artists
	Relation  Relation
	AppInfos  App
	Error     ErrorData
}

type App struct {
	AppName   string // Name of the application
	PageTitle string // Title of the page
	PageName  string // Name of the page
}

type ErrorData struct {
	Code    string
	Message string
}
