package models

type TemplateData struct {
	AllArtist   []Artists
	Artist      Artists
	AllRelation Relation
	Relation    interface{}
	AppInfos    App
	Error       ErrorData
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
