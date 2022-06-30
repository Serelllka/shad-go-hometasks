package main

func main() {
	GetCourseList(map[string][]string{
		"lol":    {"kek"},
		"kek":    {"lol"},
		"amogus": {"abobus"},
	})
}

func GetCourseList(prereqs map[string][]string) (ans []string) {
	deps := make(map[string]map[string]interface{})

	for i, content := range prereqs {
		deps[i] = make(map[string]interface{})
		for _, item := range content {
			deps[i][item] = struct{}{}
			_, ok := deps[item]
			if !ok {
				deps[item] = make(map[string]interface{})
			}
		}
	}

	for {
		f := true
		for i, item := range deps {
			if len(item) == 0 {
				ans = append(ans, i)
				delete(deps, i)
				f = false
				continue
			}
			for _, rem := range ans {
				_, ok := item[rem]
				if ok {
					f = false
				}
				delete(item, rem)
			}
		}
		if len(deps) == 0 {
			return
		}
		if f {
			panic("it's a panic!")
		}
	}
}
