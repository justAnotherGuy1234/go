package db

type Storage struct {
	UserRepo UserRepo
}

func NewStore() *Storage {
	return &Storage{
		UserRepo: &UserRepoImpl{},
	}
}
