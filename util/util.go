package util

import (
	"gocv.io/x/gocv"
	"image"
	"log"
)

/**
图像等比缩放到指定宽度
*/
func ZoomMat(mat *gocv.Mat, width float64) {
	scale := width / float64(mat.Cols())
	gocv.Resize(*mat, mat, image.Point{}, scale, scale, gocv.InterpolationArea)
}

/**
将dst图片绘制到source图片上
row col 开始绘制的端点起始值
*/
func Draw(source *gocv.Mat, dst *gocv.Mat, beginRow int, beginCol int, pngMode bool) {
	chanSource := gocv.Split(*source)
	chanDst := gocv.Split(*dst)

	for col := 0; col < (*dst).Cols(); col++ {
		for row := 0; row < (*dst).Rows(); row++ {
			for c := 0; c < (*dst).Channels(); c++ {
				if pngMode == true && chanDst[c].GetUCharAt(row, col) == 0 {
					continue
				}
				chanSource[c].SetUCharAt(beginRow+row, beginCol+col, chanDst[c].GetUCharAt(row, col))
			}
		}
	}
	gocv.Merge(chanSource, source)
}

/**
初始化一个指定大小的3通道图像 value
*/
func NewChanMat(row int, col int, value byte) gocv.Mat {
	mat := gocv.NewMat()
	// todo设置成248 248 248
	var chans []gocv.Mat

	var v []byte
	for i := 0; i < row*col; i++ {
		v = append(v, value)
	}

	for i := 0; i < 3; i++ {
		mat, err := gocv.NewMatFromBytes(row, col, gocv.MatTypeCV8U, v)
		if err != nil {
			log.Fatal(err)
		}
		chans = append(chans, mat)
	}

	gocv.Merge(chans, &mat)
	return mat
}
