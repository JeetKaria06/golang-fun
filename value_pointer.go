package main

type BigByteStruct struct {
	Buf [1 << 16]byte // 65536 bytes, 64 KBs
	// quite a big size for a variable
}

var obj BigByteStruct = BigByteStruct{}

// passing by value
func fooPBV(obj BigByteStruct) {

}

// passing by pointer
func fooPBP(obj *BigByteStruct) {

}

func foo() {
	fooPBV(obj)
	fooPBP(&obj)
}
