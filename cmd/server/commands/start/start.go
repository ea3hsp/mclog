package start

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"bitbucket.org/celsa/hydra/pkg/pb"
	"github.com/apex/log"
	"github.com/sclevine/agouti/api"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

var Command = &cli.Command{
	Name:  "start",
	Usage: "start service",
	Action: func(cCtx *cli.Context) error {
		// create router
		svc := api.New(store, log)
		apiAddr := cCtx.String("api-addr")
		apiPort := cCtx.String("api-port")
		// create url string
		url := fmt.Sprintf("%s:%s", apiAddr, apiPort)
		lis, err := net.Listen("tcp", url)
		if err != nil {
			log.Error("listener failed", err, logger.LogFields{
				"schema": "tcp",
				"url":    url,
			})
			return err
		}
		// grpc server options
		var opts []grpc.ServerOption
		// create grpc server
		grpcServer := grpc.NewServer(opts...)
		// register service
		pb.RegisterHydraServer(grpcServer, svc)
		go func() {
			err = grpcServer.Serve(lis)
			if err != nil {
				log.Error("grpc start failed", err, logger.LogFields{
					"schema": "tcp",
					"url":    url,
				})
				return
			}
		}()
		log.Info("service started", logger.LogFields{
			"app":     cCtx.App.Name,
			"address": apiAddr,
			"port":    apiPort,
		})
		// wait signal to finish
		signal := WaitSignal()
		log.Warn("service stopped", logger.LogFields{
			"received signal": signal.String(),
		})
		// return
		return nil
	},
}

// WaitSignal catching exit signal
func WaitSignal() os.Signal {
	ch := make(chan os.Signal, 2)
	signal.Notify(
		ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
			return sig
		}
	}
}
