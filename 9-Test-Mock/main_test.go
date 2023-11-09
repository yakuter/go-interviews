// mock_data_store.go
package main

import "testing"

type MockDataStore struct{}

func (m MockDataStore) GetData() string {
	return "Mock data from datastore"
}

func TestGetDataFromStore(t *testing.T) {
	mockDataStore := MockDataStore{} // Create an instance of the mock data store
	result := GetDataFromStore(mockDataStore)

	expected := "Mock data from datastore"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
