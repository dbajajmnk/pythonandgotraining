// Lab 3: Range and close
	// numCh := make(chan int)
	// go func() {
	// 	for i := 1; i <= 3; i++ {
	// 		numCh <- i
	// 	}
	// 	close(numCh)
	// }()

	// for v := range numCh {
	// 	fmt.Println("Range received:", v)
	// }

	// // Lab 4: Directional channels
	// done := make(chan bool)
	// go producer(done)
	// <-done

	// fmt.Println("=== Done ===")