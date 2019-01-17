package main

import (
	"fmt"
	"io"
	"log"

	"github.com/mtfelian/go-eccodes/native"

	"github.com/amsokol/go-errors"
	codes "github.com/mtfelian/go-eccodes"
	cio "github.com/mtfelian/go-eccodes/io"
)

func main() {
	fmt.Println("Start")
	f, err := cio.OpenFile("JUBE99 EGRR 301200", "r")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	file, err := codes.OpenFile(f, native.ProductBUFR)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	n := 0
	for {
		err = process(file, n)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("failed to get message (#%d) from index: %s", n, err.Error())
		}
		n++
	}
}

func process(file codes.File, n int) error {
	msg, err := file.Next()
	if err != nil {
		return err
	}
	defer msg.Close()
	if err := msg.SetLong("unpack", 1); err != nil {
		return errors.Wrap(err, "unpack")
	}
	printHeader(msg)

	descriptors, err := native.Ccodes_get_long_array(msg.Handle(), "bufrdcExpandedDescriptors")
	if err != nil {
		return err
	}
	fmt.Println("---------------------------------------------")
	fmt.Println("descriptors:", descriptors)
	values, err := native.Ccodes_get_double_array(msg.Handle(), "numericValues")
	if err != nil {
		return err
	}
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("values:", values)
	fmt.Println("---------------------------------------------")

	iter := native.Ccodes_bufr_keys_iterator_new(msg.Handle(), 0)
	if iter == nil {
		return errors.New("iter is null")
	}
	defer native.Ccodes_bufr_keys_iterator_delete(iter)
	for native.Ccodes_bufr_keys_iterator_next(iter) {
		name := native.Ccodes_bufr_keys_iterator_get_name(iter)
		printField(msg, name, "s")
		fmt.Printf("  ")
		printField(msg, name+"->units", "s")
		fmt.Printf("  ")
		printField(msg, name+"->code", "s")
		fmt.Printf("  ")
		printField(msg, name+"->width", "s")
	}
	fmt.Println("<<<<<<<<<<<<<<<<<<:::::::>>>>>>>>>>>>>>>>>>>")
	return nil
}

func printField(msg codes.Message, name, typ string) error {
	var v interface{}
	var err error
	switch typ {
	case "l":
		v, err = msg.GetLong(name)
	case "d":
		v, err = msg.GetDouble(name)
	case "s":
		v, err = msg.GetString(name)
	}
	if err != nil {
		return errors.Wrapf(err, "field: %s", name)
	}
	fmt.Printf("%s: %v\n", name, v)
	return nil
}

func printHeader(msg codes.Message) {
	printField(msg, "edition", "l")
	printField(msg, "masterTableNumber", "l")
	printField(msg, "dataCategory", "l")
	printField(msg, "dataSubCategory", "l")
	printField(msg, "typicalDate", "l")
	printField(msg, "typicalTime", "l")
	printField(msg, "bufrHeaderCentre", "l")
	printField(msg, "bufrHeaderSubCentre", "l")
	printField(msg, "masterTablesVersionNumber", "l")
	printField(msg, "localTablesVersionNumber", "l")
	printField(msg, "numberOfSubsets", "l")

}
