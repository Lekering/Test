package main

func OrDoneRepeat(inCh chan int, closeCh chan struct{}) chan int {
	outCh := make(chan int)

	go func() {
		defer close(outCh)
		for {
			select {
			case <-closeCh:
				return
			default:
				select {
				case v, ok := <-inCh:
					if !ok {
						return
					}
					outCh <- v
				case <-closeCh:
					return
				}
			}
		}
	}()

	return outCh
}
