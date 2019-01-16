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
	f, err := cio.OpenFile("JUVE00 EGRR 161200", "r")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	file, err := codes.OpenFile(f)
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
	fmt.Println(msg)
	if err := msg.SetLong("unpack", 1); err != nil {
		return errors.Wrap(err, "unpack")
	}

	if err := printField(msg, "dataCategory", "l"); err != nil {
		return err
	}
	if err := printField(msg, "typicalDate", "l"); err != nil {
		return err
	}
	// if err := printField(msg, "stationNumber", "l"); err != nil {
	// 	return err
	// }
	// if err := printField(msg, "airTemperatureAt2M", "d"); err != nil {
	// 	return err
	// }
	descriptors, err := native.Ccodes_get_long_array(msg.Handle(), "bufrdcExpandedDescriptors")
	if err != nil {
		return err
	}
	fmt.Println("descriptors:", descriptors)
	values, err := native.Ccodes_get_double_array(msg.Handle(), "numericValues")
	if err != nil {
		return err
	}
	fmt.Println("values:", values)

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
	}
	if err != nil {
		return errors.Wrapf(err, "field: %s", name)
	}
	fmt.Printf("%s: %v\n", name, v)
	return nil
}
