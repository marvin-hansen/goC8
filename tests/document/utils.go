package document

import (
	"fmt"
	"github.com/marvin-hansen/goC8"
)

func getTestInsertData() []byte {
	return []byte(` 
		[
		  {
			"item1": "data1",
			"item2": "data2"

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

func getTestReplaceManyData(key1, key2 string) []byte {
	str := fmt.Sprintf(` 
		[
		  {
			"_key": "%v",
			"item1": "dataReplaced"
		  },
		 {
			"_key": "%v",
			"item1": "dataReplaced"
		  }
	]
	`, key1, key2)
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

func printRes(res goC8.Responder, verbose bool) {
	if verbose {
		println(res.String())
	}
}

func printJsonRes(res goC8.JsonResponder, verbose bool) {
	if verbose {
		println(res.String())
	}
}
