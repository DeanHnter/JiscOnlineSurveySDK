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

type InnerContainer struct {
	ShortTitle             string           `json:"short_title,omitempty"`
	Class                  string           `json:"class"`
	DataQuestionVisibility string           `json:"data_question_visibility"`
	Title                  string           `json:"title"`
	DataIdentifier         string           `json:"data_identifier,omitempty"`
	Default                string           `json:"default,omitempty"`
	DisplayLegend          bool             `json:"displayLegend,omitempty"`
	DisplayOptionality     bool             `json:"displayOptionality,omitempty"`
	HasLogic               bool             `json:"has_logic,omitempty"`
	HasOther               bool             `json:"has_other,omitempty"`
	ID                     int              `json:"id"`
	Label                  string           `json:"label"`
	Layout                 string           `json:"layout,omitempty"`
	Mandatory              string           `json:"mandatory"`
	MoreInfo               string           `json:"moreInfo"`
	Text                   string           `json:"text"`
	QNo                    string           `json:"q_no"`
	Children               []InnerContainer `json:"children,omitempty"`
	Options                []Options        `json:"options,omitempty"`
	Maxsize                int              `json:"maxsize,omitempty"`
	Size                   int              `json:"size,omitempty"`
	Validationrule         string           `json:"validationrule,omitempty"`
	EndDate                string           `json:"end_date,omitempty"`
	StartDate              string           `json:"start_date,omitempty"`
	EndTime                string           `json:"end_time,omitempty"`
	StartTime              string           `json:"start_time,omitempty"`
	EndDatetime            string           `json:"end_datetime,omitempty"`
	StartDatetime          string           `json:"start_datetime,omitempty"`
	NumberPerColumn        int              `json:"number_per_column,omitempty"`
	NumberPerRow           int              `json:"number_per_row,omitempty"`
	ShowHeadings           bool             `json:"show_headings,omitempty"`
	Cols                   int              `json:"cols,omitempty"`
	Dependencies           string           `json:"dependencies"`
	Format                 string           `json:"format,omitempty"`
	IsOther                bool             `json:"is_other,omitempty"`
	Rows                   int              `json:"rows,omitempty"`
	ShowHints              bool             `json:"show_hints,omitempty"`
	MaxNumber              int              `json:"max_number,omitempty"`
	MinNumber              int              `json:"min_number,omitempty"`
}

type Options struct {
	Class           string `json:"class"`
	ID              int    `json:"id"`
	IsNotApplicable bool   `json:"is_not_applicable"`
	IsOther         bool   `json:"is_other"`
	RouteToPage     string `json:"route_to_page,omitempty"`
	ScreenToMessage bool   `json:"screen_to_message"`
	Text            string `json:"text"`
	Value           int    `json:"value"`
}

type TopContainer struct {
	Class                  string           `json:"class"`
	DataQuestionVisibility string           `json:"data_question_visibility"`
	ID                     int              `json:"id"`
	Label                  string           `json:"label"`
	Mandatory              string           `json:"mandatory"`
	Children               []InnerContainer `json:"children,omitempty"`
}

type Survey struct {
	Class         string       `json:"class"`
	ID            int          `json:"id"`
	InternalTitle string       `json:"internal_title"`
	L10N          string       `json:"l10n"`
	Runs          []Run        `json:"runs"`
	TopContainer  TopContainer `json:"top_container"`
}
