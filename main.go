package main

// import (
// 	"fmt"

// 	"github.com/Fairbrook/analizador/Assembler"
// )

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Fairbrook/analizador/Assembler"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

var (
	AppName            string
	BuiltAt            string
	VersionAstilectron string
	VersionElectron    string
)

var (
	fs    = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	debug = fs.Bool("d", false, "enables the debug mode")
	w     *astilectron.Window
)

func main() {
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	fs.Parse(os.Args[1:])

	l.Printf("Running app built at %s\n", BuiltAt)
	if err := bootstrap.Run(bootstrap.Options{
		Asset:    Asset,
		AssetDir: AssetDir,
		AstilectronOptions: astilectron.Options{
			AppName: AppName,
			// BaseDirectoryPath: wd,
			// AppIconDarwinPath:  "resources/icon.icns",
			// AppIconDefaultPath: "resources/icon.png",
			SingleInstance:     true,
			VersionAstilectron: VersionAstilectron,
			VersionElectron:    VersionElectron,
		},
		Debug:  *debug,
		Logger: l,
		MenuOptions: []*astilectron.MenuItemOptions{{
			Label: astikit.StrPtr("File"),
			SubMenu: []*astilectron.MenuItemOptions{
				{Role: astilectron.MenuItemRoleClose},
				{
					Label: astikit.StrPtr("Traducir a ensamblador"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						w.SendMessage("translate")
						return
					},
				},
			},
		}, {
			Label: astikit.StrPtr("Compilacion"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label: astikit.StrPtr("Compilar"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						w.SendMessage("compile")
						return
					},
				},
				{
					Label: astikit.StrPtr("Compilar y ejecutar"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						w.SendMessage("run")
						return
					},
				},
			},
		}},
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = ws[0]
			// w.OpenDevTools()
			return nil
		},
		RestoreAssets: RestoreAssets,
		Windows: []*bootstrap.Window{{
			Homepage:       "index.html",
			MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				BackgroundColor: astikit.StrPtr("#E5E5E5"),
				Center:          astikit.BoolPtr(true),
				Height:          astikit.IntPtr(800),
				Width:           astikit.IntPtr(1500),
			},
		}},
	}); err != nil {
		l.Fatal(fmt.Errorf("running boostrap failed: %w", err))
	}
}

func main2() {
	assembler := Assembler.Translator{
		Filename: "trans.asm",
	}
	result, err := assembler.TranslateAndOpen("int hola;float suma(float a, float b){return a +b;}\nint main()\n{float a; a=suma(10.0,5.8);int b;while(a>0.0){printF(a);printS(\"\\n\");a=a-1.0;}printI(hola);return 0;}")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}
