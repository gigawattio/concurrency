package concurrency_test

import (
	"errors"
	"testing"

	"github.com/gigawattio/go-commons/pkg/concurrency"
)

func TestMultiGoNilFuncBehavior(t *testing.T) {
	if err := concurrency.MultiGo(); err != nil {
		t.Error(err)
	}
	if err := concurrency.MultiGo(nil); err != nil {
		t.Error(err)
	}
	if err := concurrency.MultiGo(nil, nil, nil); err != nil {
		t.Error(err)
	}

	var ran bool
	fn := func() error {
		ran = true
		return nil
	}
	if err := concurrency.MultiGo(nil, fn, nil); err != nil {
		t.Error(err)
	}
	if ran != true {
		t.Errorf("Expected ran=true but actual=%v", ran)
	}
}

func TestMultiGoErrorCase(t *testing.T) {
	expectedErr := errors.New("some problematic issue")
	fn := func() error {
		return expectedErr
	}
	if err := concurrency.MultiGo(nil, fn, nil); err != expectedErr {
		t.Errorf("Expected err=%v but instead got err=%v", expectedErr, err)
	}
}
