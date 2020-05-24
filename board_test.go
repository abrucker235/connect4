package connect4

import "testing"

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
