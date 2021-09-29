module github.com/abhi311998/kafka_consumer

go 1.13

replace github.com/abhi311998/pushToDB => ../pushToDB

require (
	github.com/abhi311998/pushToDB v0.0.0-00010101000000-000000000000
	github.com/abhi311998/txnDataGen v0.0.0-00010101000000-000000000000
	github.com/confluentinc/confluent-kafka-go v1.7.0
	github.com/golang/protobuf v1.5.0
)

replace github.com/abhi311998/txnDataGen => ../txnDataGen
