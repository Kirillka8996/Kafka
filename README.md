# Kafka Consumer и Producer на Go

Этот проект демонстрирует работу с Kafka с использованием Go. Он включает в себя Producer, который отправляет сообщения в топик Kafka, и Consumer, который эти сообщения читает.

## Предварительные требования

•   **Docker и Docker Compose:**  Необходимы для запуска Kafka кластера.
•   **Go:** Для компиляции и запуска Go приложений.

## Структура проекта

•   cmd/: Содержит исходный код Consumer и Producer.
    •   cmd/consumer/main.go:  Исходный код Consumer приложения.
    •   cmd/producer/main.go:  Исходный код Producer приложения.
•   docker-compose.yml: Файл конфигурации Docker Compose для запуска кластера Kafka.
•   go.mod, go.sum: Файлы для управления зависимостями Go.

▌Запуск Kafka кластера

Для запуска Kafka кластера используйте Docker Compose. Выполните следующие шаги:

1.  Перейдите в корневой каталог проекта (там, где находится файл docker-compose.yml).
2.  Запустите кластер:

    ```bash
      docker-compose up -d
    ```
    Эта команда запустит Kafka, Zookeeper и необходимые компоненты в фоновом режиме.
3.  Убедитесь, что контейнеры запущены:
    ```bash
       docker-compose ps
    ```

▌Запуск Consumer

Consumer настроен на чтение всех необработанных сообщений из Kafka топика. Выполните следующие шаги для запуска:

1.  Перейдите в каталог Consumer:
    ```bash
    cd cmd/consumer
    ```
2.  Запустите приложение:
    ```bash
    go run .
    ```
    Consumer начнет слушать сообщения из Kafka и выводить их на экран.

▌Запуск Producer

Producer предназначен для отправки сообщений в Kafka. Выполните следующие шаги для запуска:

1.  Перейдите в каталог Producer:
    ```bash
    cd ../producer
    ```
    (если вы находитесь в каталоге consumer)
    или
    ```bash
     cd cmd/producer
    ```
    (если вы находитесь в корневом каталоге проекта)
2.  Запустите приложение:
    ```bash
    go run .
    ```
    Producer начнет отправлять сообщения в Kafka, которые будет читать запущенный Consumer. Вы сможете вводить сообщения с клавиатуры, каждое из которых будет отправлено в Kafka.

▌Конфигурация

•   Топик Kafka: И Consumer, и Producer используют один и тот же топик Kafka, который можно настроить в исходном коде. По умолчанию используется топик `"my-topic"`.
•   Адрес Kafka:  Адрес Kafka брокера по умолчанию `localhost:9092`.  Может быть изменен в исходном коде при необходимости.

▌Зависимости

Для сборки и запуска Go приложений используются следующие библиотеки:

	> github.com/confluentinc/confluent-kafka-go/kafka: Библиотека Kafka для Go.


---

Примечания:

•   Убедитесь, что у вас установлены Go и Docker.
•   go.mod и go.sum файлы могут отличаться в зависимости от того, когда вы их создали и если вы добавляли новые библиотеки.  Вы можете обновить зависимости, выполнив команду go mod tidy.
•   Вы можете изменить настройки Kafka (топик, адрес, порт) в исходном коде main.go в папках consumer и producer.
