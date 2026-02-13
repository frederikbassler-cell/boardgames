package board

// RowContains prüft, ob die angegebene Zeile den gesuchten Wert enthält.
func (b *Board) RowContains(i int, value string) bool {
	if i < 0 || i >= len(b.rows) {
		return false
	}
	for _, cell := range b.rows[i] {
		if cell == value {
			return true
		}
	}

	return false
}

// RowContainsChain prüft, ob die angegebene Zeile eine ununterbrochene Kette des gesuchten Werts enthält.
func (b *Board) RowContainsChain(rowIndex int, value string, length int) bool {
	if rowIndex < 0 || rowIndex >= len(b.rows) {
		return false
	}
	count := 0
	for _, cell := range b.rows[rowIndex] {
		if cell == value {
			count++
			if count >= length {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}

// RowContainsOnly prüft, ob die angegebene Zeile ausschließlich den gesuchten Wert enthält.
func (b *Board) RowContainsOnly(rowIndex int, value string) bool {
	if rowIndex < 0 || rowIndex >= len(b.rows) {
		return false
	}
	count := 0
	for _, cell := range b.rows[rowIndex] {
		if cell == value {
			count++
			if count >= len(b.rows) {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}

// ColContains prüft, ob die angegebene Spalte den gesuchten Wert enthält.
func (b *Board) ColContains(colIndex int, value string) bool {
	if colIndex < 0 || len(b.rows) == 0 || colIndex >= len(b.rows[0]) {
		return false
	}
	for _, row := range b.rows {
		if row[colIndex] == value {
			return true
		}
	}

	return false
}

// ColContainsChain prüft, ob die angegebene Spalte eine ununterbrochene Kette des gesuchten Werts enthält.
func (b *Board) ColContainsChain(colIndex int, value string, length int) bool {
	if colIndex < 0 || len(b.rows) == 0 || colIndex >= len(b.rows[0]) {
		return false
	}
	count := 0
	for _, row := range b.rows {
		if row[colIndex] == value {
			count++
			if count >= length {
				return true
			}
		}

		
	}
	return false
}

// ColContainsOnly prüft, ob die angegebene Spalte ausschließlich den gesuchten Wert enthält.
func (b *Board) ColContainsOnly(colIndex int, value string) bool {
	

	
	return false
}

// DiagDownRightContains prüft, ob die angegebene Diagonale von oben links nach unten rechts den gesuchten Wert enthält.
func (b *Board) DiagDownRightContains(startCol int, value string) bool {
	// TODO
	return false
}

// DiagDownRightContainsChain prüft, ob die angegebene Diagonale von oben links nach unten rechts eine ununterbrochene Kette des gesuchten Werts enthält.
func (b *Board) DiagDownRightContainsChain(startCol int, value string, length int) bool {
	// TODO
	return false
}

// DiagDownRightContainsOnly prüft, ob die angegebene Diagonale von oben links nach unten rechts ausschließlich den gesuchten Wert enthält.
func (b *Board) DiagDownRightContainsOnly(startCol int, value string) bool {
	// TODO
	return false
}

// DiagDownLeftContains prüft, ob die angegebene Diagonale von oben rechts nach unten links den gesuchten Wert enthält.
func (b *Board) DiagDownLeftContains(startCol int, value string) bool {
	// TODO
	return false
}

// DiagDownLeftContainsChain prüft, ob die angegebene Diagonale von oben rechts nach unten links eine ununterbrochene Kette des gesuchten Werts enthält.
func (b *Board) DiagDownLeftContainsChain(startCol int, value string, length int) bool {
	// TODO
	return false
}

// DiagDownLeftContainsOnly prüft, ob die angegebene Diagonale von oben rechts nach unten links ausschließlich den gesuchten Wert enthält.
func (b *Board) DiagDownLeftContainsOnly(startCol int, value string) bool {
	// TODO
	return false
}

// BoardContains prüft, ob das Spielfeld den gesuchten Wert enthält.
func (b *Board) BoardContains(value string) bool {
	// TODO
	return false
}

// BoardContainsChain prüft, ob das Spielfeld eine ununterbrochene Kette des gesuchten Werts enthält.
// Dabei werden alle Zeilen, Spalten und Diagonalen überprüft.
func (b *Board) BoardContainsChain(value string, length int) bool {
	// TODO
	return false
}

// BoardContainsOnly prüft, ob das Spielfeld ausschließlich den gesuchten Wert enthält.
func (b *Board) BoardContainsOnly(value string) bool {
	// TODO
	return true
}
