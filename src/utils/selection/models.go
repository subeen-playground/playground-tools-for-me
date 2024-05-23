package selection

type Prompt interface {
	PlainText() string
}

type Item struct {
	Label       Prompt
	LabelActive Prompt
	Help        string
}

type Selection struct {
	Items []string
}

func NewSelection(items []string) Selection {
	return Selection{items}
}

// Item - Label, LabelOnActive, LabelOnInactive
// Item Display

// String+Prompt
type StringExtended string

func (s StringExtended) PlainText() string {
	return string(s)
}
