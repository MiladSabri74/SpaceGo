package main

import (
	"sync"
	"time"
)

var applyWatermark = func(image string) string {
	time.Sleep(100 * time.Millisecond)
	return image + "_watermarked"
}

func ProcessImages(images []string) []string {
	watermarked := make([]string, len(images))
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i, image := range images {
		wg.Add(1)
		go func(index int, image string) {
			defer wg.Done()

			mu.Lock()
			watermarked[index] = applyWatermark(image)
			mu.Unlock()
		}(i, image)
	}
	wg.Wait()
	return watermarked
}
