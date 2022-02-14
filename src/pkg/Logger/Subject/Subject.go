package Subject

import "strconv"

type Type int

const (
	ApiPath Type = iota
	ApiResponse
)

type Subject struct {
	SubjectType Type
	SubjectName string
}

func SuccessfulResponses() *[]Subject {
	var subjects []Subject
	for i := 200; i <= 208; i++ {
		subject := Subject{
			SubjectType: ApiResponse,
			SubjectName: strconv.Itoa(i),
		}
		subjects = append(subjects, subject)
	}
	return &subjects
}
func ServerErrors() *[]Subject {
	var subjects []Subject
	for i := 500; i <= 511; i++ {
		subject := Subject{
			SubjectType: ApiResponse,
			SubjectName: strconv.Itoa(i),
		}
		subjects = append(subjects, subject)
	}
	return &subjects
}
func ClientErrors() *[]Subject {
	var subjects []Subject
	for i := 400; i <= 431; i++ {
		subject := Subject{
			SubjectType: ApiResponse,
			SubjectName: strconv.Itoa(i),
		}
		subjects = append(subjects, subject)
	}
	subjects = append(subjects, Subject{
		SubjectType: ApiResponse,
		SubjectName: "451",
	})
	return &subjects
}
func StatusCode(ResponseCode int) *Subject {
	return &Subject{
		SubjectType: ApiResponse,
		SubjectName: strconv.Itoa(ResponseCode),
	}
}

func Path(Path string) *Subject {
	return &Subject{
		SubjectType: ApiPath,
		SubjectName: Path,
	}
}
