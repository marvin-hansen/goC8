package main

import (
	"github.com/marvin-hansen/goC8"
	kv_req2 "github.com/marvin-hansen/goC8/requests/kv_req"
)

const (
	// client config
	apiKey   = "email.root.secretkey"
	endpoint = "https://YOUR-ID-us-west.paas.macrometa.io"
	fabric   = "SouthEastAsia"
	timeout  = 5 // http connection timeout in seconds
	// collection config
	delete         = false
	verbose        = true
	collectionName = "TestKVCollection"
	expiration     = false
)

func main() {
	println("Create new config ")
	config := goC8.NewConfig(apiKey, endpoint, fabric, timeout)

	println("Create new client with config ")
	c := goC8.NewClient(config)

	println("Create KV collection: " + collectionName)
	res, err := c.KV.CreateNewKVCollection(fabric, collectionName, expiration, nil)
	goC8.CheckErrorLog(err, "Error creatin KV collection"+collectionName)
	goC8.PrintRes(res, verbose)

	println("Get all KV collection: ")
	res, err = c.KV.GetAllKVCollections(fabric)
	goC8.CheckErrorLog(err, "Error getting KV collection"+collectionName)
	goC8.PrintRes(res, verbose)

	println("Insert values into KV collection: ")
	kvPair1 := kv_req2.NewKVPair("key1", "value1", -1)
	kvPair2 := kv_req2.NewKVPair("key2", "value2", -1)
	kvPair3 := kv_req2.NewKVPair("key3", "value3", -1)
	kvPair4 := kv_req2.NewKVPair("key4", "value4", -1)
	kvPair5 := kv_req2.NewKVPair("key5", "value5", -1)

	kvCollection := kv_req2.NewKVPairCollection(*kvPair1, *kvPair2, *kvPair3, *kvPair4, *kvPair5)
	resSet, errSet := c.KV.SetKeyValuePairs(fabric, collectionName, *kvCollection)
	goC8.CheckErrorLog(errSet, "Error setting values into KV collection"+collectionName)
	goC8.PrintRes(resSet, verbose)

	println("Count all items in KV collection: ")
	resCount, errCount := c.KV.CountKVCollection(fabric, collectionName)
	goC8.CheckErrorLog(errCount, "Error counting values into KV collection"+collectionName)
	goC8.PrintRes(resCount, verbose)

	println("Get all keys from KV collection: ")
	offset := 0
	limit := 20
	order := kv_req2.Ascending

	resGetKeys, errGetKeys := c.KV.GetAllKeys(fabric, collectionName, offset, limit, order)
	goC8.CheckErrorLog(errGetKeys, "Error getting all keys from KV collection"+collectionName)
	goC8.PrintRes(resGetKeys, verbose)

	println("Get all values from KV collection: ")
	resGetValues, errGetValues := c.KV.GetAllValues(fabric, collectionName, offset, limit, nil)
	goC8.CheckErrorLog(errGetValues, "Error getting all values from KV collection"+collectionName)
	goC8.PrintRes(resGetValues, verbose)

	println("Get single value from KV collection: ")
	key := "value1"
	resOneVal, errOneVal := c.KV.GetValue(fabric, collectionName, key)
	goC8.CheckErrorLog(errOneVal, "Error getting one value from KV collection"+collectionName)
	goC8.PrintRes(resOneVal, verbose)

	println("Delete single value from KV collection: ")
	resOneDel, errOneDel := c.KV.DeleteValue(fabric, collectionName, key)
	goC8.CheckErrorLog(errOneDel, "Error deleting one value from KV collection"+collectionName)
	goC8.PrintRes(resOneDel, verbose)

	println("Update values into KV collection: ")
	kvPairUpdate1 := kv_req2.NewKVPair("key1", "value42", -1)
	kvPairUpdate2 := kv_req2.NewKVPair("key2", "value23", -1)

	kvCollectionUpdate := kv_req2.NewKVPairCollection(*kvPairUpdate1, *kvPairUpdate2)
	resUpd, errUpd := c.KV.SetKeyValuePairs(fabric, collectionName, *kvCollectionUpdate)
	goC8.CheckErrorLog(errUpd, "Error updating values into KV collection"+collectionName)
	goC8.PrintRes(resUpd, verbose)

	println("Get updated values from KV collection: ")
	resGetValues, errGetValues = c.KV.GetAllValues(fabric, collectionName, offset, limit, nil)
	goC8.CheckErrorLog(errGetValues, "Error getting all values from KV collection"+collectionName)
	goC8.PrintRes(resGetValues, verbose)

	if delete {
		_, errTruncate := c.KV.TruncateKVCollection(fabric, collectionName)
		goC8.CheckErrorLog(errTruncate, "Error truncating all values in KV collection"+collectionName)

		_, err = c.KV.DeleteKVCollection(fabric, collectionName)
		goC8.CheckErrorLog(errTruncate, "Error deleting KV collection"+collectionName)
	}

}
