# Image Resize Tool

## Usage

```
Usage of /tmp/go-build2912109510/b001/exe/main:
  -ext string
    	image extension (default "jpg")
  -height int
    	Height of the resized image
  -input string
    	Input directory containing images (default ".")
  -output string
    	Output directory to save resized images (default ".")
  -ratio string
    	Aspect ratio of the resized image, etc auto or 4:3 (default "auto")
  -width int
    	Width of the resized image
```

## Sample

```
go run main.go --input sample --output result -height 200
```

## Output

```
prepare to resize image to width:200px height:auto
success to process sample/1/11/59a776aa35bfe.jpg => result/1/11/59a776aa35bfe.jpg
success to process sample/1/5c80e0956b20b.jpg => result/1/5c80e0956b20b.jpg
success to process sample/2/5c9475cea1c64.jpg => result/2/5c9475cea1c64.jpg
success to process sample/48f1ecc3aa64b17bdd0ccc7047288406.jpg => result/48f1ecc3aa64b17bdd0ccc7047288406.jpg
success to process sample/f41135e8527ad81f0ccdfa5ed667080a.jpg => result/f41135e8527ad81f0ccdfa5ed667080a.jpg
success walking the path sample
```

## Build

```
go build -o image_resize_tool


## Cross compile
GOOS=windows GOARCH=amd64 go build -o image_resize_tool.exe
```