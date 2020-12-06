package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	root := rootDir()
	clientDir := root + "/Client/Source/Marsettler/Message"
	serverDir := root + "/Server/Source/Message"

	filepath.Walk(root+"/Message", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		if ext != ".fbs" {
			return nil
		}

		output, err := exec.Command(flatExe(), "--cpp", "-o", clientDir, path).CombinedOutput()
		if err != nil {
			log.Fatalln(string(output))
		}

		messageName := strings.TrimSuffix(filepath.Base(path), ext)

		err = copyFile(clientDir+"/"+messageName+"_generated.h", serverDir+"/"+messageName+"_generated.h")
		if err != nil {
			return nil
		}

		clinetBuilderFile := clientDir + "/" + messageName + "Builder_generated"

		clientBuilderHeaderFile := clinetBuilderFile + ".h"
		if _, err := os.Stat(clientBuilderHeaderFile); os.IsNotExist(err) {
			log.Println("Header is not exist: " + clientBuilderHeaderFile)
		}

		clientBuilderSourceFile := clinetBuilderFile + ".h"
		if _, err := os.Stat(clientBuilderSourceFile); os.IsNotExist(err) {
			log.Println("Source is not exist: " + clientBuilderSourceFile)
		}

		return nil
	})
}

func copyFile(srcPath, dstPath string) error {
	log.Println(srcPath)
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

func rootDir() string {
	root, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	root = filepath.Dir(root)
	root = filepath.Dir(root)

	return root
}

func flatExe() string {
	root, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	return root + "/flatc"
}
