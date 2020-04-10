package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"msgsub/filtercmpt/vendor/github.com/confluentinc/confluent-kafka-go/kafka"
)

// MsgSub struct
type MsgSub struct {
	URLType string `json:"type"`
	URL     string `json:"url"`
	Task    string `json:"task"`
	Tags    string `json:"tags"`
	Content string `json:"content"`
}

func main() {

	// saveMsgSub("text/html", "https://www.baidu.com", "$..name", "Hello,World", "Hello world, yes")

	// _, err := setCache("hello", "world")
	// if err != nil {
	// 	panic(err)
	// }

	// val2, err := getCache("ssss")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%v", val2)

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    "kafka:9092",
		"group.id":             "filtercmpt",
		"default.topic.config": kafka.ConfigMap{"auto.offset.reset": "earliest"}})

	if err != nil {
		panic(err)
	}

	err = consumer.SubscribeTopics([]string{"msgsub-filter"}, nil)
	if err != nil {
		panic(err)
	}

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	topic := "msgsub-push"
	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)

	var msgSub MsgSub

	run := true

	for run == true {
		ev := consumer.Poll(1000)
		switch e := ev.(type) {
		case *kafka.Message:
			json.Unmarshal(e.Value, &msgSub)
			fmt.Printf("%v \n", msgSub)

			key := msgSub.url + msgSub.task

			val, err = getCache(key)
			if err != nil {
				panic(err)
			}

			if val != nil {
				cacheTagList = strings.Split(val, ",")
				tagList = strings.Split(msgSub.tags)

				isSame = true
				for i, x := range cacheTagList {
					isFound = false
					for j, y := range tagList {
						if x == y {
							isFound = true
							break
						}
					}
					if !isFound {
						isSame = false
						break
					}
				}

				if isSame {
					continue
				}
			}

			producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          e.Value,
			}, deliveryChan)

			e := <-deliveryChan
			m := e.(*kafka.Message)

			if m.TopicPartition.Error != nil {
				panic(m.TopicPartition.Error)
			}

		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false

		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}
}
