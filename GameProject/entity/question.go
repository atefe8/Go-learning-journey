package entity

type Question struct {
	ID              uint
	Text            string
	PossibleAnswers []string
	CorrectAnswer   uint
	Difficulty      string
	CategoryID      uint
}

type PossibleAnswer struct {
	ID     uint
	Choice PossibleAnswerChoice
	Text   string
}

type PossibleAnswerChoice uint

const (
	PossibleAnswerA PossibleAnswerChoice = 1
	PossibleAnswerB PossibleAnswerChoice = 2
	PossibleAnswerC PossibleAnswerChoice = 3
	PossibleAnswerD PossibleAnswerChoice = 4
)

type QuestionDifficulty uint

const (
	Easy   QuestionDifficulty = 1
	Medium QuestionDifficulty = 2
	Hard   QuestionDifficulty = 3
)
