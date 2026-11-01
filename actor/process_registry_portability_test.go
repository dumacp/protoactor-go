package actor

import (
	"strconv"
	"testing"

	murmur32 "github.com/twmb/murmur3"
)

func TestSliceMapHighBitHash(t *testing.T) {
	s := newSliceMap()
	for i := 0; i < 10000; i++ {
		key := "arm32-regression-" + strconv.Itoa(i)
		hash := murmur32.Sum32([]byte(key))
		if hash < 1<<31 || hash%uint32(len(s.LocalPIDs)) == 0 {
			continue
		}
		bucket := s.GetBucket(key)
		bucket.Set(key, "present")
		want := s.LocalPIDs[hash%uint32(len(s.LocalPIDs))]
		if got, ok := want.Get(key); !ok || got != "present" {
			t.Fatalf("hash %08x was stored in the wrong bucket", hash)
		}
		return
	}
	t.Fatal("could not find a hash with its high bit set")
}
