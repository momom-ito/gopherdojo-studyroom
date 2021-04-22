package go_practice

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type ImageFile struct {
	path   string
	before string
	after  string
}

func convPic(imageFile ImageFile) error {
	file, err := os.Open(imageFile.path)
	if err != nil {
		return err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}
	op := filepath.Dir(imageFile.path) + "/" + filepath.Base(imageFile.path[:len(imageFile.path)-len(filepath.Ext(imageFile.path))]) + "." + imageFile.after
	out, _ := os.Create(op)
	defer out.Close()

	switch imageFile.after {
	case "png":
		png.Encode(out, img)
	case "jpeg", "jpg":
		jpeg.Encode(out, img, &jpeg.Options{})
	case "gif":
		gif.Encode(out, img, nil)
	}

	return nil
}

func isConvertExtension(ex string) bool{
	extensionArray := []string{"png", "jpg", "jpeg", "gif"}
	for _, x := range extensionArray {
		if ex == x {
			return true
		}
	}
	return false
}

func Convpic() error {

	var (
		dirname string
		before string
		after string
	)

	flag.StringVar(&before, "b", "jpg", "type before convert")
	flag.StringVar(&after, "a", "png", "type after convert")

	flag.Parse()

	fmt.Println(before, after)
	if !isConvertExtension(before) || !isConvertExtension(after) {
		return fmt.Errorf("invalid convert type")
	}

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		dirname = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	err := filepath.Walk(dirname,
		func(path string, info os.FileInfo, err error) error {

			if filepath.Ext(path) == "."+before {
				err := convPic(ImageFile{before: before, after: after, path: path})
				if err != nil {
					return err
				}
			}
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}
