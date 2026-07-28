package highscores

import "slices"

type HighScores struct {
	scores []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	return &HighScores{scores: scores}
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.scores
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	return s.scores[len(s.scores)-1]
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	x := make([]int, len(s.scores))
	copy(x, s.scores)
	slices.Sort(x)
	slices.Reverse(x)
	return x[0]
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	x := make([]int, len(s.scores))
	copy(x, s.scores)
	slices.Sort(x)
	slices.Reverse(x)
	if len(x) <= 3 {
		return x
	}
	return x[:3]
}
