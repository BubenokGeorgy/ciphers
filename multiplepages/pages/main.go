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
		Keys:            []Key{{Label: EnterP, Visible: true, Enable: true}, {Label: EnterQ, Visible: true, Enable: true}},
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
		Keys:            []Key{{Label: EnterP, Visible: true, Enable: true}, {Label: EnterQ, Visible: true, Enable: true}, {Label: CurrentE, Visible: true, Enable: false, Const: true}},
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
		AutoKeys:        false,
		Keys:            []Key{{Label:CurrentGamma , Visible: true, Const: true}},
		VariabilityVisible: false,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher: ciphers.ShennonEncrypt,
				TextErrors: TextErrors,
				KeyErrorsHandler: ShennonKeyHandler,
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.ShennonDecrypt,
				TextErrors: CipherTextErrors,
				KeyErrorsHandler: ShennonKeyHandler,
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
		FontSize:        FontSizeTen,
		WindowConf:      WindowConfStandard,
		AutoKeys:        false,
		Keys:            []Key{{Label:CurrentGamma , Visible: true, Const: true}},
		VariabilityVisible: false,
		Ciphers: []EnDecrypt{{
			Encrypt: CiphersErrors{
				Cipher: ciphers.ShennonEncrypt,
				TextErrors: TextErrors,
				KeyErrorsHandler: ShennonKeyHandler,
			},
			Decrypt: CiphersErrors{
				Cipher: ciphers.ShennonDecrypt,
				TextErrors: CipherTextErrors,
				KeyErrorsHandler: ShennonKeyHandler,
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
