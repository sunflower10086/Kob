package ioc

import "github.com/gin-gonic/gin"

// IOC 容器层

type Service interface {
	Config()
	Name() string
}

type HTTPHandler interface {
	RouteRegistry(g *gin.Engine)
	Config() error
	Name() string
}

//  1. HostService的实例必须注册过来，HostService才会有具体的实例

// HostService HTTP模块依赖于IOC中的HostService
var (
	services    = map[string]Service{}
	httpHandler = map[string]HTTPHandler{}
)

// RegistryImpl 注册服务到Ioc中心
func RegistryImpl(service Service) {
	if _, ok := services[service.Name()]; ok {
		panic("Service is registered")
	}

	// 维护起来
	services[service.Name()] = service
}

func GetImpl(name string) interface{} {
	for s, service := range services {
		if s == name {
			return service
		}
	}

	return nil
}

// InitImpl Ioc初始化所有自己实现的服务
func InitImpl() {
	for _, service := range services {
		service.Config()
	}
}

// RegistryGin 注册HTTP handler
func RegistryGin(hdl HTTPHandler) {
	if _, ok := httpHandler[hdl.Name()]; ok {
		panic("Handler is registered")
	}

	// 维护起来
	httpHandler[hdl.Name()] = hdl
}

func LoadedGinHandler() (names []string) {
	for k, _ := range httpHandler {
		names = append(names, k)
	}
	return names
}

// InitGinHandler 初始化所有的 HTTP handler
func InitGinHandler(g *gin.Engine) error {
	// 初始化所有对象
	for _, handler := range httpHandler {
		if err := handler.Config(); err != nil {
			return err
		}
	}

	// 注册它的所有路由
	for _, handler := range httpHandler {
		handler.RouteRegistry(g)
	}
	return nil
}
