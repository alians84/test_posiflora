package listener

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"strings"
	"test_posiflora/app/service/taskService"
	"time"
)

func ItemTaskListener(con *amqp.Connection) {
	ch, err := con.Channel()
	service := taskService.TaskService{}
	if err != nil {
		log.Error(err.Error())
	}

	msg, err := ch.Consume(
		"NewTask",
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Error(err.Error())
	}

	for true {
		for d := range msg {
			// 1 emailSender 2 description
			fmt.Printf("Received a message: %s", d.Body)
			task := strings.Split(string(d.Body), " ")
			err := service.UpdateMessage(task[1])

			if err != nil {
				log.Error(err.Error())
			}

			if err == nil {
				err = d.Ack(true)
				if err != nil {
					log.Error(err)
				}
			}
		}

		time.Sleep(time.Second * 5)

	}

	err = ch.Close()
	if err != nil {
		log.Error(err.Error())
	}
}
