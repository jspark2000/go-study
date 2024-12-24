package judger

type Grader interface {
	CheckAnswer(target string, answer string) int
}

const (
	ResultAccept = iota // 0
	ResultFail          // 1
	CompileError        // 2
	SystemError         // 3
)

type BasicGrader struct{}

func NewBasicGrader() BasicGrader {
	return BasicGrader{}
}

// TODO: 정답 체크 함수 완성
func (b *BasicGrader) CheckAnswer(target string, answer string) int {}
