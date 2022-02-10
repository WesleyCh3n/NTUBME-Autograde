package utils

import (
	"archive/tar"
	"io"
	"log"
	"os"
)

func TarFile(dest string, src []string) error {
	dist, err := os.Create(dest)
	if err != nil {
		log.Fatalln(err)
	}
	defer dist.Close()

	tw := tar.NewWriter(dist)
	defer tw.Close()

	for _, file := range src {
		f, err := os.Open(file)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()

		info, err := f.Stat()
		if err != nil {
			log.Fatalln(err)
		}

		hdr, err := tar.FileInfoHeader(info, "")
		if err != nil {
			log.Fatalln(err)
		}

		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}

		_, err = io.Copy(tw, f)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return nil
}
