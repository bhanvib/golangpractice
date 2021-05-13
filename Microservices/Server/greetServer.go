package main
import (
	"context"
	"fmt"
	"log"
	"net"
	"project.com/proto/greetpb"
	"google.golang.org/grpc"
)
func (server Server) Greet(context context.Context, request *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	//GrretRequest:Greeting:FirstName,LastName
	//take data from reuest
	fname := request.GetGreeting().GetFirstName()
	lastName := request.GetGreeting().GetLastName()
	fullName := fname + "" + lastName
	//Send Response
	res := &greetpb.GreetResponse{
		Result: "Hi " + fullName + " How r u ",
	}
	return res, nil
}
type Server struct {
}
func main() {
	list, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal("Failed to listen", err)
	}
	//create grpc server
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &Server{}) //second para : object which is implementing server interface
	//func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer)
	if err := s.Serve(list); err != nil {
		fmt.Println(err)
	}
}
