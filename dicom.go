package dicom

import (
	"fmt"
	"io"
	"strings"
)

type Tag struct {
	GroupNumber   uint16
	ElementNumber uint16
}

type Element struct {
	Tag Tag
}

type Dicom struct {
	Preamble [128]byte
	Prefix   [4]byte
	Elements []*Element
}

func Parse(src io.Reader) (*Dicom, error) {
	preamble := make([]byte, 128)
	c, err := io.ReadAtLeast(src, preamble, 128)
	if err != nil {
		return nil, err
	}
	printCountAndValue("Preable", c, preamble)

	prefix := make([]byte, 4)
	c, err = io.ReadAtLeast(src, prefix, 4)
	printCountAndValue("Prefix", c, prefix)
	return nil, nil
}

func printCountAndValue(ident string, c int, val []byte) {
	lengthOfDecorator := 30
	fmt.Println(strings.Repeat("#", lengthOfDecorator))
	fmt.Printf("%s %s %s\n", strings.Repeat("#", ((lengthOfDecorator/2)-len(ident))), ident, strings.Repeat("#", ((lengthOfDecorator/2)-len(ident))))
	fmt.Println(strings.Repeat("#", lengthOfDecorator))
	fmt.Printf("Count: %d\n", c)
	strVal := strings.TrimSpace(string(val))
	if strVal == " " {
		fmt.Printf("Value: ****is empty****\n")
		return
	}
	fmt.Printf("Value: %s\n", strVal)
	fmt.Println(strings.Repeat("#", lengthOfDecorator))
	fmt.Printf("\n\n")
}
