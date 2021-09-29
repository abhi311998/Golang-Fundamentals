package main

import (
	"fmt"
	"os"
	"log"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/protobuf/proto"

	"github.com/abhi311998/txnDataGen"
	"github.com/abhi311998/pushToDB"
)

const (
	broker = "localhost:9093"
	group = "my-group"
)

func main() {
	topic := []string{"transaction"}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"broker.address.family": "v4",
		"group.id":              group,
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "earliest",
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Consumer: %v\n", c)
	err = c.SubscribeTopics(topic, nil)

	run := true
	for run {
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
				case *kafka.Message:
					if e.Headers != nil {
						fmt.Printf("%% Headers: %v\n", e.Headers)
					}
					decodedData := *decodeTxnData(e.Value)
					fmt.Print(decodedData)
					fmt.Print(" :: >> ")
					n, err := pushToDB.AddTransaction(decodedData)
					fmt.Println(n, err)
					
				case kafka.Error:
					fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
					if e.Code() == kafka.ErrAllBrokersDown {
						run = false
					}
				default:
					fmt.Printf("Ignored %v\n", e)
			}
	}

	fmt.Printf("Closing consumer\n")
	c.Close()
}

func decodeTxnData(encodedData []uint8) (*txnDataGen.Transaction) {
    res := &txnDataGen.Transaction{}
    err := proto.Unmarshal(encodedData, res)
    if err != nil {
        log.Fatal("Unmarshaling error: ", err)
    }
	// Convert to SQL suited data
    return res
}






// func main() {
// 	fmt.Println("Hello World.")
// 	txn := Transaction{"2021-08-28 04:20:12", "3IOUGWHsUYk8x9MDgldM", "user24@YES", "user24", "YES Bank", 52172}
// 	n, err := AddTransaction(txn)
// 	fmt.Println(n, err)
// }