package Surveys

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func LoadSurvey(file string) Survey {
	filebytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	survey := Survey{}
	err = json.Unmarshal([]byte(filebytes), &survey)
	if err != nil {
		panic(err)
	}
	fmt.Println(survey)
	return survey
}

func CreateSurvey(internal_title string, language LanguageClass) Survey {
	if len(internal_title) > 25 {
		panic("Survey title too long!")
	}
	internal_title = strings.ToLower(internal_title)
	internal_title = strings.ReplaceAll(internal_title, " ", "-")
	survey := Survey{Class: SurveyClass.String(), ID: rand.Intn(100000000), L10N: language.String(), InternalTitle: internal_title}
	survey.AddRun(internal_title)
	survey.AddTopContainer()
	return survey
}


func (survey *Survey) AddRun(external_title string) {
	//uses defaults, should be parameterized in the future
	contactMsg := "<p>For questions relating to this survey, <b>please contact</b> [CONTACT-PERSON]</p>"
	login_message := "Please sign in below using the credentials supplied to you"
	notification_invite_body := "Dear [EMAIL]\nAn online survey has been created for you to complete.\n\n[CUSTOM_URL]\n\n- Online surveys Team"
	notification_invite_subject := "You have been sent a survey to complete"
	notification_reminder_body := "Dear [EMAIL]\nWe have not yet detected your response to our survey.\n\n[CUSTOM_URL]\n\n- Online surveys Team"
	notification_reminder_subject := "Survey completion reminder"
	permit_back_forth := true
	permit_finish_later := false
	permit_response_printout := false
	short_name := external_title
	show_completion_receipt := false
	show_numbering := false
	show_progress := true
	timezone := 346
	NewRun := Run{Class: SurveyRun, AlertFrequency: 0, AlertLastUpdate: 0, ContactMessage: contactMsg, DiscourageSearchEngines: true, ExpectedResponses: 100, ID: rand.Intn(1000000), LoginMessage: login_message, NotificationInviteBody: notification_invite_body, NotificationInviteSubject: notification_invite_subject, NotificationReminderBody: notification_reminder_body, NotificationReminderSubject: notification_reminder_subject, PermitBackForth: permit_back_forth, PermitFinishLater: permit_finish_later, PermitResponsePrintout: permit_response_printout, ShortName: short_name, ShowCompletionReceipt: show_completion_receipt, ShowNumbering: show_numbering, ShowProgress: show_progress, Timezone: timezone, ExternalTitle: external_title}
	survey.Runs = append(survey.Runs, NewRun)
}

func (survey *Survey) AddTopContainer() {
	if survey.TopContainer.ID > 0 {
		panic("Top container already exists!")
	}
	var show = Show
	id, label := CreateID()
	tc := InnerContainer{Class: Container, ID: id, DataQuestionVisibility: show.String(), Label: label, Mandatory: JsonFalse.JBool()}
	survey.TopContainer = tc
}

func (survey *Survey) AddPage(title string, data_question_visibility Visibility) PageContainer {
	id, label := CreateID()
	page := &InnerContainer{Class: Page, Title: title, Label: label, ID: id, Mandatory: JsonFalse.JBool(), DataQuestionVisibility: data_question_visibility.String()}
	if survey.TopContainer.ID > 0 {
		survey.TopContainer.Children = append(survey.TopContainer.Children, page)
	} else {
		panic("No top container found!")
	}
	pc := PageContainer{Page: page}
	pc.Survey = survey
	return pc
}

func CreateID() (int, string) {
	id := rand.Intn(1000000)
	label := "b" + strconv.Itoa(id)
	return id, label
}

func (innerchild *PageContainer) AddNote(text string, data_question_visibility Visibility) *InnerContainer {
	id, label := CreateID()
	Note := &InnerContainer{Class: Note, Text: text, Label: label, ID: id, Mandatory: JsonFalse.JBool(), DataQuestionVisibility: data_question_visibility.String()}
	innerchild.Page.Children = append(innerchild.Page.Children, Note)
	return Note
}

func (innerchild *PageContainer) AddMultipleChoice(survey *Survey, text string, defaulttxt string, data_question_visibility Visibility, displayOptionality bool, displayLegend bool, haslogic bool, has_other bool, layout SurveyLayout, mandatory JsonBool) *InnerContainer {
	id, label := CreateID()
	survey.SurveyQNO++
	multiplechoice := &InnerContainer{Class: MultipleChoice, Text: text, Default: defaulttxt, QNo: strconv.Itoa(survey.SurveyQNO), DisplayLegend: &displayLegend, Layout: layout.String(), DisplayOptionality: &displayOptionality, HasLogic: &haslogic, HasOther: &has_other, Label: label, ID: id, Mandatory: mandatory.JBool(), DataQuestionVisibility: data_question_visibility.String()}
	innerchild.Page.Children = append(innerchild.Page.Children, multiplechoice)
	return multiplechoice
}

func (survey *Survey) ToJson() string {
	survey.SurveyQNO = 0 //remove temporary variable
	b, err := json.MarshalIndent(survey, "", "\t")
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return string(b)
}

func (survey *Survey) Save(filename string, content string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func (innerchild *InnerContainer) AddOption(text string, screen_to_message bool, routePage *PageContainer, is_other bool, is_not_applicable bool) {
	//add the classes where options can be added
	if innerchild.Class == MultipleChoice || innerchild.Class == MultipleAnswer {
		id, _ := CreateID()
		var option Option
		if routePage == nil {
			option = Option{Class: SelectionOption, ID: id, Text: text, Value: len(innerchild.Options) + 1, ScreenToMessage: screen_to_message, IsOther: is_other, IsNotApplicable: is_not_applicable}
		} else {
			option = Option{Class: SelectionOption, ID: id, Text: text, Value: len(innerchild.Options) + 1, ScreenToMessage: screen_to_message, RouteToPage: LabelToRoute(routePage.Page.Label), IsOther: is_other, IsNotApplicable: is_not_applicable}
		}
		innerchild.Options = append(innerchild.Options, option)
	}
}

func LabelToRoute(label string) string {
	return "#" + label
}

func (innerchild *PageContainer) AddScale(survey *Survey, text string, data_question_visibility Visibility, display_optionality bool, mandatory JsonBool, max_number int, min_number int, number_per_column int, number_per_row int, show_headings bool, show_hints bool) *InnerContainer {
	id, label := CreateID()
	survey.SurveyQNO++
	scale := &InnerContainer{Class: Scale, Text: text, QNo: strconv.Itoa(survey.SurveyQNO), DisplayOptionality: &display_optionality, Label: label, MaxNumber: max_number, MinNumber: min_number, NumberPerColumn: number_per_column, NumberPerRow: number_per_row, ShowHeadings: &show_headings, ShowHints: &show_hints, ID: id, Mandatory: mandatory.JBool(), DataQuestionVisibility: data_question_visibility.String()}
	innerchild.Page.Children = append(innerchild.Page.Children, scale)
	survey.SurveyQNO++
	return scale
}

func (innerchild *PageContainer) AddGrid(survey *Survey, text string, data_question_visibility Visibility, display_optionality bool, mandatory JsonBool, max_number int, min_number int, number_per_column int, number_per_row int, show_headings bool, show_hints bool) *InnerContainer {
	id, label := CreateID()
	survey.SurveyQNO++
	grid := &InnerContainer{Class: Grid, Text: text, QNo: strconv.Itoa(survey.SurveyQNO), DisplayOptionality: &display_optionality, Label: label, MaxNumber: max_number, MinNumber: min_number, NumberPerColumn: number_per_column, NumberPerRow: number_per_row, ShowHeadings: &show_headings, ShowHints: &show_hints, ID: id, Mandatory: mandatory.JBool(), DataQuestionVisibility: data_question_visibility.String()}
	innerchild.Page.Children = append(innerchild.Page.Children, scale)
	survey.SurveyQNO++
	return grid
}

func (innerchild *InnerContainer) AppendLabel(value string){
	innerchild.Label += value
}

func (innerchild *InnerContainer) AddScaleRow(survey *Survey, text string, data_question_visibility Visibility, mandatory JsonBool) *InnerContainer {
	if innerchild.Class == Scale {
		id, label := CreateID()
		scalerow := &InnerContainer{Class: ScaleRow, Text: text, QNo: strconv.Itoa(survey.SurveyQNO) + "." + strconv.Itoa(len(innerchild.Children)+1), Label: label, ID: id, Mandatory: mandatory.JBool(), DataQuestionVisibility: data_question_visibility.String()}
		innerchild.Children = append(innerchild.Children, scalerow)
		return scalerow
	}
	return nil
}

func (innerchild *InnerContainer) AddGridRow(survey *Survey, text string, data_question_visibility Visibility, mandatory JsonBool) *InnerContainer {
	if innerchild.Class == Grid {
		id, label := CreateID()
		gridrow := &InnerContainer{Class: GridRow, Text: text, QNo: strconv.Itoa(survey.SurveyQNO) + "." + strconv.Itoa(len(innerchild.Children)+1), Label: label, ID: id, Mandatory: mandatory.JBool(), DataQuestionVisibility: data_question_visibility.String()}
		innerchild.Children = append(innerchild.Children, gridrow)
		return gridrow
	}
	return nil
}

func (innerchild *InnerContainer) AddMultipleAnswer(survey *Survey, text string, show_hints bool, layout SurveyLayout, data_question_visibility Visibility, has_other bool, mandatory JsonBool) *InnerContainer {
	if innerchild == nil {
		panic("No child found!")
	}
	if innerchild.Class == ScaleRow {
		id, label := CreateID()
		survey.SurveyQNO++
		multianswer := &InnerContainer{Class: MultipleAnswer, Text: text, QNo: strconv.Itoa(survey.SurveyQNO) + "." + strconv.Itoa(len(innerchild.Children)+1) + ".i", Layout: layout.String(), HasOther: &has_other, ShowHints: &show_hints, Label: label, ID: id, Mandatory: mandatory.JBool()}
		innerchild.Children = append(innerchild.Children, multianswer)
		return multianswer
	}
	return nil
}
