package main

/*
@Author: xilin0x7f, https://github.com/xilin0x7f
*/

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

func Dcm2niixConcurrent(root string, n int, dcm2niixParameters ...string) {
	var wg sync.WaitGroup
	// 控制并行数量
	ch := make(chan int, n)
	dcmLogFile, _ := os.OpenFile(filepath.Join(root, "dcm2niix.csv"), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	dcmErrFile, _ := os.OpenFile(filepath.Join(root, "dcm2niixErr.txt"), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	_, _ = dcmLogFile.Write([]byte{0xEF, 0xBB, 0xBF})
	_, _ = dcmErrFile.Write([]byte{0xEF, 0xBB, 0xBF})
	csvWriter := csv.NewWriter(dcmLogFile)
	csvWriter.Comma = ','
	defer func() {
		_ = dcmLogFile.Close()
		_ = dcmErrFile.Close()
	}()
	groups, _ := os.ReadDir(root)
	for _, group := range groups {
		subjects, _ := os.ReadDir(filepath.Join(root, group.Name()))
		for _, subject := range subjects {
			fullPath := filepath.Join(root, group.Name(), subject.Name())
			wg.Add(1)
			ch <- 0
			dcm2niixFunParameters := append(dcm2niixParameters, fullPath)
			go Dcm2niix(csvWriter, dcmErrFile, &wg, ch, dcm2niixFunParameters...)
		}
	}
	csvWriter.Flush()
	wg.Wait()
}

func Dcm2niix(logFile *csv.Writer, errFile *os.File, wg *sync.WaitGroup, ch chan int, dcm2niixFunParameters ...string) {
	// -d 最大深度，-x 是否crop, -i ignore der, -z gz ?
	out, err := exec.Command("dcm2niix", dcm2niixFunParameters...).CombinedOutput()
	path := dcm2niixFunParameters[len(dcm2niixFunParameters)-1]
	stringOut := strings.Replace(strings.Replace(string(out), "\r\n", " ", -1), "\n", " ", -1)
	stringOut = strings.Replace(stringOut, ",", "  ", -1)
	stringErr := strings.Replace(strings.Replace(fmt.Sprint(err), "\r\n", " ", -1), "\n", " ", -1)
	stringErr = strings.Replace(stringErr, ",", "  ", -1)
	_ = logFile.Write([]string{path, stringOut, stringErr})
	logFile.Flush()
	if err != nil {
		fmt.Println(path)
		_, _ = errFile.WriteString(path + "\n")
	}
	wg.Done()
	<-ch
}
