package pages

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"main/ciphers"
	. "main/const"
	. "main/handlers"
	. "main/ui"
)

func NewAtbashPage(parent walk.Container) (Page, error) {
	page := NewComposite{
		FontSize:   FontSizeTen,
		Keys: []Key{{Visible: false}},
		WindowConf: WindowConfStandard,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher: ciphers.AtbashEncrypt,
				TextErrors:  TextErrors,
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.AtbashDecrypt,
				TextErrors:  CipherTextErrors,
			},
		},
		},
	}
	p := new(NewPage)
	if err := (GenerateComposite(p, page)).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}
	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}
	return p, nil

}

func NewCaesarPage(parent walk.Container) (Page, error) {
	page := NewComposite{
		FontSize:   FontSizeTen,
		WindowConf: WindowConfStandard,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher: ciphers.CaesarEncrypt,
				TextErrors:  TextErrors,
				KeyErrors: []Errors{{
					Regex:      RegexKeyCaesar,
					ErrorsList: KeyErrors.ErrorsList,
				}},
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.CaesarDecrypt,
				TextErrors:  CipherTextErrors,
				KeyErrors: []Errors{{
					Regex:      RegexKeyCaesar,
					ErrorsList: CipherKeyErrors.ErrorsList,
				}},
			},
		},
		},
	}
	p := new(NewPage)
	if err := (GenerateComposite(p, page)).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}
	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}
	return p, nil

}

func NewPolybiusPage(parent walk.Container) (Page, error) {
	page := NewComposite{
		FontSize:   FontSizeTen,
		Keys: []Key{{Visible: false}},
		WindowConf: WindowConfStandard,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher: ciphers.PolibiusEncrypt,
				TextErrors:  TextErrors,
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.PolibiusDecrypt,
				TextErrors:  Errors{
					Regex:      RegexTextNumCipher,
					ErrorsList: CipherTextErrors.ErrorsList,
				},
			},
		},
		},
	}
	p := new(NewPage)
	if err := (GenerateComposite(p, page)).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}
	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}
	return p, nil

}

func NewTrithemiusPage(parent walk.Container) (Page, error) {
	page := NewComposite{
		FontSize:   FontSizeTen,
		Keys: []Key{{Visible: false}},
		WindowConf: WindowConfStandard,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher: ciphers.TrithemiusEncrypt,
				TextErrors:  TextErrors,
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.TrithemiusDecrypt,
				TextErrors:  CipherTextErrors,
			},
		},
		},
	}
	p := new(NewPage)
	if err := (GenerateComposite(p, page)).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}
	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}
	return p, nil

}

func NewBelazoPage(parent walk.Container) (Page, error) {
	page := NewComposite{
		FontSize:   FontSizeTen,
		Keys: []Key{{Visible: false}},
		WindowConf: WindowConfStandard,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher: ciphers.BelazoEncrypt,
				TextErrors:  TextErrors,
				KeyErrors: []Errors{{
					Regex:      RegexKeyBelazo,
					ErrorsList: KeyErrors.ErrorsList,
				}},
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.BelazoDecrypt,
				TextErrors:  CipherTextErrors,
				KeyErrors: []Errors{{
					Regex:      RegexKeyBelazo,
					ErrorsList: CipherKeyErrors.ErrorsList,
				}},
			},
		},
		},
	}
	p := new(NewPage)
	if err := (GenerateComposite(p, page)).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}
	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}
	return p, nil

}

func NewVigenerePage(parent walk.Container) (Page, error) {
	page := NewComposite{
		FontSize:           FontSizeTen,
		VariabilityText:    "Ключ Шифртекст",
		VariabilityVisible: true,
		WindowConf: WindowConfStandard,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher: ciphers.VigenereEncrypt,
				TextErrors:  TextErrors,
				KeyErrors: []Errors{{
					Regex:      RegexKeyVigenere,
					ErrorsList: KeyErrors.ErrorsList,
				}},
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.VigenereDecrypt,
				TextErrors:  CipherTextErrors,
				KeyErrors: []Errors{{
					Regex:      RegexKeyVigenere,
					ErrorsList: CipherKeyErrors.ErrorsList,
				}},
			},
		},{
			Encrypt: CiphersErrors{
				Cipher: ciphers.VigenereCipherKeyEncrypt,
				TextErrors:  TextErrors,
				KeyErrors: []Errors{{
					Regex:      RegexKeyVigenere,
					ErrorsList: KeyErrors.ErrorsList,
				}},
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.VigenereCipherKeyDecrypt,
				TextErrors:  CipherTextErrors,
				KeyErrors: []Errors{{
					Regex:      RegexKeyVigenere,
					ErrorsList: CipherKeyErrors.ErrorsList,
				}},
			},
		},
		},
	}
	p := new(NewPage)
	if err := (GenerateComposite(p, page)).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}
	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}
	return p, nil

}

func NewMatrixPage(parent walk.Container) (Page, error) {
	page := NewComposite{
		FontSize:   FontSizeTen,
		WindowConf: WindowConfStandard,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher: ciphers.MatrixEncrypt,
				TextErrors:  TextErrors,
				KeyErrors: []Errors{{
					Regex:      RegexKeyMatrix,
				}},
				KeyErrorsHandler: MatrixKeyHandler,
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.MatrixDecrypt,
				TextErrors:  Errors{
					Regex:      RegexTextNumSpaceCipher,
					ErrorsList: CipherTextErrors.ErrorsList,
				},
				KeyErrors: []Errors{{
					Regex:      RegexKeyMatrix,
				}},
				KeyErrorsHandler: MatrixCipherKeyHandler,
			},
		},
		},
	}
	p := new(NewPage)
	if err := (GenerateComposite(p, page)).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}
	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}
	return p, nil

}

func NewPlayfairPage(parent walk.Container) (Page, error) {
	page := NewComposite{
		FontSize:   FontSizeTen,
		WindowConf: WindowConfStandard,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher: ciphers.PlayfairEncrypt,
				TextErrors:  TextErrors,
				KeyErrors: []Errors{{
					Regex:      RegexKeyCipher,
				}},
				KeyErrorsHandler: PlayfairKeyHandler,
				TextErrorsHandler: PlayfairTextHandler,
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.PlayfairDecrypt,
				TextErrors:  Errors{
					Regex:      RegexTextCipher,
				},
				KeyErrors: []Errors{{
					Regex:      RegexKeyCipher,
				}},
				KeyErrorsHandler: PlayfairCipherKeyHandler,
				TextErrorsHandler: PlayfairCipherTextHandler,
			},
		},
		},
	}
	p := new(NewPage)
	if err := (GenerateComposite(p, page)).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}
	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}
	return p, nil

}

func NewVerticalPermutationPage(parent walk.Container) (Page, error) {
	page := NewComposite{
		FontSize:   FontSizeTen,
		WindowConf: WindowConfStandard,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher: ciphers.VerticalPermutationEncrypt,
				TextErrors:  TextErrors,
				KeyErrors: []Errors{{
					Regex:      RegexKeyCipher,
				}},
				KeyErrorsHandler: VerticalPermutationKeyHandler,
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.VerticalPermutationDecrypt,
				TextErrors:  CipherTextErrors,
				KeyErrors: []Errors{{
					Regex:      RegexKeyCipher,
				}},
				KeyErrorsHandler: VerticalPermutationCipherKeyHandler,
			},
		},
		},
	}
	p := new(NewPage)
	if err := (GenerateComposite(p, page)).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}
	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}
	return p, nil

}

func NewCardanGrillePage(parent walk.Container) (Page, error) {
	page := NewComposite{
		FontSize:   FontSizeTen,
		WindowConf: WindowConfStandard,
		VariabilityText: "Автоматический ключ",
		AutoKeys: true,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher: ciphers.CardanGrilleEncrypt,
				TextErrors:  TextErrors,
				KeyErrors: []Errors{{
					Regex:      RegexKeyCardan,
				}},
				KeyErrorsHandler: CardanGrilleKeyHandler,
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.CardanGrilleDecrypt,
				TextErrors:  Errors{
					Regex:      RegexTextCardanCipher,
				},
				KeyErrors: []Errors{{
					Regex:      RegexKeyCardan,
				}},
				KeyErrorsHandler: CardanGrilleCipherKeyHandler,
				TextErrorsHandler:CardanGrilleCipherTextHandler,
			},
		},
			{
				Encrypt: CiphersErrors{
					Cipher: ciphers.CardanGrilleEncryptStart,
					TextErrors:  TextErrors,
					KeyErrors: []Errors{{
						Regex:      RegexKeyCardan,
					}},
					KeyErrorsHandler: CardanGrilleStartKeyHandler,
				},
				Decrypt: CiphersErrors{
					Cipher: ciphers.CardanGrilleDecryptStart,
					TextErrors:  Errors{
						Regex:      RegexTextCardanCipher,
					},
					KeyErrors: []Errors{{
						Regex:      RegexKeyCardan,
					}},
					KeyErrorsHandler: CardanGrilleCipherStartKeyHandler,
					TextErrorsHandler:CardanGrilleCipherTextHandler,
				},
			},
		},
	}
	p := new(NewPage)
	if err := (GenerateComposite(p, page)).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}
	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}
	return p, nil
}

func NewDiffieHellmanPage(parent walk.Container) (Page, error)  {
	page := NewComposite{
		FontSize:   FontSizeTen,
		WindowConf: WindowConf{
			InDeVisible:          false,
			OutDeVisible:         false,
			InEnVisible: false,
			DecryptButtonVisible: false,
			WriteTestTextVisible: false,
			EncryptButtonVisible: false,
			LongButtonVisible:    true,
			LongButtonText:       "Обменяться ключами.",
			InStretchFactor: 50,
			DirtyCheckVisible: false,
			InEnEnable: false,
			OutEnEnable: false,
			HeaderSecondVisible: true,
			HeaderSecondText: "Получившийся общий ключ.",
		},
		AutoKeys: true,
		Keys: []Key{{Label: "Задайте n.",Visible: true},{Label: "Задайте a.",Visible: true},{Label: "Задайте ключ виртуального пользователя.",Visible: true}, {Label: "Задайте свой ключ.",Visible: true}},
		VariabilityText: "Автоматические ключи",
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher: ciphers.DiffieHellmanEncrypt,
				KeyErrors: []Errors{{
					Regex:      RegexKeyDiffieHellman,
				}},
				TextErrorsHandler:DiffieHellmanTextHandler,
				KeyErrorsHandler: DiffieHellmanKeyHandler,
			},
		},{
			Encrypt: CiphersErrors{
				Cipher: ciphers.DiffieHellmanEncryptStart,
				KeyErrors: []Errors{{
					Regex:      RegexKeyDiffieHellman,
				}},
				TextErrorsHandler:DiffieHellmanTextHandler,
				KeyErrorsHandler: DiffieHellmanStartKeyHandler,
			},
		},
		},
	}
	p := new(NewPage)
	if err := (GenerateComposite(p, page)).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}
	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}
	return p, nil
}
