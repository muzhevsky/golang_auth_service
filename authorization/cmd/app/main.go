package main

import "authorization/internal/app"

func main() {
	app.Run()
}

//func main() {
//	lis, err := net.Listen("tcp", ":50051")
//	if err != nil {
//		log.Fatal(err)
//	}
//	s := grpc.NewServer()
//	proto.RegisterAuthServer(s, &ssov1.Server{})
//	err = s.Serve(lis)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
