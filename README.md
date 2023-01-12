# DataCleaner4MRI

## Step 0 重命名
```
DataCleaner4MRI -step 0 -root $root
```
## Step 1 格式转换
输入 根路径，并发数（默认为5）, dcm2niix 参数，数据组织格式 根路径/组别/被试，需要可以在命令行中调用dcm2niix程序。
```
DataCleaner4MRI -step 1 -root $root
```
## Step 2 功能像和结构像检查
```
DataCleaner4MRI -step 2 -root $root -regStrFun $regStrFun -regStrT1 $regStrT1
```
## Step 3 复制功能像和结构像到新路径，DPABI格式
```
DataCleaner4MRI -step 3 -root $root -dstRoot $dstRoot -regStrFun $regStrFun -regStrT1 $regStrT1
```

## 示例
```
DataCleaner4MRI -step 0 -root E:\MRIData\test
DataCleaner4MRI -step 1 -root E:\MRIData\test
DataCleaner4MRI -step 2 -root E:\MRIData\test -regStrFun EPI.*nii -regStrT1 Crop
DataCleaner4MRI -step 3 -root E:\MRIData\test -regStrFun EPI.*nii -regStrT1 Crop -dstRoot E:\MRIData\a\
```