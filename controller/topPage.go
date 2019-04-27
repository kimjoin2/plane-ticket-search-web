package controller

import (
	"fmt"
	"log"
	"net/http"
	"plane-ticket-search-web/controller/requestUtils"
	"plane-ticket-search-web/controller/viewFactory"
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
	_, err = fmt.Fprintln(w, html)
	if err != nil {
		log.Println(err)
	}
}

