package test

import (
	"fmt"
	"log"
	"path/filepath"
	"testing"
)

func TestOther(t *testing.T) {
	files, err := filepath.Glob("*")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(files)
}
