package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"debug/pe"
	"encoding/binary"
	"io/ioutil"
	"log"
	"math"
	"os"
)

// CalcEntropy is the function that calculates entropy of entered data
func CalcEntropy(data []byte) float64 {
	bytesCount := make(map[byte]int)
	dataLen := 0
	for _, b := range data {
		dataLen++
		bytesCount[b]++
	}

	bytesEntropy := make(map[byte]float64)
	var dataEntropy float64
	for k, v := range bytesCount {
		probability := float64(v) / float64(dataLen)
		bytesEntropy[k] = -math.Log2(probability) * probability
		dataEntropy += bytesEntropy[k]
	}
	return dataEntropy
}

// IsDataPacked returns true if data packed and its its probability between [0; 1]
func IsDataPacked(entropy float64) (bool, float32) {
	// TODO
	return true, 0.5
}

// PrintHashes prints SHA1, MD5 and SHA256 hashes of file byte array
func PrintHashes(file []byte) {
	log.Printf("\tMD5: %x", md5.Sum(file))
	log.Printf("\tSHA1: %x", sha1.Sum(file))
	log.Printf("\tSHA256: %x", sha256.Sum256(file))
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: ecalc.exe somePEfile.exe")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("Failed to open file:", err)
	}

	fileBuf, err := ioutil.ReadAll(file)
	log.Println("File data:")
	log.Println("\tName:", file.Name())
	log.Println("\tSize:", len(fileBuf))
	log.Println("\tEntropy:", CalcEntropy(fileBuf))
	PrintHashes(fileBuf)

	pefile, err := pe.NewFile(file)
	defer pefile.Close()
	if err != nil {
		log.Fatalln("Failed to read PE file header:", err)
	}

	var peHeaderBuf bytes.Buffer
	binary.Write(&peHeaderBuf, binary.BigEndian, pefile.FileHeader)
	binary.Write(&peHeaderBuf, binary.BigEndian, pefile.OptionalHeader)
	log.Println("PE header entropy:", CalcEntropy(peHeaderBuf.Bytes()))

	log.Println("Sections entropy:")
	for _, section := range pefile.Sections {
		data, _ := section.Data()
		log.Printf("\t[%s] %f\n", section.Name, CalcEntropy(data))
	}
}
