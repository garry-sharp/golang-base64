package main

import (
	"fmt"
	"os"
	"strings"
)

const charMap = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

//function to turn a value into a base64 value

/* The following "masks" are used here with the & operator for bitwise operators
00001111 = 15
__110000 = 48
00111111 = 63

For the remaining operators shift operators are sufficient since the LSB is set to 0 in a shift

The 1, 2 or 3 bytes in this loop are assigned to 4 x 6 bit words (represented as 4 bytes with the 2 MSBs always set to 00)

A lazy implementation is adding the '=' character as the 65th character in a map. Since the range of 6 bits is 0-63 the 64th index
will never be reached (other than manually). Since we actually have the extra 2 bits regardless we can safely set that manually in the code.
*/

func base64Encode(inp string) string {
	b := []byte(inp)
	var outP string
	for i := 0; i < len(b); i += 3 {
		gap := min(3, len(b)-i)
		sec := b[i : i+gap]
		var f1, f2, f3, f4 byte

		//F1 will always be set since there will always be 6 bits at least within the first byte
		f1 = sec[0] >> 2

		//if only 1 byte then f2 will be the last 2 digits of the byte left shifted by 4
		//otherwise it is the last 2 digits of the first world left shifted by 4 and the first 4 bits of the second byte right shifted by 4
		f2 = (sec[0] << 4) & 48
		if len(sec) >= 2 {
			f2 += (sec[1] >> 4)
		}

		//f3 and f4 defaulting to padding characters
		f3, f4 = 64, 64

		//if only 2 bytes then f3 is only the last 4 bits of byte 2 left shifted by 2
		if len(sec) == 2 {
			f3 = (sec[1] & 15) << 2
			//if 3 bytes then f2 is the last 4 bits of byte 2 left shifted by 2 plus the 2 MSBs of byte 3 right shifted by 6
			//f4 in this case is simply the last 6 bits (8 bit byte masked by & 00111111 or 63)
		} else if len(sec) == 3 {
			f3 = ((sec[1] & 15) << 2) + sec[2]>>6
			f4 = sec[2] & 63
		}

		outP += string(charMap[f1]) + string(charMap[f2]) + string(charMap[f3]) + string(charMap[f4])
	}
	return outP
}

func base64Decode(inp string) string {
	//Note this implementation requires '=' padding characters such that len(inp) % 4 == 0 is true
	if len(inp)%4 != 0 {
		panic("This implementation does not work without padding characters")
	}
	var outP string = ""
	for o := 0; o+4 <= len(inp); o += 4 {
		var b1, b2, b3 byte
		chars := inp[o : o+4]
		for i, c := range chars {
			for j, mapChar := range charMap {
				/*
					Basically just a reverse engineering of the encode function. In this case a 00111111 mask is applied to every
					6 bit word. Since the value is stored in a native golang byte type the '=' character being the 64th index in the map
					would result in a 7 bit word of 01000000. Conveniently when & with 00111111 it will create a null ascii character.
					This has the added benefit of not interfering with any number in the range of 0 - 63
				*/

				if mapChar == c {
					if i == 0 {
						b1 = (byte(j) & 63) << 2
					}
					if i == 1 {
						b1 += (byte(j) & 63 >> 4)
						b2 = (byte(j) & 63 & 15) << 4
					}
					if i == 2 {
						b2 += ((byte(j) & 63) >> 2)
						b3 = ((byte(j) & 63 & 3) << 6)
					}
					if i == 3 {
						b3 += (byte(j) & 63)
					}
				}
			}
		}
		outP += string([]byte{b1, b2, b3})

	}
	return outP
}

func main() {
	if len(os.Args) < 3 {
		panic("Application should be of the form encode/decode text eg. encode hello")
	}

	inp := strings.Join(os.Args[2:], " ")

	if os.Args[1] == "encode" {
		fmt.Println(base64Encode(inp))
	} else if os.Args[1] == "decode" {
		fmt.Println(base64Decode(inp))
	} else {
		panic(fmt.Sprintf("Unrecognized command \"%s\", only commands \"encode\",\"decode\" are accepted", os.Args[1]))
	}

}
