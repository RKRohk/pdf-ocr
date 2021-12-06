package main

func main() {
	// fileNames := listFiles()

	// var wg sync.WaitGroup

	// limiterChannel := make(chan struct{}, simultaneousFiles)
	// for range make([]struct{}, simultaneousFiles) {
	// 	limiterChannel <- struct{}{}
	// }

	// func() {
	// 	for index, fileName := range fileNames {
	// 		<-limiterChannel
	// 		wg.Add(1)
	// 		go ocr(&wg, fileName, limiterChannel)
	// 		fmt.Println(index)
	// 	}
	// }()

	// fmt.Println("Waitin for all goroutines to finish")
	// wg.Wait()

}
