package ciphers

import (
	"encoding/binary"
	"fmt"
	"github.com/lxn/walk"
	. "main/const"
	. "main/tools"
	"math"
	"math/big"
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
	if utf8.RuneCountInString(text)%len(matrix) != 0 {
		for i := 0; i < len(matrix)-utf8.RuneCountInString(text)%len(matrix); i++ {
			text += GetElement(Dictionary, rand.Intn(len(Dictionary)))
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
		text = strings.Join(masText, " ") + " " + MatrixEncrypt(enText, keys)
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
	return CardanGrilleDecrypt(text, keys)
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
	n := rand.Intn(300) + 3
	a := rand.Intn(n-2) + 2
	sV := rand.Intn(n-2) + 2
	sO := rand.Intn(n-2) + 2
	keys[0].SetText(strconv.Itoa(n))
	keys[1].SetText(strconv.Itoa(a))
	keys[2].SetText(strconv.Itoa(sV))
	keys[3].SetText(strconv.Itoa(sO))
	return DiffieHellmanEncrypt(text, keys)
}

func DiffieHellmanEncrypt(text string, keys []*walk.TextEdit) string {
	n := new(big.Int)
	n.SetString(keys[0].Text(), 10)

	a := new(big.Int)
	a.SetString(keys[1].Text(), 10)

	sV := new(big.Int)
	sV.SetString(keys[2].Text(), 10)

	sO := new(big.Int)
	sO.SetString(keys[3].Text(), 10)

	sVOpen := new(big.Int).Exp(a, sV, n)
	sOOpen := new(big.Int).Exp(a, sO, n)
	keySO := new(big.Int).Exp(sVOpen, sO, n)
	keySV := new(big.Int).Exp(sOOpen, sV, n)

	return keySO.String() + Splitter + keySV.String()
}

func GOSTR341094EncryptStart(text string, keys []*walk.TextEdit) string {
	rand.Seed(time.Now().UnixNano())
	var p, q, a int
	for {
		p = GeneratePrimeNumber()
		q = GenerateQ(p)
		for d := 2; d < p-1; d++ {
			test := int(math.Pow(float64(d), float64((p-1)/q))) % p
			if test > 1 {
				a = test
				break
			}
		}
		if a != 0 {
			break
		}
	}
	x := rand.Intn(q-1) + 1
	k := rand.Intn(q-1) + 1
	keys[0].SetText(strconv.Itoa(p))
	keys[1].SetText(strconv.Itoa(q))
	keys[2].SetText(strconv.Itoa(a))
	keys[3].SetText(strconv.Itoa(x))
	keys[4].SetText(strconv.Itoa(k))
	return GOSTR341094Encrypt(text, keys)
}

func GOSTR341094Encrypt(text string, keys []*walk.TextEdit) string {
	p, _ := new(big.Int).SetString(keys[0].Text(), 10)
	q, _ := new(big.Int).SetString(keys[1].Text(), 10)
	a, _ := new(big.Int).SetString(keys[2].Text(), 10)
	x, _ := new(big.Int).SetString(keys[3].Text(), 10)
	k, _ := new(big.Int).SetString(keys[4].Text(), 10)

	hash := GetHash(text, p)
	if hash.Mod(hash, q).Cmp(big.NewInt(0)) == 0 {
		hash.SetInt64(1)
	}

	y := new(big.Int).Exp(a, x, p)
	r := new(big.Int).Exp(a, k, p)
	r.Mod(r, q)

	s1 := new(big.Int).Mul(x, r)
	s2 := new(big.Int).Mul(k, hash)
	s := new(big.Int).Add(s1, s2)
	s.Mod(s, q)

	v := new(big.Int).ModInverse(hash, q)

	z1 := new(big.Int).Mul(s, v)
	z1.Mod(z1, q)
	z2 := new(big.Int).Sub(q, r)
	z2.Mul(z2, v)
	z2.Mod(z2, q)

	u1 := new(big.Int).Exp(a, z1, p)
	u2 := new(big.Int).Exp(y, z2, p)
	u := new(big.Int).Mul(u1, u2)
	u.Mod(u, p)
	u.Mod(u, q)

	return r.String() + Splitter + u.String()
}

func RsaSignatureEncryptStart(text string, keys []*walk.TextEdit) string {
	p := GeneratePrimeNumber()
	q := GeneratePrimeNumber()
	keys[0].SetText(strconv.Itoa(p))
	keys[1].SetText(strconv.Itoa(q))
	return RsaSignatureEncrypt(text, keys)
}

func RsaSignatureEncrypt(text string, keys []*walk.TextEdit) string {
	rand.Seed(time.Now().UnixNano())
	p, _ := new(big.Int).SetString(keys[0].Text(), 10)
	q, _ := new(big.Int).SetString(keys[1].Text(), 10)
	n := new(big.Int).Mul(p, q)
	f := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	var temp []int64
	for i := 2; i < int(f.Int64()); i++ {
		if Gcd(i, int(f.Int64())) == 1 {
			temp = append(temp, int64(i))
		}
	}
	e := temp[rand.Intn(len(temp))]
	d := big.NewInt(1)
	for {
		if new(big.Int).Mod(new(big.Int).Mul(d, big.NewInt(int64(e))), f).Cmp(big.NewInt(1)) == 0 {
			break
		}
		d.Add(d, big.NewInt(1))
	}
	hash := GetHash(text, n)
	oU := new(big.Int).Exp(hash, d, n).Int64()
	m := new(big.Int).Exp(big.NewInt(oU), big.NewInt(int64(e)), n).Int64()
	return strconv.Itoa(int(hash.Int64())) + Splitter + strconv.Itoa(int(m))
}

func RsaEncryptStart(text string, keys []*walk.TextEdit) string {

	p := GeneratePrimeNumber()
	q := GeneratePrimeNumber()
	keys[0].SetText(strconv.Itoa(p))
	keys[1].SetText(strconv.Itoa(q))
	return RsaEncrypt(text, keys)
}

func RsaEncrypt(text string, keys []*walk.TextEdit) string {
	rand.Seed(time.Now().UnixNano())
	p, _ := new(big.Int).SetString(keys[0].Text(), 10)
	q, _ := new(big.Int).SetString(keys[1].Text(), 10)
	n := new(big.Int).Mul(p, q)
	f := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	var temp []int64
	for i := 2; i < int(f.Int64()); i++ {
		if Gcd(i, int(f.Int64())) == 1 {
			temp = append(temp, int64(i))
		}
	}
	e := temp[rand.Intn(len(temp))]
	keys[3].SetText(strconv.Itoa(int(e)))
	d := big.NewInt(1)
	for {
		if new(big.Int).Mod(new(big.Int).Mul(d, big.NewInt(int64(e))), f).Cmp(big.NewInt(1)) == 0 {
			break
		}
		d.Add(d, big.NewInt(1))
	}
	var result string
	var mas []string
	for _, v := range text {
		h := IndexOf(string(v), Dictionary)
		oU := new(big.Int).Exp(big.NewInt(int64(h)), d, n).Int64()
		mas = append(mas, strconv.Itoa(int(oU)))
	}
	result = strings.Join(mas, " ")
	return result
}


func RsaDecryptStart(text string, keys []*walk.TextEdit) string {
	p := GeneratePrimeNumber()
	q := GeneratePrimeNumber()
	if len(keys[0].Text())==0 {
		keys[0].SetText(strconv.Itoa(p))
		keys[1].SetText(strconv.Itoa(q))
	}
	return RsaDecrypt(text, keys)
}

func RsaDecrypt(text string, keys []*walk.TextEdit) string {
	mas := strings.Split(text, " ")
	p, _ := new(big.Int).SetString(keys[0].Text(), 10)
	q, _ := new(big.Int).SetString(keys[1].Text(), 10)
	n := new(big.Int).Mul(p, q)
	e , _:= new(big.Int).SetString(keys[3].Text(), 10)
	var result string
	for _, v := range mas{
		num,_ := strconv.Atoi(v)
		m := new(big.Int).Exp(big.NewInt(int64(num)), e, n).Int64()
		result+=GetElement(Dictionary, int(m))
	}
	return result
}

func ShennonEncrypt(text string,keys []*walk.TextEdit ) string {
	rand.Seed(time.Now().UnixNano())
	gamma := GenerateGamma(utf8.RuneCountInString(text),rand.Intn(100)+1)
	var gammaMas []int
	for _,v:= range strings.Split(gamma," "){
		num, _ := strconv.Atoi(v)
		gammaMas = append(gammaMas, num)
	}
	fmt.Println(gamma)
	keys[0].SetText(gamma)
	encrypted := ""
	i:=0
	for _, v:= range text {
		ci := IndexOf(string(v),Dictionary)^gammaMas[i]
		encrypted += GetElement(Dictionary, ci)
		i++
	}
	fmt.Println(encrypted)
	return encrypted
}

func ShennonDecrypt(text string, keys []*walk.TextEdit) string {
	decrypted := ""
	gamma := keys[0].Text()
	if len(gamma)==0{
		gamma = GenerateGamma(utf8.RuneCountInString(text),rand.Int())
		keys[0].SetText(gamma)
	}
	var gammaMas []int
	for _,v:= range strings.Split(gamma," "){
		num, _ := strconv.Atoi(v)
		gammaMas = append(gammaMas, num)
	}
	var textMas []int
	for _, v := range text{
		textMas = append(textMas, IndexOf(string(v), Dictionary))
	}
	i := 0
	for _,v := range textMas{
		ci := v^gammaMas[i]
		decrypted += GetElement(Dictionary,ci)
		i++
	}
	return decrypted
}

func A51Encrypt(text string,keys []*walk.TextEdit ) string {
	rand.Seed(time.Now().UnixNano())
	gamma := GenerateA51Gamma(GenerateKey(),GenerateFrameNumber(), utf8.RuneCountInString(text))
	var gammaMas []int
	for _,v:= range strings.Split(gamma," "){
		num, _ := strconv.Atoi(v)
		gammaMas = append(gammaMas, num)
	}
	fmt.Println(gamma)
	keys[0].SetText(gamma)
	encrypted := ""
	i:=0
	for _, v:= range text {
		if i < len(gammaMas) {
			ci := IndexOf(string(v), Dictionary) ^ gammaMas[i]
			encrypted += GetElement(Dictionary, ci)
			i++
		}
	}
	return encrypted
}

func A51Decrypt(text string, keys []*walk.TextEdit) string {
	decrypted := ""
	gamma := keys[0].Text()
	if len(gamma)==0{
		gamma := GenerateA51Gamma(GenerateKey(),GenerateFrameNumber(), utf8.RuneCountInString(text))
		keys[0].SetText(gamma)
	}
	var gammaMas []int
	for _,v:= range strings.Split(gamma," "){
		num, _ := strconv.Atoi(v)
		gammaMas = append(gammaMas, num)
	}
	var textMas []int
	for _, v := range text{
		textMas = append(textMas, IndexOf(string(v), Dictionary))
	}
	i := 0
	for _,v := range textMas{
		ci := v^gammaMas[i]
		decrypted += GetElement(Dictionary,ci)
		i++
	}
	return decrypted
	return ""
}

var dict = []byte{
	0x32, 0x88, 0x4d, 0x41, 0x2b, 0x29, 0x4a, 0x24,
	0x14, 0x1e, 0xb, 0x7, 0x8e, 0x86, 0x3a, 0x17,
	0x8, 0x92, 0x1a, 0x3c, 0x6d, 0x46, 0x5d, 0x9d,
	0x1d, 0x3, 0x61, 0x69, 0xf, 0x7d, 0x27, 0x6f,
	0xa, 0x47, 0xf0, 0xf1, 0x2, 0xc, 0xff, 0xee,
	0x49, 0x44, 0xb8, 0x21, 0x19, 0x9, 0xbe, 0xd8,
	0x7c, 0x6, 0xe, 0x74, 0x31, 0x15, 0x1f, 0xc9,
	0x87, 0xaa, 0xba, 0x3e, 0x96, 0x5c, 0x83, 0x4b,
}

// Функция шифрования шифром магма
func MagmaEncrypt(text string,  keys []*walk.TextEdit) string {
	var ct []string
	keys[1].SetText(text)
	key := keys[0].Text()
	// Приводим текст и ключ к байтам
	pt := []byte(text)
	k := []byte(key)

	// Дополняем ключ нулями до 32 байт
	for len(k) < 32 {
		k = append(k, 0x00)
	}

	// Дополняем текст нулями до кратности 8
	for len(pt)%8 != 0 {
		pt = append(pt, 0x00)
	}

	// Шифруем блоки по 8 байт
	for i := 0; i < len(pt); i += 8 {
		// Преобразуем текст в uint32
		block := make([]uint32, 2)
		for j := 0; j < 8; j += 4 {
			block[j/4] = uint32(pt[i+j]) | uint32(pt[i+j+1])<<8 | uint32(pt[i+j+2])<<16 | uint32(pt[i+j+3])<<24
		}

		// XOR блока с первыми 32 байтами ключа
		for j := 0; j < 2; j++ {
			block[j] ^= binary.BigEndian.Uint32(k[j*4 : (j+1)*4])
		}

		// 31 раунд шифрования
		for j := 0; j < 31; j++ {
			// Переменная t получается как результат гаммирования блока и ключа
			var t uint32
			if (j+5)*4<=31{
				t = binary.BigEndian.Uint32(k[(j+4)*4:(j+5)*4]) ^ gamma(block[1], dict, j)
			}
			// Результаты t и блока смешиваются
			block[0], block[1] = block[1], block[0]^f(t)

		}


		// Добавляем зашифрованный блок к результату
		for j := 0; j < 8; j++ {
			ct = append(ct, GetElement(Dictionary,int(byte(block[j/4]>>(j%4*8)))%len(Dictionary)))
		}
	}

	// Возвращаем результат в виде строки
	return strings.Join(ct,"")
}

// Функция расшифрования шифром магма
func MagmaDecrypt(ciphertext string, keys []*walk.TextEdit) string {
	var pt []string
	key := keys[0].Text()
	// Приводим текст и ключ к байтам
	ct := []byte(ciphertext)
	k := []byte(key)

	// Дополняем ключ нулями до 32 байт
	for len(k) < 32 {
		k = append(k, 0x00)
	}

	// Расшифровываем блоки по 8 байт
	for i := 0; i < len(ct); i += 8 {
		// Преобразуем текст в uint32
		block := make([]uint32, 2)
		for j := 0; j < 8; j += 4 {
			block[j/4] = uint32(ct[i+j]) | uint32(ct[i+j+1])<<8 | uint32(ct[i+j+2])<<16 | uint32(ct[i+j+3])<<24
		}

		// 31 раунд шифрования (в обратном порядке)
		for j := 30; j >= 0; j-- {
			// Переменная t получается как результат гаммирования блока и ключа
			var t uint32
			if (j+5)*4<=31{
				t = binary.BigEndian.Uint32(k[(j+4)*4:(j+5)*4]) ^ gamma(block[0], dict, j)
			}
			// Результаты t и блока смешиваются
			block[1], block[0] = block[0], block[1]^f(t)
		}

		// XOR блока с первыми 32 байтами ключа
		for j := 0; j < 2; j++ {
			block[j] ^= binary.BigEndian.Uint32(k[j*4 : (j+1)*4])
		}

		// Добавляем расшифрованный блок к результату
		for j := 0; j < 8; j++ {
			pt = append(pt, GetElement(Dictionary,int(byte(block[j/4]>>(j%4*8)))%len(Dictionary)))
		}
	}
	ty := keys[1].Text()

	return ty
}

// Функция гаммирования
func gamma(block uint32, dict []byte, round int) uint32 {
	a := uint32(dict[4*round])
	b := uint32(dict[4*round+1])
	c := uint32(dict[4*round+2])
	d := uint32(dict[4*round+3])

	return g(a^block, b^block, c^block, d^block)
}

// Функция g
func g(a, b, c, d uint32) uint32 {
	a = g1(a, b, c, d)
	b = g2(a, b, c, d)
	c = g3(a, b, c, d)
	d = g4(a, b, c, d)
	a = g1(a, b, c, d)
	b = g2(a, b, c, d)
	c = g3(a, b, c, d)
	d = g4(a, b, c, d)

	return a | b | c | d
}

func g1(a, b, c, d uint32) uint32 {
	return f(a^b) ^ c ^ d
}

func g2(a, b, c, d uint32) uint32 {
	return f(c^d) ^ a ^ b
}

func g3(a, b, c, d uint32) uint32 {
	return f(a^c) ^ b ^ d
}

func g4(a, b, c, d uint32) uint32 {
	return f(b^d) ^ a ^ c
}

// Функция F
func f(data uint32) uint32 {
	x := (((data << 11) | (data >> 21)) ^ data)
	return (((x << 20) | (x >> 12)) ^ x)
}