package keyringservice

import (
	"errors"
	"time"

	"github.com/zalando/go-keyring"
)

type KeyringServiceInterface interface {
	Set(service, user, secret string) error
	Get(service, user string) (string, error)
	Delete(service, user string) error
}

type KeyringService struct{}

var ErrNotFound = errors.New("secret not found in keyring")

type TimeoutError struct {
	message string
}

func (e *TimeoutError) Error() string {
	return e.message
}

func NewKeyringService() *KeyringService {
	return &KeyringService{}
}

// Set secret in keyring for user.
func (ks *KeyringService) Set(service, user, secret string) error {
	ch := make(chan error, 1)
	go func() {
		defer close(ch)
		ch <- keyring.Set(service, user, secret)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(3 * time.Second):
		return &TimeoutError{"timeout while trying to set secret in keyring"}
	}
}

// Get secret from keyring given service and user name.
func (ks *KeyringService) Get(service, user string) (string, error) {
	ch := make(chan struct {
		val string
		err error
	}, 1)
	go func() {
		defer close(ch)
		val, err := keyring.Get(service, user)
		ch <- struct {
			val string
			err error
		}{val, err}
	}()
	select {
	case res := <-ch:
		if errors.Is(res.err, keyring.ErrNotFound) {
			return "", ErrNotFound
		}
		return res.val, res.err
	case <-time.After(3 * time.Second):
		return "", &TimeoutError{"timeout while trying to get secret from keyring"}
	}
}

// Delete secret from keyring.
func (ks *KeyringService) Delete(service, user string) error {
	ch := make(chan error, 1)
	go func() {
		defer close(ch)
		ch <- keyring.Delete(service, user)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(3 * time.Second):
		return &TimeoutError{"timeout while trying to delete secret from keyring"}
	}
}
