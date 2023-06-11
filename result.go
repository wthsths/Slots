package slots

type Result struct {
	Spins          []int
	WinLines       []ResultWinLines
	TotalWinAmount float64
}

type ResultWinLines struct {
	LineIndex          int
	WinAmount          float64
	SymbolMatchesCount int
}
