package node

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// Debug shows node list debug output
func Debug(n Node) {
	debugNode(n, 0)
}

func debugNode(n Node, level int) {
	for e := n; e != nil; e = e.Next() {
		fmt.Print(strings.Repeat(" | ", level))
		switch v := e.(type) {
		case *VList:
			color.Cyan("vlist (%d) wd: %s ht %s", v.ID, v.Width, v.Height)
			debugNode(v.List, level+1)
		case *HList:
			color.HiBlue("hlist (%d) wd: %s ht: %s", v.ID, v.Width, v.Height)
			debugNode(v.List, level+1)
		case *Disc:
			color.HiBlack("disc")
		case *Glyph:
			color.HiGreen("glyph: %s wd: %s cp: %d", v.Components, v.Width, v.Codepoint)
		case *Glue:
			color.HiMagenta("glue: %spt", v.Width)
		case *Image:
			color.Magenta("image: %s", v.Img.ImageFile.Filename)
		case *Lang:
			color.Magenta("lang: %s", v.Lang.Name)
		default:
			color.HiRed("Unhandled token %v", v)
		}
	}

}
