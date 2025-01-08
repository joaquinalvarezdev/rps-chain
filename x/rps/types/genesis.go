package types

// NewGenesisState creates a new genesis state with default values.
func NewGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
		Games:  DefaultGames(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
func (gs *GenesisState) Validate() error {

	if err := gs.Params.Validate(); err != nil {
		return err
	}

	unique := make(map[uint64]bool)

	for _, game := range gs.Games {
		if _, ok := unique[game.GameNumber]; ok {
			return ErrDuplicatedGameNumber
		}
		if err := game.Validate(); err != nil {
			return err
		}
		unique[game.GameNumber] = true
	}

	return nil
}
