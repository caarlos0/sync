package erronce

import (
	"errors"
	"testing"
)

func TestErrOnce(t *testing.T) {
	var once ErrOnce
	eerr := errors.New("fake err")
	if err := once.Do(func() error {
		return eerr
	}); err != eerr {
		t.Errorf("expected %v, got %v", eerr, err)
	}

	if err := once.Do(func() error {
		t.Errorf("should not have been executed again")
		return nil
	}); err != eerr {
		t.Errorf("expected %v, got %v", eerr, err)
	}
}
