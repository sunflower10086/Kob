package util

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetMap(t *testing.T) {
	playerA := Player{
		Id:       10,
		BotId:    -1,
		BotCode:  "",
		Sx:       1,
		Sy:       2,
		Steps:    nil,
		Username: "",
		Photo:    "",
	}
	tmp := []int32{0, 1, 3, 2, 1, 4, 3}

	playerA.Steps = append(playerA.Steps, tmp...)

	fmt.Println(playerA.Steps)

	fmt.Println(playerA.GetStepsString())
}

func TestStringBuilder(t *testing.T) {
	resp := strings.Builder{}

	tmp := []int32{1, 2, 3, 4, 4}
	for i := 0; i < len(tmp); i++ {
		_, err := resp.WriteString(fmt.Sprintf("%d", tmp[i]))
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(resp.String())
}
