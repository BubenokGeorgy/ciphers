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
		Keys:       []Key{{Visible: false}},
		WindowConf: WindowConfStandard,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher:     ciphers.AtbashEncrypt,
				TextErrors: TextErrors,
			},
			Decrypt: CiphersErrors{
				Cipher:     ciphers.AtbashDecrypt,
				TextErrors: CipherTextErrors,
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
				Cipher:     ciphers.CaesarEncrypt,
				TextErrors: TextErrors,
				KeyErrors: []Errors{{
					Regex:      RegexKeyCaesar,
					ErrorsList: KeyErrors.ErrorsList,
				}},
			},
			Decrypt: CiphersErrors{
				Cipher:     ciphers.CaesarDecrypt,
				TextErrors: CipherTextErrors,
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
		Keys:       []Key{{Visible: false}},
		WindowConf: WindowConfStandard,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher:     ciphers.PolibiusEncrypt,
				TextErrors: TextErrors,
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.PolibiusDecrypt,
				TextErrors: Errors{
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
		Keys:       []Key{{Visible: false}},
		WindowConf: WindowConfStandard,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher:     ciphers.TrithemiusEncrypt,
				TextErrors: TextErrors,
			},
			Decrypt: CiphersErrors{
				Cipher:     ciphers.TrithemiusDecrypt,
				TextErrors: CipherTextErrors,
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
		WindowConf: WindowConfStandard,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher:     ciphers.BelazoEncrypt,
				TextErrors: TextErrors,
				KeyErrors: []Errors{{
					Regex:      RegexKeyBelazo,
					ErrorsList: KeyErrors.ErrorsList,
				}},
			},
			Decrypt: CiphersErrors{
				Cipher:     ciphers.BelazoDecrypt,
				TextErrors: CipherTextErrors,
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
		VariabilityText:    KeyCipherText,
		VariabilityVisible: true,
		WindowConf:         WindowConfStandard,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher:     ciphers.VigenereEncrypt,
				TextErrors: TextErrors,
				KeyErrors: []Errors{{
					Regex:      RegexKeyVigenere,
					ErrorsList: KeyErrors.ErrorsList,
				}},
			},
			Decrypt: CiphersErrors{
				Cipher:     ciphers.VigenereDecrypt,
				TextErrors: CipherTextErrors,
				KeyErrors: []Errors{{
					Regex:      RegexKeyVigenere,
					ErrorsList: CipherKeyErrors.ErrorsList,
				}},
			},
		}, {
			Encrypt: CiphersErrors{
				Cipher:     ciphers.VigenereCipherKeyEncrypt,
				TextErrors: TextErrors,
				KeyErrors: []Errors{{
					Regex:      RegexKeyVigenere,
					ErrorsList: KeyErrors.ErrorsList,
				}},
			},
			Decrypt: CiphersErrors{
				Cipher:     ciphers.VigenereCipherKeyDecrypt,
				TextErrors: CipherTextErrors,
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
				Cipher:     ciphers.MatrixEncrypt,
				TextErrors: TextErrors,
				KeyErrors: []Errors{{
					Regex: RegexKeyMatrix,
				}},
				KeyErrorsHandler: MatrixKeyHandler,
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.MatrixDecrypt,
				TextErrors: Errors{
					Regex:      RegexTextNumSpaceCipher,
					ErrorsList: CipherTextErrors.ErrorsList,
				},
				KeyErrors: []Errors{{
					Regex: RegexKeyMatrix,
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
				Cipher:     ciphers.PlayfairEncrypt,
				TextErrors: TextErrors,
				KeyErrors: []Errors{{
					Regex: RegexKeyCipher,
				}},
				KeyErrorsHandler:  PlayfairKeyHandler,
				TextErrorsHandler: PlayfairTextHandler,
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.PlayfairDecrypt,
				TextErrors: Errors{
					Regex: RegexTextCipher,
				},
				KeyErrors: []Errors{{
					Regex: RegexKeyCipher,
				}},
				KeyErrorsHandler:  PlayfairCipherKeyHandler,
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
				Cipher:     ciphers.VerticalPermutationEncrypt,
				TextErrors: TextErrors,
				KeyErrors: []Errors{{
					Regex: RegexKeyCipher,
				}},
				KeyErrorsHandler: VerticalPermutationKeyHandler,
			},
			Decrypt: CiphersErrors{
				Cipher:     ciphers.VerticalPermutationDecrypt,
				TextErrors: CipherTextErrors,
				KeyErrors: []Errors{{
					Regex: RegexKeyCipher,
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
		FontSize:        FontSizeTen,
		WindowConf:      WindowConfStandard,
		VariabilityText: AutoKey,
		AutoKeys:        true,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher:     ciphers.CardanGrilleEncrypt,
				TextErrors: TextErrors,
				KeyErrors: []Errors{{
					Regex: RegexKeyCardan,
				}},
				KeyErrorsHandler: CardanGrilleKeyHandler,
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.CardanGrilleDecrypt,
				TextErrors: Errors{
					Regex: RegexTextCardanCipher,
				},
				KeyErrors: []Errors{{
					Regex: RegexKeyCardan,
				}},
				KeyErrorsHandler:  CardanGrilleCipherKeyHandler,
				TextErrorsHandler: CardanGrilleCipherTextHandler,
			},
		},
			{
				Encrypt: CiphersErrors{
					Cipher:     ciphers.CardanGrilleEncryptStart,
					TextErrors: TextErrors,
					KeyErrors: []Errors{{
						Regex: RegexKeyCardan,
					}},
					KeyErrorsHandler: CardanGrilleStartKeyHandler,
				},
				Decrypt: CiphersErrors{
					Cipher: ciphers.CardanGrilleDecryptStart,
					TextErrors: Errors{
						Regex: RegexTextCardanCipher,
					},
					KeyErrors: []Errors{{
						Regex: RegexKeyCardan,
					}},
					KeyErrorsHandler:  CardanGrilleCipherStartKeyHandler,
					TextErrorsHandler: CardanGrilleCipherTextHandler,
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

func NewDiffieHellmanPage(parent walk.Container) (Page, error) {
	page := NewComposite{
		FontSize: FontSizeTen,
		WindowConf: WindowConf{
			InEnVisible:         true,
			LongButtonVisible:   true,
			LongButtonText:      ChangeKeys,
			InStretchFactor:     50,
			HeaderFirstVisible:  true,
			HeaderFirstText:     ResultingOurFinalKey,
			HeaderSecondVisible: true,
			HeaderSecondText:    ResultingFinalKeyVirtUser,
			OutInEn:             true,
		},
		AutoKeys:        true,
		Keys:            []Key{{Label: EnterN, Visible: true, Enable: true}, {Label: EnterA, Visible: true, Enable: true}, {Label: EnterKeyVirtUser, Visible: true, Enable: true}, {Label: EnterYourKey, Visible: true, Enable: true}},
		VariabilityText: AutoKeys,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher: ciphers.DiffieHellmanEncrypt,
				KeyErrors: []Errors{{
					Regex: RegexKeyDiffieHellman,
				}},
				TextErrorsHandler: DiffieHellmanTextHandler,
				KeyErrorsHandler:  DiffieHellmanKeyHandler,
			},
		}, {
			Encrypt: CiphersErrors{
				Cipher: ciphers.DiffieHellmanEncryptStart,
				KeyErrors: []Errors{{
					Regex: RegexKeyDiffieHellman,
				}},
				TextErrorsHandler: DiffieHellmanTextHandler,
				KeyErrorsHandler:  DiffieHellmanStartKeyHandler,
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

func NewGOSTR341094(parent walk.Container) (Page, error) {
	page := NewComposite{
		FontSize: FontSizeTen,
		WindowConf: WindowConf{
			InEnVisible:          true,
			WriteTestTextVisible: true,
			EncryptButtonVisible: true,
			EncryptButtonText:    CheckWriter,
			InStretchFactor:      50,
			DirtyCheckVisible:    true,
			InEnEnable:           true,
			HeaderSecondVisible:  true,
			HeaderSecondText:     OurWriter,
			OutInEn:              true,
			ThirdEditVisible:     true,
			HeaderThirdVisible:   true,
			HeaderThirdText:      ResultingVirtUserWriter,
		},
		AutoKeys:        true,
		Keys:            []Key{{Label: EnterP, Visible: true, Enable: true}, {Label: EnterQ, Visible: true, Enable: true}, {Label: EnterA, Visible: true, Enable: true}, {Label: EnterX, Visible: true, Enable: true}, {Label: EnterK, Visible: true, Enable: true}},
		VariabilityText: AutoKeys,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher: ciphers.GOSTR341094Encrypt,
				KeyErrors: []Errors{{
					Regex: RegexKeyGOSTR341094,
				}},
				TextErrors: Errors{
					Regex:      TextErrors.Regex,
					ErrorsList: []string{ErrorWriteTextForSignature, ErrorTextInValidSym},
				},
				KeyErrorsHandler: GOSTR341094KeyHandler,
			},
		}, {
			Encrypt: CiphersErrors{
				Cipher: ciphers.GOSTR341094EncryptStart,
				KeyErrors: []Errors{{
					Regex: RegexKeyGOSTR341094,
				}},
				TextErrors: Errors{
					Regex:      TextErrors.Regex,
					ErrorsList: []string{ErrorWriteTextForSignature, ErrorTextInValidSym},
				},
				KeyErrorsHandler: GOSTR341094StartKeyHandler,
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

func NewRsaSignaturePage(parent walk.Container) (Page, error) {
	page := NewComposite{
		FontSize: FontSizeTen,
		WindowConf: WindowConf{
			InEnVisible:          true,
			WriteTestTextVisible: true,
			EncryptButtonVisible: true,
			EncryptButtonText:    CheckWriter,
			InStretchFactor:      50,
			DirtyCheckVisible:    true,
			InEnEnable:           true,
			HeaderSecondVisible:  true,
			HeaderSecondText:     OurWriterHash,
			OutInEn:              true,
			ThirdEditVisible:     true,
			HeaderThirdVisible:   true,
			HeaderThirdText:      ResultingVirtUserWriterHash,
		},
		AutoKeys:        true,
		Keys:            []Key{{Label: EnterP, Visible: true, Enable: true}, {Label: EnterQ, Visible: true, Enable: true},{Label: EnterD, Visible: true, Enable: true}},
		VariabilityText: AutoKeys,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher: ciphers.RsaSignatureEncrypt,
				KeyErrors: []Errors{{
					Regex: RegexKeyRsaSignature,
				}},
				TextErrors: Errors{
					Regex:      TextErrors.Regex,
					ErrorsList: []string{ErrorWriteTextForSignature, ErrorTextInValidSym},
				},
				KeyErrorsHandler: RsaSignatureKeyHandler,
			},
		}, {
			Encrypt: CiphersErrors{
				Cipher: ciphers.RsaSignatureEncryptStart,
				KeyErrors: []Errors{{
					Regex: RegexKeyRsaSignature,
				}},
				TextErrors: Errors{
					Regex:      TextErrors.Regex,
					ErrorsList: []string{ErrorWriteTextForSignature, ErrorTextInValidSym},
				},
				KeyErrorsHandler: RsaSignatureStartKeyHandler,
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

func NewRsaPage(parent walk.Container) (Page, error) {
	page := NewComposite{
		FontSize:        FontSizeTen,
		WindowConf:      WindowConfStandard,
		AutoKeys:        true,
		Keys:            []Key{{Label: EnterP, Visible: true, Enable: true}, {Label: EnterQ, Visible: true, Enable: true}, {Label: EnterD, Visible: true, Enable: true}},
		VariabilityText: AutoKeys,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher: ciphers.RsaEncrypt,
				KeyErrors: []Errors{{
					Regex: RegexKeyRsa,
				}},
				TextErrors: Errors{
					Regex:      TextErrors.Regex,
					ErrorsList: []string{ErrorWriteTextForSignature, ErrorTextInValidSym},
				},
				KeyErrorsHandler: RsaKeyHandler,
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.RsaDecrypt,
				KeyErrors: []Errors{{
					Regex: RegexKeyRsa,
				}},
				TextErrors: Errors{
					Regex:      RegexTextRsa,
					ErrorsList: CipherTextErrors.ErrorsList,
				},
				KeyErrorsHandler: RsaKeyHandler,
			},
		}, {
			Encrypt: CiphersErrors{
				Cipher: ciphers.RsaEncryptStart,
				KeyErrors: []Errors{{
					Regex: RegexKeyRsa,
				}},
				TextErrors: Errors{
					Regex:      TextErrors.Regex,
					ErrorsList: []string{ErrorWriteTextForSignature, ErrorTextInValidSym},
				},
				KeyErrorsHandler: RsaStartKeyHandler,
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.RsaDecryptStart,
				KeyErrors: []Errors{{
					Regex: RegexKeyRsa,
				}},
				TextErrors: Errors{
					Regex:      RegexTextRsa,
					ErrorsList: CipherTextErrors.ErrorsList,
				},
				KeyErrorsHandler: RsaStartKeyHandler,
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

func NewShennonPage(parent walk.Container) (Page, error) {
	page := NewComposite{
		FontSize:        FontSizeTen,
		WindowConf:      WindowConfStandard,
		AutoKeys:        true,
		Keys:            []Key{{Label: CurrentGamma, Visible: true, Enable: true}, {Label: EnterA, Visible: true, Enable: true}, {Label: EnterP0, Visible: true, Enable: true}, {Label: EnterC, Visible: true, Enable: true}, {Label: EnterM, Visible: true, Enable: true}},
		VariabilityText: AutoKeys,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher:           ciphers.ShennonCipher,
				TextErrors:       TextErrors,
				KeyErrorsHandler: ShennonKeyHandler,
			},
			Decrypt: CiphersErrors{
				Cipher:           ciphers.ShennonCipher,
				TextErrors:       CipherTextErrors,
				KeyErrorsHandler: ShennonKeyHandler,
			},
		}, {
			Encrypt: CiphersErrors{
				Cipher:           ciphers.ShennonEncryptStart,
				TextErrors:       TextErrors,
				KeyErrorsHandler: ShennonStartKeyHandler,
			},
			Decrypt: CiphersErrors{
				Cipher:           ciphers.ShennonDecryptStart,
				TextErrors:       CipherTextErrors,
				KeyErrorsHandler: ShennonStartKeyHandler,
			},
		},
		},
	}
	page.WindowConf.InDeEnable = false
	p := new(NewPage)
	if err := (GenerateComposite(p, page)).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}
	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}
	return p, nil
}
func NewA51Page(parent walk.Container) (Page, error) {
	page := NewComposite{
		FontSize:   FontSizeTen,
		WindowConf: WindowConfStandard,
		Keys:       []Key{{Label: EnterKey, Visible: true, Enable: true}},
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher:     ciphers.A51Encrypt,
				TextErrors: TextErrors,
				KeyErrors:  []Errors{{
					Regex: RegexKeyA51,
					ErrorsList: KeyErrors.ErrorsList}},
			},
			Decrypt: CiphersErrors{
				Cipher:           ciphers.A51Decrypt,
				TextErrors:       Errors{
					Regex:      RegexTextA51Cipher,
					ErrorsList: CipherTextErrors.ErrorsList,
				},
				KeyErrors:  []Errors{{
					Regex: RegexKeyA51,
					ErrorsList: KeyErrors.ErrorsList}},
			}},
		},
	}
	page.WindowConf.InDeEnable = false
	p := new(NewPage)
	if err := (GenerateComposite(p, page)).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}
	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}
	return p, nil
}

func MagmaPage(parent walk.Container) (Page, error) {
	page := NewComposite{
		FontSize:           FontSizeTen,
		WindowConf:         WindowConfStandard,
		VariabilityText:    "Другое представление текста",
		Keys:               []Key{{Label: EnterKey, Visible: true, Enable: true}},
		VariabilityVisible: true,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher:     ciphers.MagmaEncryptStart,
				TextErrors: TextErrors,
				KeyErrors: []Errors{{
					Regex:      RegexKeyMagma,
					ErrorsList: KeyErrors.ErrorsList,
				}},
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.MagmaDecryptStart,
				TextErrors: Errors{
					Regex:      RegexTextMagmaCipher,
					ErrorsList: CipherTextErrors.ErrorsList,
				},
				KeyErrors: []Errors{{
					Regex:      RegexKeyMagma,
					ErrorsList: CipherKeyErrors.ErrorsList,
				}},
			},
		}, {
			Encrypt: CiphersErrors{
				Cipher: ciphers.MagmaEncrypt,
				TextErrors: Errors{
					Regex:      RegexTextMagmaCipher,
					ErrorsList: CipherTextErrors.ErrorsList,
				},
				KeyErrors: []Errors{{
					Regex:      RegexKeyMagma,
					ErrorsList: KeyErrors.ErrorsList,
				}},
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.MagmaDecrypt,
				TextErrors: Errors{
					Regex:      RegexTextMagmaCipher,
					ErrorsList: CipherTextErrors.ErrorsList,
				},
				KeyErrors: []Errors{{
					Regex:      RegexKeyMagma,
					ErrorsList: CipherKeyErrors.ErrorsList,
				}},
			},
		},
		},
	}
	page.WindowConf.InDeEnable = false
	p := new(NewPage)
	if err := (GenerateComposite(p, page)).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}
	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}
	return p, nil
}

func KuzPage(parent walk.Container) (Page, error) {
	page := NewComposite{
		FontSize:           FontSizeTen,
		WindowConf:         WindowConfStandard,
		VariabilityText:    "Другое представление текста",
		Keys:               []Key{{Label: EnterKey, Visible: true, Enable: true}},
		VariabilityVisible: true,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher:     ciphers.KuzEncryptStart,
				TextErrors: TextErrors,
				KeyErrors: []Errors{{
					Regex:      RegexKeyKuz,
					ErrorsList: KeyErrors.ErrorsList,
				}},
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.KuzDecryptStart,
				TextErrors: Errors{
					Regex:      RegexTextKuzCipher,
					ErrorsList: CipherTextErrors.ErrorsList,
				},
				KeyErrors: []Errors{{
					Regex:      RegexKeyKuz,
					ErrorsList: CipherKeyErrors.ErrorsList,
				}},
			},
		}, {
			Encrypt: CiphersErrors{
				Cipher: ciphers.KuzEncrypt,
				TextErrors: Errors{
					Regex:      RegexTextKuzCipher,
					ErrorsList: CipherTextErrors.ErrorsList,
				},
				KeyErrors: []Errors{{
					Regex:      RegexKeyKuz,
					ErrorsList: KeyErrors.ErrorsList,
				}},
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.KuzDecrypt,
				TextErrors: Errors{
					Regex:      RegexTextKuzCipher,
					ErrorsList: CipherTextErrors.ErrorsList,
				},
				KeyErrors: []Errors{{
					Regex:      RegexKeyKuz,
					ErrorsList: CipherKeyErrors.ErrorsList,
				}},
			},
		},
		},
	}
	page.WindowConf.InDeEnable = false
	p := new(NewPage)
	if err := (GenerateComposite(p, page)).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}
	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}
	return p, nil
}
