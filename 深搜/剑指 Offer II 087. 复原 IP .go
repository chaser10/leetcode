func restoreIpAddresses(s string) []string {
	ret := []string{}
	var dfs func(start int, splits string, cnt int)
	dfs = func(start int, splits string, cnt int) {
        if cnt > 4 {
            return
        }
		if start == len(s) && cnt == 4{
			ret = append(ret, splits[1:])
			return
		}
        fmt.Println(splits[:])
        if start == len(s) {
            return
        }

		if s[start] == '0' {
			dfs(start + 1, splits + ".0", cnt + 1)
		} else {
			for j := 0; j < 3; j++ {
                if start + j + 1 <= len(s) {
					cur := strconv.Atoi(s[start: start + j + 1])
					if cur > 255 {
						continue
					}
                    dfs(start + j + 1, splits + "." + s[start: start + j + 1], cnt + 1)
                }
			}
		}
	}

	dfs(0, "", 0)
	return ret
}