package model

type Config struct {
	// Actions is a set of necessary actions
	Actions []Action `json:"actions"`
}

type Action struct {
	// Name is an arbitrary name for the action.
	Name        string `json:"name"`
	XCoordinate int    `json:"xCoordinate"`
	YCoordinate int    `json:"yCoordinate"`
	// DelayAfter in millisecond
	DelayAfter int `json:"delayAfter"`
	// ScrollAfterDelay takes a negative value to scroll down
	ScrollAfterDelay int `json:"scrollAfterDelay"`
}
