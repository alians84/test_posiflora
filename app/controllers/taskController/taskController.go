package taskController

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"test_posiflora/app/models"
	"test_posiflora/app/service/taskService"
	"test_posiflora/config"
	"time"
)

type TaskController struct{}

func NewTaskController() *TaskController { return &TaskController{} }

// CreateOne
// @Summary
// @Description
// @Accept   json
// @Produce  json
// @Tags     Task
// @Param    body  body      CreateOneRequest  true  "body"
// @Success  200   {object}  CreateOneResponse
// @Success  401   {object}  baseModel.ErrorResponse
// @Router   /notifications [post]
func (controller TaskController) CreateOne(c *fiber.Ctx) error {
	var response CreateOneResponse
	taskModel := models.Task{}
	request := &CreateOneRequest{}
	service := taskService.TaskService{}

	err := c.BodyParser(request)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	taskModel.EmailSender = request.Task.EmailSender
	taskModel.RecipientEmail = request.Task.RecipientEmail
	taskModel.Description = request.Task.Description
	taskModel.IsViewed = false
	taskModel.DateCreate = time.Now()
	taskModel.DateUpdate = time.Now()
	_ = service.CreateNewMessage(taskModel)

	// rabbitmq service
	conn := config.RabbitMQ
	ch, err := conn.Channel()
	if err != nil {
		log.Error(err.Error())
	}

	q, err := ch.QueueDeclare(
		"NewTask",
		false,
		false,
		false,
		false,
		nil)

	err = ch.Publish(
		"",
		q.Name, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("%v", request.Task)),
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"task": response.Task,
	})
}

// GetAll
// @Summary
// @Description
// @Accept   json
// @Produce  json
// @Tags     Task
// @Success  200   {object}  CreateOneRequestTask
// @Success  401   {object}  baseModel.ErrorResponse
// @Router   /notifications [get]
func (controller TaskController) GetAll(c *fiber.Ctx) error {
	service := taskService.TaskService{}
	coll, err := service.GetAllTask()
	if err != nil {
		log.Error(err)
	}
	return c.Status(fiber.StatusOK).JSON(coll)
}
