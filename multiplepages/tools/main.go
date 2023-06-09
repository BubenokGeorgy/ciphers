package tools

import (
	"fmt"
	"github.com/lxn/walk"
	. "main/const"
	"main/data"
	"math"
	"math/big"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func Contains(a func(string, string) string, s []CiphersErrors) bool {
	for _, v := range s {
		if reflect.ValueOf(v.Cipher) == reflect.ValueOf(a) {
			return true
		}
	}

	return false
}

func InverseMatrix(matrix [][]float64) [][]float64 {
	n := len(matrix)

	// Создаем единичную матрицу, которую будем преобразовывать
	identity := make([][]float64, n)
	for i := 0; i < n; i++ {
		identity[i] = make([]float64, n)
		identity[i][i] = 1
	}

	// Преобразуем исходную матрицу и единичную матрицу одновременно
	for i := 0; i < n; i++ {
		// Находим главный элемент в столбце
		max := i
		for j := i + 1; j < n; j++ {
			if math.Abs(matrix[j][i]) > math.Abs(matrix[max][i]) {
				max = j
			}
		}

		// Если главный элемент в столбце равен нулю, то матрица необратима
		if matrix[max][i] == 0 {
			return nil
		}

		// Обмениваем текущую строку с строкой, содержащей главный элемент
		matrix[i], matrix[max] = matrix[max], matrix[i]
		identity[i], identity[max] = identity[max], identity[i]

		// Приводим текущую строку к единичному виду
		scale := 1 / matrix[i][i]
		for j := i; j < n; j++ {
			matrix[i][j] *= scale
		}
		for j := 0; j < n; j++ {
			identity[i][j] *= scale
		}

		// Обрабатываем остальные строки
		for j := 0; j < n; j++ {
			if j != i {
				scale := matrix[j][i]
				for k := i; k < n; k++ {
					matrix[j][k] -= scale * matrix[i][k]
				}
				for k := 0; k < n; k++ {
					identity[j][k] -= scale * identity[i][k]
				}
			}
		}
	}

	return identity
}

func MultiplyMatrix(matrixA, matrixB [][]float64) [][]float64 {

	result := make([][]float64, len(matrixA))
	for i := range result {
		result[i] = make([]float64, len(matrixB[0]))
	}

	for i := 0; i < len(matrixA); i++ {
		for j := 0; j < len(matrixB[0]); j++ {
			for k := 0; k < len(matrixB); k++ {
				result[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}

	return result
}

func GenerateCardanKey(text string) string {
	var result string
	runeText := []rune(text)
	rand.Seed(time.Now().UnixNano())
	w := rand.Intn(int(math.Ceil(float64(len(runeText))/2))+1)
	if w%2!=0||w==0{
		w+=1
	}
	h := int(math.Ceil(float64(len(runeText))/float64(w)))
	if h%2!=0||h==0{
		h+=1
	}
	matrix := make([][]float64, h)
	textMatrix := make([][]string, h)
	for i, _ := range matrix {
		textMatrix[i] =make([]string, w)
		matrix[i] = make([]float64, w)
	}
	freeSpaces := h/2*w/2
	kol := 0
	for i, k := range matrix{
		for j, _:= range k{
			if matrix[i][j]==1{
				continue
			}
			if kol>freeSpaces{
				break
			}
			matrix[i][j]=1
			if HasSymmetricOnes(matrix, h, w){
				matrix[i][j]=0
			} else {
				rand := rand.Intn(5)
				kol+=1
				matrix[i][j]=0
				if rand==0{
					rand=1
				}
				switch rand {
				case 1: matrix[i][j]=1
				case 2: matrix[h-1-i][j]=1
				case 3: matrix[i][w-1-j]=1
				case 4: matrix[h-1-i][w-1-j]=1
				}
			}
		}
	}
	for i, v := range matrix{
		for j,_  := range v {
			textMatrix[i][j] = strconv.Itoa(int(matrix[i][j]))
		}
	}
	for k, i := range textMatrix {
		result += strings.Join(i, " ")
		if k != len(matrix)-1 {
			result += "\r\n"
		}
	}
	return result
}

func GetTextMatrix(text string) [][]string {
	matrix := [][]string{}
	strs := strings.Split(text, "\r\n")
	for _, num := range strs {
		var matr []string
		for _, n := range strings.Split(num, " ") {
			matr = append(matr, n)
		}
		matrix = append(matrix, matr)
	}
	return matrix
}

func GetMatrix(text string) [][]float64 {
	matrix := [][]float64{}
	strs := strings.Split(text, "\r\n")
	for _, num := range strs {
		var matr []float64
		for _, n := range strings.Split(num, " ") {
			t, _ := strconv.ParseFloat(n, 64)
			matr = append(matr, t)
		}
		matrix = append(matrix, matr)
	}
	return matrix
}

func Determinant(m [][]float64) float64 {
	size := len(m)
	if size == 0 {
		return 0
	}
	if size == 1 {
		return m[0][0]
	}
	if size == 2 {
		return m[0][0]*m[1][1] - m[0][1]*m[1][0]
	}

	mCopy := make([][]float64, size)
	for i := range m {
		mCopy[i] = make([]float64, size)
		copy(mCopy[i], m[i])
	}

	for i := 0; i < size-1; i++ {
		if math.Abs(mCopy[i][i]) < 1e-10 {
			for j := i + 1; j < size; j++ {
				if math.Abs(mCopy[j][i]) > 1e-10 {
					mCopy[i], mCopy[j] = mCopy[j], mCopy[i]
					break
				}
			}
		}
		for j := i + 1; j < size; j++ {
			factor := mCopy[j][i] / mCopy[i][i]
			for k := i + 1; k < size; k++ {
				mCopy[j][k] -= factor * mCopy[i][k]
			}
			mCopy[j][i] = 0
		}
	}

	det := 1.0
	for i := 0; i < size; i++ {
		det *= mCopy[i][i]
	}

	return det
}

func VerticalReverse(matrix [][]float64) [][]float64 {
	// Создаем копию матрицы
	reversed := make([][]float64, len(matrix))
	for i := range reversed {
		reversed[i] = make([]float64, len(matrix[i]))
		copy(reversed[i], matrix[i])
	}

	// Получаем количество строк и столбцов матрицы
	rows := len(reversed)
	cols := len(reversed[0])

	// Итерируемся по строкам матрицы
	for i := 0; i < rows; i++ {
		// Итерируемся по столбцам матрицы
		for j := 0; j < cols/2; j++ {
			// Меняем местами элементы в каждой строке
			reversed[i][j], reversed[i][cols-j-1] = reversed[i][cols-j-1], reversed[i][j]
		}
	}

	return reversed
}

func HorizontalReverse(matrix [][]float64) [][]float64 {
	// Создаем копию матрицы
	reversed := make([][]float64, len(matrix))
	for i := range reversed {
		reversed[i] = make([]float64, len(matrix[i]))
		copy(reversed[i], matrix[i])
	}

	// Получаем количество строк и столбцов матрицы
	rows := len(reversed)
	cols := len(reversed[0])

	// Итерируемся по строкам матрицы
	for i := 0; i < rows/2; i++ {
		// Итерируемся по столбцам матрицы
		for j := 0; j < cols; j++ {
			// Меняем местами элементы в каждом столбце
			reversed[i][j], reversed[rows-i-1][j] = reversed[rows-i-1][j], reversed[i][j]
		}
	}

	return reversed
}

func IndexOfStr(element, key string) int {
	for i, data := range key {
			for _, let := range strings.Split(string(element), "/") {
				if string(data) == string(let) {
					return i
				}
			}
		}
	return -1
}

// Генерация символьной гаммы
func GenerateGamma(size int, seed int) string {
	gamma := ""
	m := len(Dictionary)
	a := 3 // Простое нечетное число
	c := seed // Исходное значение порождающего числа
	var mas []string
	for i := 0; i < size; i++ {
		c = (a*c + c) % m
		if c==0{
			c = 1
		}
		mas = append(mas,strconv.Itoa(c))
	}
	gamma = strings.Join(mas," ")
	return gamma
}

func GenerateA51Gamma(key uint64, frameNumber uint32, textLen int) string {
	r1 := uint64(0)
	r2 := uint64(0)
	r3 := uint64(0)
	output := make([]byte, 0)

	// Инициализация регистров, сдвиг на первых 64 тактах
	for i := 0; i < 64; i++ {
		bit := (key>>i)&1 ^ uint64(frameNumber>>i)&1
		r1 = ShiftRegister(r1, 19, 0, true, bit)
		r2 = ShiftRegister(r2, 22, 0, true, bit)
		r3 =ShiftRegister(r3, 23, 0, true, bit)
	}

	// Следующие 22 такта XOR с номером кадра
	for i := 0; i < 22; i++ {
		bit := (frameNumber>>i)&1 ^ 1
		r1 = ShiftRegister(r1, 19, 0, true, uint64(bit))
		r2 = ShiftRegister(r2, 22, 0, true, uint64(bit))
		r3 = ShiftRegister(r3, 23, 0, true, uint64(bit))
	}

	// Управление сдвигами регистров и генерация последовательности на следующих 100 тактах
	for i := 0; i < 100; i++ {
		f := (r1>>8)&1 & (r2>>10)&1 | (r1>>8)&1 & (r3>>10)&1 | (r2>>10)&1 & (r3>>10)&1
		if f&(1<<0) != 0 {
			r1 = ShiftRegister(r1, 19, 0, false, 0)
		}
		if f&(1<<1) != 0 {
			r2 = ShiftRegister(r2, 22, 0, false, 0)
		}
		if f&(1<<2) != 0 {
			r3 = ShiftRegister(r3, 23, 0, false, 0)
		}

		// Вычисление выходного бита
		outputBit := r1&1 ^ r2&1 ^ r3&1
		output = append(output, byte(outputBit))

		r1 = ShiftRegister(r1, 19, 0, true, outputBit)
		r2 = ShiftRegister(r2, 22, 0, true, outputBit)
		r3 =ShiftRegister(r3, 23, 0, true, outputBit)
	}
	var textOut []string
	rand.Seed(time.Now().UnixNano())
	for _, v:= range output{
		num := rand.Intn(2-int(v))+int(v)
		textOut = append(textOut,strconv.Itoa(num))
	}
	for i:= len(textOut);i<textLen;i++{
		num := rand.Intn(2)
		textOut = append(textOut,strconv.Itoa(num))
	}
	result := strings.Join(textOut," ")
	return result
}

func GenerateKey() uint64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Uint64()
}

// Генерирует случайный 22-битный номер кадра
func GenerateFrameNumber() uint32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Uint32() & ((1 << 22) - 1)
}

func ShiftRegister(reg uint64, len int, pos int, leftShift bool, newBit uint64) uint64 {
	if leftShift {
		reg <<= 1
		reg |= newBit
		if len >= 64 {
			reg &= (1 << len) - 1
		}
	} else {
		reg >>= 1
		reg |= newBit << (len - 1)
	}
	return reg
}

func HasSymmetricOnes(matrix [][]float64, h, w int) bool {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if matrix[i][j] == 1 && (matrix[h-1-i][j] == 1 || matrix[i][w-1-j] == 1 || matrix[h-1-i][w-1-j] == 1) {
				return true
			}
		}
	}
	return false
}

func IndexOfMas(element string, mas[][]string) (int, int) {
	for i, data := range mas {
		for k, elem := range data{
			for _, let := range strings.Split(elem, "/") {
				if element == string(let) {
					return i, k
				}
			}
		}
	}
	return -1, -1
}

func IndexOfInt(element int, data []int) int {
	for k, v := range data {
		if v==element{
			return k
		}
	}
	return -1
}

func IndexOf(element string, data []string) int {
	for k, v := range data {
		for _, let := range strings.Split(v, "/") {
			if element == string(let) {
				return k
			}
		}
	}
	return -1
}

func GetTrithemiusTable() [][]string {
	mas := make([][]string, 0, len(Dictionary))
	for i := 0; i < len(Dictionary); i++ {
		temp := make([]string, 0, len(Dictionary))
		for j := i; j < len(Dictionary); j++ {
			temp = append(temp, Dictionary[j])
		}
		for j := 0; j < i; j++ {
			temp = append(temp, Dictionary[j])
		}
		mas = append(mas, temp)
	}
	return mas
}

func CheckText(text, regex string, errors Errors) string {
	var error string
	re := regexp.MustCompile(regex)
	if !re.MatchString(text) {
		error = errors.ErrorsList[1]
		if utf8.RuneCountInString(text) == 0 {
			error = errors.ErrorsList[0]
		}
	}
	return error
}

func GetSeq(nums []int) []int {
	result := make([]int, len(nums))

	// Создаем слайс со значениями от 1 до 9
	indexes := make([]int, len(nums))
	for i := range indexes {
		indexes[i] = i + 1
	}

	// Сортируем слайс значений по возрастанию, при этом для одинаковых чисел
	// индексы увеличиваем слева направо
	for i := 0; i < len(indexes)-1; i++ {
		for j := i + 1; j < len(indexes); j++ {
			if nums[indexes[i]-1] > nums[indexes[j]-1] ||
				(nums[indexes[i]-1] == nums[indexes[j]-1] && indexes[i] > indexes[j]) {
				indexes[i],indexes[j] = indexes[j], indexes[i]
			}
		}
	}

	// Заполняем результат значениями из nums в порядке индексов
	for i, index := range indexes {
		result[index-1] = i
	}
	fmt.Println(result)
	return result
}

func GetMinMaxMatrix(mas [][]float64, matrix [][]float64){
	for masElemId, masElem := range mas{
		masElem := masElem[0]
		sum := float64(0)
		for _, matrixElem := range matrix[masElemId]{
			sum+=matrixElem*masElem
		}

	}
}

func GetVerticalpermutationTable(seq []int, newText []rune) [][]string {
	var mas [][]string
	var tempMas []string
	checked := false
	for i:=0;;i++{
		for j:=0;j<len(seq);j++ {
			if i*len(seq)+j<len(newText){
				tempMas =  append(tempMas, string(newText[i*len(seq)+j]))
			} else {
				tempMas =  append(tempMas, "")
				checked = true
			}
		}
		mas = append(mas, tempMas)
		tempMas = []string{}
		if checked {
			break
		}
	}
	return mas
}

func GetPlayfairTable(key string, tableLen int, dictionary []string)[][]string {
	if tableLen==0{
		for k:=2;k<=len(SecondDictionary);k++{
			if len(SecondDictionary)%k==0 {
				tableLen = k
				break
			}
		}
	}
	var mas[][]string
	var tempMas []string
	keyRunes := []rune(key)
	var i,j,b int
	for i=0;; i++{
		for j=0; j<tableLen; j++{
			if len(keyRunes)>0{
				newSym := dictionary[IndexOf(string(keyRunes[0]), dictionary)]
				tempMas = append(tempMas,newSym)
				keyRunes = keyRunes[1:]
			} else{
				for b=b;b<len(SecondDictionary);b++{
					if IndexOfStr(SecondDictionary[b], key)==-1 {
						tempMas = append(tempMas,SecondDictionary[b])
						break
					}
				}
				b++
			}
		}
		mas = append(mas, tempMas)
		tempMas = []string{}
		if b>=len(SecondDictionary) {
			break
		}
	}
	return mas
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsPrimeFactor(p, q int) bool {
	if q!=0 && p % q == 0 {
		if IsPrime(q){
			return true
		}
	}
	return false
}

func GetHash(text string, mod *big.Int) *big.Int {
	h := big.NewInt(0)
	for _, v := range text {
		num := IndexOf(string(v), Dictionary) + 1
		h.Add(h, big.NewInt(int64(num)))
		h.Mul(h, h)
		h.Mod(h, big.NewInt(mod.Int64()))
	}
	return h
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func GeneratePrimeNumber() int {
	for {
		// Generating a random integer between 3 and 100
		num := rand.Intn(298) + 4
		prime := true
		// Loop through numbers from 2 to the square root of the number, to check if it's prime
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				prime = false
				break
			}
		}
		if prime {
			return num
		}
	}
}

// Function to generate a q value that is a prime factor of (p-1)
func GenerateQ(p int) int {
	for {
		q := rand.Intn(p-2) + 2
		if (p-1)%q == 0 {
			if IsPrime(q){
				return q
			}
			// Q is a prime factor of (p-1)
		}
	}
}

func GetErrorText(cipher CipherPackage) string {
	var error string
	if cipher.Cipher.TextErrorsHandler == nil {
		if error = CheckText(cipher.Text, cipher.Cipher.TextErrors.Regex, cipher.Cipher.TextErrors); utf8.RuneCountInString(error) != 0 {
			return error
		}
	} else {
		if error = cipher.Cipher.TextErrorsHandler(cipher); utf8.RuneCountInString(error) != 0 {
			return error
		}
	}
	if cipher.Cipher.KeyErrorsHandler == nil {
		for i, key := range cipher.Keys{
			if key.Visible() {
				if error = CheckText(key.Text(), cipher.Cipher.KeyErrors[i].Regex, cipher.Cipher.KeyErrors[i]); utf8.RuneCountInString(error) != 0 {
					return error
				}
			}
		}

	} else {
		if error = cipher.Cipher.KeyErrorsHandler(cipher); utf8.RuneCountInString(error) != 0 {
			return error
		}
	}

	return error
}

func GetElement(mas []string, i int) string  {
	return GetFirstOfSplitChars(mas[i])
}


func CheckPairs(str string) (bool, string) {
	// Проходим по всем парам символов в строке
	runes := []rune(str)
	var result string
	check := false
	var i int
	// Проходим по всем парам символов в слайсе рун
	for i = 0; i < len(runes)-1; i+=2 {
		if GetElement(SecondDictionary,IndexOf(string(runes[i]), SecondDictionary)) == GetElement(SecondDictionary, IndexOf(string(runes[i+1]), SecondDictionary)) {
			runes = append(runes[:i+1], append([]rune{'ф'}, runes[i+1:]...)...)
			i-=2
			// Если пара символов равна друг другу, значит строка не подходит
			check = true
		} else {
			result += string(runes[i])
			result += string(runes[i+1])
		}

	}
	if i==len(runes)-1 {
		result += string(runes[i])
	}
	// Если цикл прошел успешно, значит все парамы символов различны
	return check, result
}

func HasUniqueChars(str string, dictioanary []string) bool {
	// Создаем мапу для хранения уже встреченных символов
	chars := make(map[string]bool)

	// Проходим по всем символам строки
	for _, char := range str {
		// Если символ уже был встречен, значит строка не уникальна
		sym := GetElement(dictioanary,IndexOf(string(char), dictioanary))
		if chars[sym] {
			return false
		}

		// Добавляем символ в мапу
		chars[sym] = true
	}

	// Если цикл прошел успешно, значит все символы уникальны
	return true
}

func SetTestText(inEn *walk.TextEdit, convertText func(string) string) {
	if reflect.ValueOf(convertText) == reflect.ValueOf(CleanConvertText) {
		inEn.SetText(data.GetText())
	} else {
		inEn.SetText(data.GetProverb())
	}
}

func DirtyConvertText(text string) string {
	text = strings.ToLower(text)
	text = strings.Replace(text, " ", "", -1)
	text = strings.Replace(text, ",", "зпт", -1)
	text = strings.Replace(text, ".", "тчк", -1)
	return text
}

func CleanConvertText(text string) string {
	text = strings.Replace(text, ",", "зпт", -1)
	text = strings.Replace(text, " ", "прб", -1)
	text = strings.Replace(text, ".", "тчк", -1)
	re := regexp.MustCompile("[А-Я]")
	text = re.ReplaceAllStringFunc(text, func(s string) string {
		return "лш" + strings.ToLower(s)
	})
	return text
}

func DirtyDeConvertText(text string) string {
	return text
}

func CleanDeConvertText(text string) string {
	text = strings.Replace(text, "зпт", ",", -1)
	text = strings.Replace(text, "тчк", ".", -1)
	text = strings.Replace(text, "прб", " ", -1)
	re := regexp.MustCompile("лш([а-я])")
	text = re.ReplaceAllStringFunc(text, func(s string) string {
		return strings.ToUpper(s[len("лш"):])
	})
	return text
}

func GetFirstOfSplitChars(text string) string {
	return strings.Split(text, "/")[0]
}
