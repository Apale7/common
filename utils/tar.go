package utils

import (
	"archive/tar"
	"bytes"
	"io"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func Unpack(tarFile []byte) (files [][]byte, err error) {
	reader := bytes.NewReader(tarFile)
	tarReader := tar.NewReader(reader)
	for {
		hdr, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			logrus.Warnln(err)
			return files, err
		}
		logrus.Infof("unpack file: %s\n", hdr.Name)
		tmp, err := ioutil.ReadAll(tarReader)
		if err != nil {
			logrus.Warnln(err)
			return files, err
		}
		files = append(files, tmp)
	}
	return files, nil
}
