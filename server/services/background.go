package services

import (
	"log"
	"time"

	"github.com/manzurahammed/rm-cli/server/models"
)

type saver interface {
	save() error
}

//BackgroundSaver  use for
type BackgroundSaver struct {
	ticker *time.Ticker
	service saver
}

func newSaver(s saver) *BackgroundSaver {
	ticker := time.NewTicker(30 * time.Second)
	return &BackgroundSaver{
		ticker:ticker,
		service:s,
	}
}

func (s *BackgroundSaver) start(){
	log.Println("background saver started")

	for {
		select {
		case <-s.ticker.C:
			err := s.service.save()
			if err != nil {
				log.Printf("could not save records in background: %v", err)
			}
		}
	}
}

func (s *BackgroundSaver) stop() error {
	s.ticker.Stop()
	err := s.service.save()
	if err !=nil {
		return err
	}
	log.Println("background saver stopped")
	return nil
}

type HTTPNotifierClient interface {
	Notify(reminder models.Reminder) (NotificationResponse, error)
}

type snapshotManager interface {
	snapshot() Snapshot
	snapshotGromming(notificationReminder ...models.Reminder) 
	retry(reminder models.Reminder, time time.Duration)
}


type BackgroundNotifier struct {
	ticker *time.Ticker
	service snapshotManager
	completed chan models.Reminder
	client HTTPNotifierClient
}

func newNotifier(notifierUrl string, service snapshotManager) *BackgroundNotifier {
	ticker := time.NewTicker(1*time.Second)
	httpClient := NewHTTPClient(notifierURI)
	return &BackgroundNotifier{
		ticker: ticker,
		service: service,
		completed:make(chan models.Reminder),
		client: httpClient,
	}
}

func (s *BackgroundNotifier) start(){
	log.Println("background notifier started")

	for {
		select {
		case <-s.ticker.C:
			snapshot := s.service.snapshot()
			for id := range snapshot.UnCompleted {
				_,reminder := snapshot.UnCompleted.
			}
		case <-s.completed:
			log.Printf("reminder with with: %d was completed\n", r.ID)
		}
	}
}




