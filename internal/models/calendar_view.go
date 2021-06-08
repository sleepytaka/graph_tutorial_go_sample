package models

import "time"

type CalendarView struct {
	OdataContext  string `json:"@odata.context"`
	OdataNextLink string `json:"@odata.nextLink"`
	Value         []struct {
		OdataEtag                  string        `json:"@odata.etag"`
		Id                         string        `json:"id"`
		CreatedDateTime            time.Time     `json:"createdDateTime"`
		LastModifiedDateTime       time.Time     `json:"lastModifiedDateTime"`
		ChangeKey                  string        `json:"changeKey"`
		Categories                 []interface{} `json:"categories"`
		OriginalStartTimeZone      string        `json:"originalStartTimeZone"`
		OriginalEndTimeZone        string        `json:"originalEndTimeZone"`
		ICalUId                    string        `json:"iCalUId"`
		ReminderMinutesBeforeStart int           `json:"reminderMinutesBeforeStart"`
		IsReminderOn               bool          `json:"isReminderOn"`
		HasAttachments             bool          `json:"hasAttachments"`
		Subject                    string        `json:"subject"`
		BodyPreview                string        `json:"bodyPreview"`
		Importance                 string        `json:"importance"`
		Sensitivity                string        `json:"sensitivity"`
		IsAllDay                   bool          `json:"isAllDay"`
		IsCancelled                bool          `json:"isCancelled"`
		IsOrganizer                bool          `json:"isOrganizer"`
		ResponseRequested          bool          `json:"responseRequested"`
		SeriesMasterId             interface{}   `json:"seriesMasterId"`
		ShowAs                     string        `json:"showAs"`
		Type                       string        `json:"type"`
		WebLink                    string        `json:"webLink"`
		OnlineMeetingUrl           interface{}   `json:"onlineMeetingUrl"`
		ResponseStatus             struct {
			Response string    `json:"response"`
			Time     time.Time `json:"time"`
		} `json:"responseStatus"`
		Body struct {
			ContentType string `json:"contentType"`
			Content     string `json:"content"`
		} `json:"body"`
		Start struct {
			DateTime string `json:"dateTime"`
			TimeZone string `json:"timeZone"`
		} `json:"start"`
		End struct {
			DateTime string `json:"dateTime"`
			TimeZone string `json:"timeZone"`
		} `json:"end"`
		Location struct {
			DisplayName string `json:"displayName"`
		} `json:"location"`
		Recurrence interface{} `json:"recurrence"`
		Attendees  []struct {
			Type   string `json:"type"`
			Status struct {
				Response string    `json:"response"`
				Time     time.Time `json:"time"`
			} `json:"status"`
			EmailAddress struct {
				Name    string `json:"name"`
				Address string `json:"address"`
			} `json:"emailAddress"`
		} `json:"attendees"`
		Organizer struct {
			EmailAddress struct {
				Name    string `json:"name"`
				Address string `json:"address"`
			} `json:"emailAddress"`
		} `json:"organizer"`
	} `json:"value"`
}
