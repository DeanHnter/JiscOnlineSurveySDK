package Surveys

import "fmt"

type Visibility string

const (
	Show Visibility = "show"
	Hide            = "hide"
)

func (v Visibility) String() string {
	return fmt.Sprintf("%s", string(v))
}

type StructureClass string

const (
	SurveyClass     StructureClass = "Survey"
	SurveyRun                      = "SurveyRun"
	Container                      = "Container"
	Page                           = "Page"
	MultipleChoice                 = "MultipleChoice"
	SelectionOption                = "SelectionOption"
	FreeText                       = "FreeText"
	MultipleAnswer                 = "MultipleAnswer"
	SelectionList                  = "SelectionList"
	DatePicker                     = "DatePicker"
	TimePicker                     = "TimePicker"
	DateTimePicker                 = "DateTimePicker"
	Scale                          = "Scale"
	ScaleRow                       = "ScaleRow"
	Grid                           = "Grid"
	GridRow                        = "GridRow"
	Note                           = "Note"
)

func (sc StructureClass) String() string {
	switch sc {
	default:
		return fmt.Sprintf("%s", string(sc))
	}
}

type LanguageClass string

const (
	English LanguageClass = "en_GB"
)

func (lc LanguageClass) String() string {
	switch lc {
	default:
		return fmt.Sprintf("%s", string(lc))
	}
}

type Run struct {
	Class                       string `json:"class"`
	AlertFrequency              int    `json:"alert_frequency"`
	AlertLastUpdate             int    `json:"alert_last_update"`
	ContactMessage              string `json:"contact_message"`
	DiscourageSearchEngines     bool   `json:"discourage_search_engines"`
	ExpectedResponses           int    `json:"expected_responses"`
	ExternalTitle               string `json:"external_title"`
	ID                          int    `json:"id"`
	LoginMessage                string `json:"login_message"`
	NotificationInviteBody      string `json:"notification_invite_body"`
	NotificationInviteSubject   string `json:"notification_invite_subject"`
	NotificationReminderBody    string `json:"notification_reminder_body"`
	NotificationReminderSubject string `json:"notification_reminder_subject"`
	PermitBackForth             bool   `json:"permit_back_forth"`
	PermitFinishLater           bool   `json:"permit_finish_later"`
	PermitResponsePrintout      bool   `json:"permit_response_printout"`
	ShortName                   string `json:"short_name"`
	ShowCompletionReceipt       bool   `json:"show_completion_receipt"`
	ShowNumbering               bool   `json:"show_numbering"`
	ShowProgress                bool   `json:"show_progress"`
	Timezone                    int    `json:"timezone"`
}

type TopContainer struct {
	Class                  string `json:"class"`
	DataQuestionVisibility string `json:"data_question_visibility"`
	ID                     int    `json:"id"`
	Label                  string `json:"label"`
	Mandatory              string `json:"mandatory"`
	Children               []struct {
		Class                  string `json:"class"`
		DataQuestionVisibility string `json:"data_question_visibility"`
		ID                     int    `json:"id"`
		Label                  string `json:"label"`
		Mandatory              string `json:"mandatory"`
		Title                  string `json:"title"`
		Children               []struct {
			Class                  string `json:"class"`
			DataIdentifier         string `json:"data_identifier,omitempty"`
			DataQuestionVisibility string `json:"data_question_visibility"`
			Default                string `json:"default,omitempty"`
			DisplayLegend          bool   `json:"displayLegend,omitempty"`
			DisplayOptionality     bool   `json:"displayOptionality,omitempty"`
			HasLogic               bool   `json:"has_logic,omitempty"`
			HasOther               bool   `json:"has_other,omitempty"`
			ID                     int    `json:"id"`
			Label                  string `json:"label"`
			Layout                 string `json:"layout,omitempty"`
			Mandatory              string `json:"mandatory"`
			MoreInfo               string `json:"moreInfo"`
			Text                   string `json:"text"`
			QNo                    string `json:"q_no"`
			Options                []struct {
				Class           string `json:"class"`
				ID              int    `json:"id"`
				IsNotApplicable bool   `json:"is_not_applicable"`
				IsOther         bool   `json:"is_other"`
				RouteToPage     string `json:"route_to_page,omitempty"`
				ScreenToMessage bool   `json:"screen_to_message"`
				Text            string `json:"text"`
				Value           int    `json:"value"`
			} `json:"options,omitempty"`
			Children []struct {
				Class                  string `json:"class"`
				Cols                   int    `json:"cols,omitempty"`
				DataQuestionVisibility string `json:"data_question_visibility"`
				Dependencies           string `json:"dependencies"`
				Format                 string `json:"format,omitempty"`
				ID                     int    `json:"id"`
				IsOther                bool   `json:"is_other,omitempty"`
				Label                  string `json:"label"`
				Mandatory              string `json:"mandatory"`
				Rows                   int    `json:"rows,omitempty"`
				ShowHints              bool   `json:"show_hints,omitempty"`
				Text                   string `json:"text"`
				QNo                    string `json:"q_no"`
				Default                string `json:"default,omitempty"`
				DisplayLegend          bool   `json:"displayLegend,omitempty"`
				DisplayOptionality     bool   `json:"displayOptionality,omitempty"`
				HasOther               bool   `json:"has_other,omitempty"`
				Layout                 string `json:"layout,omitempty"`
				MaxNumber              int    `json:"max_number,omitempty"`
				MinNumber              int    `json:"min_number,omitempty"`
				MoreInfo               string `json:"moreInfo,omitempty"`
				Options                []struct {
					Class           string `json:"class"`
					ID              int    `json:"id"`
					IsNotApplicable bool   `json:"is_not_applicable"`
					IsOther         bool   `json:"is_other"`
					ScreenToMessage bool   `json:"screen_to_message"`
					Text            string `json:"text"`
					Value           int    `json:"value"`
				} `json:"options,omitempty"`
				Children []struct {
					Class                  string `json:"class"`
					DataQuestionVisibility string `json:"data_question_visibility"`
					Dependencies           string `json:"dependencies"`
					Format                 string `json:"format"`
					ID                     int    `json:"id"`
					IsOther                bool   `json:"is_other"`
					Label                  string `json:"label"`
					Mandatory              string `json:"mandatory"`
					ShowHints              bool   `json:"show_hints"`
					Size                   int    `json:"size"`
					Text                   string `json:"text"`
					QNo                    string `json:"q_no"`
				} `json:"children,omitempty"`
				DataIdentifier  string `json:"data_identifier,omitempty"`
				HasLogic        bool   `json:"has_logic,omitempty"`
				ShortTitle      string `json:"short_title,omitempty"`
				Maxsize         int    `json:"maxsize,omitempty"`
				Size            int    `json:"size,omitempty"`
				Validationrule  string `json:"validationrule,omitempty"`
				EndDate         string `json:"end_date,omitempty"`
				StartDate       string `json:"start_date,omitempty"`
				EndTime         string `json:"end_time,omitempty"`
				StartTime       string `json:"start_time,omitempty"`
				EndDatetime     string `json:"end_datetime,omitempty"`
				StartDatetime   string `json:"start_datetime,omitempty"`
				NumberPerColumn int    `json:"number_per_column,omitempty"`
				NumberPerRow    int    `json:"number_per_row,omitempty"`
				ShowHeadings    bool   `json:"show_headings,omitempty"`
			} `json:"children"`
			ShortTitle string `json:"short_title,omitempty"`
		} `json:"children,omitempty"`
	} `json:"children"`
}

type Survey struct {
	Class         string       `json:"class"`
	ID            int          `json:"id"`
	InternalTitle string       `json:"internal_title"`
	L10N          string       `json:"l10n"`
	Runs          []Run        `json:"runs"`
	TopContainer  TopContainer `json:"top_container"`
}
