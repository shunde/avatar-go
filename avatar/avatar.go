package avatar

import (
	"crypto/md5"
	"image"
	"image/color"
)

const BlockSize int = 5
const PatchSize int = 20

func NewAvatar(s string) image.Image {
	//salt := ""
	//str := fmt.Sprintf("%s%s", s, salt)
	hash := md5.Sum([]byte(s))
	avatar := image.NewRGBA(image.Rect(0, 0, BlockSize, BlockSize))

	background := color.RGBA{255, 255, 255, 0}
	foreground := color.RGBA{hash[0] & 255, hash[1] & 255, hash[2] & 255, 255}

	for x := 0; x < BlockSize; x++ {
		i := x
		if x > BlockSize/2 {
			i = BlockSize - x - 1
		}
		for y := 0; y < BlockSize; y++ {
			var pixelColor color.RGBA
			// 重点改进下面
			if (hash[i+y] >> uint32(y) & 1) == 1 {
				pixelColor = foreground
			} else {
				pixelColor = background
			}
			avatar.SetRGBA(x, y, pixelColor)
		}
	}

	m := scale(avatar)
	return m.SubImage(m.Bounds())
}

func scale(i *image.RGBA) *image.RGBA {
	size := BlockSize * PatchSize
	newImage := image.NewRGBA(image.Rect(0, 0, size, size))

	for x := 0; x < BlockSize; x++ {
		for y := 0; y < BlockSize; y++ {
			pixelColor := i.At(x, y)

			for newX := x * PatchSize; newX < (x+1)*PatchSize; newX++ {
				for newY := y * PatchSize; newY < (y+1)*PatchSize; newY++ {
					newImage.Set(newX, newY, pixelColor)
				}
			}
		}
	}
	return newImage
}
