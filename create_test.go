package godraw

import (
	"testing"
)

func Test_0T_1B_1Q_0W(t *testing.T) {
	testNum := 0
	test := Data{
		Tickets: []Ticket{},
		Bunches: []Bunch{{Data: "toy1", Nb: 1}},
	}
	d, err := CreateDraw(test)

	if nil == d.Winners {
		t.Errorf("Expected non-nil winners list")
	}
	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
	if nil != err {
		t.Errorf("Expected no errors, go: %s", err.Error())
	}
}

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

func Test_1T_0B_0Q_0W(t *testing.T) {
	testNum := 0
	test := Data{
		Tickets: []Ticket{{Data: "a", Id: "1"}},
		Bunches: []Bunch{},
	}
	d, err := CreateDraw(test)

	if nil == d.Winners {
		t.Errorf("Expected non-nil winners list")
	}
	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
	if nil != err {
		t.Errorf("Expected no errors, go: %s", err.Error())
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
		Bunches: []Bunch{{Id: "toy1", Nb: 10}, {Id: "toy2", Nb: 10}},
		Mode:    "lottery",
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_3T_3B_1Q_1Q_2W_1IT_0PM(t *testing.T) {
	testNum := 2
	test := Data{
		Tickets:           []Ticket{{Data: "a", Id: "1", Owner: ""}, {Data: "a", Id: "2", Owner: ""}, {Data: "a", Id: "3", Owner: ""}},
		Bunches:           []Bunch{{Id: "toy1", Nb: 1}, {Id: "toy2", Nb: 1}, {Id: "toy3", Nb: 1}},
		Mode:              "lottery",
		IgnoredTickets:    []Ticket{{Data: "a", Id: "2", Owner: ""}},
		PartialDraw:       false,
		PartialMaxWinners: 0,
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_3T_3B_1Q_10Q_2W_0IT_2PM(t *testing.T) {
	testNum := 2
	test := Data{
		Tickets:           []Ticket{{Data: "a", Id: "1", Owner: ""}, {Data: "a", Id: "2", Owner: ""}, {Data: "a", Id: "3", Owner: ""}},
		Bunches:           []Bunch{{Id: "toy1", Nb: 10}, {Id: "toy2", Nb: 10}, {Id: "toy3", Nb: 10}},
		Mode:              "lottery",
		IgnoredTickets:    []Ticket{},
		PartialDraw:       true,
		PartialMaxWinners: 2,
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_3T_3B_1Q_10Q_1IB(t *testing.T) {
	testNum := 2
	test := Data{
		Tickets:        []Ticket{{Data: "a", Id: "1", Owner: ""}, {Data: "", Id: "2", Owner: ""}, {Data: "a", Id: "3", Owner: ""}},
		Bunches:        []Bunch{{Id: "toy1", Nb: 1}, {Id: "toy2", Nb: 1}, {Id: "toy3", Nb: 1}},
		Mode:           "lottery",
		IgnoredBunches: []Bunch{{Id: "toy1", Nb: 1}},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}
