package RabbitMQ

import (
	"MiniIM/configs"
	"fmt"

	"github.com/streadway/amqp"
)

var exchangeName string

var ch *amqp.Channel

func Init() error {
	rabbitmqConfig := configs.GetConfig().RabbitMQ

	addr := rabbitmqConfig.Addr
	user := rabbitmqConfig.User
	password := rabbitmqConfig.Password
	exchangeName = rabbitmqConfig.ExchangeName
	exchangeType := rabbitmqConfig.ExchangeType
	// 连接 RabbitMQ
	url := fmt.Sprintf("amqp://%s:%s@%s/", user, password, addr)
	conn, err := amqp.Dial(url)
	if err != nil {
		return err
	}

	// 创建通道
	ch, err = conn.Channel()
	if err != nil {
		return err
	}

	// 声明交换机
	err = ch.ExchangeDeclare(
		exchangeName, // 交换机名称
		exchangeType, // 交换机类型
		true,         // 持久性
		false,        // 自动删除
		false,        // 内部
		false,        // 无等待
		nil,          // 参数
	)
	if err != nil {
		return err
	}

	return nil
}

func NewQueueAndBind(UserUuid string, groupUuids []string) (*amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		UserUuid, // 根据用户ID创建唯一队列名
		false,    // 非持久性
		true,     // 自动删除（没有消费者时自动删除队列）
		false,    // 非排他性
		false,    // 不等待队列声明完成
		nil,      // 无额外参数
	)
	if err != nil {
		return nil, err
	}

	// 将队列绑定到多个 Routing Key
	for _, key := range groupUuids {
		err = ch.QueueBind(
			q.Name,       // 队列名称
			key,          // Routing Key
			exchangeName, // 交换机名称
			false,        // 不等待绑定完成
			nil,          // 无额外参数
		)
		if err != nil {
			return nil, err
		}
	}

	return &q, nil
}

func NewConsume(qName string) (<-chan amqp.Delivery, error) {
	msgs, err := ch.Consume(
		qName, // 队列名称
		"",    // 消费者标识，空字符串表示由 RabbitMQ 生成
		true,  // 自动确认
		false, // 排他性
		false, // 不等待服务器响应
		false, // 无额外参数
		nil,
	)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func Publish(groupUuid string, msg []byte) error {
	err := ch.Publish(
		exchangeName, // 交换机名称
		groupUuid,    // Routing Key
		false,
		false,
		amqp.Publishing{
			ContentType: "application/octet-stream",
			Body:        msg,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
