package requestUtils

func getCrudServiceAddress() string {
	const address = "http://curl:9001"
	return address
}

func getAllTicketUrl() string {
	return getCrudServiceAddress() + "/getAll"
}
