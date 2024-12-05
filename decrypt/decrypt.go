package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

// marked.png: 隱寫圖(海報+qrcode)
// decoded.png: 浮現盲水印
func main() {
		decodeMark("marked.png", "decoded.png") // 呼叫 decodeMark 函數
}

// 從帶有隱藏訊息的圖片"marked.png"中提取出訊息，並生成黑白圖片"decoded.png"
func decodeMark(marked, decoded string) {
	// 打開隱寫圖檔案
	originalFile, err := os.Open(marked)
	if err != nil {
		log.Fatalln(err)
	}
	defer originalFile.Close()
	// 解碼隱寫圖為圖片對象
	originalImage, _, err := image.Decode(originalFile)
	if err != nil {
		log.Fatalln(err)
	}

	// 盲水印加海報黑底的長寬
	// 獲取隱寫圖的寬度和高度
	bounds := originalImage.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	// 起點(0,0) 終點(w,h)
	newImage := image.NewRGBA(image.Rect(0, 0, w, h))

	// 遍歷每個像素
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			// 獲取隱寫圖像素的顏色
			oldColor := originalImage.At(x, y)
			// 獲取 R 值
			r, _, _, _ := oldColor.RGBA()
			// 檢查 R 值的最低位元
			if r&1 == 1 {
				newImage.Set(x, y, color.RGBA{255, 255, 255, 255}) // 白色
			} else { 
				newImage.Set(x, y, color.RGBA{0, 0, 0, 255}) // 黑色
			}
		}
	}

	// 創建新圖片檔案，並將生成的盲水印加海報黑底保存為檔案
	outfile, err := os.Create(decoded)
	if err != nil {
		log.Fatalln(err)
	}
	defer outfile.Close()
	png.Encode(outfile, newImage)
}
