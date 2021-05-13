package main
import (
	"context"
	"fmt"
	"project.com/proto/calculatorpb"
	"google.golang.org/grpc"
)
func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}
	c := calculatorpb.NewCalculatorServiceClient(conn)
	fmt.Println("client is created")
	req := &calculatorpb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 20,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.GetSumResult())
}