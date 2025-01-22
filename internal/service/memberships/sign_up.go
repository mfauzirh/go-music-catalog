package memberships

import (
	"database/sql"
	"errors"

	"github.com/mfauzirh/go-music-catalog/internal/models/memberships"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(req memberships.SignUpRequest) error {
	existingUser, err := s.repository.GetUser(req.Email, req.Username, 0)
	if err != nil || !errors.Is(err, sql.ErrNoRows) {
		log.Error().Err(err).Msg("failed get user from database")
		return err
	}

	if existingUser != nil {
		return errors.New("email or username is already exists")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("failed hashing password")
		return err
	}

	model := memberships.User{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(pass),
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}

	return s.repository.CreateUser(model)
}
