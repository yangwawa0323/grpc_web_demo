package service

import (
	"context"
	"log"
	"strings"

	gorm "github.com/jinzhu/gorm"

	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/go-sql-driver/mysql"

	pb "github.com/yangwawa0323/grpc_web_demo/pb/user/v1"
)

type UserServiceServer struct {
	DB *gorm.DB
	pb.UnimplementedUserSearchServiceServer
}

func migrate(db *gorm.DB, models ...interface{}) {
	db.AutoMigrate(models...)
}

func NewUserServiceServer() *UserServiceServer {
	// Due to CGO_ENABLED=1 go-sqlite3 has build problem, we change to MySQL database
	dsn := "root:example@tcp(db:3306)/example?charset=utf8mb4&parseTime=True&loc=Local"

	// db, err := gorm.Open("sqlite3", "users.db")
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		log.Fatal("Cannot open the MySQL database.")
	}

	// migration
	migrate(db, &pb.UserORM{})

	return &UserServiceServer{DB: db}
}

func (us *UserServiceServer) SearchByName(ctx context.Context,
	req *pb.SearchByNameRequest) (*pb.SearchByNameResponse, error) {
	name := strings.TrimSpace(req.Username)
	log.Printf("[DEBUG]: user: %s", name)
	in := &pb.User{
		Username: name,
	}

	var user pb.User
	result := us.DB.Where(in).First(&user)

	return &pb.SearchByNameResponse{User: &user},
		result.Error
}

func (us *UserServiceServer) SearchByGender(ctx context.Context,
	req *pb.SearchByGenderRequest) (*pb.SearchByGenderResponse, error) {
	gender := req.Gender
	in := &pb.User{
		Gender: gender,
	}

	// should use pointer ?
	var users []*pb.User = make([]*pb.User, 0)

	result := us.DB.Where(in).Find(&users)

	response := &pb.SearchByGenderResponse{
		Users: users,
	}
	log.Printf("[DEBUG]: gender: %s , found %d\n", gender, result.RowsAffected)
	return response, nil
}
