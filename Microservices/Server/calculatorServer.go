package main
import (
	"context"
	"fmt"
	"log"
	"net"
	"project.com/proto/calculatorpb"
	"google.golang.org/grpc"
)
func (server Server1) Sum(context context.Context, request *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	no1 := request.GetFirstNumber()
	no2 := request.GetSecondNumber()
	result := no1 + no2
	res := &calculatorpb.SumResponse{
		SumResult: result,
	}
	return res, nil
}
type Server1 struct {
}
func main() {
	list, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal("Failed to listen", err)
	}
	//create grpc server
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &Server1{}) //second para : object which is implementing server interface
	//func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer)
	if err := s.Serve(list); err != nil {
		fmt.Println(err)
	}
}