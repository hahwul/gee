package printing

import (
	"fmt"
	"os"
)

// Banner is banner of gee
func Banner() {
	fmt.Fprintln(os.Stderr, "._____  ._______._______")
	fmt.Fprintln(os.Stderr, ":_ ___\\ : .____/: .____/")
	fmt.Fprintln(os.Stderr, "|   |___| : _/\\ | : _/\\ ")
	fmt.Fprintln(os.Stderr, "|   /  ||   /  \\|   /  \\")
	fmt.Fprintln(os.Stderr, "|. __  ||_.: __/|_.: __/")
	fmt.Fprintln(os.Stderr, ":/ |. |   :/      :/")
	fmt.Fprintln(os.Stderr, ":  : /   "+VERSION)
	fmt.Fprintln(os.Stderr, "   :   ")
}
