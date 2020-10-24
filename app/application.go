package app

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/emadghaffari/grpc_oauth_service/controllers"
	"github.com/emadghaffari/grpc_oauth_service/databases/cassandra"
	"github.com/emadghaffari/grpc_oauth_service/databases/protos/accesstokenpb"
	"github.com/emadghaffari/res_errors/errors"
)

const(
	certFile = "ssl/server.crt"
	keyFile = "ssl/server.pem"
)

// StartAplication func
// starter for application
func StartAplication()  {
	// if go code crashed...
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	
	// start and defer to close cassandra sesstion
	sesstion := cassandra.GetSesstion()
	defer sesstion.Close()

	lis,err := net.Listen("tcp",":50051")
	if err != nil {
		fmt.Println(errors.HandlerInternalServerError(fmt.Sprintf("Error in listen to tcp network: %v", err),err))
		return
	}
	defer lis.Close()

	// certs,err := credentials.NewClientTLSFromFile(certFile,keyFile)
	// if err != nil {
	// 	fmt.Println(errors.HandlerInternalServerError(fmt.Sprintf("Error in credential Client TLS from File: %v", err),err))
	// 	return
	// }
	
	// server := grpc.NewServer(grpc.Creds(certs))
	// defer server.Stop()
	
	server := grpc.NewServer()

	accesstokenpb.RegisterAccessTokenServer(server, controllers.ServerAccessToken)
	reflection.Register(server)

	go func ()  {
		if err := server.Serve(lis); err != nil{
			fmt.Println(errors.HandlerInternalServerError(fmt.Sprintf("Error in serve the server connection: %v", err),err))
			return
		}
	}()

	wait := make(chan os.Signal, 1)
	signal.Notify(wait,os.Interrupt)
	<- wait

	fmt.Println(errors.HandlerInternalServerError("server stoped.",nil))
}