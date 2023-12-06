package pkg

import (
	"net/http"
	"strconv"
)

func Error(code int) map[string]string {
	tab := make(map[string]string)
	switch code {
	case http.StatusNoContent:
		tab["code"] = strconv.Itoa(code)
		tab["msg"] = "No Content"
	case http.StatusBadRequest:
		tab["code"] = strconv.Itoa(code)
		tab["msg"] = "Bad request"
	case http.StatusNotFound:
		tab["code"] = strconv.Itoa(code)
		tab["msg"] = "Not Found"
	case http.StatusMethodNotAllowed:
		tab["code"] = strconv.Itoa(code)
		tab["msg"] = "Method Not Allowed"
	case http.StatusInternalServerError:
		tab["code"] = strconv.Itoa(code)
		tab["msg"] = "Internal Server Error"
	}
	return tab
}
