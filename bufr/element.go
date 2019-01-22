package bufr

import (
	"errors"

	codes "github.com/mtfelian/go-eccodes"
	"github.com/mtfelian/go-eccodes/native"
)

// Element represents bufr element
type Element struct {
	Name string
	Code string
	Unit string
}

// Codes ...
type Codes struct {
	Msg   codes.Message
	Items []Element
}

func newElementFromIter(msg codes.Message, iter native.Ccodes_keys_iterator) (*Element, error) {
	if iter == nil {
		return nil, errors.New("nil iterator")
	}
	name := native.Ccodes_bufr_keys_iterator_get_name(iter)
	code, err := msg.GetString(name + "->code")
	if err != nil {
		return nil, err
	}
	unit, err := msg.GetString(name + "->units")
	if err != nil {
		return nil, err
	}
	el := &Element{
		Name: name,
		Code: code,
		Unit: unit,
	}
	return el, nil
}

// NewCodes ...
func NewCodes(msg codes.Message) (*Codes, error) {
	if msg == nil {
		return nil, errors.New("nil msg")
	}
	iter := native.Ccodes_bufr_keys_iterator_new(msg.Handle(), 0)
	if iter == nil {
		return nil, errors.New("iter is null")
	}
	defer native.Ccodes_bufr_keys_iterator_delete(iter)

	bufr := new(Codes)

	for native.Ccodes_bufr_keys_iterator_next(iter) {
		if el, err := newElementFromIter(msg, iter); err == nil {
			bufr.items = append(bufr.Items, *el)
		}
	}
	return bufr, nil
}
