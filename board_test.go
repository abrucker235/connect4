package connect4

import "testing"

func TestMoveHorizontal(t *testing.T) {
	board := &Board{positions: [6][7]string{
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "red", "red", "red", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
	}, connected: 0}

	actual, err := board.move(4, 2, "red")
	if !actual && err != nil {
		t.Errorf("Expected true got: %t", actual)
	}
}

func TestMoveVertical(t *testing.T) {
	board := &Board{positions: [6][7]string{
		{"", "red", "", "", "", "", ""},
		{"", "red", "", "", "", "", ""},
		{"", "red", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
	}, connected: 0}

	actual, err := board.move(1, 3, "red")
	if !actual && err != nil {
		t.Errorf("Expected true got: %t", actual)
	}
}

func TestMoveDiagonal(t *testing.T) {
	board := &Board{positions: [6][7]string{
		{"red", "", "", "", "", "", ""},
		{"", "red", "", "", "", "", ""},
		{"", "", "red", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
	}, connected: 0}

	actual, err := board.move(3, 3, "red")
	if !actual && err != nil {
		t.Errorf("Expected true got: %t", actual)
	}
}

func TestMoveNonWin(t *testing.T) {
	board := &Board{positions: [6][7]string{
		{"", "red", "", "", "", "", ""},
		{"", "red", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
	}, connected: 0}

	actual, err := board.move(1, 2, "red")
	if !actual && err != nil {
		t.Errorf("Expected true got: %t", actual)
	}
}

func TestMoveAlreadyTaken(t *testing.T) {
	board := &Board{positions: [6][7]string{
		{"", "red", "", "", "", "", ""},
		{"", "red", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
	}, connected: 0}

	actual, err := board.move(1, 1, "blue")
	if actual || (err == nil && err.Error() == "Position already taken by: red") {
		t.Errorf("Expected Position already taken by: red got: %s", err.Error())
	}
}

func TestCheckHorizontalRL(t *testing.T) {
	board := &Board{positions: [6][7]string{{"blue", "red", "red", "red", "red", "blue", "blue"}}, connected: 0}

	board.CheckHorizontal(0, 4)
	if board.connected != 3 {
		t.Errorf("Expect Connected: 4 got: %d", board.connected)
	}
}

func TestCheckHorizontalLR(t *testing.T) {
	board := &Board{positions: [6][7]string{{"blue", "red", "red", "red", "red", "blue", "blue"}}, connected: 0}

	board.CheckHorizontal(0, 1)
	if board.connected != 3 {
		t.Errorf("Expect Connected: 4 got: %d", board.connected)
	}
}

func TestCheckHorizonalMiddle(t *testing.T) {
	board := &Board{positions: [6][7]string{{"blue", "red", "red", "red", "red", "blue", "blue"}}, connected: 0}

	board.CheckHorizontal(0, 3)
	if board.connected != 3 {
		t.Errorf("Expect Connected: 4 got: %d", board.connected)
	}
}

func TestCheckHorizontal(t *testing.T) {
	board := &Board{positions: [6][7]string{{"blue", "red", "red", "red", "red", "blue", "blue"}}, connected: 0}

	board.CheckHorizontal(0, 6)
	if board.connected != 1 {
		t.Errorf("Expect Connected: 2 got: %d", board.connected)
	}
}

func TestCheckVerticalBottom(t *testing.T) {
	board := &Board{positions: [6][7]string{
		{"", "blue", "red", "", "", "blue", ""},
		{"", "", "red", "", "", "", ""},
		{"", "", "red", "", "", "", ""},
		{"", "", "red", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
	}, connected: 0}

	board.CheckVertical(3, 3)
	if board.connected != 3 {
		t.Errorf("Expect Connected: 3 got: %d", board.connected)
	}
}

func TestCheckVerticalTop(t *testing.T) {
	board := &Board{positions: [6][7]string{
		{"", "", "blue", "", "", "", ""},
		{"", "", "blue", "", "", "", ""},
		{"", "", "red", "", "", "", ""},
		{"", "", "red", "", "", "", ""},
		{"", "", "red", "", "", "", ""},
		{"", "", "red", "", "", "", ""},
	}, connected: 0}

	board.CheckVertical(3, 3)
	if board.connected != 3 {
		t.Errorf("Expect Connected: 3 got: %d", board.connected)
	}
}

func TestCheckVerticalMiddle(t *testing.T) {
	board := &Board{positions: [6][7]string{
		{"", "", "blue", "", "", "", ""},
		{"", "", "red", "", "", "", ""},
		{"", "", "red", "", "", "", ""},
		{"", "", "red", "", "", "", ""},
		{"", "", "red", "", "", "", ""},
		{"", "", "", "", "", "", ""},
	}, connected: 0}

	board.CheckVertical(3, 3)
	if board.connected != 3 {
		t.Errorf("Expect Connected: 3 got: %d", board.connected)
	}
}

func TestCheckDiagonalDownLeft(t *testing.T) {
	board := &Board{positions: [6][7]string{
		{"red", "blue", "blue", "blue", "", "", ""},
		{"", "red", "blue", "blue", "", "", ""},
		{"", "", "red", "blue", "", "", ""},
		{"", "", "", "red", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
	}, connected: 0}

	board.CheckDiagonal(3, 3)
	if board.connected != 3 {
		t.Errorf("Expect Connected: 3 got: %d", board.connected)
	}
}

func TestCheckDiagonalUpRight(t *testing.T) {
	board := &Board{positions: [6][7]string{
		{"red", "blue", "blue", "blue", "", "", ""},
		{"", "red", "blue", "blue", "", "", ""},
		{"", "", "red", "blue", "", "", ""},
		{"", "", "", "red", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
	}, connected: 0}

	board.CheckDiagonal(0, 0)
	if board.connected != 3 {
		t.Errorf("Expect Connected: 3 got: %d", board.connected)
	}
}

func TestCheckDiagonalMiddleLR(t *testing.T) {
	board := &Board{positions: [6][7]string{
		{"red", "blue", "blue", "blue", "", "", ""},
		{"", "red", "blue", "blue", "", "", ""},
		{"", "", "red", "blue", "", "", ""},
		{"", "", "", "red", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
	}, connected: 0}

	board.CheckDiagonal(2, 2)
	if board.connected != 3 {
		t.Errorf("Expect Connected: 3 got: %d", board.connected)
	}
}

func TestCheckDiagonalDownRight(t *testing.T) {
	board := &Board{positions: [6][7]string{
		{"", "", "", "blue", "blue", "blue", "red"},
		{"", "", "", "blue", "blue", "red", ""},
		{"", "", "", "blue", "red", "", ""},
		{"", "", "", "red", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
	}, connected: 0}

	board.CheckDiagonal(3, 3)
	if board.connected != 3 {
		t.Errorf("Expect Connected: 3 got: %d", board.connected)
	}
}

func TestCheckDiagonalUpLeft(t *testing.T) {
	board := &Board{positions: [6][7]string{
		{"", "", "", "blue", "blue", "blue", "red"},
		{"", "", "", "blue", "blue", "red", ""},
		{"", "", "", "blue", "red", "", ""},
		{"", "", "", "red", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
	}, connected: 0}

	board.CheckDiagonal(0, 6)
	if board.connected != 3 {
		t.Errorf("Expect Connected: 3 got: %d", board.connected)
	}
}

func TestCheckDiagonalMiddleRL(t *testing.T) {
	board := &Board{positions: [6][7]string{
		{"", "", "", "blue", "blue", "blue", "red"},
		{"", "", "", "blue", "blue", "red", ""},
		{"", "", "", "blue", "red", "", ""},
		{"", "", "", "red", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
	}, connected: 0}

	board.CheckDiagonal(1, 5)
	if board.connected != 3 {
		t.Errorf("Expect Connected: 3 got: %d", board.connected)
	}
}

func TestCheckDiagonalTopRight(t *testing.T) {
	board := &Board{positions: [6][7]string{
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "red", "", "", ""},
		{"", "", "", "", "red", "", ""},
		{"", "", "", "", "", "red", ""},
		{"", "", "", "", "", "", "red"},
	}, connected: 0}

	board.CheckDiagonal(1, 5)
	if board.connected != 3 {
		t.Errorf("Expect Connected: 3 got: %d", board.connected)
	}
}

func TestCheckDiagonalTopLeft(t *testing.T) {
	board := &Board{positions: [6][7]string{
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "red", "", "", ""},
		{"", "", "red", "", "", "", ""},
		{"", "red", "", "", "", "", ""},
		{"red", "", "", "", "", "", ""},
	}, connected: 0}

	board.CheckDiagonal(1, 5)
	if board.connected != 3 {
		t.Errorf("Expect Connected: 3 got: %d", board.connected)
	}
}

func TestCheckDiagonalNonMatch(t *testing.T) {
	board := &Board{positions: [6][7]string{
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "red", "", "", "", "", ""},
		{"red", "", "", "", "", "", ""},
	}, connected: 0}

	board.CheckDiagonal(5, 0)
	if board.connected != 1 {
		t.Errorf("Expect Connected: 1 got: %d", board.connected)
	}
}
