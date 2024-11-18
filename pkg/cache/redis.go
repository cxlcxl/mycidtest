package cache

func NewRedis(host, pass string, port, db int) {
	//client := redis.NewClient(&redis.Options{
	//	Addr:               fmt.Sprintf("%s:%d", host, port),
	//	Dialer:             nil,
	//	OnConnect:          nil,
	//	Password:           pass,
	//	DB:                 db,
	//	MaxRetries:         0,
	//	MinRetryBackoff:    0,
	//	MaxRetryBackoff:    0,
	//	DialTimeout:        0,
	//	ReadTimeout:        0,
	//	WriteTimeout:       0,
	//	PoolFIFO:           false,
	//	PoolSize:           0,
	//	MinIdleConns:       0,
	//	MaxConnAge:         0,
	//	PoolTimeout:        0,
	//	IdleTimeout:        0,
	//	IdleCheckFrequency: 0,
	//	TLSConfig:          nil,
	//	Limiter:            nil,
	//})
}
