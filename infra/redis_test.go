package infra

import "testing"

func TestCreateDataBaseConnection(t *testing.T) {
	connection := CreateDataBaseConnection()
	if connection == nil {
		t.Error("FAILED. Connection not estabilish")
	} else {
		t.Logf("PASSED. Connection estabilish")
	}
}
