package api

type UserHadler struct {
	UserDO types.
}

func NewUserHandler(UserDO db.UserStore) *UserHadler {
	return &UserHadler{
		UserDO: UserDO,
	}
}
