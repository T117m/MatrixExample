package main

func InverseProduct(a, b Matrix) (Matrix, bool) {
	var (
		aInvF = InverseFuture(a)
		bInvF = InverseFuture(b)
		aInv  = <-aInvF
		bInv  = <-bInvF
	)

	return Product(aInv, bInv)
}

func InverseFuture(a Matrix) <-chan Matrix {
	future := make(chan Matrix)
	go func() { future <- Inverse(a) }()
	return future
}
