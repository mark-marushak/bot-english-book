package account

type accountService struct {
	accountRepo AccountRepository
}

func NewAccountRepository(repository AccountRepository) AccountService {
	return &accountService{
		repository,
	}
}

func (this accountService) GetID() int {
	return this.accountRepo.GetID()
}
