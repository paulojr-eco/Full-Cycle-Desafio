package cmd

import (
	"github.com/paulojr-eco/full-cycle-desafio/application/grpc"
	"github.com/paulojr-eco/full-cycle-desafio/infra/db"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "full-cycle-desafio",
	Short: "This software was created in the event Imersão Full Cycle and it makes the management of products that can be created and listed",
}

func Execute() {
	database, _ := db.ConnectDB()
	grpc.StartGrpcServer(database, 50051)
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Exec 'go run main.go' to start the server")
}
