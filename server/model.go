package server

// Represents the model for the entire game.
type Ludo struct {
	currentPlayer Player
	board         Board
	players       []Player
}

// Represents the game board
type Board struct {
	// The active spaces around the board
	activeBoard []GameSpace

	// The home rows for each player
	homes map[Player][]GameSpace

	// The final goals for each player
	goals map[Player][]GameSpace
}

// Represents a space on the board.
type GameSpace struct {
	// Pointer because this is optional
	piece *Player
}

// Represents an id of a player.
type Player int
