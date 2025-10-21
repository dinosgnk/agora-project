module github.com/dinosgnk/agora-project/internal/services/notification

go 1.24.3

toolchain go1.24.4

replace github.com/dinosgnk/agora-project/internal/pkg => ../../pkg

require github.com/dinosgnk/agora-project/internal/pkg v0.0.0-00010101000000-000000000000

require (
	github.com/caarlos0/env/v11 v11.3.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/rabbitmq/amqp091-go v1.10.0 // indirect
)
