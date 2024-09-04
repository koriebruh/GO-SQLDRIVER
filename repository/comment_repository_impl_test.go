package repository_test

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	go_mysql "go-mysql"
	"go-mysql/entity"
	"go-mysql/repository"
	"testing"
)

func TestCommentRepositoryImpl_Insert(t *testing.T) {
	commentRepository := repository.NewCommentRepository(go_mysql.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "Irohana@gmail.com",
		Comment: "test Irohana",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

func TestCommentRepositoryImpl_FindById(t *testing.T) {
	commentRepository := repository.NewCommentRepository(go_mysql.GetConnection())
	ctx := context.Background()
	id := 24

	result, err := commentRepository.FindById(ctx, int32(id))
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

func TestCommentRepositoryImpl_FindAll(t *testing.T) {
	commentRepository := repository.NewCommentRepository(go_mysql.GetConnection())
	ctx := context.Background()

	result, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, value := range result {
		fmt.Println(value.Id)
		fmt.Println(value.Email)
		fmt.Println(value.Comment)
		fmt.Println("")
	}

}
