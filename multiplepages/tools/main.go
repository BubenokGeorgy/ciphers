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

func RsaSetPQ(keys []*walk.TextEdit)  { //функция генерации и установки параметров p и q, очень больших простых чисел, не равных между собой, шифра Rsa
	rand.Seed(time.Now().UnixNano()) //инициализируем рандомайзер от текущего времени
	p := new(big.Int).SetInt64(int64(GeneratePrimeNumber())) //генерируем просто число p
	q := new(big.Int).SetInt64(int64(GeneratePrimeNumber())) //генерируем простое число q
	for { //цикл который подбирает параметр q бесконечно, пока тот не будет не равен p
		if p.Cmp(q) == 0 { //если q = p
			q.SetInt64(int64(GeneratePrimeNumber())) //генерируем q заново
		} else { //если q не равен p
			break //выходим из цикла
		}
	}
	keys[0].SetText(p.String()) //записываем в поле с ключом p - значение p
	keys[1].SetText(q.String()) //записываем в поле с ключом q - значение q
	num1 := big.NewInt(1) //записываем в переменную 1, чтобы потом добавлять или вычитать 1 из чисел без повторного создания переменной
	num2 := big.NewInt(2) //записываем в переменную 2, чтобы потом добавлять или вычитать 2 из чисел без повторного создания переменной
	f := new(big.Int).Mul(p.Sub(p,num1),q.Sub(q,num1)) // находим функцию Эйлера от p и q, путем перемножения (p-1) на (q-1)
	var temp []*big.Int //инициализуруем массив, в котором будут храниться все возможные варианты числа 1<e<f(функции Эйлера), взаимного просто с f(функцией Эйлера)
	for i := num2; i.Cmp(f)!=0; i.Add(i,num1) { //перебираем все числа e от 2 до f(функции Эйлера)
		if Gcd(i, f).Cmp(num1)==0 { //проверяем взаимно ли простые числа текущее число и f(функция Эйлера) путём сравнения их максимального общего делителя с 1
			temp = append(temp, new(big.Int).Set(i)) //если текущее число подходит, добавляем его в массив со всеми возможными вариантами
		}
	}
	e := temp[rand.Intn(len(temp))] //выбираем рандомное e из всех возможных вариантов
	keys[2].SetText(e.String()) //устанавливаем e, как значение 3 ключа, в котором указывается как раз e
}

func InverseMatrix(matrix [][]int) [][]int { //функция, генерирующая обратную матрицу для заданной
	n := len(matrix) //объявляем переменную и записываем в неё длину заданной матрицы

	// Создаем единичную матрицу, которую будем преобразовывать
	identity := make([][]int, n)
	for i := 0; i < n; i++ {
		identity[i] = make([]int, n)
		identity[i][i] = 1
	}

	// Преобразуем исходную матрицу и единичную матрицу одновременно
	for i := 0; i < n; i++ {
		// Находим главный элемент в столбце
		max := i
		for j := i + 1; j < n; j++ {
			if math.Abs(float64(matrix[j][i])) > math.Abs(float64(matrix[max][i])) {
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

func MultiplyMatrix(matrixA, matrixB [][]int) [][]int {

	result := make([][]int, len(matrixA))
	for i := range result {
		result[i] = make([]int, len(matrixB[0]))
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
	matrix := make([][]int, h)
	textMatrix := make([][]string, h)
	for i, _ := range matrix {
		textMatrix[i] =make([]string, w)
		matrix[i] = make([]int, w)
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

func GetMatrix(text string) [][]int { //функция получения матрицы из ключа
 	var matrix [][]int //объявление переменной с итоговой матрицей
	strs := strings.Split(text, "\r\n") //получение массива строк из введённого ключа
	for _, num := range strs { //проход по этому массиву строк
		var matr []int //объявление временного массива для построчного заполнения матрицы
		for _, n := range strings.Split(num, " ") { //проход по значениям в строке
			t, _ := strconv.Atoi(n) //считывание чисел из этих значений
			matr = append(matr, t) //добавление чисел в временный массив
		}
		matrix = append(matrix, matr) //добавление временного массива в общий массив
	}
	return matrix //возвращение результирующего массива
}

func Determinant(m [][]int) int {
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

	mCopy := make([][]int, size)
	for i := range m {
		mCopy[i] = make([]int, size)
		copy(mCopy[i], m[i])
	}

	for i := 0; i < size-1; i++ {
		if math.Abs(float64(mCopy[i][i])) < 1e-10 {
			for j := i + 1; j < size; j++ {
				if math.Abs(float64(mCopy[j][i])) > 1e-10 {
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

	det := 1
	for i := 0; i < size; i++ {
		det *= mCopy[i][i]
	}

	return det
}

func VerticalReverse(matrix [][]int) [][]int {
	// Создаем копию матрицы
	reversed := make([][]int, len(matrix))
	for i := range reversed {
		reversed[i] = make([]int, len(matrix[i]))
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

func HorizontalReverse(matrix [][]int) [][]int {
	// Создаем копию матрицы
	reversed := make([][]int, len(matrix))
	for i := range reversed {
		reversed[i] = make([]int, len(matrix[i]))
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

func GenerateShennonGammaParametrs() (int,int,int,int){
	cKol := 10         //Количество c - взаимно простых чисел c модулем m
	bKol := 10         //Количество b, где b = a – 1 кратно p для каждого простого p, делителя m; b кратно 4, если m кратно 4
	m := float64(len(Dictionary)) //Размер алфавита

	rand.Seed(time.Now().UnixNano()) //инициализируем рандомайзер текущим временем

	//Создание массива со значениями c - взаимно простых чисел c модулем m, и присвоение одного из них самой c для рандомизации
	var cMas []float64 //инициализируем массив с будущими значениями C
	var i float64 //инициализируем переменную для подбора этих значений
	for i = 2; ; i++ { //перебираем все числа от 2
		if AreCoprime(i, math.Abs(m)) {//если текущее подобранное число взаимно простое с модулем m(размером алфавита),
			cMas = append(cMas, i) //то добавляем это число в массив со значениями C
			if len(cMas) >= cKol { //если подобранно нужное количество значений C
				break //то выходим из цикла подбора значений
			}
		}
	}
	c := cMas[rand.Intn(len(cMas))] //рандомно выбираем одно C из массива с подобранными C

	//Создание массива со значениями p - простых делителей m
	var pMas []float64 //инициализируем массив с будущими значениями C
	for i = 2; i <= m; i++ { //перебираем все числа от 2 до m включительно
		if !IsPrime(int(i)){ //если число не простое, то переходим к началу цикла,проверяя следующее число
			continue
		}
		if math.Mod(m, i) == 0 { //если число просто и остаток деления m на это число 0, то
			pMas = append(pMas, i) //добавляем число в массив со значениями p - простых делителей m
		}
	}

	//Создание массива со значениями b и присвоение одного из них самой b для рандомизации,
	//где b = a – 1 кратно p для каждого простого p, делителя m; b кратно 4, если m кратно 4
	var bMas []float64 //инициализируем массив с будущими значениями b
	bMultiple4 := false //в чек-переменную,отвечающую за определение должна ли b быть кратна четырем, записываем false
	if math.Mod(m, 4) == 0 { //если m кратно 4
		bMultiple4 = true //то b тоже должно быть кратно 4, значит в чек-переменную записываем true
	}
	for i = 2; ; i++ { //перебираем все числа от 2
		if math.Mod(i,2)!=0{ //так как b=a-1, a - нечетное число, то b должно быть четным,
			continue //иначе идем к началу цикла, проверяя следующее число
		}
		if bMultiple4 { //если m кратно 4, то b тоже должно, провермя по чек-переменной
			if math.Mod(i, 4) != 0 {//если текущее число не кратно 4
				continue //идем к началу цикла, проверяя следующее число
			}
		}
		if !IsMultipleAllP(i, pMas) { //b  должно быть кратно p для каждого простого p, делителя m,
										// вызываем функцию, которая это проверяет, если это не так
			continue //идем к началу цикла, проверяя следующее число
		}
		bMas = append(bMas, i) //добавляем текущее число в массив с числами b = a – 1 кратно p для каждого
		                     // простого p, делителя m; b кратно 4, если m кратно 4
		if len(bMas) >= bKol { //если подобранно нужное количество значений b
			break //то выходим из цикла подбора значений
		}
	}
	b := bMas[rand.Intn(len(bMas))] //рандомно выбираем b из возможных вариантов

	//Инициализация a
	a := b+1 //инициализруем a, которая равна b+1
	return int(a),int(b),int(c),int(m) //возвращаем значения a,b,c,m
}

func ShennonParametrsSet(keys []*walk.TextEdit ){
	a,b,c,m := GenerateShennonGammaParametrs() //получаем сгенерированные параметры a,b,c,m для Блокнота Шеннона
	mas := GenerateShennonGamma(a,b,c,m) //генерируем гамму для блокнота Шеннона из полученных параметров
	var strMas []string //инициализиурем переменную, в которую будут записываться строковые представления
	                    //чисел из гаммы
	for _, elem := range mas{ //роходимся по массиву гаммы
		strMas = append(strMas,strconv.Itoa(elem)) //записываем в строковый массив строковое представление
													//чисел из массива гаммы
	}
	keys[0].SetText(strings.Join(strMas," ")) //записываем в поле параметра гаммы полученную гамму
	keys[1].SetText(strconv.Itoa(a)) //записываем в поле параметра А значение переменной A
	keys[2].SetText(strconv.Itoa(b)) //записываем в поле параметра b значение переменной b
	keys[3].SetText(strconv.Itoa(c)) //записываем в поле параметра c значение переменной c
	keys[4].SetText(strconv.Itoa(m)) //записываем в поле параметра m значение переменной m
}

// Генерация символьной гаммы
func GenerateShennonGamma(a,t0,c,m int) []int { //функция генерации гаммы для блокнота Шеннона из полученных параметров

	//Генерация гаммы
	mas := make([]int, m+1) //создаем массив длинной больше на 1 элемент, чем максимально возможная длинна гаммы
	mas[0] =t0 //первый элемент инициализируем порождающим числом t0
	for i := 0; i < m; i++ { //проходимся по всем значениям массива
		mas[i+1] = (a*mas[i]+c)%m //записывая в следующую ячейку массива
								// новое значение, получение из действий с другими параметрами и значением из предыдущей ячейки
	}
	mas = mas[1:] //обрезаем итоговую гамму, убирая порождающее число

	return mas //возвращаем итоговую гамму
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

func HasSymmetricOnes(matrix [][]int, h, w int) bool {
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

func IndexOf(element string, data []string) int { //получение индекса символа в массиве символов
	for k, v := range data { //проход по массиву символов
		for _, let := range strings.Split(v, "/") { //проход по разделенному текущему символу из массива символов при возможности, если это е/ё
			if element == let { //если символ для сравнения равен с текущим символом
				return k //то возвращаем его индекс
			}
		}
	}
	return -1 //если не было найдено равных символов, то возвращаем -1
}

func GetTrithemiusTable() [][]string { //функция генерации таблицы Тритемия
	var mas [][]string //создание результирующий массив
	for i := 0; i < len(Dictionary); i++ { //проход по всем символам словаря
		mas = append(mas, append(Dictionary[i:], Dictionary[:i]...)) //добавление массива со сдвинутыми на ещё 1 символами словаря в общий массив
	}
	return mas //возвращение результирующего массива с таблицей Тритемия
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

func GetCaesarKey(keys []*walk.TextEdit) int { //функция, возвращающая текущее значение ключа, учитывая ограничения длины словаря
	key, _ := strconv.Atoi(keys[0].Text()) //получение текущего значения ключа, введённого пользователем в окно редактирования текста
	key %=len(Dictionary) //получение остатка от деления ключа на длину словаря, чтобы ключ не превышал вышеуказанную длину
	return key //возвращение результирующего ключа
}

func TransformCaesarText(key int, text string) string { //функция, производяшая сдвиг текста для шифрования и расшифрования
	var num, secNum int //объявление переменных, в которых будут храниться индексы текущей буквы и буквы после сдвига
	var result string //объявление переменной, в которой будет храниться результирующий текст
	dicLen := len(Dictionary) //инициализирование новой переменной длинной словаря
	for _, v := range text { //проход посимвольно по тексту для сдвига
		num = IndexOf(string(v), Dictionary) //получение индекса текущей буквы в словаре
		secNum = num+key //получение индекса буквы после сдвига
		secNum = (secNum+dicLen)%dicLen //получение индекса буквы после сдвига так, чтобы он не выходил за пределы словаря
		result += GetElement(Dictionary, secNum) //добавление в результируюшую строку нового символа, полученного при сдвиге
	}
	return result //возвращение результирующую строку
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

func GetHash(text string, mod *big.Int) *big.Int { //функция хэширования
	h := big.NewInt(0) //создаем переменную нашего хэша и записываем в нее 0
	for _, v := range text { //проходимся посимвольно по тексту
		num := IndexOf(string(v), Dictionary) + 1 //получаем индекс текущего символа в нашем алфавите и
													//прибавляем единицу, так как у нас должны быть символы с индексом от 1
		fmt.Print(num," ")
		h.Add(h, big.NewInt(int64(num))) //прибавляем к нашему текущему хэшу индекс текущего символа, полученный выше
		h.Mul(h, h) //возводим получившийся выше хэш в квадрат
		h.Mod(h, big.NewInt(mod.Int64())) //получаем mod получившегося выше хэша на число для moda
	}
	return h //возвращаем итоговый хэш
}

func AreCoprime(a, b float64) bool { //функция, проверяющая взаимно ли простые числа
	k := big.NewInt(int64(a)) //инициализируем переменную со значением одного числа
	l := big.NewInt(int64(b)) //инициализируем переменную со значением другого числа
	return Gcd(k,l).Cmp(big.NewInt(1))==0 //вызываем функцию, определяющая наибольшьй общий делитель,
											//и возвращаем сравнение наибольшего общего делителя с единицей,
											//чтобы проверить взаимно ли простые числа
}

func Gcd(a, b *big.Int) *big.Int { //функция, находящая наибольшой общий делитель
	for b.Sign() != 0 { //цикл продолжается, пока второго число не станет равным нулю
		a, b = b, new(big.Int).Mod(a, b) //записываем в первое и второе число - соответственно второе и результат
											//moda первого на второе число
	}
	return a //возвращаем итоговое значение в первом числе
}


func GeneratePrimeNumber() int { //функция для генерации простых чисел
	rand.Seed(time.Now().UnixNano()) //инициализируем рандомайзер от текущего времени
	for { //бесконечный цикл для генерации простого числа
		num := rand.Intn(10) + 3 //генерирует рандомное число от 3(так как 1 не простое число) до 100
		prime := true //чек-переменную устанавливаем в значение true
		for i := 2; i*i <= num; i++ { //перебираем все числа от 2, пока квадрат текущего числа меньше либо равен нашему подобранному числу
			if num%i == 0 { //если остаток от деления подобранного рандомного числа на текущее число из цикла равен 0, то
				prime = false //чек-переменную устанавливаем в значение false
				break //выходим из цикла
			}
		}
		if prime { //если не было найдено числа, на которое разделилось бы подобранное рандомное без остатка
			return num //то возвращаем подобранное рандомное число, как простое
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
func IsMultipleAllP(a float64, pMas []float64) bool { //функция, проверяющая кратность первого числа
													// всем значениям из массива второй переменной
	for _, p := range pMas { //проходим по всем значениям p
		if math.Mod(a, p) != 0 { //если первое число не делится на p без остатка
			return false //возвращаем false, что значит, что число не кратно всем значениям
		}
	}
	return true //возвращаем true, что значит, что число кратно всем значениям
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

var transformationTable = [][]int{
	{1, 7, 14, 13, 0, 5, 8, 3, 4, 15, 10, 6, 9, 12, 11, 2},
	{8, 14, 2, 5, 6, 9, 1, 12, 15, 4, 11, 0, 13, 10, 3, 7},
	{5, 13, 15, 6, 9, 2, 12, 10, 11, 7, 8, 1, 4, 3, 14, 0},
	{7, 15, 5, 10, 8, 1, 6, 13, 0, 9, 3, 14, 11, 4, 2, 12},
	{12, 8, 2, 1, 13, 4, 15, 6, 7, 0, 10, 5, 3, 14, 9, 11},
	{11, 3, 5, 8, 2, 15, 10, 13, 14, 1, 7, 4, 12, 9, 6, 0},
	{6, 8, 2, 3, 9, 10, 5, 12, 1, 14, 4, 7, 11, 13, 0, 15},
	{12, 4, 6, 2, 10, 5, 11, 9, 14, 8, 13, 7, 0, 3, 15, 1},
}

func convertToInt(sym string) int {
	num, _ := strconv.ParseInt(sym, 2, 64)
	return int(num)
}

func overwriteMode(bitNumberIn string) string {
	var bitNumberInOut string
	for i := 0; i < 8; i++ {
		num1 := new(big.Int)
		numTemp := new(big.Int)
		num1.SetString(bitNumberIn[i*4:i*4+4], 2)
		index := convertToInt(bitNumberIn[i*4 : i*4+4])
		numTemp.SetString(strconv.Itoa(transformationTable[i][index]), 10)
		num2 := numTemp.Text(2)
		bitNumberInOut += FillZerosBeforeNumber(num2, 4)
	}
	return bitNumberInOut
}

func FillZerosBeforeNumber(num1 string, length int) string {
	for len(num1) < length {
		num1 = "0" + num1
	}
	return num1
}

func FillZerosAfterNumber(num1 string, length int) string {
	for len(num1) < length {
		num1 += "0"
	}
	return num1
}

func xorWithBase(num1, num2 string, base int) *big.Int {
	numLeft, _ := new(big.Int).SetString(num1, base)
	numRight, _ := new(big.Int).SetString(num2, base)
	result := new(big.Int)
	result.Xor(numLeft, numRight)
	return result
}

func TransformKey(key string, keyLen int) string {
	for utf8.RuneCountInString(key) < keyLen {
		key += key
	}
	return string([]rune(key)[:keyLen])
}

func CutKey(key string) []string {
	key = TransformKey(key, 64)
	keyBigInt := new(big.Int)
	keyBigInt.SetString(key, 16)
	keyStr := keyBigInt.Text(2)
	keys := make([]string, 0)
	for i := 0; i < 3; i++ {
		for j := 0; j < 8; j++ {
			if j*32+32 > 255 {
				keys = append(keys, keyStr[j*32:256])
			} else {
				keys = append(keys, keyStr[j*32:j*32+32])
			}

		}
	}
	for i := 7; i >= 0; i-- {
		if i*32+32 > 255 {
			keys = append(keys, keyStr[i*32:256])
		} else {
			keys = append(keys, keyStr[i*32:i*32+32])
		}
	}
	return keys
}

func transformation(numLeft, numRight string, key string) (string, string) {
	keyHex := new(big.Int)
	keyHex.SetString(key, 2)
	vHR := new(big.Int)
	vHR.SetString(numRight, 2)
	res := new(big.Int)
	res.Add(keyHex, vHR)
	numLeftOut := numRight
	numRightOut := FillZerosBeforeNumber(res.Text(2), 32)
	numRightOut = numRightOut[len(numRightOut)-32:]
	numRightOut = overwriteMode(numRightOut)
	numRightOut = numRightOut[11:] + numRightOut[:11]
	numRightOut = xorWithBase(numRightOut, numLeft, 2).Text(2)
	return numLeftOut, numRightOut
}

func ChainOfTransformations(numLeft, numRight string, key []string, move string) string {
	var start, stop, step int
	if strings.Index(move, "reverse") != -1 {
		start = 31
		stop = 0
		step = -1
	} else {
		start = 0
		stop = 31
		step = 1
	}
	for i := start; i != stop; i += step {
		numLeft, numRight = transformation(numLeft, numRight, key[i])
	}
	numRightLast := numRight
	numLeft, numRight = transformation(numLeft, numRight, key[stop])
	return FillZerosBeforeNumber(numRight, 32) + FillZerosBeforeNumber(numRightLast, 32)
}

func TextToHex(text string) string{
	cipherBinStr := ""
	t := new(big.Int)
	for _,v := range text {
		t.SetInt64(int64(IndexOf(string(v),Dictionary)))
		cipherBinStr+=FillZerosBeforeNumber(t.Text(2),32)
	}
	cipherHexStr := ""
	for i:=0;i*4+4<=len(cipherBinStr);i++{

		t.SetString(cipherBinStr[i*4:i*4+4],2)
		cipherHexStr+=t.Text(16)
	}
	return cipherHexStr
}

func HexToText(text string) string{
	cipherBinStr := ""
	t := new(big.Int)
	for _,v := range text {
		t.SetString(string(v),16)
		cipherBinStr+=FillZerosBeforeNumber(t.Text(2),4)
	}
	var cipherStr string
	for i:=0;i*32+32<=len(cipherBinStr);i++{
		t.SetString(cipherBinStr[i*32:i*32+32],2)
		cipherStr += GetElement(Dictionary,int(t.Int64()))
	}
	return cipherStr
}

var pi = []int64{
	252, 238, 221, 17, 207, 110, 49, 22, 251, 196, 250, 218, 35, 197, 4, 77,
	233, 119, 240, 219, 147, 46, 153, 186, 23, 54, 241, 187, 20, 205, 95, 193,
	249, 24, 101, 90, 226, 92, 239, 33, 129, 28, 60, 66, 139, 1, 142, 79,
	5, 132, 2, 174, 227, 106, 143, 160, 6, 11, 237, 152, 127, 212, 211, 31,
	235, 52, 44, 81, 234, 200, 72, 171, 242, 42, 104, 162, 253, 58, 206, 204,
	181, 112, 14, 86, 8, 12, 118, 18, 191, 114, 19, 71, 156, 183, 93, 135,
	21, 161, 150, 41, 16, 123, 154, 199, 243, 145, 120, 111, 157, 158, 178, 177,
	50, 117, 25, 61, 255, 53, 138, 126, 109, 84, 198, 128, 195, 189, 13, 87,
	223, 245, 36, 169, 62, 168, 67, 201, 215, 121, 214, 246, 124, 34, 185, 3,
	224, 15, 236, 222, 122, 148, 176, 188, 220, 232, 40, 80, 78, 51, 10, 74,
	167, 151, 96, 115, 30, 0, 98, 68, 26, 184, 56, 130, 100, 159, 38, 65,
	173, 69, 70, 146, 39, 94, 85, 47, 140, 163, 165, 125, 105, 213, 149, 59,
	7, 88, 179, 64, 134, 172, 29, 247, 48, 55, 107, 228, 136, 217, 231, 137,
	225, 27, 131, 73, 76, 63, 248, 254, 141, 83, 170, 144, 202, 216, 133, 97,
	32, 113, 103, 164, 45, 43, 9, 91, 203, 155, 37, 208, 190, 229, 108, 82,
	89, 166, 116, 210, 230, 244, 180, 192, 209, 102, 175, 194, 57, 75, 99, 182,
}

var piInv = []int64{
	165, 45, 50, 143, 14, 48, 56, 192, 84, 230, 158, 57, 85, 126, 82, 145,
	100, 3, 87, 90, 28, 96, 7, 24, 33, 114, 168, 209, 41, 198, 164, 63,
	224, 39, 141, 12, 130, 234, 174, 180, 154, 99, 73, 229, 66, 228, 21, 183,
	200, 6, 112, 157, 65, 117, 25, 201, 170, 252, 77, 191, 42, 115, 132, 213,
	195, 175, 43, 134, 167, 177, 178, 91, 70, 211, 159, 253, 212, 15, 156, 47,
	155, 67, 239, 217, 121, 182, 83, 127, 193, 240, 35, 231, 37, 94, 181, 30,
	162, 223, 166, 254, 172, 34, 249, 226, 74, 188, 53, 202, 238, 120, 5, 107,
	81, 225, 89, 163, 242, 113, 86, 17, 106, 137, 148, 101, 140, 187, 119, 60,
	123, 40, 171, 210, 49, 222, 196, 95, 204, 207, 118, 44, 184, 216, 46, 54,
	219, 105, 179, 20, 149, 190, 98, 161, 59, 22, 102, 233, 92, 108, 109, 173,
	55, 97, 75, 185, 227, 186, 241, 160, 133, 131, 218, 71, 197, 176, 51, 250,
	150, 111, 110, 194, 246, 80, 255, 93, 169, 142, 23, 27, 151, 125, 236, 88,
	247, 31, 251, 124, 9, 13, 122, 103, 69, 135, 220, 232, 79, 29, 78, 4,
	235, 248, 243, 62, 61, 189, 138, 136, 221, 205, 11, 19, 152, 2, 147, 128,
	144, 208, 36, 52, 203, 237, 244, 206, 153, 16, 68, 64, 146, 58, 1, 38,
	18, 26, 72, 104, 245, 129, 139, 199, 214, 32, 10, 8, 0, 76, 215, 116,
}

func S(x *big.Int) *big.Int {
	xStr := FillZerosBeforeNumber(x.Text(2),128)
	var resStr string
	for i := 0; i <= 15; i++ {
		t,_:= new(big.Int).SetString(xStr[i*8:i*8+8],2)
		t.SetInt64(pi[int(t.Int64())])
		resStr+=FillZerosBeforeNumber(t.Text(2),8)
	}
	res, _ := new(big.Int).SetString(resStr,2)
	return res
}

func SInv(x *big.Int) *big.Int {
	xStr := FillZerosBeforeNumber(x.Text(2),128)
	var resStr string
	for i := 0; i <= 15; i++ {
		t,_:= new(big.Int).SetString(xStr[i*8:i*8+8],2)
		t.SetInt64(piInv[int(t.Int64())])
		resStr+=FillZerosBeforeNumber(t.Text(2),8)
	}
	res, _ := new(big.Int).SetString(resStr,2)
	return res
}

func MultiplyIntsAsPolynomials(x, y int) int {
	if x == 0 || y == 0 {
		return 0
	}
	z := 0
	for x != 0 {
		if x&1 == 1 {
			z ^= y
		}
		y <<= 1
		x >>= 1
	}
	return z
}

func NumberBits(x int) int {
	nb := 0
	for x != 0 {
		nb++
		x >>= 1
	}
	return nb
}

func ModIntAsPolynomial(x, m int) int {
	nbm := NumberBits(m)
	for {
		nbx := NumberBits(x)
		if nbx < nbm {
			return x
		}
		mshift := m << (nbx - nbm)
		x ^= mshift
	}
}


func KuznyechikMultiplication(x, y int) int {
	z := MultiplyIntsAsPolynomials(x, y)
	m := int64(0b111000011)
	return ModIntAsPolynomial(z, int(m))
}

func KuznyechikLinearFunctional(x *big.Int) *big.Int {
	C := []int{148, 32, 133, 16, 194, 192, 1, 251, 1, 192, 194, 16, 133, 32, 148, 1}
	y := new(big.Int)
	xStr := FillZerosBeforeNumber(x.Text(2),128)
	for i:=15;i>=0;i-- {
		num,_ := new(big.Int).SetString(xStr[i*8:i*8+8],2)
		res := new(big.Int).SetInt64(int64(C[i]))
		res.SetInt64(int64(KuznyechikMultiplication(int(num.Int64()),int(res.Int64()))))
		y.Xor(y,res)
	}
	return y
}

func R(x *big.Int) *big.Int {
	a := KuznyechikLinearFunctional(x)
	numText := FillZerosBeforeNumber(a.Text(2),8)+FillZerosBeforeNumber(x.Text(2),128)[:120]
	res,_ := new(big.Int).SetString(numText,2)
	return res
}


func RInv(x *big.Int) *big.Int {
	xStr := FillZerosBeforeNumber(x.Text(2),128)
	a,_ := new(big.Int).SetString(xStr[:8],2)
	x.SetString(xStr[8:]+"00000000",2)
	res := new(big.Int).Xor(x,a)
	b := KuznyechikLinearFunctional(res)
	res.Xor(x,b)
	return res
}


func L(x *big.Int) *big.Int {
	for i := 0; i < 16; i++ {
		x = R(x)
	}
	return x
}

func LInv(x *big.Int) *big.Int {
	for i := 0; i < 16; i++ {
		x = RInv(x)

	}
	return x
}

func KuznyechikKeySchedule(k *big.Int) []*big.Int {
	var keys []*big.Int
	kStr := k.Text(2)
	aStr := kStr[:len(kStr)/2]
	bStr := kStr[len(kStr)/2:]
	a,_ := new(big.Int).SetString(aStr,2)
	b,_ := new(big.Int).SetString(bStr,2)
	keys = append(keys, a)
	keys = append(keys, b)
	for i := 0; i < 4; i++ {
		for j := 0; j < 8; j++ {
			n := new(big.Int).SetInt64(int64(8*i + j + 1))
			c := L(n)
			res := new(big.Int).Xor(a,c)
			res.Xor(L(S(res)),b)
			a, b = res, a
		}
		keys = append(keys, a)
		keys = append(keys, b)
	}
	return keys
}

func FormatTextGetParts(text string, partLen int) []string{
	textStr := ""
	t := new(big.Int)
	for _,v := range text {
		t.SetString(string(v),16)
		textStr+=FillZerosBeforeNumber(t.Text(2),4)
	}

	if len(textStr)%8 != 0 {
		textStr = FillZerosBeforeNumber(textStr, (len(textStr)/8)*8+8)
	}
	textArray := make([]string, 0)

	count := len(textStr) / partLen
	if len(textStr)%partLen != 0 {
		count++
	}
	for i := 0; i < count; i++ {
		var textForAppend string
		if i*partLen+partLen >= len(textStr) {
			textForAppend = textStr[i*partLen:]
		} else {
			textForAppend = textStr[i*partLen : i*partLen+partLen]
		}
		textForAppend = FillZerosAfterNumber(textForAppend, partLen)
		textArray = append(textArray, textForAppend)
	}
	return textArray
}

func KeyToBits(key string) string {
	var result string
	for _, keyChar := range key {
		result += transformBit(strconv.FormatInt(int64(IndexOf(string(keyChar),Dictionary)),2))
	}
	return result
}

func transformBit(alpha string) string {
	if len(alpha) < 5 {
		alpha = strings.Repeat("0", 5-len(alpha)) + alpha
	}
	return alpha
}


func ToInt(str string) int {
	if str == "1" {
		return 1
	}
	return 0
}

func RemoveDuplicates(strSlice []string) []string {
	seen := make(map[string]struct{})
	result := make([]string, 0)
	for _, str := range strSlice {
		if _, ok := seen[str]; !ok {
			seen[str] = struct{}{}
			result = append(result, str)
		}
	}
	return result
}

func ChunkString(str string, chunkSize int) []string {
	var chunks []string
	for i := 0; i < len(str); i += chunkSize {
		end := i + chunkSize
		if end > len(str) {
			end = len(str)
		}
		chunks = append(chunks, str[i:end])
	}
	return chunks
}