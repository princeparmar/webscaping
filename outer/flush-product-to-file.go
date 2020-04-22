package outer

import (
	"encoding/json"
	"os"
)

func SaveFile(baseDir string, data interface{}) error {

	f, err := os.Create(baseDir)
	if err != nil {
		return err
	}

	return json.NewEncoder(f).Encode(data)

}
