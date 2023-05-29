package ciphers

import (
	"fmt"
	"github.com/lxn/walk"
	. "main/const"
	. "main/tools"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func AtbashEncrypt(text string, keys []*walk.TextEdit) string {
	var num, secNum int
	var result string
	for _, v := range text {
		num = IndexOf(string(v), Dictionary)
		if num != -1 {
			secNum = len(Dictionary) - 1 - num
			result += GetElement(Dictionary, secNum)
		} else {
			result += string(v)
		}

	}
	return result
}

func AtbashDecrypt(text string, keys []*walk.TextEdit) string {
	var num, secNum int
	var result string
	for _, v := range text {
		num = IndexOf(string(v), Dictionary)
		if num != -1 {
			secNum = len(Dictionary) - 1 - num
			result += GetElement(Dictionary, secNum)
		} else {
			result += string(v)
		}

	}
	return result
}

func CaesarEncrypt(text string, keys []*walk.TextEdit) string {
	userKey := keys[0].Text()
	key, _ := strconv.Atoi(userKey)
	key = key % len(Dictionary)
	var num, secNum int
	var result string
	for _, v := range text {
		num = IndexOf(string(v), Dictionary)
		if num+key > len(Dictionary)-1 {
			secNum = num + key - len(Dictionary)
		} else if num+key < 0 {
			secNum = len(Dictionary) + num + key
		} else {
			secNum = num + key
		}
		result += GetElement(Dictionary, secNum)
	}
	return result
}

func CaesarDecrypt(text string, keys []*walk.TextEdit) string {
	userKey := keys[0].Text()
	key, _ := strconv.Atoi(userKey)
	key = key % len(Dictionary)
	key *= -1
	var num, secNum int
	var result string
	for _, v := range text {
		num = IndexOf(string(v), Dictionary)
		if num+key > len(Dictionary)-1 {
			secNum = num + key - len(Dictionary)
		} else if num+key < 0 {
			secNum = len(Dictionary) + num + key
		} else {
			secNum = num + key
		}
		result += GetElement(Dictionary, secNum)
	}
	return result
}

func PolibiusEncrypt(text string, keys []*walk.TextEdit) string {
	var mas []string
	var num int
	for _, v := range text {
		for j := 0; j < 6; j++ {
			for i := 0; i < 6; i++ {
				if j*6+i >= len(Dictionary) {
					break
				}
				for _, let := range strings.Split(Dictionary[j*6+i], "/") {
					if let == string(v) {
						num = (j+1)*10 + (i + 1)
						mas = append(mas, strconv.Itoa(num))
					}
				}
			}
		}
	}
	return strings.Join(mas, "")
}

func PolibiusDecrypt(text string, keys []*walk.TextEdit) string {
	var j string
	var mas []string
	var result string
	for i, sym := range text + "0" {
		if i%2 == 0 && i != 0 {
			mas = append(mas, j)
			j = ""
		}
		j += string(sym)
	}
	for _, v := range mas {
		a, _ := strconv.Atoi(string(v[0]))
		b, _ := strconv.Atoi(string(v[1]))
		a -= 1
		b -= 1
		result += GetElement(Dictionary, a*6+b)
	}
	return result
}

func TrithemiusEncrypt(text string, keys []*walk.TextEdit) string {
	j := 0
	var result string
	mas := GetTrithemiusTable()
	for _, v := range text {
		result += GetElement(mas[j], IndexOf(string(v), mas[0]))
		j++
		if j > len(Dictionary)-1 {
			j = 0
		}
	}
	return result
}

func TrithemiusDecrypt(text string, keys []*walk.TextEdit) string {
	j := 0
	var result string
	mas := GetTrithemiusTable()
	for _, v := range text {
		result += GetElement(mas[0], IndexOf(string(v), mas[j]))
		j++
		if j > len(Dictionary)-1 {
			j = 0
		}
	}
	return result
}

func BelazoEncrypt(text string, keys []*walk.TextEdit) string {
	key := keys[0].Text()
	i := 0
	var result string
	mas := GetTrithemiusTable()
	for _, v := range text {
		num := IndexOf(string(v), mas[0])
		for j := 0; ; j++ {
			if GetElement(mas[j], 0) == string([]rune(key)[i]) {
				result += GetElement(mas[j], num)
				i++
				if i > utf8.RuneCountInString(key)-1 {
					i = 0
				}
				break
			}
		}
	}
	return result
}

func BelazoDecrypt(text string, keys []*walk.TextEdit) string {
	key := keys[0].Text()
	var result string
	mas := GetTrithemiusTable()
	i := 0
	for _, v := range text {
		t := IndexOf(string([]rune(key)[i]), mas[0])
		num := IndexOf(string(v), mas[t])
		if num == -1 {
			fmt.Println(string(v), mas[t], text)
		}
		result += GetElement(mas[0], num)
		i++
		if i > utf8.RuneCountInString(key)-1 {
			i = 0
		}
	}
	return result
}

func VigenereCipherKeyEncrypt(text string, keys []*walk.TextEdit) string {
	key := keys[0].Text()
	i := 0
	var result string
	mas := GetTrithemiusTable()
	for _, v := range text {
		num := IndexOf(string(v), mas[0])
		for j := 0; ; j++ {
			if GetElement(mas[j], 0) == string([]rune(key)[i]) {
				result += GetElement(mas[j], num)
				key += GetElement(mas[j], num)
				i++
				if i > utf8.RuneCountInString(key)-1 {
					i = 0
				}
				break
			}
		}
	}
	return result
}

func VigenereCipherKeyDecrypt(text string, keys []*walk.TextEdit) string {
	key := keys[0].Text()
	var result string
	i := 0
	mas := GetTrithemiusTable()
	for _, v := range text {

		for j := 0; ; j++ {
			if GetElement(mas[j], 0) == string([]rune(key)[i]) {
				num := IndexOf(string(v), mas[j])
				result += GetElement(mas[0], num)
				key += string(v)
				i++
				break
			}
		}
	}
	return result
}

func VigenereEncrypt(text string, keys []*walk.TextEdit) string {
	key := keys[0].Text()
	i := 0
	key += text
	var result string
	mas := GetTrithemiusTable()
	for _, v := range text {
		num := IndexOf(string(v), mas[0])
		for j := 0; ; j++ {
			if GetElement(mas[j], 0) == string([]rune(key)[i]) {
				result += GetElement(mas[j], num)
				i++
				if i > utf8.RuneCountInString(key)-1 {
					i = 0
				}
				break
			}
		}
	}
	return result
}

func VigenereDecrypt(text string, keys []*walk.TextEdit) string {
	key := keys[0].Text()
	var result string
	mas := GetTrithemiusTable()
	i := 0
	for _, v := range text {
		t := IndexOf(string([]rune(key)[i]), mas[0])
		num := IndexOf(string(v), mas[t])
		if num == -1 {
		}
		result += GetElement(mas[0], num)
		key += GetElement(mas[0], num)
		i++
	}
	return result
}

func MatrixEncrypt(text string, keys []*walk.TextEdit) string {
	key := keys[0].Text()
	resultAr := []string{}
	matrix := GetMatrix(key)
	rand.Seed(time.Now().UnixNano())
	if utf8.RuneCountInString(text)%len(matrix)!=0{
		for i:=0;i<len(matrix)-utf8.RuneCountInString(text)%len(matrix);i++{
			text+= GetElement(Dictionary, rand.Intn(len(Dictionary)))
		}
	}
	i := 0
	matr := [][]float64{}
	for _, v := range text {
		t := float64(IndexOf(string(v), Dictionary) + 1)
		matr = append(matr, []float64{t})
		i++
		if i > len(matrix[0])-1 {
			i = 0
			for _, elem := range MultiplyMatrix(matrix, matr) {
				resultAr = append(resultAr, strconv.Itoa(int(elem[0])))
			}
			matr = [][]float64{}
		}
	}
	return strings.Join(resultAr, " ")
}

func MatrixDecrypt(text string, keys []*walk.TextEdit) string {
	key := keys[0].Text()
	var result string
	rand.Seed(time.Now().UnixNano())
	matrix := GetMatrix(key)
	inMatrix := InverseMatrix(matrix)
	matr := [][]float64{}
	masText := strings.Split(text, " ")
	textStartLen := len(strings.Split(text, " "))
	manyElems := textStartLen % len(matrix)
	enText := ""
	var i int
	var tempMat []float64
	if manyElems != 0 {
		testMat := append(tempMat, float64(i))
		testMatLen := len(testMat)
		for j := 0; j < len(matrix)-testMatLen; j++ {
			testMat = append(testMat, 0)
		}
		var rowTestMatr [][]float64
		for h := 0; h < len(testMat); h++ {
			rowTestMatr = append(rowTestMatr, []float64{0})
			rowTestMatr[h][0] = testMat[h]
		}

		newMultMatr := MultiplyMatrix(inMatrix, rowTestMatr)
		for _, elem := range newMultMatr {
			if elem[0] != 0 {
				enText += GetElement(Dictionary, int(elem[0]-1))
			} else {
				enText += GetElement(Dictionary, rand.Intn(len(Dictionary)))
			}
		}
		masText = masText[:len(masText)-1]
		text = strings.Join(masText," ")+" "+MatrixEncrypt(enText, keys)
	}
	i = 0
	for _, v := range strings.Split(text, " ") {
		t, _ := strconv.ParseFloat(v, 64)
		matr = append(matr, []float64{t})
		i++
		if i > len(inMatrix[0])-1 {
			i = 0
			for _, elem := range MultiplyMatrix(inMatrix, matr) {
				result += GetElement(Dictionary, int(elem[0])-1)
			}
			matr = [][]float64{}
		}
	}
	return result
}

func PlayfairEncrypt(text string, keys []*walk.TextEdit) string {
	key := keys[0].Text()
	tableWightLen := 6
	playfairTable := GetPlayfairTable(key, tableWightLen, SecondDictionary)
	tableHeightLen := len(playfairTable)
	_, textRunes := CheckPairs(text)
	var x1, y1, x2, y2 int
	var result string
	runes := []rune(textRunes)
	for i := 0; i < len(runes)-1; i += 2 {
		x1, y1 = IndexOfMas(string(runes[i]), playfairTable)
		x2, y2 = IndexOfMas(string(runes[i+1]), playfairTable)
		if x1 == x2 {
			result += GetElement(playfairTable[x1], (y1+1)%(tableWightLen))
			result += GetElement(playfairTable[x1], (y2+1)%(tableWightLen))
		} else if y1 == y2 {
			result += GetElement(playfairTable[(x1+1)%(tableHeightLen)], y1)
			result += GetElement(playfairTable[(x2+1)%(tableHeightLen)], y2)
		} else {
			result += GetElement(playfairTable[x1], y2)
			result += GetElement(playfairTable[x2], y1)
		}
	}
	return result
}

func PlayfairDecrypt(text string, keys []*walk.TextEdit) string {
	key := keys[0].Text()
	tableWightLen := 6
	playfairTable := GetPlayfairTable(key, tableWightLen, SecondDictionary)
	tableHeightLen := len(playfairTable)
	_, textRunes := CheckPairs(text)
	var x1, y1, x2, y2, minY1, minY2, minX1, minX2 int
	var result string
	runes := []rune(textRunes)
	for i := 0; i < len(runes)-1; i += 2 {
		x1, y1 = IndexOfMas(string(runes[i]), playfairTable)
		x2, y2 = IndexOfMas(string(runes[i+1]), playfairTable)
		if x1 == x2 {
			if minY1 = y1 - 1; minY1 < 0 {
				minY1 = tableWightLen - (y1-1)*-1
			}
			if minY2 = y2 - 1; minY2 < 0 {
				minY2 = tableWightLen - (y2-1)*-1
			}
			result += GetElement(playfairTable[x1], minY1)
			result += GetElement(playfairTable[x1], minY2)
		} else if y1 == y2 {
			if minX1 = x1 - 1; minX1 < 0 {
				minX1 = tableHeightLen - (x1-1)*-1
			}
			if minX2 = x2 - 1; minX2 < 0 {
				minX2 = tableHeightLen - (x2-1)*-1
			}
			result += GetElement(playfairTable[minX1], y1)
			result += GetElement(playfairTable[minX2], y2)
		} else {
			result += GetElement(playfairTable[x1], y2)
			result += GetElement(playfairTable[x2], y1)
		}
	}
	return result
}

func VerticalPermutationEncrypt(text string, keys []*walk.TextEdit) string {
	key := keys[0].Text()
	var seq []int
	for _, v := range key {
		seq = append(seq, IndexOf(string(v), Dictionary))
	}
	seq = GetSeq(seq)
	var result string
	newText := []rune(text)
	mas := GetVerticalpermutationTable(seq, newText)
	for j, _ := range seq {
		i := IndexOfInt(j, seq)
		for _, elem := range mas {
			result += elem[i]
		}
	}
	return result
}

func VerticalPermutationDecrypt(text string, keys []*walk.TextEdit) string {
	key := keys[0].Text()
	var seq []int
	for _, v := range key {
		seq = append(seq, IndexOf(string(v), Dictionary))
	}
	newText := []rune(text)
	seq = GetSeq(seq)
	di := int(math.Ceil(float64(utf8.RuneCountInString(text)) / float64(len(seq))))
	re := di*len(seq) - utf8.RuneCountInString(text)
	var mas [][]string
	var tempMas []string
	var result string
	for i := 0; i < di; i++ {
		for i := 0; i < len(seq); i++ {
			tempMas = append(tempMas, "")
		}
		mas = append(mas, tempMas)
		tempMas = []string{}
	}
	var kol []int
	for i := 0; i < len(seq); i++ {
		if re+i >= len(seq) {
			kol = append(kol, di-1)
		} else {
			kol = append(kol, di)
		}
	}
	l := 0
	for j, _ := range seq {
		i := IndexOfInt(j, seq)
		for k := 0; k < kol[i]; k++ {
			if l < len(newText) {
				mas[k][i] = string(newText[l])
			}
			l++
		}
	}
	for _, i := range mas {
		result += strings.Join(i, "")
	}
	return result
}

func CardanGrilleEncryptStart(text string, keys []*walk.TextEdit) string {
	keys[0].SetText(GenerateCardanKey(text))
	return CardanGrilleEncrypt(text, keys)
}

func CardanGrilleDecryptStart(text string, keys []*walk.TextEdit) string {
	return CardanGrilleDecrypt(text,keys)
}

func CardanGrilleEncrypt(text string, keys []*walk.TextEdit) string {
	key := keys[0].Text()
	matrix := GetMatrix(key)
	var result string
	rand.Seed(time.Now().UnixNano())
	runeText := []rune(text)
	newMatr := make([][]string, len(matrix))
	for i, _ := range newMatr {
		newMatr[i] = make([]string, len(matrix[0]))
	}
	var i int
	for x, xKeyMatr := range matrix {
		for y, yKeyMatr := range xKeyMatr {
			if yKeyMatr == 1 {
				if i < len(runeText) {
					newMatr[x][y] = string(runeText[i])
				} else {
					newMatr[x][y] = GetElement(Dictionary, rand.Intn(len(Dictionary)))
				}
				i++
			}
		}
	}
	matrix = VerticalReverse(matrix)
	for x, xKeyMatr := range matrix {
		for y, yKeyMatr := range xKeyMatr {
			if yKeyMatr == 1 {
				if i < len(runeText) {
					newMatr[x][y] = string(runeText[i])
				} else {
					newMatr[x][y] = GetElement(Dictionary, rand.Intn(len(Dictionary)))
				}
				i++
			}
		}
	}
	matrix = HorizontalReverse(matrix)
	for x, xKeyMatr := range matrix {
		for y, yKeyMatr := range xKeyMatr {
			if yKeyMatr == 1 {
				if i < len(runeText) {
					newMatr[x][y] = string(runeText[i])
				} else {
					newMatr[x][y] = GetElement(Dictionary, rand.Intn(len(Dictionary)))
				}
				i++
			}
		}
	}
	matrix = VerticalReverse(matrix)
	for x, xKeyMatr := range matrix {
		for y, yKeyMatr := range xKeyMatr {
			if yKeyMatr == 1 {
				if i < len(runeText) {
					newMatr[x][y] = string(runeText[i])
				} else {
					newMatr[x][y] = GetElement(Dictionary, rand.Intn(len(Dictionary)))
				}
				i++
			}
		}
	}
	for k, i := range newMatr {
		result += strings.Join(i, " ")
		if k != len(newMatr)-1 {
			result += "\r\n"
		}
	}
	return result
}

func CardanGrilleDecrypt(text string, keys []*walk.TextEdit) string {
	key := keys[0].Text()
	matrix := GetMatrix(key)
	textMatrix := GetTextMatrix(text)
	var result string
	rand.Seed(time.Now().UnixNano())
	for x, xKeyMatr := range matrix {
		for y, yKeyMatr := range xKeyMatr {
			if yKeyMatr == 1 {
				result += textMatrix[x][y]
			}
		}
	}
	matrix = VerticalReverse(matrix)
	for x, xKeyMatr := range matrix {
		for y, yKeyMatr := range xKeyMatr {
			if yKeyMatr == 1 {
				result += textMatrix[x][y]
			}
		}
	}
	matrix = HorizontalReverse(matrix)
	for x, xKeyMatr := range matrix {
		for y, yKeyMatr := range xKeyMatr {
			if yKeyMatr == 1 {
				result += textMatrix[x][y]
			}
		}
	}
	matrix = VerticalReverse(matrix)
	for x, xKeyMatr := range matrix {
		for y, yKeyMatr := range xKeyMatr {
			if yKeyMatr == 1 {
				result += textMatrix[x][y]
			}
		}
	}
	return result
}

func DiffieHellmanEncryptStart(text string, keys []*walk.TextEdit) string {
	rand.Seed(time.Now().UnixNano())
	n:= rand.Intn(16)+3
	a:= rand.Intn(n-2)+2
	sV :=rand.Intn(n-2)+2
	sO := rand.Intn(n-2)+2
	keys[0].SetText(strconv.Itoa(n))
	keys[1].SetText(strconv.Itoa(a))
	keys[2].SetText(strconv.Itoa(sV))
	keys[3].SetText(strconv.Itoa(sO))
	return DiffieHellmanEncrypt(text, keys)
}

func DiffieHellmanEncrypt(text string, keys []*walk.TextEdit) string {
	n, _ := strconv.Atoi(keys[0].Text())
	a, _ := strconv.Atoi(keys[1].Text())
	sV,_ := strconv.Atoi(keys[2].Text())
	sO,_ := strconv.Atoi(keys[3].Text())
	sVOpen := int(math.Pow(float64(a), float64(sV)))%n
	sOOpen := int(math.Pow(float64(a), float64(sO)))%n
	keySO := int(math.Pow(float64(sVOpen), float64(sO)))%n
	keySV := int(math.Pow(float64(sOOpen), float64(sV)))%n
	if keySO==keySV{
		return strconv.Itoa(keySO)
	}
	return ""
}
