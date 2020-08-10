package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"math"
	"net/url"
	"path"
	"reflect"
	"sort"
	"strings"
	"time"
)

func main1() {

	stop := make(chan int, 1)
	go pr(stop)

	time.Sleep(3 * time.Second)
	stop <- 1
}

func pr(stop <-chan int) {
	tick := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-stop:
			// Request finished.
			tick.Stop()
			return
		case <-tick.C:
			print("...\n")
		}
	}
}

func main3() {
	flushed := make(chan struct{})

	go func() {
		defer close(flushed)
		// Force a log flush, because with very short requests we
		// may not ever flush logs.
		time.Sleep(3 * time.Second)
		flushed <- struct{}{}
		time.Sleep(2 * time.Second)
		print("----")
	}()

	print("b")
	<-flushed
	print("a")
}

func main5() {
	now := time.Now()
	fmt.Println(now.Nanosecond())
	utc := now.UTC()
	fmt.Println(utc.Nanosecond())
	fmt.Println(now)
	fmt.Println(utc)
	fmt.Println(ptypes.TimestampProto(now))
	fmt.Println(ptypes.TimestampProto(utc))
}

func main4() {
	a := [][]int{
		{10, 21, 0},
		{10, 21, 1},
		{16, 24, 2},

		{13, 25, 4},
		{16, 24, 3},
	}
	ca := func(i, j int) bool {
		if a[i][0] < a[j][0] {
			return true
		}

		if (a[i][0] == a[j][0]) && a[i][1] < a[j][1] {
			return true
		}

		return false
	}

	fmt.Println(ca(3, 4))
	fmt.Println(ca(4, 3))

	fmt.Println(a)
	sort.Slice(a, ca)
	fmt.Println(a)
}

func main6() {
	var x, y float64
	z := 0
	x = 0.0000000000000003847341387443574 * math.Cos(0)
	y = 0.0000000000000003847341387443574 * math.Sin(0)
	fmt.Println(z, x, y)
}

func main7() {
	//fileName := "aa.d.f.jpg"
	fileName := "ajpg"
	dotIx := strings.LastIndex(fileName, ".")
	fmt.Println(dotIx)
	fmt.Printf("%s_thumbnail.%s", fileName[:dotIx], fileName[dotIx+1:])
}

func main8() {
	filePath, _ := extractFilePath("https://storage.cloud.google.com/bucketName/ur/89899/sdf/5555/121")
	fmt.Println(filePath)
	fmt.Println(extractOwnerID(filePath))
}

func extractOwnerID(filePath string) (string, error) {
	// BucketName/UploadRoot/OwnerId/UID/Timestamp/FileName.
	const partsLength = 6

	sp := strings.Split(filePath, "/")
	if len(sp) != partsLength {
		return "", errors.New("not a validated image file path")
	}

	return sp[2], nil
}

func extractFilePath(iurl string) (string, error) {
	p, err := url.Parse(iurl)
	if err != nil {
		return "", err
	}

	fp := p.Path
	if strings.HasPrefix(fp, "/") {
		fp = fp[1:]
	}
	return fp, nil
}

func constructStorageFileAccessURL(bucket, fileName string) string {
	fmt.Println(bucket[:1])
	if bucket[:1] == "/" {
		bucket = bucket[1:]
	}
	const host = "https://storage.cloud.google.com"
	return fmt.Sprintf("%s/%s", host, path.Join(bucket, fileName))
}

func main10() {
	url := "https://storage.cloud.google.com/fieldbrowser-dummy-data/upload_image_dummy/Public/107095795860879954475/1595488270458915249/微信截图_20200412150755_thumbnail.png"
	fmt.Println(base64.RawURLEncoding.EncodeToString([]byte(url)))
}

type ci interface {
	ciFun()
}

type pti int

func (c *pti) ciFun() {}

func main() {

	fmt.Printf("no permission to batch mutate data for %v")

	//f := -101.3051015841
	//f1 := -101.30507674235011
	//
	//fmt.Println(truncate(f, 9))
	//fmt.Println(truncate(f1, 9))

}

func truncate(some float64, in float64) float64 {
	pow := math.Pow(10, in)
	return float64(int(some*pow)) / pow
}

func round(num float64) int {
	return int(num + math.Copysign(0, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func encodeFileName(fileName string) string {
	ldi := strings.LastIndex(fileName, ".")
	ex := fileName[ldi:]

	return fmt.Sprintf("%d%s", time.Now().UnixNano(), ex)

}

func getImageMimeType(filePath string) string {
	ldi := strings.LastIndex(filePath, ".")
	ex := filePath[ldi:]
	if strings.ToLower(ex) == ".png" {
		return "png"
	}

	return "jpg"
}

func main12() { //nolint:whitespace
	for i := 1; i < 3-1; i++ {
		fmt.Println("dsfa")
	}
	//
	//// bStcXZwG
	//c := []byte{
	//	'b',
	//	'S',
	//	't',
	//	'c',
	//	'X',
	//	'Z',
	//	'w',
	//	'G',
	//}
	//
	//for _, b := range c {
	//	fmt.Println(string(b-1))
	//}
	//fmt.Println(time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC))
}

func main0() {
	var pi *pti
	fmt.Println(pi == nil) // true

	c2 := c()
	fmt.Println(c2 == nil)                   // false
	fmt.Println(reflect.ValueOf(c2).IsNil()) // true
	fmt.Println(reflect.ValueOf(c2).Kind())  // ptr
}

func c() ci {
	var pi *pti
	return pi
}
