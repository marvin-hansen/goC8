package document

import (
	"fmt"
)

func getTestInsertData() []byte {
	return []byte(` 
		[
		  {
			"item1": "data1"
		  },
		  {
			"item2": "data2"
		  }
		]
	`)
}

func getTestReplaceSingleData(key string) []byte {
	str := fmt.Sprintf(`
		  {
			"_key": "%v",
			"item1": "dataReplaced"
		  }
	`, key)
	return []byte(str)
}

func getTestReplaceManyData(key1 string) []byte {
	str := fmt.Sprintf(` 
		[
		  {
			"_key": "%v",
			"item1": "dataReplaced"
		  }
	]
	`, key1)
	return []byte(str)
}

func getKeysToDelete(key1, key2 string) []byte {
	str := fmt.Sprintf(` 
		[
		"%v",
		"%v"
		]
	`, key1, key2)
	return []byte(str)
}

func getTestUpdateSingleData(key string) []byte {
	s := fmt.Sprintf("{\"_key\": \"%v\", \"item1\": \"updatedData\"}\n", key)
	return []byte(s)
}

func getTestUpdateManyData(key string) []byte {
	s := fmt.Sprintf("[ {\n   \"_key\": \"%v\",\n   \"item1\": \"data42\"\n} ]", key)
	return []byte(s)
}
