package painter

import "pictures/pkg/model"

type Painter interface {
	DrawPicture(draft model.PictureDraft, canvas model.Canvas)
}

func NewPainter() Painter {
	return &painter{}
}

type painter struct{}

func (p *painter) DrawPicture(draft model.PictureDraft, canvas model.Canvas) {
	for i := 0; i < draft.GetShapesCount(); i++ {
		shape := draft.GetShape(i)
		shape.Draw(canvas)
	}
}
