package main

/*
@Author: xilin0x7f, https://github.com/xilin0x7f
*/

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func CheckFile(regStrFun, regStrT1, root, z string) ([][]string, [][]string) {
	postN := 3
	if z == "y" {
		postN = 6
	}
	regFun := regexp.MustCompile(regStrFun)
	regT1 := regexp.MustCompile(regStrT1)
	var resultsFunJson, resultsT1Json []string
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
				fmt.Println(fmt.Sprintf("Group %v subject %v has %v Fun images.", group.Name(), subject.Name(), len(filesFunReg)))
				fmt.Println(fmt.Sprintf("Group %v subject %v has %v T1 images.", group.Name(), subject.Name(), len(filesT1Reg)))
			} else {
				_ = funCSVWriter.Write([]string{root, group.Name(), subject.Name(), filesFunReg[0]})
				funCSVWriter.Flush()
				resultsFun = append(resultsFun, []string{root, group.Name(), subject.Name(), filesFunReg[0]})
				_ = t1CSVWriter.Write([]string{root, group.Name(), subject.Name(), filesT1Reg[0]})
				t1CSVWriter.Flush()
				resultsT1 = append(resultsT1, []string{root, group.Name(), subject.Name(), filesT1Reg[0]})
			}
			for funIdx := range filesFunReg {
				resultsFunJson = append(resultsFunJson, filepath.Join(root, group.Name(), subject.Name(),
					filesFunReg[funIdx][:len(filesFunReg[funIdx])-postN]+"json"))
			}
			for t1Idx := range filesT1Reg {
				t1Origin := strings.Replace(filesT1Reg[t1Idx], "_Crop_1", "", -1)
				resultsT1Json = append(resultsT1Json, filepath.Join(root, group.Name(), subject.Name(),
					t1Origin[:len(t1Origin)-postN]+"json"))
			}

		}
	}
	if err := WriteJson2XLSX(resultsFunJson, filepath.Join(root, "Fun.xlsx"), "Sheet1", "A"); err != nil {
		log.Fatal(err)
	}
	if err := WriteJson2XLSX(resultsT1Json, filepath.Join(root, "T1.xlsx"), "Sheet1", "A"); err != nil {
		log.Fatal(err)
	}
	return resultsFun, resultsT1
}

func WriteJson2XLSX(filesName []string, dstFileName, sheetName, start string) error {
	keysMap := make(map[string]int)
	for _, fileName := range filesName {
		file, _ := os.Open(fileName)
		var jsonData map[string]interface{}
		reader := io.Reader(file)
		decoder := json.NewDecoder(reader)
		_ = decoder.Decode(&jsonData)
		_ = file.Close()
		for key := range jsonData {
			keysMap[key]++
		}
	}
	var keys []string
	for key := range keysMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	resMap := make(map[string][]interface{})
	for _, fileName := range filesName {
		file, _ := os.Open(fileName)
		var jsonData map[string]interface{}
		reader := io.Reader(file)
		decoder := json.NewDecoder(reader)
		_ = decoder.Decode(&jsonData)
		_ = file.Close()
		for _, key := range keys {
			resMap[key] = append(resMap[key], jsonData[key])
		}
	}
	res := make([][]interface{}, len(resMap[keys[0]])+1)
	for idx := range res {
		res[idx] = make([]interface{}, len(keys)+1)
	}
	res[0][0] = ""
	for idx, key := range keys {
		res[0][idx+1] = key
	}
	for rowIdx := range resMap[keys[0]] {
		res[rowIdx+1][0] = filesName[rowIdx]
		for colIdx := range keys {
			res[rowIdx+1][colIdx+1] = resMap[keys[colIdx]][rowIdx]
		}
	}
	err := Write2XLSX(dstFileName, sheetName, start, res)
	return err
}
func Write2XLSX(fileName, sheetName, start string, data [][]interface{}) error {
	f := excelize.NewFile()
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return err
	}
	f.SetActiveSheet(index)
	for idx := range data {
		if err := f.SetSheetRow(sheetName, fmt.Sprint(start, idx+1), &data[idx]); err != nil {
			return err
		}
	}
	if err := f.SaveAs(fileName); err != nil {
		return err
	}
	return nil
}
