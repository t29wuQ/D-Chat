package test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kindai-csg/D-Chat/domain"
	mock "github.com/kindai-csg/D-Chat/test/mock_usecase"
	"github.com/kindai-csg/D-Chat/usecase"
)

func TestMessageInteractorCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	messageRepository := mock.NewMockMessageRepository(ctrl)
	interactor := usecase.NewMessageInteractor(messageRepository)

	message := domain.Message{}
	messageRepository.EXPECT().Create(message).Return(message, errors.New(""))

	_, err := interactor.Create(message)
	if err == nil {
		t.Errorf("Expectation: return error")
	}

	message = domain.Message{UserId: "testid", Body: "body", CreatedAt: "11:20"}
	messageRepository.EXPECT().Create(message).Return(message, nil)
	m, err := interactor.Create(message)
	if err != nil {
		t.Errorf("Expectation: return nil")
	}
	if m != message {
		t.Errorf("Expectation: return nil")
	}
}
