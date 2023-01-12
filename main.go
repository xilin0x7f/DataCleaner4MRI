package main

func main() {
	root := "D:\\SY\\Project\\2022012\\002Processing\\001Data\\a"
	// Step 0 重命名
	RenameSubjectFolder(root, "rename.csv")

	// Step 1 格式转换
	// 输入 根路径，并发数, dcm2niix 参数
	// 数据组织格式 根路径/组别/被试
	dcm2niixParameters := []string{"-f", "%f_%p_%t_%s", "-d", "9", "-x", "y", "-i", "y", "-z", "y"}
	Dcm2niixConcurrent(root, 5, dcm2niixParameters...)

	// Step 2 功能像和结构像检查
	regStrFun := ""
	regStrT1 := ""
	CheckFile(regStrFun, regStrT1, root)

	// Step 3 复制功能像和结构像到新路径，DPABI格式
	funFiles, t1Files := CheckFile(regStrFun, regStrT1, root)
	dstRoot := "D:\\SY\\Project\\2022012\\002Processing\\001Data"
	CopyFile2DPABIFormat(funFiles, t1Files, dstRoot)
}
