package model

type UserKnowledge struct {
	UserID  uint `db:"user_id" gorm:"primaryKey"`
	WordID  uint `db:"word_id" gorm:"primaryKey"`
	Learned float32
	Attempt int
	Success int
}

type userKnowledgeService struct {
	repo UserKnowledgeRepository
}

type UserKnowledgeRepository interface {
	UpdateLearned(knowledge UserKnowledge) error
	GetLearned(knowledge UserKnowledge) (float32, error)
	StoreAttempt(knowledge UserKnowledge) error
	StoreSuccess(knowledge UserKnowledge) error
	Create(knowledge UserKnowledge) (UserKnowledge, error)
	Get(knowledge UserKnowledge) (UserKnowledge, error)
	GetUserKnowledge(userID uint) ([]UserKnowledge, error)
}

type UserKnowledgeService interface {
	UpdateLearned(knowledge UserKnowledge) error
	GetLearned(knowledge UserKnowledge) (float32, error)
	StoreAttempt(knowledge UserKnowledge) error
	StoreSuccess(knowledge UserKnowledge) error
	Create(knowledge UserKnowledge) (UserKnowledge, error)
	Get(knowledge UserKnowledge) (UserKnowledge, error)
	GetUserKnowledge(userID uint) ([]UserKnowledge, error)
}

func NewUserKnowledgeService(repo UserKnowledgeRepository) UserKnowledgeService {
	return &userKnowledgeService{repo: repo}
}

func (u userKnowledgeService) UpdateLearned(knowledge UserKnowledge) error {
	return u.repo.UpdateLearned(knowledge)
}

func (u userKnowledgeService) GetLearned(knowledge UserKnowledge) (float32, error) {
	return u.repo.GetLearned(knowledge)
}

func (u userKnowledgeService) StoreAttempt(knowledge UserKnowledge) error {
	return u.repo.StoreAttempt(knowledge)
}

func (u userKnowledgeService) StoreSuccess(knowledge UserKnowledge) error {
	return u.repo.StoreSuccess(knowledge)
}

func (u userKnowledgeService) Create(knowledge UserKnowledge) (UserKnowledge, error) {
	return u.repo.Create(knowledge)
}

func (u userKnowledgeService) Get(knowledge UserKnowledge) (UserKnowledge, error) {
	return u.repo.Get(knowledge)
}

func (u userKnowledgeService) GetUserKnowledge(userID uint) ([]UserKnowledge, error) {
	return u.repo.GetUserKnowledge(userID)
}
