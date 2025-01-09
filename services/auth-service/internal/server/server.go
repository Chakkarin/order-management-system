package server

import (
	"net/http"
	"services/auth-service/internal/config"
	dbconn "services/auth-service/internal/infrastructure/db_conn"
	mqconn "services/auth-service/internal/infrastructure/mq_conn"
	redisconn "services/auth-service/internal/infrastructure/redis_conn"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type (
	ginServer struct {
		app          *gin.Engine
		grpcServer   *grpc.Server
		dbconn       *gorm.DB
		redisConn    *redis.Client
		rabbitMqConn *amqp.Channel
		cfg          *config.Config
	}
)

var (
	server *ginServer
	once   sync.Once
)

func NewServer(cfg *config.Config) *ginServer {

	gin.SetMode(gin.ReleaseMode)

	ginApp := gin.Default()
	once.Do(func() {
		server = &ginServer{
			app:          ginApp,
			grpcServer:   grpc.NewServer(),
			cfg:          cfg,
			dbconn:       dbconn.ConnectDB(&cfg.PgDatabase),
			redisConn:    redisconn.ConnectRedis(&cfg.Redis),
			rabbitMqConn: mqconn.ConnectMQ(&cfg.RabbitMq),
		}
	})

	return server
}

func (s *ginServer) Start() {

	s.app.GET("/v1/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})

	// start domain
	s.initAuthenRouter()

	go s.startHttp()

	s.startGrpc()

}
