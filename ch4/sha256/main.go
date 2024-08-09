// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256"
	"fmt"

	"gopl.io/ch2/popcount"
) //!+

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8

        fmt.Printf("There are %d bits difference between the two hashes.\n", compareHashes(&c1, &c2))
}

func compareHashes(a *[32]byte, b *[32]byte) int {
    var diff [32]byte
    for i := range diff {
        diff[i] = a[i] & b[i]
    }

    count := 0
    for i := 0; i < 4; i++ {
        var number uint64
        for j := 0; j < 8; j++ {
            idx := 8 * i + j
            number += uint64(diff[idx]) << (8*(7-j))
        }
        count += popcount.PopCount(number)
    }
    return count
}

//!-
