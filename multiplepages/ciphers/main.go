package ciphers

import (
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

func CaesarEncrypt(text string, keys []*walk.TextEdit) string { //шифрование шифра Цезаря с любым ключом
	key := GetCaesarKey(keys) //получение ключа, введённого пользователем
	result := TransformCaesarText(key, text) //получение результирующей строки
	return result //возвращение результирующей строки
}

func CaesarDecrypt(text string, keys []*walk.TextEdit) string { //расшифрование шифра Цезаря с любым ключом
	key := GetCaesarKey(keys) //получение ключа, введённого пользователем
	key *= -1 //умножение ключа на -1, чтобы для расшифрования двигать строку в обратную сторону
	result := TransformCaesarText(key, text) //получение результирующей строки
	return result //возвращение результирующей строки
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
		}
		result += GetElement(mas[0], num)
		i++
		if i > utf8.RuneCountInString(key)-1 {
			i = 0
		}
	}
	return result
}

func VigenereCipherKeyEncrypt(text string, keys []*walk.TextEdit) string { //шифрование шифра Виженера с ключом-шифртекстом
	key := keys[0].Text() //считывание ключа
	var result string //объявляение результирующей строки
	mas := GetTrithemiusTable() //получение таблицы Тритемия
	for i, v := range []rune(text) { //проход по всем символам текста для шифрования
		num := IndexOf(string(v), mas[0]) //получение индекса текущей буквы в первой строки таблицы Тритемия
		for j := 0; ; j++ { //проход по всем строкам таблицы Тритемия
			if GetElement(mas[j], 0) == string([]rune(key)[i]) { //проверка на то, что строка начинается с текущего символа ключа
				newSym := GetElement(mas[j], num) //получение нового символа из текущей строки
				result += newSym //добавление нового символа в результирующую строку
				key += newSym //добавление нового символа в ключ, чтобы дальше шифрование шло уже по полученному символу
				break //выход из цикла к следующему символу текста для шифрования

			}
		}
	}
	return result //возвращение результирующей строки
}

func VigenereCipherKeyDecrypt(text string, keys []*walk.TextEdit) string {
	key := keys[0].Text() //считывание ключа
	var result string //объявляение результирующей строки
	mas := GetTrithemiusTable() //получение таблицы Тритемия
	for i, v := range []rune(text) { //проход по всем символам текста для шифрования
		for j := 0; ; j++ { //проход по всем строкам таблицы Тритемия
			if GetElement(mas[j], 0) == string([]rune(key)[i]) { //проверка на то, что строка начинается с текущего символа ключа
				num := IndexOf(string(v), mas[j]) //получение индекса текущей буквы в текущей строке таблицы Тритемия
				result += GetElement(mas[0], num) //добавление символа, полученного из первой строки по найденному индексу, в результирующую строку
				key += string(v) //добавление нового символа в ключ, чтобы дальше расшифрование шло уже по полученному символу
				break //выход из цикла к следующему символу текста для расшифрования
			}
		}
	}
	return result //возвращение результирующей строки
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

func MatrixEncrypt(text string, keys []*walk.TextEdit) string { //функция шифрования Матричным шифром
	key := keys[0].Text() //получение ключа
	var resultAr []string //объявление результирующего массива
	matrix := GetMatrix(key) //получение матрицы из ключа
	matrixLen := len(matrix) //объявление переменной и запись в неё текущей длины матрицы
	runeTextLen := utf8.RuneCountInString(text) //объявление переменной и запись в неё текущей длины текста
	rand.Seed(time.Now().UnixNano()) //инициализирование рандомайзера текущим временем
	if runeTextLen%matrixLen != 0 { //если длина текст на кратна длине матрицы
		for i := 0; i < matrixLen-runeTextLen%matrixLen; i++ { //проход столько раз, сколько еще нужно символов для кратности
			text += GetElement(Dictionary, rand.Intn(len(Dictionary))) //добавление в текст рандомного символа из словаря
		}
	}
	runeText := []rune(text) //объявление массива и запись в него всех символов текста
	var matr [][]int //объявление временного массива для хранения нужного количества значений для матричного перемножения
	for i:=0;i<len(runeText);i+=matrixLen{ //проход по всему тексту с шагом равным длине матрицы
		for _, v := range runeText[i:i+matrixLen]{ //проход для считывания новых символов количеством равных длине матрицы
													//и записью их индексов в временный массив
			t := IndexOf(string(v), Dictionary) + 1 //объявление переменной и запись в неё индекса текущего символа
			matr = append(matr, []int{t}) //добавление вышеуказанного индекса в временный массив
		}
		for _, elem := range MultiplyMatrix(matrix, matr) { //проход по значениям, полученным после
															// перемножения матриц - ключа и временной матрицы
			resultAr = append(resultAr, strconv.Itoa(elem[0])) //добавление в результирующий массив текстового представления
																//одного из полученных выше значений
		}
		matr = [][]int{} //обнуление временного массива
	}
	return strings.Join(resultAr, " ") //возвращение результирующей строки со значениями результируюшего массива, соединенными " "
}

func MatrixDecrypt(text string, keys []*walk.TextEdit) string { //функция расшифрования Матричным шифром
	key := keys[0].Text() //получение ключа
	var result string //объявление результирующей строки
	matrix := GetMatrix(key) //получение матрицы из ключа
	inMatrix := InverseMatrix(matrix) //получение обратной матрицы относительно матрицы из ключа
	var matr [][]int ////объявление временного массива для хранения нужного количества значений для матричного перемножения
	i := 0
	for _, v := range strings.Split(text, " ") {
		t, _ := strconv.Atoi(v)
		matr = append(matr, []int{t})
		i++
		if i > len(inMatrix[0])-1 {
			i = 0
			for _, elem := range MultiplyMatrix(inMatrix, matr) {
				result += GetElement(Dictionary, elem[0]-1)
			}
			matr = [][]int{}
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

func DiffieHellmanEncryptStart(text string, keys []*walk.TextEdit) string { //функция генерации значений для обмена ключами по Диффи-Хеллману
	rand.Seed(time.Now().UnixNano()) //устанавливаем рандомайзер по текущему времени
	n := rand.Intn(298) + 3 //подбираем рандомно число n, от 3 до 300. от 3, чтобы удовлетворяло равенству (1< a < n)
	a := rand.Intn(n-2) + 2 //подбираем рандомно число a от 2 до n не включительно, чтобы удовлетворяло (1< a < n)
	sV := rand.Intn(n-2) + 2 //подбираем рандомно секретный ключ одному пользователю от 2 до n не включительно, чтобы удовлетворяло [2, n-1]
	sO := rand.Intn(n-2) + 2 //подбираем рандомно секретный ключ другому пользователю от 2 до n не включительно, чтобы удовлетворяло [2, n-1]
	keys[0].SetText(strconv.Itoa(n)) //выводим на экран значение n
	keys[1].SetText(strconv.Itoa(a)) //выводим на экран значение a
	keys[2].SetText(strconv.Itoa(sV)) //выводим на экран значение секретного ключа первого пользователя
	keys[3].SetText(strconv.Itoa(sO)) //выводим на экран значение секретного ключа второго пользователя
	return DiffieHellmanEncrypt(text, keys) //вызываем непосредственно функцию обмена ключа по Диффи-Хеллману
}

func DiffieHellmanEncrypt(text string, keys []*walk.TextEdit) string { //функция обмена ключами по Диффи-Хеллману
	n, _ := new(big.Int).SetString(keys[0].Text(), 10) //считываем введенное n, удовлетворяющее равенству (1< a < n)
	a, _ := new(big.Int).SetString(keys[1].Text(), 10) //считываем введеное a от 2 до n не включительно, чтобы удовлетворяло (1< a < n)
	sV, _ := new(big.Int).SetString(keys[2].Text(), 10) //считываем секретный ключ одного пользователя от 2 до n не включительно, чтобы удовлетворял [2, n-1]
	sO, _ := new(big.Int).SetString(keys[3].Text(), 10) //считываем секретный ключ другого пользователя от 2 до n не включительно, чтобы удовлетворял [2, n-1]
	sVOpen := new(big.Int).Exp(a, sV, nil)//возводим a в степень секретного ключа первого пользователя
	sVOpen =  sVOpen.Mod(sVOpen,n) //делаем mod возведенного выше числа на n, получая открытый ключ первого пользователя
	sOOpen := new(big.Int).Exp(a, sO, nil) //возводим a в степень секретного ключа второго пользователя
	sOOpen =  sOOpen.Mod(sOOpen,n) //делаем mod возведенного выше числа на n, получая открытый ключ второго пользователя
	keySO := new(big.Int).Exp(sVOpen, sO, nil) //возводим открытый ключ первого пользователя в степень секретного ключа второго пользователя
	keySO =  keySO.Mod(keySO,n) //делаем mod возведенного выше числа на n, получая общий секретный ключ
	keySV := new(big.Int).Exp(sOOpen, sV, nil) //возводим открытый ключ второго пользователя в степень секретного ключа первого пользователя
	keySV =  keySV.Mod(keySV,n) //делаем mod возведенного выше числа на n, получая общий секретный ключ
	error := "Секретные ключи не могут быть равны 0 или 1, введите другие открытые ключи" //ошибка, которую будем выводить
	if keySO.Int64()*1<2||keySV.Int64()*1<2{ //если секретные итоговые ключи равны 0 или 1, то
		return error+Splitter+error //возвращаем ошибку
	}
	return keySO.String() + Splitter + keySV.String() //возвращаем получившееся общие секретные ключи
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
	fmt.Println("hash", hash)
	if hash.Mod(hash, q).Cmp(big.NewInt(0)) == 0 { //если хэш mod q равно 0
		hash.Add(hash,big.NewInt(1)) //то добавляем к значению хэша 1
	}



	y := new(big.Int).Exp(a, x, nil)
	y = y.Mod(y,p)
	r := new(big.Int).Exp(a, k, nil)
	r = r.Mod(r,p)
	r.Mod(r, q)

	error := "Введите другое k, так как r и s не должны быть равны 0" //сообщение об ошибке
	if r.Int64()==0{ //если r равно 0
		return error+Splitter+error //то возвращаем ошибку
	}

	s1 := new(big.Int).Mul(x, r)
	s2 := new(big.Int).Mul(k, hash)
	s := new(big.Int).Add(s1, s2)
	s.Mod(s, q)

	if s.Int64()==0{ //если s равно 0
		return error+Splitter+error //то возвращаем ошибку
	}

	v := new(big.Int).ModInverse(hash, q)

	z1 := new(big.Int).Mul(s, v)
	z1.Mod(z1, q)
	z2 := new(big.Int).Sub(q, r)
	z2.Mul(z2, v)
	z2.Mod(z2, q)

	u1 := new(big.Int).Exp(a, z1, nil)
	u1 = u1.Mod(u1,p)
	u2 := new(big.Int).Exp(y, z2, nil)
	u2 = u2.Mod(u2,p)
	u := new(big.Int).Mul(u1, u2)
	u.Mod(u, p)
	u.Mod(u, q)
	str1 := "r = "+r.String()+" s = "+s.String()
	str2 := "r = "+r.String()+" u = "+u.String()
	return str1+Splitter+str2
}

func RsaSignatureEncryptStart(text string, keys []*walk.TextEdit) string {//функция, которая нужна для генерации параметров p и q, очень больших простых чисел, не равных между собой, для алгоритма цифровой подписи Rsa
	RsaSetPQ(keys) //Вызываем функцию генерации и установки параметров p и q, очень больших простых чисел, не равных между собой
	return RsaSignatureEncrypt(text, keys) //вызываем непосредственно функцию генерации цифровой подписи
}

func RsaSignatureEncrypt(text string, keys []*walk.TextEdit) string {
	rand.Seed(time.Now().UnixNano()) //инициализируем рандомайзер от текущего времени
	p, _ := new(big.Int).SetString(keys[0].Text(), 10) //считываем p - введённое очень большое просто десятичное число
	q, _ := new(big.Int).SetString(keys[1].Text(), 10) //считываем q - введённое очень большое простое десятичное число, отличное от p
	e, _ := new(big.Int).SetString(keys[2].Text(), 10) //считываем e - введёное число 1<e<f(функции Эйлера), взаимно простое с f(функцией Эйлера)
	n := new(big.Int).Mul(p, q) //получаем n, как результат перемножения простых чисел p и q
	num1 := big.NewInt(1) //записываем в переменную 1, чтобы потом добавлять или вычитать 1 из чисел без повторного создания переменной
	f := new(big.Int).Mul(p.Sub(p,num1),q.Sub(q,num1)) // находим функцию Эйлера от p и q, путем перемножения (p-1) на (q-1)
	d := new(big.Int) //инициализируем переменную d
	for i:=big.NewInt(1);;i.Add(num1,i){ //проходимся циклом по значениям i от 1 до бесконечности, прибавляя к i по единице, чтобы подобрать d
		//подбираем d=(1 mod f)/e, либо (e*d) mod f=1
		j := new(big.Int) //инициализируем переменную, в которой будем подбирать d
		j = j.Mul(i,e) //находим произведение текущего i и подобранного e
		j = j.Mod(j,f) //находим mod произведения выше на f(функцию Эйлера)
		if j.Cmp(num1)==0{ //если mod(остаток от деления) выше равен 1, то мы нашли d
			d = i //записываем в d текущую подобранную i
			break //выходим из цикла
		}
	}
	//Генерация подписи
	hash := GetHash(text, n) //получаем хэш нашего сообщения с модулем n
	if hash.Mod(hash, n).Cmp(big.NewInt(0)) == 0 { //если hash mod n == 0
		hash.Add(hash,big.NewInt(1)) //то добавляем к хэшу единицу, так как хэш должен быть от (1,N)
	}
	oU := new(big.Int).Exp(hash, d, nil) //вовзодим получившийся хэш в степень d
	oU = oU.Mod(oU,n) //берем mod получившегося выше возведения на n
	//Првоерка подписи
	m := new(big.Int).Exp(oU, e, nil) //возводим получившуюся подпись в степень e
	m = m.Mod(m,n) //делаем mod возведенного выше числа на n
	return hash.String() + Splitter + m.String() //возвращаем как результат - полученный hash текста и полученную при проверке подпись
}

func RsaEncryptStart(text string, keys []*walk.TextEdit) string { //функция, которая нужна для генерации параметров p и q, очень больших простых чисел, не равных между собой, для шифрования шифра Rsa
	RsaSetPQ(keys) //Вызываем функцию генерации и установки параметров p и q, очень больших простых чисел, не равных между собой
	return RsaEncrypt(text, keys) //вызываем непосредственно функцию шифрования Rsa
}

func RsaEncrypt(text string, keys []*walk.TextEdit) string { //функция шифрования шифра RSA из лабкрипта
	rand.Seed(time.Now().UnixNano()) //инициализируем рандомайзер от текущего времени
	p, _ := new(big.Int).SetString(keys[0].Text(), 10) //считываем p - введённое очень большое просто десятичное число
	q, _ := new(big.Int).SetString(keys[1].Text(), 10) //считываем q - введённое очень большое простое десятичное число, отличное от p
	e, _ := new(big.Int).SetString(keys[2].Text(), 10) //считываем e - введёное число 1<e<f(функции Эйлера), взаимно простое с f(функцией Эйлера)
	n := new(big.Int).Mul(p, q) //получаем n, как результат перемножения простых чисел p и q
	num1 := big.NewInt(1) //записываем в переменную 1, чтобы потом добавлять или вычитать 1 из чисел без повторного создания переменной
	f := new(big.Int).Mul(p.Sub(p,num1),q.Sub(q,num1)) // находим функцию Эйлера от p и q, путем перемножения (p-1) на (q-1)
	d := new(big.Int) //инициализируем переменную d
	for i:=big.NewInt(1);;i.Add(num1,i){ //проходимся циклом по значениям i от 1 до бесконечности, прибавляя к i по единице, чтобы подобрать d
		//подбираем d=(1 mod f)/e, либо (e*d) mod f=1
		j := new(big.Int) //инициализируем переменную, в которой будем подбирать d
		j = j.Mul(i,e) //находим произведение текущего i и подобранного e
		j = j.Mod(j,f) //находим mod произведения выше на f(функцию Эйлера)
		if j.Cmp(num1)==0{ //если mod(остаток от деления) выше равен 1, то мы нашли d
			d = i //записываем в d текущую подобранную i
			break //выходим из цикла
		}
	}
	var result string //инициализируем переменную, в которую будет записана результирующая строка
	var mas []string //инициализируем переменную, в которой будет хранится массив с данными, которые будут потом записаны в результ
	for _, v := range text { //проходимся посимвольно по тексу для шифрования
		h := IndexOf(string(v), Dictionary)+1 //получаем индекс текущей буквы из текста в нашем словаре
		oU := new(big.Int).Exp(big.NewInt(int64(h)), d, nil) //возводим этот индекс в степень d
		oU = oU.Mod(oU,n) //получаем mod от результата возведения на n
		mas = append(mas, oU.String()) //записываем получившееся значение в результирующий массив
	}
	result = strings.Join(mas, " ") //склеиваем результируюущую строку из результирующего массива
	return result //возвращаем результируюущую(зашифрованную строку)
}


func RsaDecryptStart(text string, keys []*walk.TextEdit) string { //функция, которая нужна для генерации параметров p и q, очень больших простых чисел, не равных между собой, для расшифрования шифра Rsa
	if len(keys[0].Text())*len(keys[1].Text())*len(keys[2].Text())==0{ //если p или q или e не указано, что значит, что ключи не были введены
		RsaSetPQ(keys) //вызываем функцию генерации и установки параметров p и q, очень больших простых чисел, не равных между собой
	}
	return RsaDecrypt(text, keys) //вызываем непосредственно функцию для расшифрования Rsa
}

func RsaDecrypt(text string, keys []*walk.TextEdit) string { ////функция расшифрования шифра RSA из лабкрипта
	mas := strings.Split(text, " ") //строку переводим в массив, делая её по пробелу, таким образом записываем зашифрованные значения в массив
	p, _ := new(big.Int).SetString(keys[0].Text(), 10) //считываем p - введённое очень большое просто десятичное число
	q, _ := new(big.Int).SetString(keys[1].Text(), 10) //считываем q - введённое очень большое простое десятичное число, отличное от p
	n := new(big.Int).Mul(p, q) //получаем n, как результат перемножения простых чисел p и q
	e , _:= new(big.Int).SetString(keys[2].Text(), 10) //считываем число 1<e<f(функции Эйлера), взаимного просто с f(функцией Эйлера)
	var result string //инициализируем перменную с результирующей строкой
	for _, v := range mas{ //проходим по массиву с зашифрованными значенияеми
		num,_ := new(big.Int).SetString(v,10) //считываем текущее значение
		m := new(big.Int).Exp(num, e, nil) //возводим это значение в степень числа e
		m = new(big.Int).Mod(m,n) //делаем mod результата возведения на n
		result+=GetElement(Dictionary, int(m.Int64())-1) //получившееся число считаем индексом символа из нашего алфавита
														//и записываем нужный символ в результирующую строку
	}
	return result //возвращаем результирующую строку
}

func ShennonEncryptStart(text string,keys []*walk.TextEdit ) string { //функция, которая подбирает значения для шифрования Блокнота Шеннона
	ShennonParametrsSet(keys)       //вызываем функцию подбора параметров для Блокнота Шеннона
	return ShennonCipher(text,keys) //вызываем непосредственно функцию шифрования Блокнота Шеннона
}

func ShennonDecryptStart(text string,keys []*walk.TextEdit ) string {//функция, которая подбирает значения для расшифрования Блокнота Шеннона
	if len(keys[0].Text())*len(keys[1].Text())*
		len(keys[2].Text())*len(keys[3].Text())*len(keys[4].Text())==0 { //если один из нужных параметров не указан, то
		ShennonParametrsSet(keys) //вызываем функцию подбора параметров для Блокнота Шеннона
	}
	return ShennonCipher(text, keys) ////вызываем непосредственно функцию расшифрования Блокнота Шеннона
}


func ShennonCipher(text string,keys []*walk.TextEdit ) string { //фукнция шифрования Блокнота Шеннона

	if len(keys[0].Text())==0{
		a,_ := strconv.Atoi(keys[1].Text())
		t0,_ := strconv.Atoi(keys[2].Text())
		c,_ := strconv.Atoi(keys[3].Text())
		m,_ := strconv.Atoi(keys[4].Text())
		mas := GenerateShennonGamma(a,t0,c,m) //генерируем гамму для блокнота Шеннона из полученных параметров
		var strMas []string //инициализиурем переменную, в которую будут записываться строковые представления
		//чисел из гаммы
		for _, elem := range mas{ //роходимся по массиву гаммы
			strMas = append(strMas,strconv.Itoa(elem)) //записываем в строковый массив строковое представление
			//чисел из массива гаммы
		}
		keys[0].SetText(strings.Join(strMas," ")) //записываем в поле параметра гаммы полученную гамму
	}
	m,_ := strconv.Atoi(keys[4].Text()) //считываем значение m, которое равно размеру алфавита
	var mas []int //инициализируем переменную массива с числовыми представлениями гаммы
	for _, elem := range strings.Split(keys[0].Text()," "){ //проходимя по строковым представлениям гаммы
		num,_ := strconv.Atoi(elem) //конвертируем строковые представления гаммы в числовые
		mas = append(mas, num) //добавляем числовые представления в массив
	}
	encrypted := "" //инициализируем переменную для хранения зашифрованного текста
	var j int //инициаилизируем счетчик, который помогает проходиться по всей гамме параллельно
			// с текстом для шифрования
	var t []int
	for _, v:= range text { //проходимся по тексту для шифрования

		ci := IndexOf(string(v),Dictionary)^(mas[j%m]) //индекс текущей буквы в нашем словаре возводим
		t = append(t, ci)												// в степень соответствующего числа из гаммы.
														// используем mod, чтобы гамма шла по кругу
		encrypted += GetElement(Dictionary, ci) //записываем в строку с шифрованным текстом символ из
												//нашего словаря, чей индекс равен полученному выше числу
		j++ //увеличиваем счетчик
	}
	fmt.Println(len(t),t)
	return encrypted //возвращаем зашифрованный текст
}

func A51Encrypt(text string,keys []*walk.TextEdit ) string {
	a51key := keys[0].Text()
	if strings.Join(RemoveDuplicates(strings.Split(a51key, "")), "") == "А" {
		return "Invalid key"
	}
	a51key = TransformKey(a51key,64)
	var keyBits string
	for _, keyChar := range a51key {
		keyBits += KeyToBits(string(keyChar))
	}
	var bitStockText string
	for _, textChar := range text {
		bitStockText += KeyToBits(string(textChar))
	}
	bitStockTextArray := ChunkString(bitStockText, 114)
	var result string

	for j, bitStockTextChunk := range bitStockTextArray {
		var registr1 = make([]int, 19)
		var registr2 = make([]int, 22)
		var registr3 = make([]int, 23)

		for i := 0; i < 64; i++ {
			registr1[0] = registr1[0]^ToInt(string(keyBits[i]))
			registr1 = append([]int{registr1[len(registr1)-1]}, registr1[:len(registr1)-1]...)

			registr2[0] = registr2[0]^ToInt(string(keyBits[i]))
			registr2 = append([]int{registr2[len(registr2)-1]}, registr2[:len(registr2)-1]...)

			registr3[0] = registr3[0]^ToInt(string(keyBits[i]))
			registr3 = append([]int{registr3[len(registr3)-1]}, registr3[:len(registr3)-1]...)
		}
		jB := new(big.Int).SetInt64(int64(j))
		jS := FillZerosBeforeNumber(jB.Text(2),22)
		for i := 0; i < 22; i++ {
			registr1[0] = registr1[0]^ToInt(string(jS[i]))
			registr1 = append([]int{registr1[len(registr1)-1]}, registr1[:len(registr1)-1]...)

			registr2[0] = registr2[0]^ToInt(string(jS[i]))
			registr2 = append([]int{registr2[len(registr2)-1]}, registr2[:len(registr2)-1]...)

			registr3[0] = registr3[0]^ToInt(string(jS[i]))
			registr3 = append([]int{registr3[len(registr3)-1]}, registr3[:len(registr3)-1]...)
		}
		for i := 0; i < 100; i++ {
			F := (registr1[8]&registr2[10])^(registr1[8]&registr3[10])^(registr2[10]&registr3[10])
			if registr1[8] == F {
				registr1 = append([]int{registr1[len(registr1)-1]}, registr1[:len(registr1)-1]...)
			}
			if registr2[10] == F {
				registr2 = append([]int{registr2[len(registr2)-1]}, registr2[:len(registr2)-1]...)
			}
			if registr3[10] == F {
				registr3 = append([]int{registr3[len(registr3)-1]}, registr3[:len(registr3)-1]...)
			}
		}
		var tempResult string
		var tempKey string
		for _, bit := range bitStockTextChunk {

			tempResult += strconv.Itoa((ToInt(string(bit))^(registr1[18]^registr2[21]^registr3[22]))%2)
			tempKey+=strconv.Itoa(registr1[18]^registr2[21]^registr3[22])
			registr1[0] = registr1[18]^registr1[17]^registr1[16]^registr1[13]
			registr1 = append([]int{registr1[len(registr1)-1]}, registr1[:len(registr1)-1]...)

			registr2[0] = registr2[21]^registr2[20]
			registr2 = append([]int{registr2[len(registr2)-1]}, registr2[:len(registr2)-1]...)

			registr3[0] = registr3[22]^registr3[21]^registr3[20]^registr3[7]
			registr3 = append([]int{registr3[len(registr3)-1]}, registr3[:len(registr3)-1]...)
		}
		result += tempResult
	}
	return result
}

func A51Decrypt(text string,keys []*walk.TextEdit ) string {
	a51key := keys[0].Text()
	if strings.Join(RemoveDuplicates(strings.Split(a51key, "")), "") == "А" {
		return "Invalid key"
	}
	a51key = TransformKey(a51key,64)
	var keyBits string
	for _, keyChar := range a51key {
		keyBits += KeyToBits(string(keyChar))
	}
	bitStockTextArray := ChunkString(text, 114)
	var result string

	for j, bitStockTextChunk := range bitStockTextArray {
		var registr1 = make([]int, 19)
		var registr2 = make([]int, 22)
		var registr3 = make([]int, 23)

		for i := 0; i < 64; i++ {
			registr1[0] = registr1[0]^ToInt(string(keyBits[i]))
			registr1 = append([]int{registr1[len(registr1)-1]}, registr1[:len(registr1)-1]...)

			registr2[0] = registr2[0]^ToInt(string(keyBits[i]))
			registr2 = append([]int{registr2[len(registr2)-1]}, registr2[:len(registr2)-1]...)

			registr3[0] = registr3[0]^ToInt(string(keyBits[i]))
			registr3 = append([]int{registr3[len(registr3)-1]}, registr3[:len(registr3)-1]...)
		}
		jB := new(big.Int).SetInt64(int64(j))
		jS := FillZerosBeforeNumber(jB.Text(2),22)
		for i := 0; i < 22; i++ {
			registr1[0] = registr1[0]^ToInt(string(jS[i]))
			registr1 = append([]int{registr1[len(registr1)-1]}, registr1[:len(registr1)-1]...)

			registr2[0] = registr2[0]^ToInt(string(jS[i]))
			registr2 = append([]int{registr2[len(registr2)-1]}, registr2[:len(registr2)-1]...)

			registr3[0] = registr3[0]^ToInt(string(jS[i]))
			registr3 = append([]int{registr3[len(registr3)-1]}, registr3[:len(registr3)-1]...)
		}
		for i := 0; i < 100; i++ {
			F := (registr1[8]&registr2[10])^(registr1[8]&registr3[10])^(registr2[10]&registr3[10])
			if registr1[8] == F {
				registr1 = append([]int{registr1[len(registr1)-1]}, registr1[:len(registr1)-1]...)
			}
			if registr2[10] == F {
				registr2 = append([]int{registr2[len(registr2)-1]}, registr2[:len(registr2)-1]...)
			}
			if registr3[10] == F {
				registr3 = append([]int{registr3[len(registr3)-1]}, registr3[:len(registr3)-1]...)
			}
		}
		var tempResult string
		for _, bit := range bitStockTextChunk { //проход по 114 битам текста
			tempResult += strconv.Itoa(ToInt(string(bit))^(registr1[18]^registr2[21]^registr3[22])) //xor между текущим битом текста и 18 битом первого регистра, 21 битом второго регистра,
																											//22 битом третьего регистра
			registr1[0] = registr1[18]^registr1[17]^registr1[16]^registr1[13] //записываем в нулевой бит первого регистра xor между 18 битом первого регистра, 17 битом первого регистра,
																				//16 битом первого регистра, 14 битом первого регистра
			registr1 = append([]int{registr1[len(registr1)-1]}, registr1[:len(registr1)-1]...) //сдвигаем первый регистр на 1 бит вправо

			registr2[0] = registr2[21]^registr2[20]
			registr2 = append([]int{registr2[len(registr2)-1]}, registr2[:len(registr2)-1]...)

			registr3[0] = registr3[22]^registr3[21]^registr3[20]^registr3[7]
			registr3 = append([]int{registr3[len(registr3)-1]}, registr3[:len(registr3)-1]...)
		}
		result += tempResult
	}
	var alphaResult string
	tempArray := ChunkString(result, 5)
	for _, bitChunk := range tempArray {
		alphaIndex,_ := strconv.ParseInt(bitChunk,2,0)
		alphaResult += GetElement(Dictionary, int(alphaIndex))
	}

	return alphaResult
}

func MagmaEncryptStart(text string, keys []*walk.TextEdit) string {
	cipherText := TextToHex(text)
	return MagmaEncrypt(cipherText,keys)
}

func MagmaEncrypt(text string,  keys []*walk.TextEdit) string {
	newKey := CutKey(keys[0].Text())
	textEncrypt := ""
	textArray := FormatTextGetParts(text,64)
	for _, textPart := range textArray {
		textEncrypt += ChainOfTransformations(textPart[:32], textPart[32:], newKey, "straight")
	}
	cipherHexStr := ""
	t := new(big.Int)
	for i:=0;i*4+4<=len(textEncrypt);i++{
		t.SetString(textEncrypt[i*4:i*4+4],2)
		cipherHexStr+=t.Text(16)
	}

	return cipherHexStr
}


func MagmaDecryptStart(text string, keys []*walk.TextEdit) string {
	decryptText := MagmaDecrypt(text, keys)
	return HexToText(decryptText)
}

func MagmaDecrypt(text string, keys []*walk.TextEdit) string {
	newKey := CutKey(keys[0].Text())
	textDecrypt := ""
	textArray := FormatTextGetParts(text,64)
	for _, textPart := range textArray {
		textDecrypt += ChainOfTransformations(textPart[:32], textPart[32:], newKey, "reverse")
	}
	cipherHexStr := ""
	t := new(big.Int)
	for i:=0;i*4+4<=len(textDecrypt);i++{
		t.SetString(textDecrypt[i*4:i*4+4],2)
		cipherHexStr+=t.Text(16)
	}
	return cipherHexStr
}

func KuzEncryptStart(text string, keys []*walk.TextEdit) string {
	cipherText := TextToHex(text)
	return KuzEncrypt(cipherText,keys)
}

func KuzDecryptStart(text string, keys []*walk.TextEdit) string {
	decryptText := KuzDecrypt(text,keys)
	return HexToText(decryptText)
}

func KuzEncrypt(text string, keys []*walk.TextEdit) string {
	textEncrypt := ""
	key := TransformKey(keys[0].Text(),64)
	k,_ := new(big.Int).SetString(key,16)
	keysKuz := KuznyechikKeySchedule(k)
	textArray := FormatTextGetParts(text,128)
	for _, textPart := range textArray {
		t,_ := new(big.Int).SetString(textPart,2)
		for round := 0; round < 9; round++ {
			res := new(big.Int).Xor(t,keysKuz[round])
			t = L(S(res))
		}
		textEncrypt += FillZerosBeforeNumber(new(big.Int).Xor(t, keysKuz[9]).Text(16),32)
	}

	return textEncrypt
}

func KuzDecrypt(text string, keys []*walk.TextEdit) string {
	textDecrypt := ""
	key := TransformKey(keys[0].Text(),64)
	k,_ := new(big.Int).SetString(key,16)
	keysKuz := KuznyechikKeySchedule(k)
	textArray := FormatTextGetParts(text,128)
	for _, textPart := range textArray {
		t,_ := new(big.Int).SetString(textPart,2)
		for round := 0; round < 9; round++ {
			t = SInv(LInv(new(big.Int).Xor(t, keysKuz[9-round])))
		}
		textDecrypt += FillZerosBeforeNumber(new(big.Int).Xor(t, keysKuz[0]).Text(16),32)
	}
	return textDecrypt
}

