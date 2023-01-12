/*
* @Author: xilin0x7f
* @Date:   2023/1/12 15:37
 */
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func CheckFile(regStrFun, regStrT1, root string) ([][]string, [][]string) {
	regFun := regexp.MustCompile(regStrFun)
	regT1 := regexp.MustCompile(regStrT1)
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
				resultsFun = append(resultsFun, []string{root, group.Name(), subject.Name(), filesFunReg[0]})
				resultsT1 = append(resultsT1, []string{root, group.Name(), subject.Name(), filesT1Reg[0]})
			}
		}
	}
	return resultsFun, resultsT1
}
