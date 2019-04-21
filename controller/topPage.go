package controller

import (
	"fmt"
	"log"
	"net/http"
	"simpleWebServer/controller/requestUtils"
	"simpleWebServer/controller/viewFactory"
)

func TopPageController(w http.ResponseWriter, r *http.Request) {
	_ = r

	allTicketData, err := requestUtils.GetAllTicketData()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	html := ""

	html += viewFactory.GetTableHeader()
	for i:=0; i<len(allTicketData); i++ {
		rowData := allTicketData[i]
		html += viewFactory.GetTableRow(rowData, i)
	}
	html += viewFactory.GetTableFooter()
	i, err := fmt.Fprintln(w, html)
	log.Println(i, err)
}

