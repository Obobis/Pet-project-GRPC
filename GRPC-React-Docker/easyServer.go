package main

import (
	"GRPC-React-Docker/app/proto"
	"context"
	"log"
	"net"
	"time"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	goproto "google.golang.org/protobuf/proto"
)

type server struct {
	proto.UnimplementedUsrServer
}

var (
	cacheDB *redis.Client
	TTL     = 30 * time.Minute
)

func (*server) GetUser(ctx context.Context, in *proto.UserRequest) (*proto.UserResponse, error) {
	cacheKey := "user" + in.GetName() //Ключ в redis.

	data, err := cacheDB.Get(ctx, cacheKey).Bytes()
	//Проверяем есть ли ключ.
	if err == nil {
		user := &proto.User{}
		if err := goproto.Unmarshal(data, user); err == nil {
			return &proto.UserResponse{
				User:   user,
				Status: 200,
				Error:  "",
			}, nil
		} else {
			log.Printf("Failed to load from cache: %v", err)
		}
	}
	//Если нет, создаем и далее кидаем в кэш

	others := make(map[string]string)
	others["secondary"] = "123456"
	phone := &proto.PhoneNumber{Primary: "0123456789", Others: others}
	user := &proto.User{Name: "LevDemchenko", Age: 23,
		Adress: &proto.Address{Street: "Pune", City: "Pune",
			State: "MAHARASHTRA", Zip: "201223"}, Phone: phone,
		UpdatedAt: time.Now().UTC().String(),
		CreatedAt: time.Now().UTC().String()}

	//Создаем горутину для асинхронного добавления в БД.

	go func() {
		data, err := goproto.Marshal(user)
		if err != nil {
			log.Printf("Failed to marshal user: %v", err)
			return
		}
		if err := cacheDB.Set(context.Background(), cacheKey, data, TTL).Err(); err != nil {
			log.Printf("Failed to load user to cache: %v", err)
		}
	}()

	return &proto.UserResponse{User: user, Status: 200, Error: ""}, nil

}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	grpcServer := &server{}
	proto.RegisterUsrServer(s, grpcServer)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
