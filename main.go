package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/nfnt/resize"
)

func main() {
	inputDir := flag.String("input", ".", "Input directory containing images")
	outputDir := flag.String("output", ".", "Output directory to save resized images")
	width := flag.Int("width", 0, "Width of the resized image")
	height := flag.Int("height", 0, "Height of the resized image")
	ratio := flag.String("ratio", "auto", "Aspect ratio of the resized image, etc auto or 4:3")
	extension := flag.String("ext", "jpg", "image extension")

	flag.Usage = func() {
		fmt.Printf(`  ====================================
  welcome to use image_resize_tool!!!
  Author: Chad
  Email: chad_lwm@hotmail.com
  Source Code: https://github.com/chadlwm/image_resize_tool
  ====================================

`)
		flag.PrintDefaults()
	}

	flag.Parse()

	if (*inputDir == "" && *outputDir == "") || *inputDir == *outputDir {
		fmt.Printf("please specify inputDir and outputDir, and not same\n\n")
		flag.Usage()
		os.Exit(1)
	}

	if *width == 0 && *height == 0 {
		fmt.Printf("please specify either width or height\n\n")
		flag.Usage()
		os.Exit(1)
	}

	if *extension == "" {
		fmt.Printf("please specify extension\n\n")
		flag.Usage()
		os.Exit(1)
	}

	var w, h uint
	if *width != 0 && *height != 0 {
		w = uint(*width)
		h = uint(*height)
	} else {
		switch *ratio {
		case "auto":
			w = uint(*width)
			h = uint(*height)
		default:
			rationVals := strings.Split(*ratio, ":")
			if len(rationVals) != 2 {
				fmt.Printf("invalid aspect ratio: %s\n\n", *ratio)
				flag.Usage()
				os.Exit(1)
			}

			wr, err := strconv.Atoi(rationVals[0])
			if err != nil {
				fmt.Printf("invalid ratio: %s\n\n", *ratio)
				flag.Usage()
				os.Exit(1)
			}
			hr, err := strconv.Atoi(rationVals[1])
			if err != nil {
				fmt.Printf("invalid aspect ratio: %s\n\n", *ratio)
				flag.Usage()
				os.Exit(1)
			}

			if *width != 0 {
				w = uint(*width)
				h = uint(*width * wr / hr)
			} else if *height != 0 {
				h = uint(*width)
				w = uint(*height * hr / wr)
			}

		}
	}

	htip, wtip := strconv.Itoa(int(h))+"px", strconv.Itoa(int(w))+"px"
	if htip == "0px" {
		htip = "auto"
	}
	if wtip == "0px" {
		wtip = "auto"
	}

	fmt.Printf("prepare to resize image to width:%v height:%v\n", htip, wtip)

	err := filepath.Walk(*inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), "."+*extension) {
			outPath, err := processImage(path, *inputDir, *outputDir, w, h)
			if err != nil {
				fmt.Printf("failed to process %s: %v\n", path, err)
			} else {
				fmt.Printf("success to process %s => %s\n", path, outPath)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %v: %v\n", *inputDir, err)
	} else {
		fmt.Printf("success walking the path %v\n", *inputDir)
	}
}

func processImage(path, inputDir, outputDir string, width, height uint) (outPath string, errRet error) {
	file, err := os.Open(path)
	if err != nil {
		errRet = err
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		errRet = err
		return
	}

	m := resize.Resize(width, height, img, resize.Lanczos3)

	outPath = strings.Replace(path, inputDir, outputDir, 1)
	err = os.MkdirAll(filepath.Dir(outPath), 0755)
	if err != nil {
		errRet = err
		return
	}

	outFile, err := os.Create(outPath)
	if err != nil {
		errRet = err
		return
	}
	defer outFile.Close()

	errRet = jpeg.Encode(outFile, m, &jpeg.Options{Quality: 90})
	return
}
