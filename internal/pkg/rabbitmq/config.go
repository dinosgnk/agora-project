package rabbitmq

type RabbitMQConfig struct {
	Host     string `env:"RABBITMQ_HOST" envDefault:"localhost"`
	Port     string `env:"RABBITMQ_PORT" envDefault:"5672"`
	User     string `env:"RABBITMQ_USER" envDefault:"guest"`
	Password string `env:"RABBITMQ_PASS" envDefault:"guest"`
}
