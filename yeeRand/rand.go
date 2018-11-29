/**
 * Created by angelina on 2017/5/3.
 */

package yeeRand

import (
	cryptoRand "crypto/rand"
	mathRand "math/rand"
	"time"
)

const Seed = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const NumSeed = "0123456789"
const LowerAlphaSeed = "abcdefghijklmnopqrstuvwxyz"
const UpperAlphaSeed = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandString
// return the rand string by given length.
func RandString(length int) string {
	var bytes = make([]byte, 2*length)
	var outBytes = make([]byte, length)
	_, err := cryptoRand.Read(bytes)
	if err != nil {
		panic(err)
	}
	mapLen := len(Seed)
	for i := 0; i < length; i++ {
		outBytes[i] = Seed[(int(bytes[2*i])*256+int(bytes[2*i+1]))%(mapLen)]
	}
	return string(outBytes)
}

// RandNum
// return the rand num string by given length.
func RandNum(length int) string {
	var bytes = make([]byte, 2*length)
	var outBytes = make([]byte, length)
	_, err := cryptoRand.Read(bytes)
	if err != nil {
		panic(err)
	}
	mapLen := len(NumSeed)
	for i := 0; i < length; i++ {
		outBytes[i] = NumSeed[(int(bytes[2*i])*256+int(bytes[2*i+1]))%(mapLen)]
	}
	return string(outBytes)
}

// RandIntBetween
// return random int between two given int.
func RandIntBetween(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	mathRand.Seed(time.Now().UnixNano())
	return mathRand.Intn(max-min) + min
}

// RandIntBetween
// return random int64 between two given int.
func RandInt64Between(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	mathRand.Seed(time.Now().UnixNano())
	return mathRand.Int63n(max-min) + min
}

// RandNumber
// return no-repeat number by given start end and count[min,max).
func RandNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}
	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}
