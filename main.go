package main

// import "fmt"

import (
	"flag"
	"fmt"
	"log"
	"os"

	// "time"
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
				// {
				// 	Label: astikit.StrPtr("About"),
				// 	OnClick: func(e astilectron.Event) (deleteListener bool) {
				// 		if err := bootstrap.SendMessage(w, "about", htmlAbout, func(m *bootstrap.MessageIn) {
				// 			// Unmarshal payload
				// 			var s string
				// 			if err := json.Unmarshal(m.Payload, &s); err != nil {
				// 				l.Println(fmt.Errorf("unmarshaling payload failed: %w", err))
				// 				return
				// 			}
				// 			l.Printf("About modal has been displayed and payload is %s!\n", s)
				// 		}); err != nil {
				// 			l.Println(fmt.Errorf("sending about event failed: %w", err))
				// 		}
				// 		return
				// 	},
				// },
				{Role: astilectron.MenuItemRoleClose},
			},
		}},
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			// w = ws[0]
			// go func() {
			// 	time.Sleep(5 * time.Second)
			// 	if err := bootstrap.SendMessage(w, "check.out.menu", "Don't forget to check out the menu!"); err != nil {
			// 		l.Println(fmt.Errorf("sending check.out.menu event failed: %w", err))
			// 	}
			// }()
			return nil
		},
		RestoreAssets: RestoreAssets,
		Windows: []*bootstrap.Window{{
			Homepage:       "index.html",
			MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				BackgroundColor: astikit.StrPtr("#E5E5E5"),
				Center:          astikit.BoolPtr(true),
				Height:          astikit.IntPtr(700),
				Width:           astikit.IntPtr(700),
			},
		}},
	}); err != nil {
		l.Fatal(fmt.Errorf("running boostrap failed: %w", err))
	}

}
