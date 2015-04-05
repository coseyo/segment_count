package segment_count

import (
	"fmt"
	"math"
	"time"

	"github.com/huichen/sego"
)

var (
	dataobj   data
	segmenter sego.Segmenter
)

func Exec(srcTable string, srcField string, distTable string) error {
	t1 := time.Now()
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
		offset := i * GROUP_COUNT
		worker(offset, GROUP_COUNT)
	}

	t2 := time.Now()
	d := t2.Sub(t1)
	fmt.Println("used time:", d)

	return nil
}

func worker(offset int, limit int) error {
	titleArray, err := dataobj.read(offset, limit)
	if err != nil {
		return err
	}

	for _, title := range titleArray {
		text := []byte(title)
		segments := segmenter.Segment(text)
		segmentArray := sego.SegmentsToSlice(segments, false)
		for _, segment := range segmentArray {
			if err := dataobj.write(segment); err != nil {
				return err
			}
		}
	}
	return nil
}
