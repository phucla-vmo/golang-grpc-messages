package main

import (
	"context"
	"github.com/golang-grpc-messages/messagepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

func main() {
	cc, err := grpc.Dial("localhost:50070", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("err while dial %v", err)
	}
	defer cc.Close()
	client := messagepb.NewMessageServiceClient(cc)
	//go CallCombineMessage(client)
	//CallSplitMessageIntoWords(client)
	CallMakeParagraphByWords(client)
	//CallAverage(client)
	//time.Sleep(3 * time.Second)
}

func CallCombineMessage(c messagepb.MessageServiceClient) {
	log.Println("CombineMessage calling...")
	res, err := c.CombineMessage(context.Background(), &messagepb.MessageRequest{
		Message1: "Hello",
		Message2: "World",
	})

	if err != nil {
		log.Fatalf("CallCombineMessage err %v", err)
	}
	log.Printf("CombineMessage Response: %v", res.GetResult())
}

func CallSplitMessageIntoWords(c messagepb.MessageServiceClient) {
	log.Println("SplitMessageIntoWords calling...")
	stream, err := c.SplitMessageIntoWords(context.Background(), &messagepb.SplitMessage{
		Message: "Just the way you are",
	})
	if err != nil {
		log.Fatalf("CallSplitMessageIntoWords err %v", err)
	}
	for {
		res, receiverErr := stream.Recv()
		if receiverErr == io.EOF {
			log.Println("Server finish streaming")
			return
		}
		log.Printf("Word: %v", res.Word)
	}
}

func CallMakeParagraphByWords(c messagepb.MessageServiceClient) {
	log.Println("CallMakeParagraphByWords calling...")
	stream, err := c.MakeParagraphByWord(context.Background())
	if err != nil {
		log.Fatalf("CallMakeParagraphByWords err %v", err)
	}
	listReq := []messagepb.ParagraphRequest{
		{
			Word: "Hello",
		},
		{
			Word: "from",
		},
		{
			Word: "the other",
		},
		{
			Word: "side",
		},
	}
	for _, req := range listReq {
		err := stream.Send(&req)
		if err != nil {
			log.Fatalf("Send CallMakeParagraphByWords request err %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Receive CallMakeParagraphByWords response error %v", err)
	}
	log.Println("Paragraph: ", res)
}
