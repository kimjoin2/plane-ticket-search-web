package requestUtils

import "testing"

func TestGetAllTicketData(t *testing.T) {
	_, err := GetAllTicketData()
	if err != nil {
		t.Error(err)
	}
}