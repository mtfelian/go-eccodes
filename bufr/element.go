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

// BufrCodes ...
type BufrCodes struct {
	msg   codes.Message
	items []Element
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

// NewBufrCodes ...
func NewBufrCodes(msg codes.Message) (*BufrCodes, error) {
	if iter == nil {
		return nil, errors.New("nil msg")
	}
	iter := native.Ccodes_bufr_keys_iterator_new(msg.Handle(), 0)
	if iter == nil {
		return nil, errors.New("iter is null")
	}
	defer native.Ccodes_bufr_keys_iterator_delete(msg)

	bufr := new(BufrCodes)

	for native.Ccodes_bufr_keys_iterator_next(msg) {
		if el, err := newElementFromIter(msg, iter); err == nil {
			bufr.items = append(bufr.items, el)
		}
	}
	return bufr, nil
}
