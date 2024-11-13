# Huffnan Encoder in Golang

### Encoder Struct
```golang
type HuffmanEncoder struct{ 
	Tree string
	Code string
} 
```

### Encoder Function
```golang
func (encoder *HuffmanEncoder)Encode(src string, binaryEncodingLength int) string
```

### Decoder Function
```golang
func (encoder *HuffmanEncoder)Decode(src string, Tree string, binaryEncodingLength int) string 
```

### Usage 
```golang
type HuffmanEncoder = hme.HuffmanEncoder

func main() {
	var encoder HuffmanEncoder
	binary:=encoder.Encode("this huffman encoder works", 8)
	code:=encoder.Code
	original:=encoder.Decode(encoder.Code, encoder.Tree, 8)
  
	fmt.Println(binary)
	fmt.Println(code)
	fmt.Println(encoder.Tree)
	fmt.Println(original)
}
```