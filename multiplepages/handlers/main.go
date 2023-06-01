package handlers

import (
	. "main/const"
	. "main/tools"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func MatrixKeyHandler(cipher CipherPackage)  string {
	regex := cipher.Cipher.KeyErrors[0].Regex
	keys := cipher.Keys
	key := keys[0].Text()
	re := regexp.MustCompile(regex)
	if utf8.RuneCountInString(key)==0{
		return ErrorKeyWrite
	}
	if !re.MatchString(key){
		return ErrorKeyInValidSym
	}
	matrix := GetMatrix(key)
	firstNumLen := len(matrix[0])
	for _, num := range matrix {
		if len(num)!=firstNumLen{
			return ErrorMatrixEmptyElem
		}
	}
	if len(matrix)!=firstNumLen{
		return ErrorMatrixNotSquare
	}
	if Determinant(matrix)==0{
		return ErrorMatrixDegenerate
	}

	return ""
}

func MatrixCipherKeyHandler(cipher CipherPackage)  string {
	regex := cipher.Cipher.KeyErrors[0].Regex
	keys := cipher.Keys
	key := keys[0].Text()
	re := regexp.MustCompile(regex)
	if utf8.RuneCountInString(key)==0{
		return ErrorKeyWriteCipher
	}
	if !re.MatchString(key){
		return ErrorKeyInValidSymCipher
	}
	matrix := GetMatrix(key)
	firstNumLen := len(matrix[0])
	for _, num := range matrix {
		if len(num)!=firstNumLen{
			return ErrorMatrixEmptyElem
		}
	}
	if len(matrix)!=firstNumLen{
		return ErrorMatrixNotSquare
	}
	if Determinant(matrix)==0{
		return ErrorMatrixDegenerate
	}

	inMatrix := InverseMatrix(matrix)
	matr := [][]float64{}
	i := 0
	text := cipher.Text
	for _, v := range strings.Split(text, " ") {
		t, _ := strconv.ParseFloat(v, 64)
		matr = append(matr, []float64{t})
		i++
		if i > len(inMatrix[0])-1 {
			i = 0
			for _, elem := range MultiplyMatrix(inMatrix, matr) {
				if int(elem[0])-1<0 {
					return ErrorTextMatrixNegativeSym
				}
				if int(elem[0])-1>= len(Dictionary) {
					return ErrorTextMatrixExtraSym
				}
			}
			matr = [][]float64{}
		}
	}
	return ""
}

func PlayfairKeyHandler(cipher CipherPackage)  string {
	regex := cipher.Cipher.KeyErrors[0].Regex
	keys := cipher.Keys
	key := keys[0].Text()
	re := regexp.MustCompile(regex)
	if utf8.RuneCountInString(key)==0{
		return ErrorKeyWrite
	}
	if !re.MatchString(key){
		return ErrorKeyInValidSym
	}
	if !HasUniqueChars(key, SecondDictionary){
		return ErrorKeyNotUniqueChars
	}
	return ""
}

func PlayfairTextHandler(cipher CipherPackage)  string {
	regex := cipher.Cipher.TextErrors.Regex
	text := cipher.Text
	_, text = CheckPairs(text)
	re := regexp.MustCompile(regex)
	if utf8.RuneCountInString(text)==0{
		return ErrorTextWrite
	}
	if !re.MatchString(text){
		return ErrorTextInValidSym
	}
	if utf8.RuneCountInString(text)%2!=0{
		return ErrorTextRepeatedTwo
	}
	return ""
}


func PlayfairCipherTextHandler(cipher CipherPackage)  string {
	regex := cipher.Cipher.TextErrors.Regex
	text := cipher.Text
	re := regexp.MustCompile(regex)
	if utf8.RuneCountInString(text)==0{
		return ErrorTextWriteCipher
	}
	if !re.MatchString(text){
		return ErrorTextInValidSymCipher
	}
	if utf8.RuneCountInString(text)%2!=0{
		return ErrorTextRepeatedTwoCipher
	}
	check, _ := CheckPairs(text)
	if check{
		return ErrorTextRepeatedPairsCipher
	}
	return ""
}

func PlayfairCipherKeyHandler(cipher CipherPackage)  string {
	regex := cipher.Cipher.KeyErrors[0].Regex
	keys := cipher.Keys
	key := keys[0].Text()
	re := regexp.MustCompile(regex)
	if utf8.RuneCountInString(key)==0{
		return ErrorKeyWriteCipher
	}
	if !re.MatchString(key){
		return ErrorKeyInValidSymCipher
	}
	if !HasUniqueChars(key,SecondDictionary){
		return ErrorKeyNotUniqueCharsCipher
	}
	return ""
}

func VerticalPermutationKeyHandler(cipher CipherPackage)  string {
	regex := cipher.Cipher.KeyErrors[0].Regex
	keys := cipher.Keys
	key := keys[0].Text()
	re := regexp.MustCompile(regex)
	if utf8.RuneCountInString(key)==0{
		return ErrorKeyWrite
	}
	if !re.MatchString(key){
		return ErrorKeyInValidSym
	}
	return ""
}

func VerticalPermutationCipherKeyHandler(cipher CipherPackage)  string {
	regex := cipher.Cipher.KeyErrors[0].Regex
	keys := cipher.Keys
	key := keys[0].Text()
	re := regexp.MustCompile(regex)
	if utf8.RuneCountInString(key)==0{
		return ErrorKeyWriteCipher
	}
	if !re.MatchString(key){
		return ErrorKeyInValidSymCipher
	}
	return ""
}

func CardanGrilleCipherTextHandler(cipher CipherPackage)  string {
	regex := cipher.Cipher.TextErrors.Regex
	text := cipher.Text
	re := regexp.MustCompile(regex)
	if utf8.RuneCountInString(text)==0{
		return ErrorTextWriteCipher
	}
	if !re.MatchString(text){
		return ErrorTextInValidSymCipher
	}
	matrix := GetMatrix(text)
	firstNumLen := len(matrix[0])
	for _, num := range matrix {
		if len(num)!=firstNumLen{
			return ErrorTextCardanEmptyElementCipher
		}
	}
	if len(matrix)%2!=0||len(matrix[0])%2!=0{
		return ErrorTextCardanRepeatedTwoCipher
	}
	return ""
}

func CardanGrilleKeyHandler(cipher CipherPackage)  string {
	regex := cipher.Cipher.KeyErrors[0].Regex
	keys := cipher.Keys
	key := keys[0].Text()
	re := regexp.MustCompile(regex)
	if utf8.RuneCountInString(key)==0{
		return ErrorKeyWrite
	}
	if !re.MatchString(key){
		return ErrorKeyInValidSym
	}
	matrix := GetMatrix(key)
	firstNumLen := len(matrix[0])
	for _, num := range matrix {
		if len(num)!=firstNumLen{
			return ErrorKeyCardanEmptyElement
		}
	}
	if len(matrix)%2!=0||len(matrix[0])%2!=0{
		return ErrorKeyCardanRepeatedTwo
	}
	matrixLen := len(matrix)*len(matrix[0])
	if matrixLen%4!=0 {
		return ErrorKeyCardanRepeatedFour
	}
	kol1 := 0
	for _, sym := range matrix{
		for _, sym2 := range sym {
			if sym2==1{
				kol1+=1
			}
		}
	}
	if kol1*4!=matrixLen{
		return ErrorKeyCardanNumHoles
	}
	if HasSymmetricOnes(matrix, len(matrix), len(matrix[0])){
		return ErrorKeySymmetric
	}
	runeText := []rune(cipher.Text)
	if len(runeText)>matrixLen{
		return ErrorKeyCardanSmall
	}
	return ""
}

func CardanGrilleCipherKeyHandler(cipher CipherPackage)  string {
	regex := cipher.Cipher.KeyErrors[0].Regex
	keys := cipher.Keys
	key := keys[0].Text()
	re := regexp.MustCompile(regex)
	if utf8.RuneCountInString(key)==0{
		return ErrorKeyWriteCipher
	}
	if !re.MatchString(key){
		return ErrorKeyInValidSymCipher
	}
	matrix := GetMatrix(key)
	firstNumLen := len(matrix[0])
	for _, num := range matrix {
		if len(num)!=firstNumLen{
			return ErrorKeyCardanEmptyElementCipher
		}
	}
	if len(matrix)%2!=0||len(matrix[0])%2!=0{
		return ErrorKeyCardanRepeatedTwoSipher
	}
	matrixLen := len(matrix)*len(matrix[0])
	if matrixLen%4!=0 {
		return ErrorKeyCardanRepeatedFourCipher
	}
	kol1 := 0
	for _, sym := range matrix{
		for _, sym2 := range sym {
			if sym2==1{
				kol1+=1
			}
		}
	}
	if kol1*4!=matrixLen{
		return ErrorKeyCardanNumHolesCipher
	}
	if HasSymmetricOnes(matrix, len(matrix), len(matrix[0])){
		return ErrorKeySymmetricSipher
	}
	text := GetTextMatrix(cipher.Text)
	if len(text)*len(text[0])>matrixLen{
		return ErrorKeyCardanSmallSipher
	}
	return ""
}

func CardanGrilleStartKeyHandler(cipher CipherPackage)  string {
	return ""
}

func CardanGrilleCipherStartKeyHandler(cipher CipherPackage)  string {
	return ""
}

func DiffieHellmanKeyHandler(cipher CipherPackage)  string {
	regex := cipher.Cipher.KeyErrors[0].Regex
	keys := cipher.Keys
	re := regexp.MustCompile(regex)
	for i, keyEdit := range keys{
		if keyEdit.Visible() {
			key := keyEdit.Text()
			if utf8.RuneCountInString(key) == 0 {
				return ErrorParameterWrite + " " + strconv.Itoa(i+1)
			}
			if !re.MatchString(key) {
				return strconv.Itoa(i+1) + " " + ErrorParameterInValidSym
			}
			_, err := strconv.Atoi(key)
			if err != nil {
				return strconv.Itoa(i+1) + " " + ErrorKeyBigNum
			}
		}
	}
	n, _ := strconv.Atoi(keys[0].Text())
	a, _ := strconv.Atoi(keys[1].Text())
	sV,_ := strconv.Atoi(keys[2].Text())
	sO,_ := strconv.Atoi(keys[3].Text())

	if n<3 {
		return ErrorParameterDiffieHellmanSmallN
	}
	if a<2 {
		return ErrorParameterSmallA
	}
	if a>=n {
		return ErrorParameterDiffieHellmanBigA
	}
	if sV<2 {
		return ErrorKeyDiffieHellmanSmallSV
	}
	if sV>=n {
		return ErrorKeyDiffieHellmanBigSV
	}
	if sO<2 {
		return ErrorKeyDiffieHellmanSmallSO
	}
	if sO>=n {
		return ErrorKeyDiffieHellmanBigSO
	}

	sVOpen := int(math.Pow(float64(a), float64(sV)))%n
	if sVOpen<0{
		return ErrorParameterASVOverBig
	}
	sOOpen := int(math.Pow(float64(a), float64(sO)))%n
	if sOOpen<0{
		return ErrorParameterASOOverBig
	}
	key := int(math.Pow(float64(sVOpen), float64(sO)))
	if key<0{
		return ErrorParameterKeyBig
	}
	return ""
}

func DiffieHellmanTextHandler(cipher CipherPackage)  string {
	return ""
}

func DiffieHellmanStartKeyHandler(cipher CipherPackage)  string {
	return ""
}

func GOSTR341094KeyHandler(cipher CipherPackage)  string {
	regex := cipher.Cipher.KeyErrors[0].Regex
	keys := cipher.Keys
	re := regexp.MustCompile(regex)
	for i, keyEdit := range keys{
		if keyEdit.Visible() {
			key := keyEdit.Text()
			if utf8.RuneCountInString(key) == 0 {
				return ErrorParameterWrite + " " + strconv.Itoa(i+1)
			}
			if !re.MatchString(key) {
				return strconv.Itoa(i+1) + " " + ErrorParameterInValidSym
			}
			_, err := strconv.Atoi(key)
			if err != nil {
				return strconv.Itoa(i+1) + " " + ErrorKeyBigNum
			}
		}
	}
	p, _ := strconv.Atoi(keys[0].Text())
	q, _ := strconv.Atoi(keys[1].Text())
	a,_ := strconv.Atoi(keys[2].Text())
	x,_ := strconv.Atoi(keys[3].Text())
	k,_ := strconv.Atoi(keys[4].Text())

	if p<4 {
		return ErrorParameterPSmall
	}
	if !IsPrime(p){
		return ErrorParameterPNotPrime
	}
	if !IsPrimeFactor(p-1, q){
		return ErrorParameterQNotPrimeFactor
	}
	if a < 2 {
		return ErrorParameterSmallA
	}
	if a >= p-1 {
		return ErrorParameterABiggerP1
	}
	check := false
	for d:=2;d<p-1;d++{
		test := int(math.Pow(float64(d), float64((p-1)/q)))%p
		if test==a{
			check = true
			break
		}
	}
	if !check{
		return ErrorParameterAQMODPOverBig
	}

	if x>=q{
		return ErrorParameterXBiggerQ
	}
	y := int(math.Pow(float64(a), float64(x)))%p
	if y<0{
		return ErrorParameterAXOverBig
	}
	if k>=q{
		return ErrorParameterKBiggerQ
	}
	return ""
}


func GOSTR341094StartKeyHandler(cipher CipherPackage)  string {
	return ""
}

func RsaSignatureKeyHandler(cipher CipherPackage)  string {
	regex := cipher.Cipher.KeyErrors[0].Regex
	keys := cipher.Keys
	re := regexp.MustCompile(regex)
	for i, keyEdit := range keys{
		if keyEdit.Visible() {
			key := keyEdit.Text()
			if utf8.RuneCountInString(key) == 0 {
				return ErrorParameterWrite + " " + strconv.Itoa(i+1)
			}
			if !re.MatchString(key) {
				return strconv.Itoa(i+1) + " " + ErrorParameterInValidSym
			}
			_, err := strconv.Atoi(key)
			if err != nil {
				return strconv.Itoa(i+1) + " " + ErrorKeyBigNum
			}
		}
	}
	p, _ := strconv.Atoi(keys[0].Text())
	q, _ := strconv.Atoi(keys[1].Text())

	if !IsPrime(p){
		return ErrorParameterPNotPrime
	}
	if !IsPrime(q){
		return ErrorParameterQNotPrime
	}
	return ""
}

func RsaStartKeyHandler(cipher CipherPackage)  string {
	return ""
}

func RsaKeyHandler(cipher CipherPackage)  string {
	regex := cipher.Cipher.KeyErrors[0].Regex
	keys := cipher.Keys
	re := regexp.MustCompile(regex)
	for i, keyEdit := range keys{
		if keyEdit.Visible() {
			key := keyEdit.Text()
			if utf8.RuneCountInString(key) == 0 {
				return ErrorParameterWrite + " " + strconv.Itoa(i+1)
			}
			if !re.MatchString(key) {
				return strconv.Itoa(i+1) + " " + ErrorParameterInValidSym
			}
			_, err := strconv.Atoi(key)
			if err != nil {
				return strconv.Itoa(i+1) + " " + ErrorKeyBigNum
			}
		}
	}
	p, _ := strconv.Atoi(keys[0].Text())
	q, _ := strconv.Atoi(keys[1].Text())

	if !IsPrime(p){
		return ErrorParameterPNotPrime
	}
	if !IsPrime(q){
		return ErrorParameterQNotPrime
	}
	return ""
}

func RsaSignatureStartKeyHandler(cipher CipherPackage)  string {
	return ""
}

func ShennonKeyHandler(cipher CipherPackage)  string {
	return ""
}