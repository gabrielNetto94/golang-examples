package src

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"os"
)

func CsvToJson(path string) string {

	// Open the SCV filepost
	//csvFile, err := os.Open("../files/candy.csv")
	csvFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	// Read the SCV file
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Convert the SCV records to JSON
	jsonData, err := json.Marshal(records)
	if err != nil {
		panic(err)
	}

	output := "../out/file.json"
	// Write the JSON data to a file
	file, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		panic(err)
	}

	return output

}

func ReadAndWriteFile() {

	filename := "hotel"
	// Open the SCV file
	csvFile, err := os.Open("../files/" + filename + ".csv")
	//csvFile, err := os.Open("../files/candy.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	// Create a new JSON file
	jsonFile, err := os.Create("../out/" + filename + ".json")
	//jsonFile, err := os.Create("../out/candy.json")

	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	// Create a JSON encoder
	//encoder := json.NewEncoder(jsonFile)

	// Create a CSV reader
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Read the CSV file in chunks
	for {
		records, err := reader.Read()
		if err != nil {
			break
		}
		// Convert the chunk to JSON
		jsonData, err := json.Marshal(records)
		if err != nil {
			panic(err)
		}

		// Write the JSON chunk to the file
		_, err = jsonFile.Write(jsonData)
		if err != nil {
			panic(err)
		}

		// Write a comma separator after each JSON chunk
		_, err = jsonFile.WriteString(",")
		if err != nil {
			panic(err)
		}
	}
}
