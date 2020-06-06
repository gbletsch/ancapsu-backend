package main

type eachItem struct {
	ID           string     `json:"id"`
	Authors      authors    `json:"authors"`
	Categories   categories `json:"categories"`
	StartingText string     `json:"starting-text"`
	Status       int        `json:"status"`
	StatusName   string     `json:"status-name"`
	Text         string     `json:"text"`
	Paragraphs   []string   `json:"paragraphs"`
	Title        string     `json:"title"`
	UserVote     int        `json:"user-vote"`
	VoteFake     int        `json:"vote-fake"`
	VoteGood     int        `json:"vote-good"`
	VoteNot      int        `json:"vote-not"`
	VoteOld      int        `json:"vote-old"`
	VoteVery     int        `json:"vote-very"`
	VoteWrite    int        `json:"vote-write"`
}

type authors struct {
	Authored author `json:"authored"`
	Narrated author `json:"narrated"`
	Produced author `json:"produced"`
	Revised  author `json:"revised"`
	Sugested author `json:"sugested"`
}

type author struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Date  string `json:"date"`
	Label string `json:"label"`
}

type categories struct {
	MainCategory string     `json:"main-category"`
	Categories   []category `json:"categories"`
}

type category struct {
	Category string `json:"category"`
	Label    string `json:"label"`
}
