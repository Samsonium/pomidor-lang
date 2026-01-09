package lexer

func (l *Lexer) skipWhitespace() {
	for {
		ch := l.peek()

		if !isWhitespace(ch) {
			return
		}

		l.next()
	}
}

func (l *Lexer) skipComment() {
	l.next()
	l.next()

	for {
		ch := l.peek()

		if ch == '\n' || ch == 0 {
			return
		}

		l.next()
	}
}
