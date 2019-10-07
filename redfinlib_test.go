package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func loadData(num int) (string, error) {
	dat, err := ioutil.ReadFile(fmt.Sprintf("testdata/%v.html", num))
	if err != nil {
		return "", err
	}

	return string(dat), err
}

func TestGetCurrentPrice(t *testing.T) {
	data, err := loadData(706712)
	if err != nil {
		t.Fatalf("Can't even load the data: %v", err)
	}

	stats, err := extract(data)
	if stats.CurrentPrice != 799000 {
		t.Errorf("Price has been read incorrectly: %v", stats)
	}
}

func TestMain(t *testing.T) {
	main()
}
