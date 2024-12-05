package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

// original.png: 原始圖像(海報)
// mark.png: 隱寫訊息(qrcode)
// marked.png: 隱寫圖(原始圖像+隱藏圖像)
func main() {
	fileList := []string{"filmStudio.png", "movieDirector.png", "printingHouse.png"}
	for _, fileName := range fileList {
		addMark("original.png", fileName, "marked-"+fileName) // 呼叫 addMark 函數
	}
}

// 進行隱寫
// 將 "original.png" 和 "mark.png" 結合，生成 "marked.png"
func addMark(original, mark, marked string) {
	// 打開海報
	originalFile, err := os.Open(original)
	if err != nil {
		log.Fatalln(err)
	}
	defer originalFile.Close()
	// 解碼海報為圖片對象
	originalImage, _, err := image.Decode(originalFile)
	if err != nil {
		log.Fatalln(err)
	}

	// 打開qrcode
	markFile, err := os.Open(mark)
	if err != nil {
		log.Fatalln(err)
	}
	defer markFile.Close()
	// 解碼 qrcode 為圖片對象
	markImage, _, err := image.Decode(markFile)
	if err != nil {
		log.Fatalln(err)
	}

	// 準備隱寫圖，(0, 0) 是左上角的座標，(w, h) 是右下角的座標)
	bounds := originalImage.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	markedImage := image.NewRGBA(image.Rect(0, 0, w, h))

	// 遍歷每個像素(0,0)->(w,h)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			// 獲取海報的顏色
			oldColor := originalImage.At(x, y)
			r, g, b, a := oldColor.RGBA()

			// 獲取 qrcode 對應位置的 R 值
			markR, _, _, _ := markImage.At(x, y).RGBA()
			var mark uint32 = 1
			if markR < 128 {
				mark = 0
			}
			// 先將海報r值的最低位元設為0
			// 再將qrcode的R值的最低位元插入到海報的R值的最低位元中
			r = (r & 0b11111110) + mark
			// 設定新圖片的像素
			markedImage.SetRGBA(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		}
	}

	// 創建新圖片檔案，並將帶有隱藏訊息的新圖片保存為檔案
	outfile, err := os.Create(marked)
	if err != nil {
		log.Fatalln(err)
	}
	defer outfile.Close()
	png.Encode(outfile, markedImage)
}
