package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println()
}

type Codec struct {
	store map[string]string
}

func Constructor() Codec {
	return Codec{store: make(map[string]string)}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	tiny := fmt.Sprintf("%s", sha256.New().Sum([]byte(longUrl)))
	this.store[tiny] = longUrl
	return tiny
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.store[shortUrl]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
