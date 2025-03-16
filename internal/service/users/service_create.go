package users

import (
	"context"
	"diana-quiz/internal/model"
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (s service) Create(ctx context.Context, req *model.CreateUserReq) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "can't generate password hash")
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(req.Password))
	if err != nil {
		return errors.Wrap(err, "can't compare generated hash with password")
	}

	rows, err := s.db.Query(ctx, fmt.Sprintf("insert into users (username, email, full_name, password_hash) values ('%s', '%s', '%s', '%s')", req.Username, req.Email, req.FullName, string(hash)))
	if err != nil {
		return errors.Wrap(err, "can't insert into db")
	}
	rows.Close()

	return nil
}
