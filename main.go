package main

import (
	"os"
	"path/filepath"

	"gitee.com/rocket049/mycrypto"

	"github.com/rocket049/gettext-go/gettext"

	"github.com/therecipe/qt/gui"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

var T = gettext.T

func MainDialog() int {
	gettext.BindTextdomain("mycryptoqt", "locale.zip", localeData)
	gettext.Textdomain("mycryptoqt")

	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQDialog(nil, core.Qt__Window)
	window.SetWindowTitle(T("MyCrypto QT"))

	box := widgets.NewQVBoxLayout()

	font := gui.NewQFont2("", 18, 16, false)

	buttonEnc := widgets.NewQPushButton2(T("Encrypt File By AES256"), window)
	buttonEnc.SetFont(font)
	box.AddWidget(buttonEnc, 1, 0)
	buttonEnc.ConnectClicked(func(b bool) {
		EncryptDialog()
	})

	buttonDec := widgets.NewQPushButton2(T("Decrypt File By AES256"), window)
	buttonDec.SetFont(font)
	box.AddWidget(buttonDec, 1, 0)
	buttonDec.ConnectClicked(func(b bool) {
		DecryptDialog()
	})

	window.SetLayout(box)

	app.SetActiveWindow(window)

	window.Show()
	return app.Exec()
}

func DecryptDialog() {
	window := widgets.NewQDialog(nil, core.Qt__Window)
	window.SetWindowTitle(T("Decrypt File By AES256"))

	box := widgets.NewQGridLayout(window)

	var pwd, name string
	var err error

	msgLabel := widgets.NewQLabel2(T("Message Box"), window, core.Qt__Widget)
	box.AddWidget3(msgLabel, 0, 0, 1, 2, 0)

	label1 := widgets.NewQLabel2(T("Input:"), window, core.Qt__Widget)
	box.AddWidget(label1, 1, 0, 0)

	edit1 := widgets.NewQLineEdit(window)
	edit1.SetPlaceholderText(T("Double click to select file."))
	edit1.SetMinimumWidth(400)
	box.AddWidget(edit1, 1, 1, 0)

	edit1.ConnectMouseDoubleClickEvent(func(e *gui.QMouseEvent) {
		path1 := widgets.QFileDialog_GetOpenFileName(window, T("Select a file to encrypt"), "", "", "", widgets.QFileDialog__ReadOnly)
		if len(path1) == 0 {
			return
		}

		for {
			var b bool
			pwd = widgets.QInputDialog_GetText(window, T("File Password"), T("Password:"), widgets.QLineEdit__Password, "", &b, core.Qt__Dialog, core.Qt__ImhHiddenText)
			if b == false {
				msgLabel.SetText(T("Message Box"))
				edit1.SetText("")
				pwd = ""
				name = ""
				break
			}

			name, err = mycrypto.ReadNameFromFile(path1, pwd)
			if err == nil {
				msgLabel.SetText(T("FileName:") + name)
				edit1.SetText(path1)
				break
			} else {
				widgets.QMessageBox_About(window, T("Error Password!"), T("Error Password!"))
			}
		}

	})

	label2 := widgets.NewQLabel2(T("Output:"), window, core.Qt__Widget)
	box.AddWidget(label2, 2, 0, 0)

	edit2 := widgets.NewQLineEdit(window)
	edit2.SetPlaceholderText(T("Double click to select directory."))
	edit2.SetMinimumWidth(400)
	box.AddWidget(edit2, 2, 1, 0)

	edit2.ConnectMouseDoubleClickEvent(func(e *gui.QMouseEvent) {
		path1 := widgets.QFileDialog_GetExistingDirectory(window, T("Select a file to encrypt"), "", widgets.QFileDialog__ShowDirsOnly)

		edit2.SetText(path1)
	})

	button1 := widgets.NewQPushButton2(T("Decrypt"), window)
	box.AddWidget3(button1, 3, 0, 1, 2, 0)

	button1.ConnectClicked(func(b bool) {
		src := edit1.Text()
		dst := edit2.Text()
		if pwd == "" || name == "" || src == "" || dst == "" {
			return
		}

		dst = filepath.Join(dst, name)
		err := mycrypto.DecryptoFromFile(src, dst, pwd)
		if err != nil {
			widgets.QMessageBox_About(window, T("Error"), err.Error())
			window.Close()
		} else {
			widgets.QMessageBox_About(window, T("Successful!"), T("Successful!"))
			window.Close()
		}
	})

	window.SetLayout(box)
	window.Show()
}

func EncryptDialog() {
	window := widgets.NewQDialog(nil, core.Qt__Window)
	window.SetWindowTitle(T("Encrypt File By AES256"))

	box := widgets.NewQGridLayout(window)

	labelPwd := widgets.NewQLabel2(T("Password:"), window, core.Qt__Widget)
	box.AddWidget(labelPwd, 0, 0, 0)

	editPwd := widgets.NewQLineEdit(window)
	editPwd.SetEchoMode(widgets.QLineEdit__Password)
	box.AddWidget(editPwd, 0, 1, 0)

	labelCfm := widgets.NewQLabel2(T("Confirm:"), window, core.Qt__Widget)
	box.AddWidget(labelCfm, 1, 0, 0)

	editCfm := widgets.NewQLineEdit(window)
	editCfm.SetEchoMode(widgets.QLineEdit__Password)
	box.AddWidget(editCfm, 1, 1, 0)

	label1 := widgets.NewQLabel2(T("Input:"), window, core.Qt__Widget)
	box.AddWidget(label1, 2, 0, 0)

	edit1 := widgets.NewQLineEdit(window)
	edit1.SetPlaceholderText(T("Double click to select file."))
	edit1.SetMinimumWidth(400)
	box.AddWidget(edit1, 2, 1, 0)

	edit1.ConnectMouseDoubleClickEvent(func(e *gui.QMouseEvent) {
		path1 := widgets.QFileDialog_GetOpenFileName(window, T("Select a file to encrypt"), "", "", "", widgets.QFileDialog__ReadOnly)

		edit1.SetText(path1)

	})

	label2 := widgets.NewQLabel2(T("Output:"), window, core.Qt__Widget)
	box.AddWidget(label2, 3, 0, 0)

	edit2 := widgets.NewQLineEdit(window)
	edit2.SetPlaceholderText(T("Double click to select directory."))
	edit2.SetMinimumWidth(400)
	box.AddWidget(edit2, 3, 1, 0)

	edit2.ConnectMouseDoubleClickEvent(func(e *gui.QMouseEvent) {
		path1 := widgets.QFileDialog_GetExistingDirectory(window, T("Select a file to encrypt"), "", widgets.QFileDialog__ShowDirsOnly)

		edit2.SetText(path1)
	})

	button1 := widgets.NewQPushButton2(T("Encrypt"), window)
	box.AddWidget3(button1, 4, 0, 1, 2, 0)

	button1.ConnectClicked(func(b bool) {
		src := edit1.Text()
		dst := filepath.Join(edit2.Text(), filepath.Base(src)+".e")
		pwd := editPwd.Text()
		cfm := editCfm.Text()

		if pwd == "" || cfm == "" || src == "" || dst == edit2.Text() {
			return
		}

		if pwd != cfm {
			widgets.QMessageBox_About(window, T("Passwords are inconsistent"), T("Passwords are inconsistent"))
			return
		}

		err := mycrypto.EncryptoToFile(src, dst, pwd)
		if err != nil {
			widgets.QMessageBox_About(window, T("Error"), err.Error())
			window.Close()
		} else {
			widgets.QMessageBox_About(window, T("Successful!"), T("Successful!"))
			window.Close()
		}
	})

	window.SetLayout(box)
	window.Show()
}

func main() {
	os.Exit(MainDialog())
}
