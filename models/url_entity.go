package models

import (
	"bytes"
	"net/url"
)

type Url struct {
	ID       string `json:"id" db:"id"`
	Original string `json:"original" db:"original"`
	Short    string `json:"short,omitempty" db:"short"`
}

type UrlStore interface {
	Create(u *Url) error
	Read(id string) (Url, error)
	Update(u *Url) error
	Delete(id string) error
}

type Errors map[string]string

type UrlForm struct {
	Url
	Err Errors
}

func (f *UrlForm) Validate() bool {
	f.Err = Errors{}
	_, err := url.ParseRequestURI(f.Url.Original)
	if err != nil {
		f.Err["original"] = "please enter a valid original link"
	}
	if f.Url.Short != "" && !ShortValidate(f.Url.Short) {
		f.Err["short"] = "please enter a valid short link"
	}

	return len(f.Err) == 0
}

func ShortValidate(s string) bool {
	sToLower := bytes.ToLower([]byte(s))
	var i int
	for _, ch := range sToLower {
		if ('a' <= ch && ch <= 'z') || ('0' <= ch && ch <= '9') || (ch == '-') {
			i++
		} else {
			return false
		}
	}
	return true
}
