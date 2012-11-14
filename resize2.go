package main

import (
    "github.com/rif/resize"
    "image/jpeg"
    "log"
    "os"
    "flag"
)

const(
	THUMBNAIL_MAX_WIDTH  = 800
	THUMBNAIL_MAX_HEIGHT = THUMBNAIL_MAX_WIDTH
)

func main() {
    // open "test.jpg"
    flag.Parse()    
    file, err := os.Open(flag.Arg(0))
    if err != nil {
        log.Fatal(err)
    }

    // decode jpeg into image.Image
    img, err := jpeg.Decode(file)
    if err != nil {
        log.Fatal(err)
    }
    file.Close()
    bounds:=img.Bounds()
    w, h := THUMBNAIL_MAX_WIDTH, THUMBNAIL_MAX_HEIGHT
    if bounds.Dx() > bounds.Dy() {
            h = bounds.Dy() * h / bounds.Dx()
    } else {
            w = bounds.Dx() * w / bounds.Dy()
    }
    img = resize.Resize(img, img.Bounds(), w, h)       

    out, err := os.Create("test_resized.jpg")
    if err != nil {
        log.Fatal(err)
    }
    defer out.Close()

    // write new image to file
    jpeg.Encode(out, img, nil)
}
