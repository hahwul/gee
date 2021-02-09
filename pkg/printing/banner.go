package printing

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

// Banner is banner of gee
func Banner() {
	fmt.Fprintln(os.Stderr, aurora.White("._____  ._______._______").String())
	fmt.Fprintln(os.Stderr, aurora.White(":_ ___\\ : .____/: .____/").String())
	fmt.Fprintln(os.Stderr, aurora.White("|   |___| : _/\\ | : _/\\ ").String())
	fmt.Fprintln(os.Stderr, aurora.White("|   /  ||   /  \\|   /  \\").String())
	fmt.Fprintln(os.Stderr, "|. __  ||_.: __/|_.: __/")
	fmt.Fprintln(os.Stderr, ":/ |. |   :/      :/")
	fmt.Fprintln(os.Stderr, ":  : /   "+aurora.BrightYellow("{"+VERSION+"}").String())
	fmt.Fprintln(os.Stderr, "   :   ")
}
