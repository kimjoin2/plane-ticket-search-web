package requestUtils

func getCrudServiceAddress() string {
	const address = "http://127.0.0.1:9001"
	return address
}

func getAllTicketUrl() string {
	return getCrudServiceAddress() + "/getAll"
}
