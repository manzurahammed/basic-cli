package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/manzurahammed/rm-cli/server/models"
	"github.com/pkg/errors"
)

type HTTPClient struct {
	notifierUrl string
	client      *http.Client
}

func NewHTTPClient(url string) HTTPClient {
	return HTTPClient{
		notifierUrl: url,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

type NotificationResponse struct {
	completed bool
	duration  time.Duration
}

func (c HTTPClient) notify(reminder models.Reminder) (NotificationResponse, error) {
	var notifierResponse struct {
		ActivationType  string `json:"activationtype"`
		ActivationValue string `json:"activationvalue"`
	}

	bs, err := json.Marshal(reminder)
	if err != nil {
		e := models.WrapError("could not marshal json", err)
		return NotificationResponse{}, e
	}

	res, err := c.client.Post(
		c.notifierUrl+"/notify",
		"application/json",
		bytes.NewReader(bs),
	)

	if err != nil {
		e := models.WrapError("notifier service is not available", err)
		return NotificationResponse{}, e
	}

	err = json.NewDecoder(res.Body).Decode(&notifierResponse)
	if err != nil && err != io.EOF {
		e := models.WrapError("could not decode notifier response", err)
		return NotificationResponse{}, e
	}

	t := notifierResponse.ActivationType
	v := notifierResponse.ActivationValue

	if t == "closed" {
		return NotificationResponse{
			completed: true,
		}, nil
	}
	d, err := time.ParseDuration(v)
	if err != nil && d != 0 {
		e := models.WrapError("could not parse notifier duration", err)
		return NotificationResponse{}, e
	}

	if d == 0 {
		return NotificationResponse{}, errors.New("notification duration must be > 0s")
	}

	return NotificationResponse{
		duration: d,
	}, nil
}
