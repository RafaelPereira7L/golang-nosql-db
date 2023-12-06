package db

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/gookit/color"
)

type Data struct {
	ID    string                 `json:"_id"`
	Fields map[string]interface{} `json:"fields"`
}

type Collection struct {
	Data []Data `json:"data"`
}

func CreateCollection(name string) Collection {
	color.Cyan.Printf("Creating Collection %s...\n", name)

	file, err := os.Create("./storage/collections/"+name+".json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write([]byte("{\"data\":[]}"))
	if err != nil {
		panic(err)
	}

	color.Green.Printf("Collection %s is Created Successfully!\n", name)

	return Collection{}
}

func NewData(fields map[string]interface{}) Data {
	newUUID := uuid.New()
	idString := newUUID.String()
	return Data{
		ID:    idString,
		Fields: fields,
	}
}

func (c *Collection) Add(data Data) {
	c.Data = append(c.Data, data)
}

func (c *Collection) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Collection) SaveToFile(filename string) error {
	jsonData, err := json.Marshal(c.Data)
	if err != nil {
		return err
	}

	file, err := os.Create("./storage/collections/"+filename + ".json")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	return err
}

func (c *Collection) LoadFromFile(filename string) error {
	file, err := os.Open("./storage/collections/"+filename + ".json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c.Data)
	return err
}

func main() {
	fmt.Println("Database Initialized!")

	collection := Collection{}

	collectionFilename := "example.json"

	err := collection.LoadFromFile(collectionFilename)
	if err != nil {
		fmt.Println("Error loading collection:", err)
		collection = CreateCollection(collectionFilename)
	}

	dataToInsert := map[string]interface{}{"fodase": 1234}
	collection.Add(NewData(dataToInsert))

	err = collection.SaveToFile(collectionFilename)
	if err != nil {
		fmt.Println("Error saving collection:", err)
	}

	jsonData, err := collection.ToJSON()
	if err != nil {
		fmt.Println("Error converting collection to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}
