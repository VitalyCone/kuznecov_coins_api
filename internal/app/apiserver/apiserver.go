package apiserver

import (
	"github.com/VitalyCone/kuznecov_coins_api/internal/app/apiserver/endpoints"
	"github.com/VitalyCone/kuznecov_coins_api/internal/app/store"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/VitalyCone/kuznecov_coins_api/docs"
	// _ "github.com/swaggo/swag/example/basic/docs"
)

var (
	mainPath string = "/usercoins"
)

type APIServer struct {
	config *Config
	router *gin.Engine
	store  *store.Store
}

func NewAPIServer(config *Config, store *store.Store) *APIServer {
	return &APIServer{
		config: config,
		router: gin.Default(),
		store:  store,
	}
}

func (s *APIServer) Start() error {

	s.configureEndpoints()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.router.MaxMultipartMemory = 8 << 20

	return s.router.Run(s.config.ApiAddr)
}

func (s *APIServer) configureEndpoints() {
	endpoint := endpoints.NewEndpoints(s.store)
	
	s.router.GET("/", endpoint.Ping) 
	docs.SwaggerInfo.BasePath = mainPath
	path := s.router.Group(mainPath)
	{
		path.GET("/user", endpoint.GetCoin)
		path.PUT("/user", endpoint.PutCoin)
		path.POST("/user", endpoint.PostNewUser)
	}
	// path.GET("/publication/:id", endpoint.GetPublication) 
	// path.GET("/news", endpoint.GetNews) 
	// path.POST("/publication", endpoint.PostPublication) 
	// path.DELETE("/publication/:id", endpoint.DeletePublication)
	// {
	// 	path.GET("/news") //получение новостей определенного пользователя

	// 	path.POST("/post") //создание поста
	// 	path.DELETE("/post/:id") //удаление поста
	// 	path.PUT("/post/:id") //изменение поста по id
	// }


	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (s *APIServer) configureStore() error{
	if err:= s.store.Open(); err != nil{
		return err
	}

	return nil
}