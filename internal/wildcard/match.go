package wildcard

// MatchSimple - 查找文本是否与模式字符串匹配/满足。
// 模式字符串支持'*'和'?'通配符。
// 与Match的唯一区别在于，`?`在末尾是可选的，这意味着`a?`模式将匹配名称`a`。
func MatchSimple(pattern, name string) bool {
	if pattern == "" {
		return name == pattern
	}
	if pattern == "*" {
		return true
	}
	// Do an extended wildcard '*' and '?' match.
	return deepMatchRune([]rune(name), []rune(pattern), true)
}

// Match - 查找文本是否与模式字符串匹配/满足。
// 模式字符串支持'*'和'?'通配符。
// 与path.Match()不同，它将路径视为匹配模式的扁平名称空间。
// 差异在此处说明：https://play.golang.org/p/Ega9qgD4Qz。
func Match(pattern, name string) (matched bool) {
	if pattern == "" {
		return name == pattern
	}
	if pattern == "*" {
		return true
	}
	// Do an extended wildcard '*' and '?' match.
	return deepMatchRune([]rune(name), []rune(pattern), false)
}

func deepMatchRune(str, pattern []rune, simple bool) bool {
	for len(pattern) > 0 {
		switch pattern[0] {
		default:
			if len(str) == 0 || str[0] != pattern[0] {
				return false
			}
		case '?':
			if len(str) == 0 {
				return simple
			}
		case '*':
			return deepMatchRune(str, pattern[1:], simple) ||
				(len(str) > 0 && deepMatchRune(str[1:], pattern, simple))
		}
		str = str[1:]
		pattern = pattern[1:]
	}
	return len(str) == 0 && len(pattern) == 0
}

// MatchAsPatternPrefix - 将文本作为给定模式的前缀进行匹配。示例：
//
//	| Pattern | Text    | Match Result |
//	====================================
//	| abc*    | ab      | True         |
//	| abc*    | abd     | False        |
//	| abc*c   | abcd    | True         |
//	| ab*??d  | abxxc   | True         |
//	| ab*??d  | abxc    | True         |
//	| ab??d   | abxc    | True         |
//	| ab??d   | abc     | True         |
//	| ab??d   | abcxdd  | False        |
//
// 此函数仅在某些特殊情况下有用。
func MatchAsPatternPrefix(pattern, text string) bool {
	return matchAsPatternPrefix([]rune(pattern), []rune(text))
}

func matchAsPatternPrefix(pattern, text []rune) bool {
	for i := 0; i < len(text) && i < len(pattern); i++ {
		if pattern[i] == '*' {
			return true
		}
		if pattern[i] == '?' {
			continue
		}
		if pattern[i] != text[i] {
			return false
		}
	}
	return true
}
