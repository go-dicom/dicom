package dicom_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-dicom/dicom"
)

func TestParse(t *testing.T) {
	f, err := os.Open("./samples/D0001.dcm")
	if err != nil {
		t.Logf("error: %v", err)
		t.FailNow()
	}
	d, err := dicom.Parse(f)
	if err != nil {
		t.Logf("error: %v", err)
		t.FailNow()
	}
	fmt.Println(d)
	t.FailNow()
}
