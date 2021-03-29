package applicationserver

import "github.com/andrewflash/benam-lora-network-server/internal/api/client/asclient"

var pool asclient.Pool

// SetPool sets the given Pool.
func SetPool(p asclient.Pool) {
	pool = p
}

// Pool returns teh configured application-server Pool.
func Pool() asclient.Pool {
	return pool
}

// Setup sets up the application-server pool.
func Setup() error {
	pool = asclient.NewPool()
	return nil
}
