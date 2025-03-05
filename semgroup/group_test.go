package semgroup_test

import (
	"testing"
	"time"

	"github.com/caarlos0/sync/semgroup"
)

func TestGroupNoLimit(t *testing.T) {
	var g semgroup.Group
	for i := 0; i < 10; i++ {
		i := i
		g.Go(func() {
			t.Log(i)
			time.Sleep(time.Duration(i) * 100 * time.Millisecond)
		})
	}
	g.Wait()
}

func TestGroupLimit(t *testing.T) {
	var g semgroup.Group
	g.SetLimit(2)
	for i := 0; i < 10; i++ {
		i := i
		g.Go(func() {
			t.Log(i)
			time.Sleep(time.Duration(i) * 100 * time.Millisecond)
		})
	}
	g.Wait()
}
