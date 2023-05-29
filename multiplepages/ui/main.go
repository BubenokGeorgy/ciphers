package ui

import (
	"bytes"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	. "main/const"
	"main/errors"
	. "main/tools"
	"reflect"
	"unicode/utf8"
)

type MultiPageMainWindowConfig struct {
	Name                 string
	Enabled              Property
	Visible              Property
	Font                 Font
	MinSize              Size
	MaxSize              Size
	ContextMenuItems     []MenuItem
	OnKeyDown            walk.KeyEventHandler
	OnKeyPress           walk.KeyEventHandler
	OnKeyUp              walk.KeyEventHandler
	OnMouseDown          walk.MouseEventHandler
	OnMouseMove          walk.MouseEventHandler
	OnMouseUp            walk.MouseEventHandler
	OnSizeChanged        walk.EventHandler
	OnCurrentPageChanged walk.EventHandler
	Title                string
	Size                 Size
	MenuItems            []MenuItem
	ToolBar              ToolBar
	PageCfgs             []PageConfig
}

type PageConfig struct {
	Title   string
	Image   string
	NewPage PageFactoryFunc
}

type PageFactoryFunc func(parent walk.Container) (Page, error)

type Page interface {
	// Provided by Walk
	walk.Container
	Parent() walk.Container
	SetParent(parent walk.Container) error
}

type MultiPageMainWindow struct {
	*walk.MainWindow
	navTB                       *walk.ToolBar
	pageCom                     *walk.Composite
	action2NewPage              map[*walk.Action]PageFactoryFunc
	pageActions                 []*walk.Action
	currentAction               *walk.Action
	currentPage                 Page
	currentPageChangedPublisher walk.EventPublisher
}

func NewMultiPageMainWindow(cfg *MultiPageMainWindowConfig) (*MultiPageMainWindow, error) {
	mpmw := &MultiPageMainWindow{
		action2NewPage: make(map[*walk.Action]PageFactoryFunc),
	}

	if err := (MainWindow{
		AssignTo:         &mpmw.MainWindow,
		Name:             cfg.Name,
		Title:            cfg.Title,
		Enabled:          cfg.Enabled,
		Visible:          cfg.Visible,
		Font:             cfg.Font,
		MinSize:          cfg.MinSize,
		MaxSize:          cfg.MaxSize,
		MenuItems:        cfg.MenuItems,
		ToolBar:          cfg.ToolBar,
		ContextMenuItems: cfg.ContextMenuItems,
		OnKeyDown:        cfg.OnKeyDown,
		OnKeyPress:       cfg.OnKeyPress,
		OnKeyUp:          cfg.OnKeyUp,
		OnMouseDown:      cfg.OnMouseDown,
		OnMouseMove:      cfg.OnMouseMove,
		OnMouseUp:        cfg.OnMouseUp,
		OnSizeChanged:    cfg.OnSizeChanged,
		Layout:           HBox{MarginsZero: true, SpacingZero: true},
		Children: []Widget{
			ScrollView{
				HorizontalFixed: true,
				Layout:          VBox{MarginsZero: true},
				Children: []Widget{
					Composite{
						Layout: VBox{MarginsZero: true},
						Children: []Widget{
							ToolBar{
								AssignTo:    &mpmw.navTB,
								Orientation: Vertical,
								ButtonStyle: ToolBarButtonImageAboveText,
								MaxTextRows: 2,
							},
						},
					},
				},
			},
			Composite{
				AssignTo: &mpmw.pageCom,
				Name:     "pageCom",
				Layout:   HBox{MarginsZero: true, SpacingZero: true},
			},
		},
	}).Create(); err != nil {
		return nil, err
	}

	succeeded := false
	defer func() {
		if !succeeded {
			mpmw.Dispose()
		}
	}()

	for _, pc := range cfg.PageCfgs {
		action, err := mpmw.newPageAction(pc.Title, pc.Image, pc.NewPage)
		if err != nil {
			return nil, err
		}

		mpmw.pageActions = append(mpmw.pageActions, action)
	}

	if err := mpmw.updateNavigationToolBar(); err != nil {
		return nil, err
	}

	if len(mpmw.pageActions) > 0 {
		if err := mpmw.setCurrentAction(mpmw.pageActions[0]); err != nil {
			return nil, err
		}
	}

	if cfg.OnCurrentPageChanged != nil {
		mpmw.CurrentPageChanged().Attach(cfg.OnCurrentPageChanged)
	}

	succeeded = true

	return mpmw, nil
}

func (mpmw *MultiPageMainWindow) CurrentPage() Page {
	return mpmw.currentPage
}

func (mpmw *MultiPageMainWindow) CurrentPageTitle() string {
	if mpmw.currentAction == nil {
		return ""
	}

	return mpmw.currentAction.Text()
}

func (mpmw *MultiPageMainWindow) CurrentPageChanged() *walk.Event {
	return mpmw.currentPageChangedPublisher.Event()
}

func (mpmw *MultiPageMainWindow) newPageAction(title, image string, newPage PageFactoryFunc) (*walk.Action, error) {
	img, err := walk.Resources.Bitmap(image)
	if err != nil {
		return nil, err
	}

	action := walk.NewAction()
	action.SetCheckable(true)
	action.SetExclusive(true)
	action.SetImage(img)
	action.SetText(title)

	mpmw.action2NewPage[action] = newPage

	action.Triggered().Attach(func() {
		mpmw.setCurrentAction(action)
	})

	return action, nil
}

func (mpmw *MultiPageMainWindow) setCurrentAction(action *walk.Action) error {
	defer func() {
		if !mpmw.pageCom.IsDisposed() {
			mpmw.pageCom.RestoreState()
		}
	}()

	mpmw.SetFocus()

	if prevPage := mpmw.currentPage; prevPage != nil {
		mpmw.pageCom.SaveState()
		prevPage.SetVisible(false)
		prevPage.(walk.Widget).SetParent(nil)
		prevPage.Dispose()
	}

	newPage := mpmw.action2NewPage[action]

	page, err := newPage(mpmw.pageCom)
	if err != nil {
		return err
	}

	action.SetChecked(true)

	mpmw.currentPage = page
	mpmw.currentAction = action

	mpmw.currentPageChangedPublisher.Publish()

	return nil
}

func (mpmw *MultiPageMainWindow) updateNavigationToolBar() error {
	mpmw.navTB.SetSuspended(true)
	defer mpmw.navTB.SetSuspended(false)

	actions := mpmw.navTB.Actions()

	if err := actions.Clear(); err != nil {
		return err
	}

	for _, action := range mpmw.pageActions {
		if err := actions.Add(action); err != nil {
			return err
		}
	}

	if mpmw.currentAction != nil {
		if !actions.Contains(mpmw.currentAction) {
			for _, action := range mpmw.pageActions {
				if action != mpmw.currentAction {
					if err := mpmw.setCurrentAction(action); err != nil {
						return err
					}

					break
				}
			}
		}
	}

	return nil
}

type AppMainWindow struct {
	*MultiPageMainWindow
}

func (mw *AppMainWindow) UpdateTitle(prefix string) {
	var buf bytes.Buffer

	if prefix != "" {
		buf.WriteString(prefix)
		buf.WriteString(" - ")
	}

	buf.WriteString("Задания по криптографии")

	mw.SetTitle(buf.String())
}

type NewPage struct {
	*walk.Composite
}

func SetTextAfterDecrypt(inDe, outDe *walk.TextEdit, keys []*walk.TextEdit, deConvertText func (string) string, decrypt CiphersErrors) {
	textDe, errorDe := errors.CheckErrors(inDe.Text(), keys, decrypt)
	if utf8.RuneCountInString(errorDe) != 0 {
		outDe.SetText(errorDe)
	} else {
		outDe.SetText(deConvertText(textDe))
	}
}
func SetTextAfterEncrypt(inEn, inDe, outEn *walk.TextEdit, keys []*walk.TextEdit, convertText func (string) string, encrypt CiphersErrors) bool {
	for _, key := range keys{
		if !key.Enabled(){
			key.SetText("")
		}
	}
	textEn, errorEn := errors.CheckErrors(convertText(inEn.Text()), keys, encrypt)
	if utf8.RuneCountInString(errorEn) != 0 {
		outEn.SetText(errorEn)
		return false
	} else {
		outEn.SetText(textEn)
		inDe.SetText(textEn)
	}
	return true
}

func GenerateComposite(p *NewPage, newComposite NewComposite) Composite {
	var inEn, outEn, inDe, outDe *walk.TextEdit
	var convertText = CleanConvertText
	var deConvertText = CleanDeConvertText
	var encrypt = newComposite.Ciphers[0].Encrypt
	var decrypt = newComposite.Ciphers[0].Decrypt
	if len(newComposite.Keys)==0{
		newComposite.Keys = []Key{{Label: "Задайте ключ.", Visible: true}}
	}
	if len(newComposite.Keys)<4{
		for i:= len(newComposite.Keys);i<4;i++{
			newComposite.Keys = append(newComposite.Keys,Key{Visible: false})
		}
	}
	keys :=make([]*walk.TextEdit, len(newComposite.Keys))
	return Composite{
		AssignTo: &p.Composite,
		Font: Font{
			Bold:      true,
			PointSize: newComposite.FontSize,
		},
		Layout: VBox{
			Spacing: 20},
		Children: []Widget{
			HSplitter{
				StretchFactor: newComposite.WindowConf.InStretchFactor,
				Children: []Widget{
					HSplitter{Children: []Widget{
						VSplitter{
							Children: []Widget{
								HSplitter{
									Visible: newComposite.WindowConf.HeaderFirstVisible,
									Children: []Widget{
										Label{
											Text: newComposite.WindowConf.HeaderFirstText},
									}},
								HSplitter{
									Visible: newComposite.WindowConf.InEnVisible,
									Children: []Widget{
										TextEdit{
											Enabled: newComposite.WindowConf.InEnEnable,
											VScroll:  true,
											AssignTo: &inEn},
									}},
								HSplitter{
									Visible: newComposite.WindowConf.WriteTestTextVisible,
									Children: []Widget{
										PushButton{
											MinSize: Size{
												Height: 60,
											},
											Text: "Вставить тестовый текст",
											OnClicked: func() {
												SetTestText(inEn, convertText)
											},
										},
									}},
							}},
						VSplitter{
							Children: []Widget{
								HSplitter{
									Visible: newComposite.WindowConf.HeaderSecondVisible,
									Children: []Widget{
										Label{
											Text: newComposite.WindowConf.HeaderSecondText},
									}},
								HSplitter{
									Children: []Widget{
										TextEdit{
											Enabled: newComposite.WindowConf.OutEnEnable,
											VScroll:  true,
											AssignTo: &outEn},
									}},
								HSplitter{
									Visible: newComposite.WindowConf.EncryptButtonVisible,
									Children: []Widget{
										PushButton{
											MinSize: Size{
												Height: 60,
											},
											Text: "Зашифровать.",
											OnClicked: func() {
												if SetTextAfterEncrypt(inEn, inDe, outEn, keys, convertText, encrypt){
													SetTextAfterDecrypt(inDe, outDe, keys,deConvertText,decrypt)
												}
											},
										},
									}},
							}},
					}},
				},
			},
			HSplitter{
				Visible: newComposite.WindowConf.LongButtonVisible,
				Children: []Widget{
					PushButton{
						MinSize: Size{
							Height: 60,
						},
						Text: newComposite.WindowConf.LongButtonText,
						OnClicked: func() {
							SetTextAfterEncrypt(inEn, inDe, outEn, keys, convertText, encrypt)
						},
					},
				}},
			HSplitter{
				StretchFactor: 1,
				Children: []Widget{
					HSplitter{Children: []Widget{
						VSplitter{Children: []Widget{
							HSplitter{
								Visible: newComposite.WindowConf.InDeVisible,
								StretchFactor: 10,
								Children: []Widget{
									TextEdit{
										VScroll:  true,
										AssignTo: &inDe},
								}},
							HSplitter{
								MinSize: Size{
									Height: 60,
								},
								StretchFactor: 1,
								Children: []Widget{
									VSplitter{
										MinSize: Size{Height: 60},
										MaxSize: Size{Height: 60},
										Children: []Widget{
											CheckBox{
												Visible: newComposite.WindowConf.DirtyCheckVisible,
												Text: "Черновая проверка",
												OnCheckedChanged: func() {
													if reflect.ValueOf(convertText) == reflect.ValueOf(CleanConvertText) {
														convertText = DirtyConvertText
														deConvertText = DirtyDeConvertText
													} else {
														convertText = CleanConvertText
														deConvertText = CleanDeConvertText
													}
												},
											},
											CheckBox{
												Text:    newComposite.VariabilityText,
												Visible: newComposite.VariabilityVisible,
												OnCheckedChanged: func() {
													if reflect.ValueOf(encrypt.Cipher) == reflect.ValueOf(newComposite.Ciphers[0].Encrypt.Cipher) {
														encrypt = newComposite.Ciphers[1].Encrypt
														decrypt = newComposite.Ciphers[1].Decrypt
													} else {
														encrypt = newComposite.Ciphers[0].Encrypt
														decrypt = newComposite.Ciphers[0].Decrypt
													}
												},
											},
											CheckBox{
												Text:    newComposite.VariabilityText,
												Visible: newComposite.AutoKeys,
												OnCheckedChanged: func() {
													for _, key := range keys{
														if key.Enabled(){
															key.SetEnabled(false)
															encrypt = newComposite.Ciphers[1].Encrypt
															decrypt = newComposite.Ciphers[1].Decrypt
														} else {
															encrypt = newComposite.Ciphers[0].Encrypt
															decrypt = newComposite.Ciphers[0].Decrypt
															key.SetEnabled(true)
														}
													}
												},
											},
										}},
									VSplitter{
										Visible: newComposite.Keys[0].Visible,
										Children: []Widget{
											Label{
												Text:newComposite.Keys[0].Label,
											},
											TextEdit{
												VScroll: true,
												AssignTo: &keys[0],
											},
										}},
									VSplitter{
										Visible: newComposite.Keys[1].Visible,
										Children: []Widget{
											Label{
												Text:newComposite.Keys[1].Label,
											},
											TextEdit{
												VScroll: true,
												AssignTo: &keys[1],
											},
										}},
									VSplitter{
										Visible: newComposite.Keys[2].Visible,
										Children: []Widget{
											Label{
												Text:newComposite.Keys[2].Label,
											},
											TextEdit{
												VScroll: true,
												AssignTo: &keys[2],
											},
										}},
									VSplitter{
										Visible: newComposite.Keys[3].Visible,
										Children: []Widget{
											Label{
												Text:newComposite.Keys[3].Label,
											},
											TextEdit{
												VScroll: true,
												AssignTo: &keys[3],
											},
										}},
								}},
						}},
						VSplitter{
							Children: []Widget{
							HSplitter{
								Visible: newComposite.WindowConf.OutDeVisible,
								Children: []Widget{
									TextEdit{
										VScroll:  true,
										AssignTo: &outDe},
								}},
							HSplitter{
								Visible: newComposite.WindowConf.DecryptButtonVisible,
								Children: []Widget{
									PushButton{
										MinSize: Size{
											Height: 60,
										},
										Text: "Расшифровать",
										OnClicked: func() {
											SetTextAfterDecrypt(inDe, outDe,keys,deConvertText,decrypt)
										},
									},
								}},
						}},
					}},
				},
			},
		},
	}
}
