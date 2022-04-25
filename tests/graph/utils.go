package graph

func getGraphDefinition() []byte {
	// Graph named lectureteacher.
	return []byte(`{
  "edgeDefinitions": [
    {
      "collection": "teach",
      "from": [
        "teachers"
      ],
      "to": [
        "lectures"
      ]
    }
  ],
  "name": "lectureteacher",
  "options": {}
}`)
}
