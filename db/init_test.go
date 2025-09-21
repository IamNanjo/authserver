package db

import (
	"testing"
)

func TestConnection(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("%+v", err)
		}
	}()

	Q()
}

func TestTransaction(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("%+v", err)
		}
	}()

	Tx()
}
