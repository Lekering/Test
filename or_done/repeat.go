package main

func OrDoneRepeat[T any](inCh chan T, closeCh chan struct{}) chan T {
	outCh := make(chan T)

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
