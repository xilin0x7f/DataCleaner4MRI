/*
* @Author: xilin0x7f
* @Date:   2023/1/12 15:37
 */
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func CheckFile(regStrFun, regStrT1, root string) ([][]string, [][]string) {
	regFun := regexp.MustCompile(regStrFun)
	regT1 := regexp.MustCompile(regStrT1)
	funFile, _ := os.OpenFile(filepath.Join(root, "FunImg.csv"), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 644)
	t1File, _ := os.OpenFile(filepath.Join(root, "T1Img.csv"), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 644)
	_, _ = funFile.Write([]byte{0xEF, 0xBB, 0xBF})
	_, _ = t1File.Write([]byte{0xEF, 0xBB, 0xBF})
	funCSVWriter := csv.NewWriter(funFile)
	t1CSVWriter := csv.NewWriter(t1File)
	defer func() {
		_ = funFile.Close()
		_ = t1File.Close()
	}()

	groups, _ := os.ReadDir(root)
	var resultsFun [][]string
	var resultsT1 [][]string
	for _, group := range groups {
		subjects, _ := os.ReadDir(filepath.Join(root, group.Name()))
		for _, subject := range subjects {
			files, _ := os.ReadDir(filepath.Join(root, group.Name(), subject.Name()))
			filesFunReg := make([]string, 0, len(files))
			filesT1Reg := make([]string, 0, len(files))
			for _, file := range files {
				if regFun.MatchString(file.Name()) {
					filesFunReg = append(filesFunReg, file.Name())
				}
				if regT1.MatchString(file.Name()) {
					filesT1Reg = append(filesT1Reg, file.Name())
				}
			}
			if len(filesFunReg) != 1 || len(filesT1Reg) != 1 {
				fmt.Println(fmt.Sprintf("Please check %v group %v subject Fun and T1 image", group.Name(), subject.Name()))
			} else {
				_ = funCSVWriter.Write([]string{root, group.Name(), subject.Name(), filesFunReg[0]})
				funCSVWriter.Flush()
				resultsFun = append(resultsFun, []string{root, group.Name(), subject.Name(), filesFunReg[0]})
				_ = t1CSVWriter.Write([]string{root, group.Name(), subject.Name(), filesT1Reg[0]})
				t1CSVWriter.Flush()
				resultsT1 = append(resultsT1, []string{root, group.Name(), subject.Name(), filesT1Reg[0]})
			}
		}
	}
	return resultsFun, resultsT1
}
