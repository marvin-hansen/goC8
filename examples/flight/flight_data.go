package main

func GetCityData() []byte {
	// Document collection named cities.
	return []byte(`[
    {"_key": "sanfrancisco", "location": [-122.416667, 37.783333]},
    {"_key": "newyork", "location": [-74.0059, 40.7127]},
    {"_key": "detroit", "location": [-83.045833, 42.331389]},
    {"_key": "paris", "location": [2.3508, 48.8567]},
    {"_key": "dublin", "location": [-6.260278, 53.349722]},
    {"_key": "singapore", "location": [103.8, 1.3]},
    {"_key": "houston", "location": [-95.383056, 29.762778]},
    {"_key": "seattle", "location": [-122.333056, 47.609722]}
]
`)
}

func GetFlightData() []byte {
	return []byte(` 
		[
		{"_from": "cities/sanfrancisco", "_to": "cities/singapore", "distance": 13600},
		{"_from": "cities/sanfrancisco", "_to": "cities/newyork", "distance": 4000},
		{"_from": "cities/sanfrancisco", "_to": "cities/detroit", "distance": 3300},
		{"_from": "cities/sanfrancisco", "_to": "cities/houston", "distance": 2600},
		{"_from": "cities/sanfrancisco", "_to": "cities/seattle", "distance": 807},
		{"_from": "cities/detroit", "_to": "cities/sanfrancisco", "distance": 3300},
		{"_from": "cities/detroit", "_to": "cities/newyork", "distance": 800},
		{"_from": "cities/newyork", "_to": "cities/sanfrancisco", "distance": 4000},
		{"_from": "cities/newyork", "_to": "cities/detroit", "distance": 800},
		{"_from": "cities/newyork", "_to": "cities/dublin", "distance": 5100},
		{"_from": "cities/newyork", "_to": "cities/paris", "distance": 5800},
		{"_from": "cities/newyork", "_to": "cities/houston", "distance": 2300},
		{"_from": "cities/newyork", "_to": "cities/seattle", "distance": 2831},
		{"_from": "cities/dublin", "_to": "cities/newyork", "distance": 5100},
		{"_from": "cities/dublin", "_to": "cities/paris", "distance": 800},
		{"_from": "cities/paris", "_to": "cities/newyork", "distance": 5800},
		{"_from": "cities/paris", "_to": "cities/dublin", "distance": 800},
		{"_from": "cities/paris", "_to": "cities/singapore", "distance": 10700},
		{"_from": "cities/singapore", "_to": "cities/sanfrancisco", "distance": 13600},
		{"_from": "cities/singapore", "_to": "cities/paris", "distance": 10700},
		{"_from": "cities/houston", "_to": "cities/sanfrancisco", "distance": 2600},
		{"_from": "cities/houston", "_to": "cities/newyork", "distance": 2300}
]
`)
}

func GetAirlineGraph() []byte {
	return []byte(`
{
  "edgeDefinitions": [
    {
      "collection": "flights",
      "from": [
        "cities"
      ],
      "to": [
        "cities"
      ]
    }
  ],
  "name": "airline",
  "options": {}
}

`)
}

// GetAllCitiesQuery
// Get all cities. Same as SELECT * FROM cities.
func GetAllCitiesQuery() string {
	return "for u in cities  return u._key"
}

// GetBreadthFirstQuery
// Get all cities with a direct flight to New York.
func GetBreadthFirstQuery() string {
	return `
	WITH cities
     FOR city IN INBOUND "cities/newyork" flights
     RETURN city._key
`
}

// GetShortestPathQuery
// Get the shortest path from San Francisco to Paris.
func GetShortestPathQuery() string {
	return `
	WITH cities
	LET path = (
    FOR city IN OUTBOUND SHORTEST_PATH "cities/sanfrancisco" TO "cities/paris"
        GRAPH "airline"
        OPTIONS {"weightAttribute": "distance"}
        RETURN city._key
     )
     RETURN path
`
}

// GetShortestDistanceQuery
// Get the distance on the shortest path from San Francisco to Paris.
func GetShortestDistanceQuery() string {
	return `
	WITH cities
     LET path = (
    FOR city, e IN OUTBOUND SHORTEST_PATH "cities/sanfrancisco" TO "cities/paris"
        GRAPH "airline"
        OPTIONS {"weightAttribute": "distance"}
        RETURN e.distance
    )
    RETURN SUM(path)
`
}

// GetNearestCities
// Get the 2 nearest cities to a specified latitude and longitude.
func GetNearestCities() string {
	return `
	FOR loc IN NEAR(cities, 53.35, -6.26, 2, "distance")
	RETURN {
    name: loc._key,
    distance: loc.distance / 1000
    }
`
}

// GetCitiesMaxDistance
// Get the cities that are no more than 2500km away from houston.
func GetCitiesMaxDistance() string {
	return `
	LET city = DOCUMENT("cities/houston")
	FOR loc IN WITHIN(cities, city.location[1], city.location[0], 2500 * 1000, "distance")
	RETURN {
    name: loc._key,
    distance: loc.distance / 1000
    }
`
}
