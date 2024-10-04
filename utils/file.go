package utils

import (
	"base-api/constants"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func DeleteDownloadedDocument(filename string) error {
	err := os.Remove(fmt.Sprintf("%s%s", constants.TempDownloadedFileDir, filename))
	if err != nil {
		log.Errorln(err)
		return err
	}
	return nil
}
