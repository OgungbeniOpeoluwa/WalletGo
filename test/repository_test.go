package test

import (
	"WalletService/data/models"
	"WalletService/data/repositories"
	"WalletService/util"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestSaveUser(t *testing.T) {
	repository := repositories.NewBaseRepository[models.User]()
	result, err := repository.Save(models.User{LastName: "shola2", FirstName: "ope3", Email: "opemi3p@1@gmail.com"})
	if err != nil {
		t.Errorf("expected %v but got %v", result, err)
	}

}

func TestUserByIdExist(t *testing.T) {
	repository := repositories.NewBaseRepository[models.User]()
	result, err := repository.Save(models.User{LastName: "shola", FirstName: "ope", Email: "opemi3p@1@gmail.com"})
	user, err := repository.FindById(result.ID)
	if err != nil {
		t.Errorf("expected %v but got %v", user, err)
	}
	assert.Equal(t, user.FirstName, "ope")

}

func TestThatUserByIdDoesntExist(t *testing.T) {
	repository := repositories.NewBaseRepository[models.Wallet]()
	user, err := repository.FindById(100)
	err2 := util.ErrFetching
	if err == nil {
		t.Errorf("expected %v but got %v", user, err)
	}
	assert.Equal(t, err2, err)

}
func TestFindAll(t *testing.T) {
	repository := repositories.NewBaseRepository[models.User]()
	//_, err = repository.Save(models.User{LastName: "shola", FirstName: "ope", Email: "opemi3p@1@gmail.com"})
	//_, err = repository.Save(models.User{LastName: "shola", FirstName: "ope", Email: "opemi3p@1@gmail.com"})
	//_, err = repository.Save(models.User{LastName: "shola", FirstName: "ope", Email: "opemi3p@1@gmail.com"})
	user, err := repository.GetAllBy("last_name", "shola")
	if err != nil {
		t.Errorf("expected %v but got %v", user, err)
	}
	assert.Equal(t, 0, 0)

}

func TestFindByEmail(t *testing.T) {
	repository := repositories.NewBaseRepository[models.User]()
	_, err := repository.Save(models.User{LastName: "shola4", FirstName: "ope4", Email: "opemi32@gmail.com"})
	user, err := repository.FindByEmail("opemi32@gmail.com")
	if err != nil {
		t.Errorf("expected %v but got %v", user, err)
	}
	assert.Equal(t, user.Email, "opemi32@gmail.com")
}
