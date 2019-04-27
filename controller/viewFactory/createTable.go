package viewFactory

import (
	"plane-ticket-search-web/dto"
	"strconv"
)

func GetTableHeader() string {
	html := ""
	html += "<table>"
	html += "<tr>"
	html += "<th>" +"index"+ "</th>"
	html += "<th>" +"base nation"+ "</th>"
	html += "<th>" +"base airport"+ "</th>"
	html += "<th>" +"target nation "+ "</th>"
	html += "<th>" +"target airport"+ "</th>"
	html += "<th>" +"base depart time"+ "</th>"
	html += "<th>" +"target arrival time"+ "</th>"
	html += "<th>" +"target depart time"+ "</th>"
	html += "<th>" +"base arrival time"+ "</th>"
	html += "<th>" +"base airline"+ "</th>"
	html += "<th>" +"target airline"+ "</th>"
	html += "<th>" +"price"+ "</th>"
	html += "</tr>"
	return html
}
func GetTableRow(data dto.TicketPriceDTO, index int) string {
	html := ""
	html += "<tr>"
	html += "<td>" +strconv.Itoa(index)+ "</td>"
	html += "<td>" +data.BaseNation+ "</td>"
	html += "<td>" +data.BaseAirport+ "</td>"
	html += "<td>" +data.TargetNation+ "</td>"
	html += "<td>" +data.TargetAirport+ "</td>"
	html += "<td>" +data.BaseDepartTime+ "</td>"
	html += "<td>" +data.TargetArrivalTime+ "</td>"
	html += "<td>" +data.TargetDepartTime+ "</td>"
	html += "<td>" +data.BaseArrivalTime+ "</td>"
	html += "<td>" +data.BaseAirline+ "</td>"
	html += "<td>" +data.TargetAirline+ "</td>"
	html += "<td>" +strconv.Itoa(int(data.Price))+ "</td>"
	html += "</tr>"
	return html
}
func GetTableFooter() string {
	html := "</table>"
	return html
}
