
package hme

import (
	"sort"
)

type HuffmanEncoder struct{ 
	Tree string
	Code string
} 

func bin2rune(src string)rune{
	value,power:=0,1
	for i,l:=len(src)-1,(len(src)-1); i>=0; i-- {
		if src[i]=='1'{
			power = 1<<(l-i)
			value+=power
		}
	}
	return rune(value)
}

func rune2bin(src rune, binlen int)string{
	var res []byte = make([]byte,binlen)
	for i,bit:=0,byte(0); i<binlen; i++ {
		bit = '0'
		if src&1==1 {bit='1'}
		res[(binlen-1)-i] = bit
		src>>=1
	}
	return string(res)
}



func (encoder *HuffmanEncoder)Encode(src string, binaryEncodingLength int) string {
	var freqTable map[string]int = make(map[string]int)
	var CodeTable map[string]string = make(map[string]string)
	var keyTable []string
	
	for _,v:= range src {
		if freqTable[string(v)]==0 {keyTable = append(keyTable, string(v))} 
		freqTable[string(v)]+=1
	}
	sort.SliceStable(keyTable, func(i,j int)bool{return freqTable[keyTable[i]] < freqTable[keyTable[j]]})
	
	CodeTable[keyTable[0]] = "1"
	for i,bin:=0,""; i<len(keyTable); i++ {
		encoder.Tree+= keyTable[i]
		if i<len(keyTable)-1 {
			bin+= "0"
			CodeTable[keyTable[i+1]] = bin+"1"
		}
	}
	var bin, binary string 
	for i:= 0; i<len(src); i++ {
		bin+=CodeTable[string(src[i])]
		binary+=CodeTable[string(src[i])]
		if len(bin)>=binaryEncodingLength {
			encoder.Code+= string(bin2rune(bin[:binaryEncodingLength]))
			bin = bin[binaryEncodingLength:len(bin)]
		}
	}
	 if len(bin)>0{
		var bits string
		for i:=0; i<len(bin); i++ {
			bits+= string(bin[i])
			if len(bits)==binaryEncodingLength{
				encoder.Code+= string(bin2rune(bits))
				bits = ""
			}  
		}
		if len(bits)>0 {
			s,l:="",binaryEncodingLength-len(bits)
			for i:=0; i<l; i++{s+="0"}
			encoder.Code+= string(bin2rune(bits+s))
		}
	}
	return binary
}

func (encoder *HuffmanEncoder)Decode(src string, Tree string, binaryEncodingLength int) string {
	res, i:= "", 0
	for _,char:=range src{
		bin:= rune2bin(rune(char), binaryEncodingLength) 
		for _,bit:=range bin {
			if bit=='1' {
				if i>len(Tree)-1 {return res}
				res+=string(Tree[i])
				i=0
			} else {i+=1}
		}
	}
	return res
}