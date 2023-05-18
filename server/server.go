package main

import (
	"context"
	"fmt"
	"github.com/golang-grpc-messages/messagepb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

type server struct {
	messagepb.MessageServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50070")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()

	messagepb.RegisterMessageServiceServer(s, &server{})
	fmt.Println("server connecting...")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}

func (*server) CombineMessage(ctx context.Context, req *messagepb.MessageRequest) (*messagepb.CombineMessageResponse, error) {
	log.Println("CombineMessage is running...")
	res := messagepb.CombineMessageResponse{
		Result: req.GetMessage1() + " " + req.GetMessage2(),
	}
	return &res, nil
}

func (*server) SplitMessageIntoWords(req *messagepb.SplitMessage, stream messagepb.MessageService_SplitMessageIntoWordsServer) error {
	log.Println("SplitMessageIntoWords is running...")
	message := req.GetMessage()
	messageArr := strings.Fields(message)
	for _, word := range messageArr {
		err := stream.Send(&messagepb.ResponseWords{
			Word: word,
		})
		if err != nil {
			log.Fatalf("Send to client err: %v", err)
		}
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

func (*server) MakeParagraphByWord(stream messagepb.MessageService_MakeParagraphByWordServer) error {
	log.Println("CombineMessageToParagraph is running...")
	var combineMessage string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			res := messagepb.ParagraphResponse{
				Result: combineMessage,
			}
			return stream.SendAndClose(&res)
		}
		if err != nil {
			log.Fatalf("CombineMessageToParagraph call err: %v", err)
		}
		log.Println("received word: ", req)
		combineMessage += req.GetWord() + " "
		fmt.Println("combineMessage", combineMessage)
	}
}
