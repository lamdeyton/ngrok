// +build release

package static

import (
	"bytes"
	"compress/gzip"
	"io"
	"reflect"
	"unsafe"
)

var _PageHtml = ""+
"\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x58\x51\x8f\xe2\x36"+
"\x10\x7e\xef\xaf\xb0\x7c\xfb\x00\xea\x62\x77\x6f\xd5\xab\xca\x01"+
"\x2f\x27\x55\x77\xd2\x55\x77\xdd\x6d\x7f\x80\x89\x0d\x31\xeb\xd8"+
"\xa9\xed\xb0\xa0\x15\xff\xbd\xe3\x98\x84\x00\x01\xc2\xde\x3d\xf4"+
"\xe1\x22\xa1\x8d\xed\x99\x6f\xc6\xe3\x6f\x66\x9c\x1d\xa5\x3e\x53"+
"\x93\x9f\x10\x3c\xa3\x54\x30\x1e\x5f\xcb\xa1\x92\xfa\x09\xa5\x56"+
"\xcc\xc6\x98\x52\x2d\x3c\xd7\x8c\x4c\x8d\xf1\xce\x5b\x96\x27\x5c"+
"\x93\xc4\x64\xd4\x3f\x4b\xef\x85\x1d\xd4\x0b\xf4\x2d\xb9\x27\x77"+
"\x34\x71\x8e\xd6\x73\x03\x90\x9c\x4a\x2d\x38\xc9\x24\xa8\x39\x87"+
"\x91\x15\x6a\x8c\x9d\x5f\x2b\xe1\x52\x21\x3c\x3e\xb4\x7b\xb8\xbe"+
"\x75\x24\xf5\x3e\x1f\x52\xba\x66\x9a\x8b\x15\x71\x9e\xa6\x72\x9e"+
"\x2a\xf8\xf9\x85\xa3\xbf\x91\x7b\x1a\x55\x28\x17\x33\x56\x28\x5f"+
"\xdb\x6b\xc0\xbb\xc4\xca\xdc\x23\x67\x93\x0e\x70\xf5\xb8\x44\x5a"+
"\x00\xd0\x88\x46\x80\xf3\x88\x89\xe1\x82\x2c\xfe\x2d\x84\x5d\x97"+
"\x61\x8a\xaf\x83\x3b\xf2\x3b\xb9\xbb\x0c\x35\x49\xd5\xc2\x11\xa9"+
"\xa5\xff\x58\xd9\x97\x7a\xfe\x45\x7f\x36\x8c\xf7\xfa\xef\x4f\xbb"+
"\xe0\xd7\xb9\x18\x63\x2f\x56\x9e\x2e\xd8\x92\xc5\xd9\xc6\xe6\xc3"+
"\x73\xd3\x9b\x15\x3a\xf1\xd2\xe8\x5e\x1f\xbd\xec\x2d\xc5\x65\x5c"+
"\x28\x62\x05\x38\xec\x3c\xd1\x6c\x89\x18\xee\x93\x44\xc9\xe4\xe9"+
"\xbc\x62\x54\xf6\xa9\x74\x41\xdc\x38\x50\xef\x61\x2e\x97\xa0\x3d"+
"\x93\x9a\xf7\x30\xb1\xec\xf9\x16\x91\xc0\x32\x61\x1d\xbc\xb9\x22"+
"\xcb\x98\x5d\x83\x40\x2a\xb9\x80\x8d\xbd\x02\x13\xa3\x9f\x6b\x09"+
"\xe6\xbd\xed\x61\xcf\xec\x1c\x08\xd3\xef\x13\x97\x9a\xe7\xce\xa8"+
"\x85\xaa\x41\x95\x24\x0c\xf6\xb9\x14\x30\x63\x45\x66\x96\xe2\x83"+
"\x62\xce\xf5\x70\x35\x7b\x1e\x32\x67\x56\x68\xdf\x03\x7f\x38\xef"+
"\xa4\x68\x85\x2f\xac\x46\x33\xa6\x9c\x38\x96\xd8\x80\x56\xdb\x29"+
"\x81\x6b\xb9\x62\xeb\x2b\x0e\x87\xb0\x05\x5b\xf5\xda\x17\x51\xc9"+
"\x9d\x21\xc2\x5f\xbf\x3c\xfe\x8d\x6f\x4f\xc8\x14\x56\x81\x08\x0d"+
"\x24\xa7\x52\xd3\xad\x07\xa7\xa4\x39\xf3\x6c\x88\x5e\x90\x5f\x69"+
"\xc9\x87\xc7\x21\x27\xb0\x80\xeb\x63\x0b\x42\xb8\x8f\x36\xad\x60"+
"\x9b\x13\xa1\x7b\x65\x3c\x4f\x8d\xdb\x32\x2b\x14\x94\x66\x62\xed"+
"\x97\x93\xf0\x4c\x0d\x5f\xc3\x36\x81\xcb\x73\xa9\x07\xde\xe4\x43"+
"\xf4\xeb\x2f\xf9\xea\xfd\xc1\x5e\x20\xad\x60\xff\xde\xd8\x35\x9a"+
"\x20\x25\x41\x43\x1b\x2d\xde\x57\x7a\x53\xe3\xbd\xc9\x86\xe8\x6d"+
"\xa9\x9a\x03\x79\x20\xe7\x0f\x66\xa7\xc6\xf2\xb2\xd8\xc6\xc9\xbb"+
"\x7c\x85\x9c\x51\x92\xa3\x37\x49\x92\x1c\xda\x6b\x64\x1a\x64\x1e"+
"\x98\xe3\xd2\x85\xf3\x1a\x6e\xed\xee\x4b\x7b\x36\x55\x22\x70\x97"+
"\x65\x0e\x64\x67\x46\xfb\xc1\x8c\x65\x52\x81\xfc\x07\x53\x58\x29"+
"\xec\x2d\xca\x8c\x36\x2e\x67\xc9\x9e\xf6\x28\x16\xdd\x6d\x13\xa1"+
"\xb1\x8b\xc4\x41\x88\x4c\x23\x96\x90\xb9\x28\x09\x19\x31\xc6\x09"+
"\xe0\x33\x68\x09\xf6\x20\x96\x4d\x19\x28\x3e\x53\x66\x51\xfc\x33"+
"\x90\x7a\x09\x9b\x11\xd5\x70\x26\x57\x82\x87\x60\x1f\x00\xb4\x83"+
"\x80\xf6\xb1\xad\x6b\xfc\xda\x93\x67\x95\xf4\xd4\x42\xe7\xa8\x7a"+
"\xd3\x1b\x3c\xd1\x73\x6b\x9e\x46\x94\x9d\xd1\x2d\x54\xc3\xb3\x33"+
"\x46\x4a\x61\x60\xc9\x56\x78\x5b\x41\x26\x60\xbb\xb6\xf6\x49\x4f"+
"\x4d\xa1\x39\x7a\x88\xc5\xda\x05\xc3\x23\xaa\xe4\x45\xd0\x3d\x94"+
"\x2f\x85\xff\x1e\x30\x1f\x8c\x9e\xc9\x79\x61\x59\xa8\x41\x97\x21"+
"\x46\xb4\x50\x27\x4e\x83\xc2\x71\xb4\x9c\xe9\xf1\x74\xdb\xd4\x2e"+
"\xbc\x55\xae\x15\xba\xa4\x27\x3f\x88\xf5\xcb\x0b\x82\xc3\x9b\x0b"+
"\x44\x3e\x42\x35\xab\xf6\x4e\x1e\xa1\x96\x0a\xb4\x39\xae\x42\x8d"+
"\xb3\x08\x65\x2b\x16\xb5\x31\x06\x18\xf2\x89\x83\x42\x07\x6e\x59"+
"\xf3\x7c\x8e\x55\x0d\x49\xc8\x31\xfd\xee\x02\x39\xc0\x32\xdc\xbd"+
"\x52\x44\xc0\xf7\x36\x87\x8f\xf0\xd3\xfb\x49\xf0\xf6\x4f\xe1\x53"+
"\x13\x3c\x0e\x08\xe4\x9f\x87\xcf\xe4\x2b\x03\x98\xcd\x06\x52\xf7"+
"\xfe\xbc\xc9\x12\x66\x17\xe1\xed\x25\x21\x64\x64\xf8\x0d\x72\xa9"+
"\xd4\x61\x6d\x3c\x09\xd3\x4a\xed\xd8\xbb\x61\xff\xdb\x9b\xc1\x8e"+
"\x5e\x8f\x71\xa6\x1b\x37\x1b\x46\x9a\x1c\xad\xe1\xb7\x85\x11\x4f"+
"\x3e\xc6\x97\xef\x83\x0a\x45\x16\x4f\x1e\xd8\xf3\xd5\x68\x55\x20"+
"\xf2\x42\xa9\x81\x0d\xb7\x3d\x08\xc6\xb4\x80\x1a\xaf\x77\xa1\x0e"+
"\x95\x1b\x4d\xbd\x0e\xbf\x41\x6e\x65\x19\x9f\xc9\x43\x39\x3f\xa2"+
"\x51\xba\xa3\xd9\x98\x7b\x97\xc5\x9a\xd4\x65\xe7\xa8\xdb\x7c\x2a"+
"\x5a\xde\x84\x9e\x33\x1c\x23\x5e\x64\x55\x76\x21\xd2\x85\xa7\xb5"+
"\xf9\xdc\x8a\xc9\x28\x5c\xa4\xeb\x8c\x86\x44\xc5\x81\xc3\x25\x76"+
"\x20\x6c\x58\x85\x4d\x07\xc9\xae\xce\x09\xcd\x3b\x65\x4b\xac\x2d"+
"\x57\x05\xa9\xa6\xd5\x55\x81\x22\x91\x84\x57\x85\x26\x7d\xb7\xa3"+
"\x2e\xbc\x77\x57\x2c\xbb\x7c\x5d\xc7\xca\x41\x6c\xf9\x1d\x9d\xde"+
"\x3a\x1e\x2b\xe7\xcd\x93\x58\xdf\xa2\x9b\x25\x53\x85\x08\x47\x4d"+
"\xae\xd8\x42\xf4\xc6\x76\xb7\xba\x53\x4a\x27\xc1\x85\x60\x3c\x14"+
"\xb1\x11\x85\x89\x57\xa0\xf0\x88\x12\x9d\x8f\x38\xfc\x3a\x1c\xd0"+
"\xb8\xc2\xff\xee\xd4\x6b\xe0\x87\x03\xfa\x3f\x30\xbb\xaa\xc7\x57"+
"\x32\x3b\x34\x97\xbf\xc2\x87\xef\xb5\xe4\x8e\x4a\x5f\x4b\x66\xfe"+
"\x60\xf8\x0f\x86\xef\xc3\x77\x15\xf7\x22\x83\xee\xe8\x05\xc2\xe1"+
"\x3b\x04\x77\x6b\x3f\x27\xae\x9f\xed\x9e\xbc\x12\xe4\x1b\xae\x7a"+
"\x2e\xef\x78\xd7\xab\xf0\x41\xb7\x7c\xfb\xc3\xd8\x47\xcf\x7c\xe1"+
"\x10\xd9\xfe\x0d\x77\xd6\x70\x05\xac\x87\xf1\x02\x78\x11\xbc\xbd"+
"\xc3\xbb\xdc\x68\xf8\x3a\xeb\xdc\xe2\xbf\x77\x7b\xef\xc6\xbe\x6f"+
"\x3a\xba\x93\x5f\x26\x87\x97\xae\x76\x2b\xfb\x9f\x3c\x0d\x30\xb8"+
"\xbe\x95\xdf\xc9\x10\xff\xf2\x9f\xb1\xff\x05\x00\x00\xff\xff\x1b"+
"\x0a\x19\xd7\x94\x15\x00\x00"

// PageHtml returns raw, uncompressed file data.
func PageHtml() []byte {
	var empty [0]byte
	sx := (*reflect.StringHeader)(unsafe.Pointer(&_PageHtml))
	b := empty[:]
	bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bx.Data = sx.Data
	bx.Len = len(_PageHtml)
	bx.Cap = bx.Len

	gz, err := gzip.NewReader(bytes.NewBuffer(b))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var buf bytes.Buffer
	io.Copy(&buf, gz)
	gz.Close()

	return buf.Bytes()
}
