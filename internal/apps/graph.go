package apps

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
	"github.com/sleepytaka/graph_tutorial_go_sample/internal/models"
)

type Server struct {
	cli   *http.Client
	token string
}

func New(appId string, scopes []string) *Server {
	app, err := public.New(appId, public.WithAuthority("https://login.microsoftonline.com/common"))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	devCode, err := app.AcquireTokenByDeviceCode(ctx, scopes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Device Code is: %s\n", devCode.Result.Message)
	result, err := devCode.AuthenticationResult(ctx)
	if err != nil {
		panic(fmt.Sprintf("got error while waiting for user to input the device code: %s", err))
	}
	s := &Server{
		cli:   new(http.Client),
		token: result.AccessToken,
	}
	return s
}

func (g *Server) GetMe() (*models.User, error) {
	req, _ := http.NewRequest("GET", "https://graph.microsoft.com/v1.0/me", nil)
	req.Header.Set("Authorization", "Bearer "+g.token)

	resp, err := g.cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Status Error : " + resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err.Error())
	}

	return &user, nil
}

func (g *Server) ListCalendarEvents() error {
	t := time.Now().UTC()
	t1 := t.AddDate(0, 0, int(t.Weekday())*-1)
	t2 := t1.AddDate(0, 0, 7)
	url := fmt.Sprintf("https://graph.microsoft.com/v1.0/me/calendarview?startDateTime=%s&endDateTime=%s", t1.Format("2006-01-02T00:00:00Z"), t2.Format("2006-01-02T00:00:00Z"))
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+g.token)

	resp, err := g.cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Status Error : " + resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	events := models.CalendarView{}
	err = json.Unmarshal(body, &events)
	if err != nil {
		return err
	}

	for _, v := range events.Value {
		fmt.Println("Subject:" + v.Subject)
		fmt.Println("Organizer:" + v.Organizer.EmailAddress.Name)
		fmt.Printf("Start:%s (%s)\n", v.Start.DateTime, v.Start.TimeZone)
		fmt.Printf("End:%s (%s)\n", v.End.DateTime, v.End.TimeZone)
	}

	return nil
}

func (g *Server) CreateEvent(subject string, start time.Time, end time.Time, content string) error {
	event := models.Event{}
	event.Subject = subject
	event.Start.DateTime = start.Format("2006-01-02T15:04:05Z")
	event.Start.TimeZone = "UTC"
	event.End.DateTime = end.Format("2006-01-02T15:04:05Z")
	event.End.TimeZone = "UTC"
	event.Body.Content = content
	event.Body.ContentType = "Text"
	event.Attendees = []models.Attendee{}
	j, _ := json.Marshal(event)
	req, _ := http.NewRequest("POST", "https://graph.microsoft.com/v1.0/me/events", bytes.NewBuffer(j))
	req.Header.Set("Authorization", "Bearer "+g.token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := g.cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("Status Error : " + resp.Status)
	}

	return nil
}

func (g *Server) GetAccessToken() string {
	return g.token
}
