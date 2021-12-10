package main

import (
	"context"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	pb "github.com/leslesnoa/chatbox-grpc/pb"
	"github.com/leslesnoa/chatbox-grpc/server/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// gRPCサーバ
type server struct {
}

const (
	port = ":9090"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatBoxServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %c", err)
	}
	log.Println("starting on gRPC server on " + port)
}

// -------------------------------------------
func (s *server) SendChat(ctx context.Context, r *pb.ChatData) (*pb.Empty, error) {
	log.Printf("Recieved SendChatRequest : %s", r)
	log.Println(r.GetName())
	log.Println(r.GetText())
	// DB接続
	conn := db.Connect()
	db.CreateRow(conn, r)
	return &pb.Empty{}, nil
}

func (s *server) GetChat(ctx context.Context, r *pb.Empty) (*pb.ChatList, error) {
	log.Printf("Recieved GetChatRequest : %s", r)

	// DB接続
	conn := db.Connect()

	// 接続確認
	err := conn.Ping()
	if err != nil {
		log.Println("connection failed.")
		panic(err)
	} else {
		log.Println("connection success.")
	}

	// 行データ取得
	rows := db.GetRows(conn)
	log.Println(rows)
	// var result []*pb.ChatData
	// for _, p := range rows {
	// 	log.Println(p.Id, p.Name, p.Text)
	// 	result = append(result, p)
	// }
	defer conn.Close()
	return &pb.ChatList{ChatList: rows}, nil

	// return &pb.ChatList{ChatList: []*pb.ChatData{
	// 	{
	// 		Id:   1,
	// 		Name: "user1",
	// 		Text: "ChatText1",
	// 	},
	// 	{
	// 		Id:   2,
	// 		Name: "user2",
	// 		Text: "ChatText2",
	// 	},
	// 	{
	// 		Id:   3,
	// 		Name: "user3",
	// 		Text: "ChatText3",
	// 	}},
	// }, nil
}
