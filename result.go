package slots

type Result struct {
	Bet            float64
	Lines          int
	LastSpin       []int
	Spins          []int
	WinLines       []ResultWinLines
	TotalWinAmount float64
}

type ResultWinLines struct {
	LineIndex          int
	WinAmount          float64
	SymbolMatchesCount int
}
