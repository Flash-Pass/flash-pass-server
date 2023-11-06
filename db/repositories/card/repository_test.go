package card

import (
	"errors"
	"github.com/Flash-Pass/flash-pass-server/db/model"
	CardRepositoryMocks "github.com/Flash-Pass/flash-pass-server/db/repositories/card/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	card := model.NewCard("1", "2", "3", "4")

	t.Run("create success", func(t *testing.T) {
		mockCard := CardRepositoryMocks.NewMockICard(ctrl)
		mockCard.EXPECT().Create(card).Return(nil)

		err := mockCard.Create(card)
		require.Nil(t, err)
	})

	t.Run("create defeat", func(t *testing.T) {
		mockCard := CardRepositoryMocks.NewMockICard(ctrl)
		mockCard.EXPECT().Create(card).Return(errors.New("error"))

		require.NotNil(t, mockCard.Create(card))
	})
}
