package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/pkg/errors"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile) // ログの出力書式を設定する
	for i, f := range os.Args {
		if i == 0 {
			continue
		}
		ext := strings.ToLower(filepath.Ext(f))

		//log.Printf("file : %v", f)
		//log.Printf("\text : %v", ext)
		//log.Printf("file: %v, ext: %v", f, ext)
		if true && //
			ext != ".gif" && //
			ext != ".jpg" && //
			ext != ".jpeg" && //
			ext != ".png" && //
			true {
			//log.Printf("\tError: 拡張子がgif,jpg,pngの何れかではありません。")
			continue
		}
		if i != 1 {
			fmt.Printf("\n")
		}
		//log.Printf("file: %v, ext: %v", f, ext)
		fmt.Printf("file:\n%v\n", f)

		file, err := os.Open(f)
		if err != nil {
			log.Printf("\t%v", err)
			continue
		}
		img, _, err := image.Decode(file)
		if err != nil {
			log.Printf("\t%v", err)
			continue
		}

		// prepare BinaryBitmap
		bmp, err := gozxing.NewBinaryBitmapFromImage(img)
		if err != nil {
			log.Printf("\t%v", err)
			continue
		}

		// decode image
		qrReader := qrcode.NewQRCodeReader()
		result, err := qrReader.Decode(bmp, nil)
		if err != nil {
			log.Printf("\t%v", err)
			continue
		}
		np := filepath.Join(filepath.Dir(f), filepath.Base(f)+".txt")
		if IsExist(np) {
			os.RemoveAll(np)
		}
		WriteText(np, fmt.Sprintf("%v", result))
		fmt.Printf("result:\n%v\n", result)
	}
}

func IsExist(path string) bool {
	//info, err := os.Stat(path)
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func WriteText(filepath, str string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("panic!:", err)
		}
	}()
	f, err := os.Create(filepath)
	defer f.Close()
	if err != nil {
		panic(errors.Errorf("%v", err))
	} else {
		if _, err := f.Write([]byte(str)); err != nil {
			panic(errors.Errorf("%v", err))
		}
	}
}
