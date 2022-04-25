package main

func GetTeacherData() []byte {
	// Document collection named teachers.
	return []byte(`
[
        {
            "_key": "Jean",
            "firstname": "Jean",
            "lastname": "Picard",
            "email": "jean.picard@macrometa.io"
        },
        {
            "_key": "James",
            "firstname": "James",
            "lastname": "Kirk",
            "email": "james.kirk@macrometa.io"
        },
        {
            "_key": "Han",
            "firstname": "Han",
            "lastname": "Solo",
            "email": "han.solo@macrometa.io"
        },
        {
            "_key": "Bruce",
            "firstname": "Bruce",
            "lastname": "Wayne",
            "email": "bruce.wayne@macrometa.io"
        }
]
`)
}

func GetLecturesData() []byte {
	// Document collection named lectures.
	return []byte(`[
        {"_id": "lectures/CSC101", "difficulty": "easy", "_key":"CSC101", "firstname":"Jean"},
        {"_id": "lectures/CSC102", "difficulty": "hard", "_key":"CSC102","firstname":"Jean"},
        {"_id": "lectures/CSC103", "difficulty": "hard", "_key":"CSC103","firstname":"Jean"},
        {"_id": "lectures/CSC104", "difficulty": "moderate", "_key":"CSC104","firstname":"Jean"},
        {"_id": "lectures/CSC105", "difficulty": "hard", "_key":"CSC105","firstname":"Jean"}
    ]`)
}

func GetTeachEdgeData() []byte {
	// Document collection named teach.
	str := `
[
        {
            "_key": "Jean-CSC101",
            "_from": "teachers/Jean",
            "_to": "lectures/CSC101",
            "online": false
        },
        {
            "_key": "Jean-CSC102",
            "_from": "teachers/Jean",
            "_to": "lectures/CSC102",
            "online": true
        },
        {
            "_key": "Jean-CSC103",
            "_from": "teachers/Jean",
            "_to": "lectures/CSC103",
            "online": false
        },
        {
            "_key": "Jean-CSC104",
            "_from": "teachers/Jean",
            "_to": "lectures/CSC104",
            "online": false
        },
        {
            "_key": "Han-CSC105",
            "_from": "teachers/Han",
            "_to": "lectures/CSC105",
            "online": false
        },
        {
            "_key": "Bruce-CSC101",
            "_from": "teachers/Bruce",
            "_to": "lectures/CSC101",
            "online": true
        },
 		{
            "_key": "James-CSC105",
            "_from": "teachers/James",
            "_to": "lectures/CSC105",
            "online": false
        }
]`

	return []byte(str)
}

func GetUniversityGraphDefinition() []byte {
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

func GetTutorialsData() []byte {
	// Document collection named tutorials.
	return []byte(`[
		{"_id": "tutorials/TCS101", "difficulty": "easy", "_key": "TCS101", "firstname": "Han"},
        {"_id": "tutorials/TCS102", "difficulty": "hard", "_key": "TCS102", "firstname": "Han"},
        {"_id": "tutorials/TCS103", "difficulty": "hard", "_key": "TCS103", "firstname": "Han"},
        {"_id": "tutorials/TCS104", "difficulty": "moderate", "_key": "TCS104", "firstname": "Han"},
        {"_id": "tutorials/TCS105", "difficulty": "moderate", "_key": "TCS105", "firstname": "Han"}
]`)
}

func GetTutorsEdgeData() []byte {
	// Document collection named tutors.
	return []byte(`[		 
		{
            "_key": "Han-TCS101",
            "_from": "teachers/Han",
            "_to": "tutorials/TCS101",
            "mandatory": false
        },
        {
            "_key": "Han-TCS102",
            "_from": "teachers/Han",
            "_to": "tutorials/TCS102",
            "mandatory": true
        },
        {
            "_key": "Han-TCS103",
            "_from": "teachers/Han",
            "_to": "tutorials/TCS103",
            "online": false
        },
        {
            "_key": "Han-TCS104",
            "_from": "teachers/Han",
            "_to": "tutorials/TCS104",
            "mandatory": false
        },
        {
            "_key": "Han-TCS105",
            "_from": "teachers/Han",
            "_to": "tutorials/TCS105",
            "mandatory": false
        }
]`)
}
