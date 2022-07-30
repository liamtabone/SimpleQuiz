package Main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	protoc "https://github.com/liamtabone/SimpleQuiz/proto"

	//simpleQuiz "SimpleQuiz.pb.go"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

type Answer struct {
	id     string `json:"id"`
	Answer string `json:"answer"`
}

type Questions struct {
	Question  string   `json:"question"`
	Answers   []Answer `json:"answers"`
	CorrectId string   `json:"correctId"`
}

var (
	rootCmd = &cobra.Command{
		Use:   "example",
		Short: "An example cobra program",
		Long:  "This is a looooooong description of the example",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from the root command")
		},
	}
)

func main() {
	file, _ := ioutil.ReadFile("questions.json")

	data := []Questions{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data); i++ {
		fmt.Println("Product Id: ", data[i].Question)
	}

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen to port 9000: %v", err)
	}

	s := protoc.Server{}
	grpcServer := grpc.NewServer()

	protoc.RegisterSimpleQuizService(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Hello")
	}

}
