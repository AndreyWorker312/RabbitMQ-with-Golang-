package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // имя очереди
		false,   // не сохранять сообщения после перезагрузки
		false,   // очередь не удаляется, если нет потребителей
		false,   // доступ только для одного подключения
		false,   // без ожидания
		nil,     // дополнительные аргументы отсутствуют
	)
	if err != nil {
		log.Fatal(err)
	}

	body := "Hello RabbitMQ!"
	err = ch.Publish(
		"",     // обменник не указан, используется стандартный
		q.Name, // имя очереди
		false,  // не обязательное сообщение
		false,  // нет переадресации сообщений
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(" [x] Sent %s", body)
}
