package language

type Termer interface {
	Name() bool
}

type Term struct {
	Termer
}

func (t Term) Name() {

}
