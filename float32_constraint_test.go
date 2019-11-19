package govalid_test

import (
	"testing"

	"github.com/twharmon/govalid"
)

func init() {
	govalid.Register(f32{}, f32Req{}, f32ReqMin{}, f32Max{})
}

type f32 struct {
	F32 float32
}

type f32Req struct {
	F32 float32 `govalid:"req"`
}

type f32ReqMin struct {
	F32 float32 `govalid:"req|min:5.5"`
}

type f32Max struct {
	F32 float32 `govalid:"max:5.5"`
}

func TestFloat32(t *testing.T) {
	assertNoViolation(t, "no validation rules with empty field", &f32{})
	assertNoViolation(t, "no validation rules with non-empty field", &f32{5})

	assertViolation(t, "`req` with empty field", &f32Req{})
	assertNoViolation(t, "`req` with non-empty field", &f32Req{5.5})

	assertViolation(t, "`req|min` with empty field", &f32ReqMin{})
	assertViolation(t, "`req|min` with field too less", &f32ReqMin{3.5})
	assertNoViolation(t, "`req|min` with valid field", &f32ReqMin{5.5})

	assertNoViolation(t, "`max` with empty field", &f32Max{})
	assertViolation(t, "`max` with field too great", &f32Max{7.5})
	assertNoViolation(t, "`max` with valid field", &f32Max{3.5})
}
