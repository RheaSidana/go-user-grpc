package users

import (
	pb "gRPC-go/proto"
	"gRPC-go/server/db"
)

var DB db.Db

func InitDB() {
	DB.Users = make(map[int32]*pb.User)
	DB.Users[1] = &pb.User{
		Id:      1,
		Fname:   "Steve",
		City:    "LA",
		Phone:   1234567890,
		Height:  5.8,
		Married: true,
	}

	DB.Users[2] = &pb.User{
		Id:      2,
		Fname:   "Alice",
		City:    "New York",
		Phone:   1234567890,
		Height:  5.6,
		Married: true,
	}

	// invalid user: 
	DB.Users[3] = &pb.User{
		Id:      3,
		Fname:   "Bob",
		City:    "San Francisco",
		Phone:   9876543210,
		Height:  0.5,
		Married: false,
	}

	DB.Users[4] = &pb.User{
		Id:      4,
		Fname:   "Charlie",
		City:    "Los Angeles",
		Phone:   5555555555,
		Height:  5.10,
		Married: true,
	}
	DB.Users[5] = &pb.User{
		Id:      5,
		Fname:   "David",
		City:    "Chicago",
		Phone:   1112223333,
		Height:  5.8,
		Married: false,
	}

}
