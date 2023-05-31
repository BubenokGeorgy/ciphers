package _const

import "github.com/lxn/walk"

//Structs
type EnDecrypt struct {
	Encrypt CiphersErrors
	Decrypt CiphersErrors
}

type Key struct {
	Label string
	Visible bool
}

type WindowConf struct {
	InDeVisible bool
	OutDeVisible bool
	InEnVisible bool
	OutEnVisible bool
	DecryptButtonVisible bool
	WriteTestTextVisible bool
	EncryptButtonVisible bool
	LongButtonVisible bool
	LongButtonText string
	InStretchFactor int
	DirtyCheckVisible bool
	InEnEnable bool
	OutEnEnable bool
	HeaderFirstVisible bool
	HeaderSecondVisible bool
	HeaderFirstText string
	HeaderSecondText string
}

type NewComposite struct {
	Keys        []Key
	AutoKeys bool
	Id                 string
	FontSize           int
	VariabilityVisible bool
	VariabilityText    string
	WindowConf WindowConf
	Ciphers            []EnDecrypt
}
type Errors struct {
	Regex      string
	ErrorsList []string
}
type CiphersErrors struct {
	Cipher            func(string, []*walk.TextEdit) string
	TextErrors        Errors
	KeyErrors         []Errors
	TextErrorsHandler Handler
	KeyErrorsHandler  Handler
}
type CipherPackage struct {
	Text   string
	Keys   []*walk.TextEdit
	Cipher CiphersErrors
}

//Factories
type Handler func(cipher CipherPackage) string

//Variables
var Dictionary = []string{"а", "б", "в", "г", "д", "е/ё", "ж", "з", "и", "й", "к", "л", "м", "н", "о", "п", "р", "с", "т", "у", "ф", "х", "ц", "ч", "ш", "щ", "ъ", "ы", "ь", "э", "ю", "я"}
var SecondDictionary = []string{"а", "б", "в", "г", "д", "е/ё", "ж", "з", "и/й", "к", "л", "м", "н", "о", "п", "р", "с", "т", "у", "ф", "х", "ц", "ч", "ш", "щ", "ы", "ь/ъ", "э", "ю", "я"}
var CiphersFunctions []CiphersErrors
var TextErrors = Errors{
	Regex:      RegexText,
	ErrorsList: []string{ErrorTextWrite, ErrorTextInValidSym},
}
var CipherTextErrors = Errors{
	Regex:      RegexTextCipher,
	ErrorsList: []string{ErrorTextWriteCipher, ErrorTextInValidSymCipher},
}
var KeyErrors = Errors{
	ErrorsList: []string{ErrorKeyWrite, ErrorKeyInValidSym},
}
var CipherKeyErrors = Errors{
	ErrorsList: []string{ErrorKeyWriteCipher, ErrorKeyInValidSymCipher},
}
var WindowConfStandard = WindowConf{
	InDeVisible:          true,
	OutDeVisible:         true,
	DecryptButtonVisible: true,
	WriteTestTextVisible: true,
	EncryptButtonVisible: true,
	InEnVisible: true,
	OutEnVisible: true,
	LongButtonVisible:    false,
	LongButtonText:       "",
	InStretchFactor: 1,
	DirtyCheckVisible: true,
	InEnEnable: true,
	OutEnEnable: true,
	HeaderFirstVisible: false,
	HeaderSecondVisible: false,
}

//Paths
const PathText = "text.txt"
const PathProverb = "proverb.txt"

//Errors
const ErrorTextInValidSym = "в тексте для шифрования содержатся недопустимые символы"
const ErrorTextWrite = "введите текст для шифрования"
const ErrorKeyInValidSym = "в ключе для шифрования содержатся недопустимые символы"
const ErrorKeyWrite = "введите ключ для шифрования"
const ErrorMatrixNotSquare = "матрица должна быть квадратной"
const ErrorMatrixEmptyElem = "в матрице в каждой строке должно быть равное количество значений"
const ErrorMatrixDegenerate = "матрица не должна быть вырождена"
const ErrorKeyNotUniqueChars = "в ключе для шифрования не должно быть повторяющихся букв"
const ErrorTextRepeatedTwo = "длина текста для шифрования должна быть кратна двум"
const ErrorKeyCardanEmptyElement = "в ключе для шифрования в каждой строке должно быть равное количество значений"
const ErrorKeyCardanRepeatedTwo = "количество элементов в строках и столбцах ключа для шифрования должно быть кратно двум"
const ErrorKeyCardanRepeatedFour = "количество значений в ключе для шифрования должно быть кратно 4"
const ErrorKeySymmetric = "дырки в ключе для шифрования не должны быть симметричны"
const ErrorKeyCardanNumHoles = "количество отверствий в ключе для шифрования должно равняться 1/4 от всех значений"
const ErrorKeyCardanSmall = "ключ для шифрования слишком маленький"

//Cipher
const ErrorTextMatrixNegativeSym = "в шифротексте содержатся значения, которые дают отрицательные результаты при расшифровке"
const ErrorTextMatrixExtraSym = "в шифротексте содержатся значения, которые превышают максимально возможное число при расшифровке"
const ErrorTextInValidSymCipher = "в шифротексте содержатся недопустимые символы"
const ErrorTextWriteCipher = "введите текст для расшифрования"
const ErrorKeyInValidSymCipher = "в ключе для расшифрования содержатся недопустимые символы"
const ErrorKeyWriteCipher = "введите ключ для расшифрования"
const ErrorMatrixRepeatedCipher = "размер матрицы должен быть кратен длине шифртекста"
const ErrorKeyNotUniqueCharsCipher = "в ключе для расшифрования не должно быть повторяющихся букв"
const ErrorTextRepeatedTwoCipher = "длина шифртекста должна быть кратна двум"
const ErrorTextRepeatedPairsCipher = "в шифротексте не должно быть повторяющихся букв в парах"
const ErrorTextTenValuesCipher = "в шифротексте должно быть по 10 значений в строке"
const ErrorTextCardanEmptyElementCipher = "в шифротексте в каждой строке должно быть равное количество значений"
const ErrorKeyCardanEmptyElementCipher = "в ключе для расшифрования в каждой строке должно быть равное количество значений"
const ErrorKeyCardanRepeatedFourCipher = "количество значений в ключе для расшифрования должно быть кратно 4"
const ErrorKeyCardanNumHolesCipher = "количество отверствий в ключе для расшифрования должно равняться 1/4 от всех значений"
const ErrorKeySymmetricSipher = "дырки в ключе для расшифрования не должны быть симметричны"
const ErrorTextCardanRepeatedTwoCipher = "количество элементов в строках и столбцах текста для расшифрования должно быть кратно двум"
const ErrorKeyCardanRepeatedTwoSipher = "количество элементов в строках и столбцах ключа для расшифрования должно быть кратно двум"
const ErrorKeyCardanSmallSipher = "ключ для расшифрования слишком маленький"
const ErrorParameterDiffieHellmanBigNum = "параметр слишком большое число"
const ErrorKeyDiffieHellmanBigNum = "ключ слишком большое число"
const ErrorParameterDiffieHellmanSmallN = "параметр n должен быть больше или равен 3"
const ErrorParameterDiffieHellmanSmallA = "параметр a должен быть больше 1"
const ErrorParameterDiffieHellmanBigA = "параметр a должен быть меньше n"
const ErrorKeyDiffieHellmanSmallSV = "секретный ключ виртуального пользователя должен быть не меньше 2"
const ErrorKeyDiffieHellmanBigSV = "секретный ключ виртуального пользователя должен быть меньше n"
const ErrorKeyDiffieHellmanSmallSO = "наш секретный ключ должен быть не меньше 2"
const ErrorKeyDiffieHellmanBigSO = "наш секретный ключ должен быть меньше n"
const ErrorParameterWrite = "введите параметр №"
const ErrorParameterInValidSym = "параметр содержит недопустимые символы"
const ErrorParameterASVOverBig = "нужно уменьшить либо параметр а либо секретный ключ виртуального пользователя"
const ErrorParameterASOOverBig = "нужно уменьшить либо параметр а либо наш секретный ключ"
const ErrorParameterKeyBig = "нужно уменьшить параметры,общий ключ выходит слишком большим"
//Fonts
const FontSizeTen = 10

//Regex
const RegexText = `^[А-Яа-яёЁ,.\s]+$`
const RegexTextCipher = `^[а-яё]+$`
const RegexTextNumCipher = `^[0-9]+$`
const RegexTextNumSpaceCipher = `^-?\d+(\s-?\d+)*$`
const RegexKeyCaesar = `^-?[0-9]+$`
const RegexKeyBelazo = `^[а-я]+$`
const RegexKeyVigenere = `^[а-я]+$`
const RegexKeyMatrix = `^(-?\d+(\s-?\d+)*[\r]?[\n]?)*$`
const RegexKeyCipher = `^[а-яё]+$`
const RegexTextCardanCipher = `^([а-яё]+(\s[а-яё]+)*[\r]?[\n]?)*$`
const RegexKeyCardan = `^([01](\s[01])*[\r]?[\n]?)*$`
const RegexKeyDiffieHellman = `^[0-9]+$`
