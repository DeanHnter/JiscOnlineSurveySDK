package main

import (
	"fmt"

	Surveys "github.com/DeanHnter/JiscOnlineSurveySDK"
)

func main() {
	survey := Surveys.CreateSurvey("Helloworld", Surveys.English)
	page1 := survey.AddPage("Page 1", Surveys.Show)
	page1.AddNote("Hello world!", Surveys.Show)
	survey.AddPage("Final page", Surveys.Show)
	json := survey.ToJson()
	fmt.Println(json)
	survey.Save("helloworld-survey.json", json)
}
