// Copyright 2017 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	. "main/pages"
	. "main/ui"
)

func main() {
	walk.Resources.SetRootDirPath("img")

	mw := new(AppMainWindow)
	cfg := &MultiPageMainWindowConfig{
		Name:    "mainWindow",
		MinSize: Size{700, 640},
		OnCurrentPageChanged: func() {
			mw.UpdateTitle(mw.CurrentPageTitle())
		},
		PageCfgs: []PageConfig{
			{"Шифр АТБАШ", "key.png", NewAtbashPage},
			{"Шифр Цезаря", "key.png", NewCaesarPage},
			{"Квадрат Полибия", "key.png", NewPolybiusPage},
			{"Шифр Тритемия", "key.png", NewTrithemiusPage},
			{"Шифр Белазо", "key.png", NewBelazoPage},
			{"Шифр Виженера", "key.png", NewVigenerePage},
			{"Матричный шифр", "key.png", NewMatrixPage},
			{"Шифр Плейфера", "key.png", NewPlayfairPage},
			{"Вертикальная перестановка", "key.png", NewVerticalPermutationPage},
			{"Решётка Кардано", "key.png", NewCardanGrillePage},
			{"Диффи-Хеллман", "key.png", NewDiffieHellmanPage},
		},
	}

	mpmw, err := NewMultiPageMainWindow(cfg)
	if err != nil {
		fmt.Println(err)
	}

	mw.MultiPageMainWindow = mpmw

	mw.UpdateTitle(mw.CurrentPageTitle())
	//mw.SetFullscreen(true)
	mw.Run()
}



