package segment_count

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/huichen/sego"
)

var (
	dataobj   data
	segmenter sego.Segmenter
)

func Exec(srcDB string, srcTable string, srcField string, distTable string) error {
	tStart := time.Now()
	dataobj.srcDB = srcDB
	dataobj.SrcTable = srcTable
	dataobj.SrcField = srcField
	dataobj.DistTable = distTable

	if err := dataobj.init(); err != nil {
		return err
	}
	defer dataobj.db.Close()

	// 载入词典
	segmenter.LoadDictionary(DICT_FILE)

	count, err := dataobj.count()
	if err != nil {
		return err
	}
	fmt.Println("total count: ", count)

	num := int(math.Ceil(float64(count) / float64(GROUP_COUNT)))

	for i := 0; i < num; i++ {
		log.Println("worker :", (i + 1), " / ", num)
		offset := i * GROUP_COUNT
		worker(offset, GROUP_COUNT)
	}

	tEnd := time.Now()
	duration := tEnd.Sub(tStart)
	fmt.Println("Used time:", duration)

	return nil
}

func worker(offset int, limit int) error {
	titleArray, err := dataobj.read(offset, limit)
	if err != nil {
		return err
	}

	var wordSlice []string

	for _, title := range titleArray {
		text := []byte(title)
		segments := segmenter.Segment(text)
		segmentArray := sego.SegmentsToSlice(segments, false)
		for _, segment := range segmentArray {
			wordSlice = append(wordSlice, segment)
		}
	}
	if err := dataobj.write(wordSlice); err != nil {
		return err
	}
	return nil
}
