/*
* @Author: xilin0x7f
* @Date:   2023/1/12 15:37
 */
package main

import "flag"

func main() {
	step := flag.Int("step", -1, "0: 重命名, 1: 格式转换, 2: 检查, 3：复件文件")
	root := flag.String("root", "", "数据存放路径，root/组别/被试")
	n := flag.Int("n", 5, "格式转换时的并发数量")
	regStrFun := flag.String("regStrFun", "", "功能像正则匹配表达式")
	regStrT1 := flag.String("regStrT1", "", "结构像正则匹配表达式")
	dstRoot := flag.String("dstRoot", "", "DPABI格式数据存放位置")

	dcm2niixParF := flag.String("f", "%f_%p_%t_%s", "dcm2niix -f")
	dcm2niixParD := flag.String("d", "9", "dcm2niix -d")
	dcm2niixParX := flag.String("x", "y", "dcm2niix -x")
	dcm2niixParI := flag.String("i", "y", "dcm2niix -i")
	dcm2niixParZ := flag.String("z", "n", "dcm2niix -z")
	dcm2niixParameters := []string{"-f", *dcm2niixParF, "-d", *dcm2niixParD, "-x", *dcm2niixParX, "-i",
		*dcm2niixParI, "-z", *dcm2niixParZ}
	flag.Parse()

	switch *step {
	case 0:
		RenameSubjectFolder(*root, "rename.csv")
	case 1:
		Dcm2niixConcurrent(*root, *n, dcm2niixParameters...)
	case 2:
		CheckFile(*regStrFun, *regStrT1, *root)
	case 3:
		funFiles, t1Files := CheckFile(*regStrFun, *regStrT1, *root)
		CopyFile2DPABIFormat(funFiles, t1Files, *dstRoot)
	}
}
