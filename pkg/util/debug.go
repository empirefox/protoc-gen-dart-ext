package util

import (
	"encoding/json"
	"os"
)

func DumpJsonToSdtout(v interface{}) error {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(b)
	if err != nil {
		_, err = os.Stdout.Write([]byte{'\n'})
	}
	return err
}
