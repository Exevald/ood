package model

type Paragraph interface {
	GetText() string
	SetText(text string)
}

func NewParagraph(text string) Paragraph {
	return &paragraph{text: text}
}

type paragraph struct {
	text string
}

func (p *paragraph) GetText() string {
	return p.text
}

func (p *paragraph) SetText(text string) {
	p.text = text
}
