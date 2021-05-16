package fxtabs

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pierrec/lz4/v4"
)

func getRawJSON(filePath string) ([]byte, error) {
	fp, err := os.Open(filePath)
	if err != nil {
		return []byte{}, fmt.Errorf("error opening %v: %v", filePath, err)
	}
	defer fp.Close()

	// Checking file magic for jsonlz4.
	buf := make([]byte, 8)
	_, err = fp.Read(buf)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading magic from %v: %v", filePath, err)
	}
	if !bytes.Equal(buf, []byte("mozLz40\x00")) {
		return []byte{}, fmt.Errorf("invalid header magic from %v: %v", filePath, err)
	}

	// Getting uncompressed size from next 4 bytes.
	buf = make([]byte, 4)
	_, err = fp.Read(buf)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading uncompressed size from %v: %v", filePath, err)
	}
	decompressedSize := binary.LittleEndian.Uint32(buf)

	// Getting compressed blob.
	src, err := ioutil.ReadAll(fp)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading compressed blob from %v: %v", filePath, err)
	}

	// Uncompressing blob.
	dst := make([]byte, decompressedSize)
	_, err = lz4.UncompressBlock(src, dst)
	if err != nil {
		return []byte{}, fmt.Errorf("error uncompressing %v: %v", filePath, err)
	}

	return dst, nil
}

// OpenTabs collect Firefox open tabs from a recovery file (recovery.jsonlz4).
// This file is written almost in real time by Firefox (there are some seconds'
// delay).
func OpenTabs(filePath string) ([]FirefoxTab, error) {
	tabsJson, err := getRawJSON(filePath)
	if err != nil {
		return nil, err
	}

	var data jsonData
	err = json.Unmarshal(tabsJson, &data)
	if err != nil {
		return nil, fmt.Errorf("error decoding json: %v", err)
	}

	var tabs []FirefoxTab
	for _, w := range data.Windows {
		for _, t := range w.Tabs {
			index := t.Index - 1
			var cur uint = 0
			for _, e := range t.Entries {
				if cur == index {
					tabs = append(tabs, FirefoxTab{Title: e.Title, URL: e.URL})
					break
				}
				cur++
			}
		}
	}

	return tabs, nil
}
