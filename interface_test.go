package gomocktst_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	mockData "github.com/nawa/gomocktst/mock_data"
)

func Test_Method(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockData.NewMockInterface(mockCtrl)

}
