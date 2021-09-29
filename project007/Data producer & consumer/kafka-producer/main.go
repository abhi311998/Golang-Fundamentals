package main

import (
    "fmt"
	"log"
	"github.com/confluentinc/confluent-kafka-go/kafka"
    "github.com/golang/protobuf/proto"

    "github.com/abhi311998/txnDataGen"
)

const (
	broker = "localhost:9093"
	group = "my-group"
	topics = "transaction"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	// Produce messages to topic (asynchronously)
	topic := topics
	for {
        message := txnDataGen.GenTxnData(1000)
        fmt.Println("Original message: ", message)

		dataSent := []byte(encodeTxnData(message))
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          dataSent,
		}, nil)
		// time.Sleep(time.Second)
		fmt.Println(dataSent)
	}
	// Wait for message deliveries before shutting down
	p.Flush(5 * 1000)
}

func encodeTxnData(txn txnDataGen.Transaction) []uint8 {
    data := txnDataGen.Transaction {
        TimeStamp: txn.TimeStamp,
        UserName: txn.UserName,
        BankName: txn.BankName,
        UpiId: txn.UpiId,
        TxnId: txn.TxnId,
        TxnAmount: txn.TxnAmount,
    }

    res, err := proto.Marshal(&data)
    if err != nil {
        log.Fatal("Marshaling error: ", err)
    }
    return res
}
















