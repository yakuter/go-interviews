// data_store.go
package main

type DataStore interface {
	GetData() string
}

func GetDataFromStore(ds DataStore) string {
	return ds.GetData()
}
