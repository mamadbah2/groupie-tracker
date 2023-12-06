package models

type TemplateData struct {
	StringData map[string]string
	IntData    map[string]int
	AppInfos   App
	Error      ErrorData
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
