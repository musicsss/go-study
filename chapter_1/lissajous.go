package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// 定义颜色调色板
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // 调色板中的第一种颜色
	blackIndex = 1 // 调色板中的下一种颜色
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	// 定义常量
	const (
		cycles  = 5     // 完整的x振荡器变化的个数
		res     = 0.001 // 角度分辨率
		size    = 100   // 图像画布包含[-size..+size]
		nframes = 64    // 动画中的帧数
		delay   = 8     // 以10ms为单位的帧间延迟
	)
	// 随机生成频率
	freq := rand.Float64() * 3
	anim := gif.GIF{LoopCount: nframes} // 创建一个新的gif动画
	phase := 0.0                        // 相位差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) // 创建一个新的图像矩形
		img := image.NewPaletted(rect, palette)      // 创建一个新的调色板图像
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			// 根据正弦函数生成x和y的坐标
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// 设置像素颜色
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
			phase += 0.1
			anim.Delay = append(anim.Delay, delay) // 添加帧间延迟
			anim.Image = append(anim.Image, img)   // 添加图像帧
		}
	}
	gif.EncodeAll(out, &anim) // 将动画写入输出流
}
