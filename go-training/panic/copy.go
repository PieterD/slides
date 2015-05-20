package something

import (
	"io"
	"os"
)

// START OMIT
func Copy(from, to string) error {
	fromHandle, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fromHandle.Close() // HL

	toHandle, err := os.Create(to)
	if err != nil {
		return err
	}
	defer toHandle.Close() // HL

	_, err = io.Copy(toHandle, fromHandle)
	return err
}

// END OMIT
