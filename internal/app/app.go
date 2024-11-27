package app

import (
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
	"github.com/valek177/chat-server/internal/config"
	"github.com/valek177/chat-server/internal/interceptor"
	"github.com/valek177/chat-server/internal/tracing"
	"github.com/valek177/platform-common/pkg/closer"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

// App contains application object
type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

// NewApp creates new App object
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Run runs application
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initJaegerTracing,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	flag.Parse()

	err := config.Load(configPath)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initJaegerTracing(_ context.Context) error {
	cfg, err := a.serviceProvider.JaegerConfig()
	if err != nil {
		return err
	}

	return tracing.Init(cfg)
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	grpcCfg, err := a.serviceProvider.GRPCConfig()
	if err != nil {
		return err
	}

	creds, err := credentials.NewServerTLSFromFile(grpcCfg.TLSCertFile(), grpcCfg.TLSKeyFile())
	if err != nil {
		return err
	}

	client, err := a.serviceProvider.AuthClient(ctx)
	if err != nil {
		return err
	}

	auth := interceptor.NewAuthInterceptor(client)

	a.grpcServer = grpc.NewServer(
		grpc.Creds(creds),
		grpc.ChainUnaryInterceptor(
			interceptor.ServerTracingInterceptor,
			auth.Interceptor(ctx),
		),
	)

	reflection.Register(a.grpcServer)

	chatImpl, err := a.serviceProvider.ChatImpl(ctx)
	if err != nil {
		return err
	}

	chat_v1.RegisterChatV1Server(a.grpcServer, chatImpl)

	return nil
}

func (a *App) runGRPCServer() error {
	grpcConfig, err := a.serviceProvider.GRPCConfig()
	if err != nil {
		return err
	}
	log.Printf("GRPC server is running on %s", grpcConfig.Address())

	list, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
