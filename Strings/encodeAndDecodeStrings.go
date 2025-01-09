package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	var encoded strings.Builder

	for _, str := range strs {
		encoded.WriteString(strconv.Itoa(len(str)) + "#" + str)
	}
	return encoded.String()
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	var decoded []string
	i := 0
	for i < len(strs) {
		j := strings.Index(strs[i:], "#")
		if j == -1 {
			break
		}

		j += i
		length, _ := strconv.Atoi(strs[i:j])
		i = j + 1

		decoded = append(decoded, strs[i:i+length])
		i += length
	}
	return decoded
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));

func main() {
	var codec Codec

	fmt.Println(codec.Decode(codec.Encode([]string{"Hello", "World"})))

	fmt.Println(codec.Decode(codec.Encode([]string{""})))
}
