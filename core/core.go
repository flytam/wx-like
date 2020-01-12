package core

import (
	"gocv.io/x/gocv"
	"image"
	"log"
	"wx-like/util"
)

func Process() {
	// 窗口展示
	window := gocv.NewWindow("Tracking")

	// 原始图像
	originImg := gocv.IMRead("./2.png", gocv.IMReadColor)

	// 头像个数。
	size := 2

	// 原始图像处理成 435宽度

	// img缩放成435*953

	// result是匹配的临时用的
	result := gocv.NewMat()

	// target最后绘制到原图像上
	// 主要在这个上面绘制

	// 匹配
	// 用于匹配位置的图像

	pickImg := gocv.IMRead("./target.png", gocv.IMReadColor)
	gocv.MatchTemplate(originImg, pickImg, &result, gocv.TmCcoeffNormed, gocv.NewMat())

	_, _, _, maxLoc := gocv.MinMaxLoc(result)

	// 假定435
	myWidth := 435

	// 距离边上的距离 20。也就是x坐标
	pasteCol := originImg.Cols() - (maxLoc.X + pickImg.Cols())
	pasteRow := maxLoc.Y + pickImg.Rows()

	// 创建背景
	// 背景的 rows
	// 背景的 cols
	targetImg := util.NewChanMat((size/7+1)*(40+5)+1, myWidth-pasteCol*2, 248)

	//-----top处理
	// 读取top素材
	topImg := gocv.IMRead("./top.png", gocv.IMReadColor)
	// 先等比处理成435宽度
	util.ZoomMat(&topImg, float64(myWidth))
	// 绘制，从 0 0 开始
	util.Draw(&targetImg, &topImg, 0, 0, false)

	// 处理点赞的心型
	leftImg := gocv.IMRead("./left.png", gocv.IMReadColor)
	// 等比处理成40*40。。TODO: 换成直接绘制
	util.ZoomMat(&leftImg, 40)
	// 绘制 第0列 第topImg下部行开始
	util.Draw(&targetImg, &leftImg, topImg.Rows(), 0, true)

	// 绘制头像
	// 拉取头像url
	//urlStrings := api.GetAvaterUrls(size)

	// 读取头像生成矩阵，并绘制上去
	for i := 0; i < size; i++ {
		// x, err := gocv.OpenVideoCapture(urlStrings[i])
		// https://source.unsplash.com/user/erondu/40x40
		x, err := gocv.OpenVideoCapture("https://source.unsplash.com/user/erondu/40x40")
		if err != nil {
			log.Fatal(err)
		}

		// 头像
		tMat := gocv.NewMat()

		x.Read(&tMat)
		//处理成40x40
		gocv.Resize(tMat, &tMat, image.Point{
			X: 40,
			Y: 40,
		}, 0, 0, gocv.InterpolationArea)
		util.ZoomMat(&tMat, 40)
		// 当前的行数
		currentRow := i/7 + 1
		// 当前行的第几个
		currentCol := i % 7

		// row坐标
		row := topImg.Rows() + (currentRow-1)*(40+5)

		// col坐标
		col := leftImg.Cols() + (currentCol-1)*(40+5)

		// 把头像绘制
		util.Draw(&targetImg, &tMat, row, col, false)
	}

	util.Draw(&originImg, &targetImg, pasteRow, pasteCol, false)

	//gocv.Rectangle(&originImg, image.Rectangle{
	//	Min: image.Point{
	//		X: 300, Y: 300,
	//	},
	//	Max: image.Point{60, 60},
	//}, color.RGBA{R: 248, G: 248, B: 248}, 5)
	window.IMShow(originImg)
	window.WaitKey(30000)
}
