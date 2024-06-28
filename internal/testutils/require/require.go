package require

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func New(t *testing.T) *Require {
	return &Require{
		t: t,
	}
}

type Require struct {
	t *testing.T
}

func (r *Require) Equal(expected, actual any, msgAndArgs ...any) {
	r.t.Helper()
	Equal(r.t, expected, actual, msgAndArgs...)
}

func (r *Require) NoError(err error, msgAndArgs ...any) {
	r.t.Helper()
	NoError(r.t, err, msgAndArgs...)
}

func (r *Require) Error(err error, msgAndArgs ...any) {
	r.t.Helper()
	Error(r.t, err, msgAndArgs...)
}

func (r *Require) ErrorIs(expected, actual error, msgAndArgs ...any) {
	r.t.Helper()
	ErrorIs(r.t, expected, actual, msgAndArgs...)
}

func (r *Require) NotNil(a any, msgAndArgs ...any) {
	r.t.Helper()
	NotNil(r.t, a, msgAndArgs...)
}

func (r *Require) True(b bool, msgAndArgs ...any) {
	r.t.Helper()
	True(r.t, b, msgAndArgs...)
}

func (r *Require) False(b bool, msgAndArgs ...any) {
	r.t.Helper()
	False(r.t, b, msgAndArgs...)
}

func (r *Require) GreaterOrEqual(expected, actual any, msgAndArgs ...any) {
	r.t.Helper()
	GreaterOrEqual(r.t, expected, actual, msgAndArgs...)
}
func (r *Require) Greater(expected, actual any, msgAndArgs ...any) {
	r.t.Helper()
	Greater(r.t, expected, actual, msgAndArgs...)
}
func (r *Require) LessOrEqual(expected, actual any, msgAndArgs ...any) {
	r.t.Helper()
	LessOrEqual(r.t, expected, actual, msgAndArgs...)
}
func (r *Require) Less(expected, actual any, msgAndArgs ...any) {
	r.t.Helper()
	Less(r.t, expected, actual, msgAndArgs...)
}

func (r *Require) Zero(a any, msgAndArgs ...any) {
	r.t.Helper()
	Zero(r.t, a, msgAndArgs...)
}

func (r *Require) NotZero(a any, msgAndArgs ...any) {
	r.t.Helper()
	NotZero(r.t, a, msgAndArgs...)
}

func Equal(t *testing.T, expected, actual any, msgAndArgs ...any) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		FailNow(t, fmt.Sprintf("expected: %v, got: %v", expected, actual), msgAndArgs...)
	}
}

func NoError(t *testing.T, err error, msgAndArgs ...any) {
	t.Helper()
	if err != nil {
		FailNow(t, fmt.Sprintf("expected no error, got: %v", err), msgAndArgs...)
	}
}

func Error(t *testing.T, err error, msgAndArgs ...any) {
	t.Helper()
	if err != nil {
		return
	}
	FailNow(t, "expected error, got nil", msgAndArgs...)
}

func ErrorIs(t *testing.T, expected, actual error, msgAndArgs ...any) {
	t.Helper()
	if errors.Is(actual, expected) {
		return
	}
	FailNow(t, fmt.Sprintf("expected error: %v, got: %v", expected, actual), msgAndArgs...)
}

func NotNil(t *testing.T, a any, msgAndArgs ...any) {
	t.Helper()
	if a != nil {
		return
	}
	FailNow(t, "expected not nil, got nil", msgAndArgs...)
}

func True(t *testing.T, b bool, msgAndArgs ...any) {
	t.Helper()
	if b {
		return
	}
	FailNow(t, "expected true, got false", msgAndArgs...)
}

func False(t *testing.T, b bool, msgAndArgs ...any) {
	t.Helper()
	if !b {
		return
	}
	FailNow(t, "expected false, got true", msgAndArgs...)
}

func Zero(t *testing.T, a any, msgAndArgs ...any) {
	t.Helper()
	if reflect.ValueOf(a).IsZero() {
		return
	}
	FailNow(t, fmt.Sprintf("expected zero value, got %v", a), msgAndArgs...)
}

func NotZero(t *testing.T, a any, msgAndArgs ...any) {
	t.Helper()
	if !reflect.ValueOf(a).IsZero() {
		return
	}
	FailNow(t, "expected not zero value, got zero", msgAndArgs...)
}
