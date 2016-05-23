package bitmap
import (
    "testing"
    "time"
    "math/rand"
)

func TestBitMap(t *testing.T) {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    bm := NewBitMap(100000)
    inData := make([]int64, 2000)
    for i := 0; i < 2000; i++ {
        inData = append(inData, r.Int63n(100000))
    }

    for _, num := range inData {
        bm.Set(uint64(num))
    }
    for _, num :=range inData {
        res := bm.IsSet(uint64(num))
        if res != true {
            t.Errorf("bitmap exists some problems. Got %d, exprcted true", res)
        }
    }
}
