package manha

import (
	"fmt"
	"net/http"
	"os"
)

type myError struct {
	status int
	msg    string
}

// type error interface {
//     Error() string
// }

func (m *myError) Error() string {
	return fmt.Sprintf("%d | %v", m.status, m.msg)
}

func myCustomErrorTest(status int) (int, error) {
	if status >= http.StatusBadRequest {
		return http.StatusInternalServerError, &myError{
			status: status,
			msg:    "Algo deu errado!!",
		}
	}
	return http.StatusOK, nil
}

func Exemplo1() {
	status, err := myCustomErrorTest(400)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Status %d, Funciona!", status)
}
