package utils

import (
	"archive/tar"
	"bytes"
	"io"

	"github.com/sirupsen/logrus"
)

func Unpack(tarFile []byte) (files [][]byte, err error) {
	reader := bytes.NewReader(tarFile)
	tarReader := tar.NewReader(reader)
	for {
		hdr, err := tarReader.Next()
		logrus.Infof("unpack file: %s\n", hdr.Name)
		if err == io.EOF {
			break
		}
		if err != nil {
			logrus.Warnln(err)
			return files, err
		}
		tmp := []byte{}
		_, err = tarReader.Read(tmp)
		if err != nil {
			logrus.Warnln(err)
			return files, err
		}
		files = append(files, tmp)
	}
	return files, nil
}
