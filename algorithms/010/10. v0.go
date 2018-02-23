package algorithms
//
//type myRegularMatch struct {
//	posStr, posPattern, lenStr, lenPattern int
//	str, pattern                           string
//	lastPattern                            byte
//}
//
//// 一开始思路想着parser那一套用状态机实现 然后卡在实现非递归回溯上了 最后还是用了递归回溯
//func isMatchV0(s string, p string) bool {
//	return (&myRegularMatch{
//		lenStr:     len(s),
//		lenPattern: len(p),
//		str:        s,
//		pattern:    p,
//		posStr:     -1,
//		posPattern: -1,
//	}).parseStart()
//}
//
//func (m myRegularMatch) charStr() byte {
//	if m.posStr == -1 {
//		return 0
//	}
//	return m.str[m.posStr]
//}
//
//func (m myRegularMatch) charPattern() byte {
//	if m.posPattern == -1 {
//		return 0
//	}
//	return m.pattern[m.posPattern]
//}
//
//func (m myRegularMatch) peekStr() byte {
//	return (&m).nextStr()
//}
//
//func (m myRegularMatch) peekPattern() byte {
//	return (&m).nextPattern()
//}
//
//func (m *myRegularMatch) nextStr() byte {
//	m.posStr++
//	if m.posStr >= m.lenStr {
//		return 0
//	}
//	return m.charStr()
//}
//
//func (m *myRegularMatch) nextPattern() byte {
//	m.lastPattern = m.charPattern()
//	m.posPattern++
//	if m.posPattern == m.lenPattern {
//		return 0
//	}
//	return m.charPattern()
//}
//
//func (m *myRegularMatch) parseStart() bool {
//	if m.posPattern == m.lenPattern-1 {
//		return m.posStr == m.lenStr-1
//	} else if m.posStr == m.lenStr-1 {
//		if m.lenPattern > 1 && m.pattern[1] == '*' {
//			return isMatchV0(m.str, m.pattern[2:])
//		}
//		return false
//	}
//	switch m.peekPattern() {
//	case '*':
//		return m.parseStar()
//	default:
//		return m.parseDefault()
//	}
//}
//
//func (m *myRegularMatch) parseStar() bool {
//	m.nextPattern()
//	if isMatchV0(m.str[m.posStr+1:], m.pattern[m.posPattern+1:]) {
//		return true
//	} else if m.peekStr() == m.lastPattern || m.lastPattern == '.' {
//		m.nextStr()
//		// 回溯
//		return isMatchV0(m.str[m.posStr+1:], m.pattern[m.posPattern-1:])
//	}
//	return false
//}
//
//func (m *myRegularMatch) parseDefault() bool {
//	p := m.nextPattern()
//	if m.peekPattern() == '*' {
//		return m.parseStar()
//	}
//	if m.peekStr() == p || p == '.' {
//		m.nextStr()
//		return isMatchV0(m.str[m.posStr+1:], m.pattern[m.posPattern+1:])
//	}
//	return false
//}
