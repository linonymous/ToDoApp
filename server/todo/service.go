package todo

import (
	context "context"
)

type Service struct {
	todos []*ToDo
}

func NewService() Service {
	return Service{}
}

func (s *Service) Create(ctx context.Context, do *ToDo) (*Response, error) {
	s.todos = append(s.todos, do)
	return &Response{
		StatusCode: 0,
		Message:    "Created",
	}, nil
}

func (s Service) List(ctx context.Context, request *ListRequest) (*ToDoList, error) {
	return &ToDoList{
		TodoList: s.todos,
	}, nil
}

func (s Service) mustEmbedUnimplementedToDoServiceServer() {
	//TODO implement me
	panic("implement me")
}
