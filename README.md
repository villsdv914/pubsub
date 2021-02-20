# Pubsub
***

Please Install rabbitmq and let it run with default configuration.
## Packages
Please get all the packages related to project to run it properly.

* sqlite (gorm.io/driver/sqlite)
* logrus (github.com/sirupsen/logrus)
* ampq  (github.com/streadway/amqp)
* gorm (gorm.io/gorm)
## Run
```bash
go build main.go
```


It will create Related table in eastern.db and insert all data after inserting data it will be stopped.
