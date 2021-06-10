package main

import (
	"encoding/binary"
	"io"
	"os"
)

type database struct {
	table *table
}

const (
	pageSize    int64 = 4096
	rowSize     int64 = 291
	rowsPerPage int64 = pageSize / rowSize
	maxPages    int64 = 100
)

type row struct {
	id    uint32
	name  string
	email string
}

type page []byte

type pager struct {
	*os.File
	fileSize int64
	pages    []page
}

type table struct {
	rows  int64
	pager *pager
}

func (r *row) serialize(wr io.Writer) error {
	if r == nil {
		return nil
	}

	var nameArr [32]byte
	var emailArr [255]byte

	copy(nameArr[:], r.name)
	copy(emailArr[:], r.email)

	if err := binary.Write(wr, binary.BigEndian, r.id); err != nil {
		return err
	}

	if err := binary.Write(wr, binary.BigEndian, nameArr); err != nil {
		return err
	}

	if err := binary.Write(wr, binary.BigEndian, emailArr); err != nil {
		return err
	}

	return nil
}

func (r *row) deserialize(re io.Reader) error {
	if r == nil {
		return nil
	}

	var nameArr [32]byte
	var emailArr [255]byte

	if err := binary.Read(re, binary.BigEndian, &r.id); err != nil {
		return err
	}

	if err := binary.Read(re, binary.BigEndian, &nameArr); err != nil {
		return err
	}

	if err := binary.Read(re, binary.BigEndian, &emailArr); err != nil {
		return err
	}

	r.name = string(nameArr[:])
	r.email = string(emailArr[:])

	return nil
}

func newPager(f string) (*pager, error) {
	file, err := os.OpenFile(f, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	stats, _ := file.Stat()
	fileSize := stats.Size()

	p := pager{
		File:     file,
		fileSize: fileSize,
		pages:    make([]page, maxPages, maxPages),
	}

	return &p, nil
}
