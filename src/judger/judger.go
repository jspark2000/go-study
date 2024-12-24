package judger

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/google/uuid"
)

type JudgeResult struct {
	Result int `json:"result"`
}

type Judger interface {
	Judge(code string) JudgeResult
}

type BasicJudger struct {
	grader Grader
}

func NewBasicJudger() BasicJudger {
	grader := NewBasicGrader()
	return BasicJudger{grader: &grader}
}

func (b *BasicJudger) Judge(code string) JudgeResult {
	uuidStr := uuid.New().String()
	cFileName := fmt.Sprintf("user_code_%s.c", uuidStr)
	outputFileName := fmt.Sprintf("user_program_%s", uuidStr)

	err := os.WriteFile(cFileName, []byte(code), 0644)
	if err != nil {
		log.Fatalf("C 코드 파일 저장 실패: %v", err)
	}
	defer func() {
		if err := os.Remove(cFileName); err != nil {
			log.Printf("C 코드 파일 삭제 실패: %v", err)
		}
		if err := os.Remove(outputFileName); err != nil {
			log.Printf("실행 파일 삭제 실패: %v", err)
		}
	}()

	// gcc로 컴파일
	cmd := exec.Command("gcc", cFileName, "-o", outputFileName)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("gcc 컴파일 실패: %v\n%s", err, string(output))
		return JudgeResult{
			Result: CompileError,
		}
	}

	// 컴파일한 파일 실행
	execCmd := exec.Command("./" + outputFileName)
	execOutput, err := execCmd.CombinedOutput()

	if err != nil {
		log.Printf("코드 실행 실패: %v\n%s", err, string(execOutput))
		return JudgeResult{
			Result: SystemError,
		}
	}

	// TODO: 채점 결과를 JudgeResult 구조체 반환

}
