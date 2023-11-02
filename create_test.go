package godraw

import (
	"testing"
)

func Test_1T_1B_1Q_1W(t *testing.T) {
	testNum := 1
	test := Data{
		Tickets: []Ticket{{Data: "a", Id: "1"}},
		Bunches: []Bunch{{Data: "toy1", Nb: 1}},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_3T_3B_3Q_3W(t *testing.T) {
	testNum := 3
	test := Data{
		Tickets: []Ticket{{Data: "a", Id: "1"}, {Data: "a", Id: "1"}, {Data: "a", Id: "1"}},
		Bunches: []Bunch{{Data: "toy1", Nb: 1}, {Data: "toy2", Nb: 1}, {Data: "toy3", Nb: 1}},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_3T_2B_2Q_2W(t *testing.T) {
	testNum := 2
	test := Data{
		Tickets: []Ticket{{Data: "a", Id: "1"}, {Data: "a", Id: "1"}, {Data: "a", Id: "1"}},
		Bunches: []Bunch{{Id: "toy1", Nb: 1}, {Id: "toy2", Nb: 1}},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_2T_3B_3Q_2W(t *testing.T) {
	testNum := 2
	test := Data{
		Tickets: []Ticket{{Data: "a", Id: "1"}, {Data: "a", Id: "1"}},
		Bunches: []Bunch{{Id: "toy1", Nb: 1}, {Id: "toy2", Nb: 1}, {Id: "toy3", Nb: 1}},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_0T_3B_3Q_0W(t *testing.T) {
	testNum := 0
	test := Data{
		Tickets: []Ticket{},
		Bunches: []Bunch{{Id: "toy1", Nb: 1}, {Id: "toy2", Nb: 1}, {Id: "toy3", Nb: 1}},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_3T_0B_0Q_0W(t *testing.T) {
	testNum := 0
	test := Data{
		Tickets: []Ticket{{Data: "a", Id: "1"}, {Data: "a", Id: "1"}, {Data: "a", Id: "1"}},
		Bunches: []Bunch{},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_3T_1B_4Q_3W(t *testing.T) {
	testNum := 3
	test := Data{
		Tickets: []Ticket{{Data: "a", Id: "1"}, {Data: "a", Id: "1"}, {Data: "a", Id: "1"}},
		Bunches: []Bunch{{Id: "toy1", Nb: 4}},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_3T_2B_1Q_2Q_3W(t *testing.T) {
	testNum := 3
	test := Data{
		Tickets: []Ticket{{Data: "a", Id: "1"}, {Data: "a", Id: "1"}, {Data: "a", Id: "1"}},
		Bunches: []Bunch{{Id: "toy1", Nb: 1}, {Id: "toy2", Nb: 2}},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_3T_2B_1Q_10Q_3W(t *testing.T) {
	testNum := 3
	test := Data{
		Tickets: []Ticket{{Data: "a", Id: "1"}, {Data: "a", Id: "1"}, {Data: "a", Id: "1"}},
		Bunches: []Bunch{{Id: "toy1", Nb: 1}, {Id: "toy2", Nb: 10}},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_3T_2B_1Q_10Q_2W(t *testing.T) {
	testNum := 2
	test := Data{
		Tickets: []Ticket{{Data: "a", Id: "1", Owner: ""}, {Data: "a", Id: "1", Owner: "test"}, {Data: "a", Id: "1", Owner: "test"}},
		Bunches: []Bunch{{Id: "toy1", Nb: 1}, {Id: "toy2", Nb: 10}},
		Mode:    "lottery",
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}
