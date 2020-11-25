package windows

import (
	vlc "github.com/adrg/libvlc-go/v3"
	"github.com/gotk3/gotk3/gdk"
)

func SetPlayerWindow(player *vlc.Player, window *gdk.Window) error {
	return player.SetXWindow(window.GetXID())
}