package authentication

import "sinarmas/kredit-sinarmas/utils/authToken"

type Service interface {
	Login(req DataRequest) (DataResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) Login(req DataRequest) (DataResponse, error) {
	if err := req.Validation(); err != nil {
		return DataResponse{}, err
	}

	user, err := s.repo.Login(req)
	if err != nil {
		return DataResponse{}, err
	}

	token, err := authToken.GenerateToken(user)
	if err != nil {
		return DataResponse{}, err
	}

	res := DataResponse{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Token:    "Bearer " + token,
	}

	return res, nil
}
