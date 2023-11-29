# Kafka-go
Apache Kafka is a open source & distributed data store that optimized for ingesting and processing streaming data in real time. Kafka operates based on a publish-subscribe model, where data is organized into topics. Producers publish data records to specific topics, and consumers subscribe to these topics to receive and process the data.

Overall, Kafka is primarily used to build real-time streaming data pipelines and applications that adapt to data streams. It combines messaging, storage, and stream processing to store historical and real-time data.

# Connet to Kafka
let’s get started to run Kafka in go programming language, first of all here’s the step that you need to do:
- There are multiple ways to connect go with kafka
  1. Download & instal kafka in local then connect with go application.
  2. Run kafka instances in docker then connect with go application.
- Here, I will use docker to run kafka cluser up locally.
- There are mutiple docker images are availabe in docker hub like
  1. bitnami/kafka
  2. ubuntu/kafka
  3. vmware/kafka
  4. rancher/kafka
  5. confluentinc/cp-kafka
- You can directly pull images in docker and good to go.
- I will pull bitmani/kafka.

1. ### Setup Docker Compose
    - Add below code into `docker-compose.yml` in your folder & this file using below command. it will initialize the kafka and zookeeper in your docker container.
    - docker command:
     ```
    docker-compose up -d
    ```
    - `docker-compose.yml` file
    ```
    ersion: "3"
    services:
      zookeeper:
        image: 'bitnami/zookeeper:latest'
        ports:
          - '2181:2181'
        environment:
          - ALLOW_ANONYMOUS_LOGIN=yes
      kafka:
        image: 'bitnami/kafka:2.8.0'
        ports:
          - '9092:9092'
        environment:
          - KAFKA_BROKER_ID=1
          - KAFKA_CFG_LISTENERS=PLAINTEXT://:9092
          - KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092
          - KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181
          - ALLOW_PLAINTEXT_LISTENER=yes
          - KAFKA_CFG_LOG_RETENTION_HOURS=6
        depends_on:
          - zookeeper
    ```
    - once instance are running then it looks like below :
      <img width="1007" alt="Screenshot 2023-11-29 at 8 00 17 PM" src="https://github.com/Sufiyan33/Kafka-go/assets/133976350/300b5491-741e-474d-9b0d-3243faf895b4">

2. ### Create module
   - initialize your module based on the project struct, you can use below command:
   ```
    go mod init github.com/Sufiyan33/kafka-go
   ```
   -  Your project structure looks like below:
   <img width="304" alt="Screenshot 2023-11-29 at 7 42 41 PM" src="https://github.com/Sufiyan33/Kafka-go/assets/133976350/bf4f24cb-c158-4f0b-874d-2ee0fef18a74">

   - Now download library from github:
     ```
      go get "github.com/IBM/sarama"
       OR
      go get "github.com/Shopify/sarama"
     ```

   - Create config file & mention boostrap.server and topic name like below :
     ```
      package config

      const (
      	CONST_HOST  = "localhost:9092"
      	CONST_TOPIC = "my-topic"
      )
     ```
3. ### Create Producer
   - Create NewSyncProducer or NewAsyncProducer.
   - Producer messages in topics.
4. ### Create Consumer
   - Create new consumer to consume messages produce by producer.
   - Create consumerPartitions.
   - Now consumer messages and display it.
5. ### Run
   - Run consumer
     
     <img width="588" alt="Screenshot 2023-11-29 at 7 52 10 PM" src="https://github.com/Sufiyan33/Kafka-go/assets/133976350/da875a23-07f5-4e55-9052-7e486f47257d">

   - Run producer
     
     <img width="609" alt="Screenshot 2023-11-29 at 7 51 48 PM" src="https://github.com/Sufiyan33/Kafka-go/assets/133976350/79616b01-dca8-466b-8423-dbab73dd9d57">

   - Run in seprate terminal.
  
6. ### Congratulation
   - Now you have connected your go application with Apache kafka.
   - If you liked it then follow me.
    
