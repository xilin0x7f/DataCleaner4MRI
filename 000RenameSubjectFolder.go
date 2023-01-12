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

func RenameSubjectFolder(root, renameFile string) {
	file, _ := os.OpenFile(filepath.Join(root, renameFile), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 644)
	_, _ = file.Write([]byte{0xEF, 0xBB, 0xBF})
	defer func() {
		_ = file.Close()
	}()
	groups, _ := os.ReadDir(root)
	matchStr := "[0-9_a-zA-Z]+"
	matchReg, _ := regexp.Compile(matchStr)
	for _, group := range groups {
		subjects, _ := os.ReadDir(filepath.Join(root, group.Name()))
		idx := 0
		for _, subject := range subjects {
			originPath := filepath.Join(root, group.Name(), subject.Name())
			newPath := filepath.Join(root, group.Name(), fmt.Sprintf("%03d%v", idx, matchReg.FindString(subject.Name())))
			_, _ = file.WriteString(originPath + "," + newPath + "\n")
			_ = os.Rename(originPath, newPath)
			idx++
		}
	}
}
