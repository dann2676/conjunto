package domain

import (
	"fmt"
	"net/http"
)

type DError struct {
	Status int
	Msg    string
	Entity string
	Type   string
}

func (e DError) Error() string {
	return fmt.Sprintf("error %s: %s", e.Entity, e.Msg)
}

func (e DError) Is(target error) bool {
	t, ok := target.(DError)
	if !ok {
		return false
	}
	return e.Msg == t.Msg
}

func NewDErr(msg, t string, s int) func(e string) DError {
	return func(e string) DError {
		return DError{Entity: e, Status: s, Msg: msg, Type: t}
	}
}

var (
	NotFounErr    = NewDErr("no encontrado", "warning", http.StatusConflict)
	SavingErr     = NewDErr("no se pudo guardar", "danger", http.StatusInternalServerError)
	DuplicatedErr = NewDErr("ya existe", "warning", http.StatusConflict)
	DeletingErr   = NewDErr("no se pudo borrar", "danger", http.StatusInternalServerError)
	AsociationErr = NewDErr("tiene información asociada", "warning", http.StatusConflict)
)
