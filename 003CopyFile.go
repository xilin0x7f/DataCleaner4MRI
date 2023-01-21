package main

/*
@Author: xilin0x7f, https://github.com/xilin0x7f
*/

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyFile2DPABIFormat(funFiles, t1Files [][]string, dstRoot string) {
	for _, subject := range funFiles {
		_ = os.MkdirAll(filepath.Join(dstRoot, "FunImg", subject[1]+subject[2]), 644)
		_ = CopyFile(filepath.Join(dstRoot, "FunImg", subject[1]+subject[2], "Fun.nii"),
			filepath.Join(subject...))
	}
	for _, subject := range t1Files {
		_ = os.MkdirAll(filepath.Join(dstRoot, "T1Img", subject[1]+subject[2]), 644)
		_ = CopyFile(filepath.Join(dstRoot, "T1Img", subject[1]+subject[2], "T1.nii"),
			filepath.Join(subject...))
	}
}

func CopyFile(dst, src string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(fmt.Sprintf("open file %v fail", src), err)
		return err
	}
	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 644)
	if err != nil {
		fmt.Println(fmt.Sprintf("open file %v fail", dst), err)
		return err
	}
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		fmt.Println(fmt.Sprintf("copy file %v to %v fail", src, dst), err)
	}
	return err
}
