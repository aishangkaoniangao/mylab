package local

import (
	"os"
	"fmt"
	"github.com/golang/freetype"
	"image"
	"image/color"
	"image/jpeg"
	"image/draw"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
	_ "errors"
)

var (
	ttfFile = "/home/luwenwei/labs/go/project/captcha/src/github.com/freetype/testdata/luxirr.ttf"
	demoFile = "/home/luwenwei/tmp/a.jpg"
)

func InitDemo() {
	var textLen int = 4
	
	file := demoFile
	f1, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	var x0, y0, x1, y1 int

	dst := image.NewRGBA(image.Rect(0, 0, 80, 40))
	draw.Draw(dst, dst.Bounds(), image.White, image.ZP, draw.Src)
	
	FuncDrawNoise(dst)

	rand.Seed(time.Now().Unix())
	for i := 0; i < 2; i++ {
		x0 = rand.Intn(dst.Bounds().Max.X)
		x1 = dst.Bounds().Max.X
		y0 = rand.Intn(dst.Bounds().Max.Y)
		y1 = rand.Intn(dst.Bounds().Max.Y)

		FuncDrawLine(x0, y0, x1, y1, dst)
	}
	
	FuncDrawLetterNoise(dst, GetText(textLen))
	FuncDrawLetter(dst, GetText(textLen))

	err = jpeg.Encode(f1, dst, &jpeg.Options{90})
	if err != nil {
		panic(err)
	}
	fmt.Println("OK")
}

func FuncDrawLine(x0, y0, x1, y1 int, m *image.RGBA) {
	var k, t float32

	if (x1 == x0) {
		return
	}

	k = (float32)(y1 - y0) / (float32)(x1 - x0)
	t = (float32)(y0) - k * (float32)(x0)

	var y int
	for x := x0; x < x1; x++ {
		y = (int)(k * (float32)(x) + t)
		m.Set(x, y, color.Black)
	}
}

func FuncDrawLetter(dst *image.RGBA, letters string) {
	//加载字体文件
	ttf := ttfFile
	data, err := ioutil.ReadFile(ttf)
	if err != nil {
		panic(err)
	}
	f, err := freetype.ParseFont(data)
	if err != nil {
		panic(err)
	}
	
	ss := strings.Split(letters, "")
	c := freetype.NewContext()
	c.SetDst(dst)
	c.SetClip(dst.Bounds())
	c.SetSrc(image.Black)
	c.SetFont(f)
	c.SetFontSize(20)
	for t := 0; t < len(ss); t++ {
		_, err = c.DrawString(ss[t], freetype.Pt(t * 20, (int)((float32)(dst.Bounds().Max.Y) * 0.75)))
    }
}

func FuncDrawNoise(dst *image.RGBA) {
	bound := dst.Bounds()
	area := bound.Max.X * bound.Max.Y

	density := 20
	num := area / density

	var x, y int
	rand.Seed(time.Now().Unix())
	for i := 0; i < num; i++ {
		x = rand.Intn(bound.Max.X)
		y = rand.Intn(bound.Max.Y)
		dst.Set(x, y, RandomColor())
		//dst.Set(x, y, color.Black)
		//dst.Set(x, y, color.RGBA{R:255, G:255, B:255})
		//dst.Set(x, y, color.RGBA{R:0, G:0, B:0})
	}
}

func FuncDrawLetterNoise(dst *image.RGBA, letters string) {
	//加载字体文件
	ttf := ttfFile
	data, err := ioutil.ReadFile(ttf)
	if err != nil {
		panic(err)
	}
	f, err := freetype.ParseFont(data)
	if err != nil {
		panic(err)
	}
	
	ss := strings.Split(letters, "")
	c := freetype.NewContext()
	c.SetDst(dst)
	c.SetClip(dst.Bounds())
	c.SetFont(f)
	c.SetFontSize(10)
	for t := 0; t < len(ss); t++ {
		c.SetSrc(image.NewUniform(LightGreenColor()))
		_, err = c.DrawString(ss[t], freetype.Pt(t * 20, (int)((float32)(dst.Bounds().Max.Y) * 0.25)))
	}
}

func RandomColor() (color.RGBA) {
	r := rand.Intn(100)
	g := rand.Intn(100) + 100
	b := 400 - r - g
	return color.RGBA{R:uint8(r), G:uint8(g), B:uint8(b)}
}

func LightGreenColor() (color.RGBA) {
	r, g, b := 255, 0, 0
	return color.RGBA{R:uint8(r), G:uint8(g), B:uint8(b)}
}



//***********************
//===This is demo func===
//***********************

//mask demo
func CopyToF3() {
	dir := "/home/luwenwei/labs/go/project/captcha/tmp/"
	f1, err := os.Open(dir + "1.jpg")
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	f2, err := os.Open(dir + "2.jpg")
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	f3, err := os.Create(dir + "3.jpg")
	if err != nil {
		panic(err)
	}
	defer f3.Close()

	//解码jpg文件
	m1, err := jpeg.Decode(f2)
	if err != nil {
		panic(err)
	}
	//获取jpg的边界
	bounds := m1.Bounds()
	m2, err := jpeg.Decode(f1)
	if err != nil {
		panic(err)
	}
	//创建空白图片，仅定义了边界
	//m := image.NewRGBA(bounds)
	m := image.NewRGBA(image.Rect(0, 0, 800, 800))
	//创建颜色句柄
	white := color.RGBA{255, 255, 255, 255}
	//1=目标文件 2=图片矩形域(目标两个顶点的矩形域) 3=image.Image 4=image.Point(src的mask顶点) 5=Op
	//用白颜色覆盖到目标图片
	draw.Draw(m, bounds, &image.Uniform{white}, image.ZP, draw.Src)
	var _ = m2
	//用m1的覆盖到目标图片
	draw.Draw(m, bounds, m1, image.ZP, draw.Src)
	//用m2的起点是250,60的矩形覆盖到目标图片的矩形上
	draw.Draw(m, image.Rect(350, 350, 550, 550), m2, image.Pt(100, 100), draw.Src)
	//Options=quality 90%的质量贴到f3上
	err = jpeg.Encode(f3, m, &jpeg.Options{90})
	if err != nil {
		panic(err)
	}
	fmt.Printf("ok\n")
}

//line demo
func DrawLine() {
	file := demoFile
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m := image.NewRGBA(image.Rect(0, 0, 200, 200))
	draw.Draw(m, image.Rect(0, 0, 200, 200), &image.Uniform{color.White}, image.ZP, draw.Src)
	
	var x0, x1, y0, y1 int
	var k, t float32

	x0 = 0
	y0 = 0
	x1 = 200
	y1 = 100
	k = (float32)(y1 - y0) / (float32)(x1 - x0)
	t = (float32)(y0) - k * (float32)(x0)

	var _ = t
	var y int
	for x := x0; x < x1; x++ {
		y = (int)(k * (float32)(x) + t)
		m.Set(x, y, color.Black)
	}
	err = jpeg.Encode(f, m, &jpeg.Options{90})
	if err != nil {
		panic(err)
	}
	fmt.Println("OK")
}

//letter demo
func DrawLetter() {
	//加载字体文件
	ttf := ttfFile
	data, err := ioutil.ReadFile(ttf)
	if err != nil {
		panic(err)
	}
	f, err := freetype.ParseFont(data)
	if err != nil {
		panic(err)
	}
	
	dst := image.NewRGBA(image.Rect(0, 0, 800, 600))
	draw.Draw(dst, dst.Bounds(), image.White, image.ZP, draw.Src)

	var line string = "N"
	c := freetype.NewContext()
	c.SetDst(dst)
	c.SetClip(dst.Bounds())
	c.SetSrc(image.Black)
	c.SetFont(f)
	c.SetFontSize(20)
	for t := 0; t < 4; t++ {
		_, err = c.DrawString(line, freetype.Pt(t * 50, 100))
    }

	file := demoFile
	f1, _ := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	err = jpeg.Encode(f1, dst, &jpeg.Options{90})
	if err != nil {
		panic(err)
	}
	fmt.Println("OK")
}
