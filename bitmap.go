//Author: Liuzekun
//Date  : 2016.06.23
//Desc  : a bitmap

package bitmap

import (
    "errors"
)

const BitMapMaxSize = 0x01 << 63

type BitMap struct {
    bitSeq []byte     //比特序列
    bitSize uint64    //比特序列的大小
    maxPos uint64     //标记为1的最大位置
}

func NewBitMap(size uint64) *BitMap {
    if size > BitMapMaxSize {
        size = BitMapMaxSize
    }
    byteBlocks :=  (size / 8) + 1
    return &BitMap{bitSeq: make([]byte, byteBlocks), bitSize: 8 * byteBlocks}
}

//desc:   置位
//para:   @num 给定数值
//return: @置位成功返回nil
func (bm *BitMap) Set(num uint64) error {
    if num > BitMapMaxSize {
        return errors.New("the given num is beyong the range")
    }
    index, pos := num / 8, num % 8
    bm.bitSeq[index]  |= 0x01 << pos
    if bm.maxPos < num {
        bm.maxPos = num
    }
    return nil
}

func (bm *BitMap) IsSet(num uint64) bool {
    index, pos := num /8, num % 8
    if num > BitMapMaxSize {
        return false
    }
    return ((bm.bitSeq[index] >> pos) & 0x01 == 0x01)
}

func (bm *BitMap) GetMaxPos() uint64 {
    return bm.maxPos
}
