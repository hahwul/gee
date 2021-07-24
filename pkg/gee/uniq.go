package gee

func Uniq(line string, dtable []string) (bool, []string) {
	for _, v := range dtable {
		if v == line {
			return true, dtable
		}
	}
	tmp := append(dtable, line)
	return false, tmp
}
