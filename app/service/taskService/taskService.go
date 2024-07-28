package taskService

import (
	log "github.com/sirupsen/logrus"
	"test_posiflora/app/models"
	"test_posiflora/config"
)

type TaskService struct {
}

func (t *TaskService) CreateNewMessage(task models.Task) error {
	db := config.Db

	err := db.Create(&task).Error
	if err != nil {
		log.Error(err.Error())
	}

	return nil
}

func (t *TaskService) GetAllTask() ([]models.Task, error) {
	db := config.Db
	task := []models.Task{}

	err := db.Find(&task).Error
	if err != nil {
		log.Error()
	}
	return task, nil
}

func (t *TaskService) UpdateMessage(EmailSender string) error {
	db := config.Db
	task := models.Task{}
	task.EmailSender = EmailSender
	task.IsViewed = true

	err := db.Where("email_sender = ? and is_viewed = ?", EmailSender, false).First(&task).Error
	if err != nil {
		log.Error(err.Error())
	}

	err = db.Model(task).Select("email_sender", "is_viewed").Update("is_viewed", true).Where("id = ?", task.ID).Find(&task).Error
	if err != nil {
		log.Error(err.Error())
	}

	return nil
}
