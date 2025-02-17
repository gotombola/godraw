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

func Test_10T_2B_9T_SAME_OWNER_FEATURE_MAX_1_PER_OWNER(t *testing.T) {
	testNum := 2
	multipleOwnerId := "o1"
	test := Data{
		Tickets: []Ticket{
			{Data: "a", Id: "1", Owner: multipleOwnerId},
			{Data: "a", Id: "2", Owner: "o2"},
			{Data: "a", Id: "3", Owner: multipleOwnerId},
			{Data: "a", Id: "4", Owner: multipleOwnerId},
			{Data: "a", Id: "5", Owner: multipleOwnerId},
			{Data: "a", Id: "6", Owner: multipleOwnerId},
			{Data: "a", Id: "7", Owner: multipleOwnerId},
			{Data: "a", Id: "8", Owner: multipleOwnerId},
			{Data: "a", Id: "9", Owner: multipleOwnerId},
			{Data: "a", Id: "10", Owner: multipleOwnerId},
		},
		Bunches: []Bunch{{Id: "toy1", Nb: 1}, {Id: "toy2", Nb: 1}},
		Mode:    "raffle",
		Features: []string{
			"max_1_per_owner",
		},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
	m := map[string]int{
		"o1": 0,
		"o2": 0,
	}
	mT := map[string]string{}

	for _, t := range test.Tickets {
		mT[t.Id] = t.Owner
	}
	for _, w := range d.Winners {
		m[mT[w.T]]++
	}
	for k, v := range m {
		if v > 1 {
			t.Errorf("Owner [%s] has [%d] winning tickets, not allowed", k, v)
		}
	}
}
func Test_9T_2B_9T_SAME_OWNER_NO_FEATURE_MAX_1_PER_OWNER(t *testing.T) {
	testNum := 2
	multipleOwnerId := "o1"
	test := Data{
		Tickets: []Ticket{
			{Data: "a", Id: "1", Owner: multipleOwnerId},
			{Data: "a", Id: "3", Owner: multipleOwnerId},
			{Data: "a", Id: "4", Owner: multipleOwnerId},
			{Data: "a", Id: "5", Owner: multipleOwnerId},
			{Data: "a", Id: "6", Owner: multipleOwnerId},
			{Data: "a", Id: "7", Owner: multipleOwnerId},
			{Data: "a", Id: "8", Owner: multipleOwnerId},
			{Data: "a", Id: "9", Owner: multipleOwnerId},
			{Data: "a", Id: "10", Owner: multipleOwnerId},
		},
		Bunches:  []Bunch{{Id: "toy1", Nb: 1}, {Id: "toy2", Nb: 1}},
		Mode:     "raffle",
		Features: []string{},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
	m := map[string]int{
		"o1": 0,
	}
	mT := map[string]string{}

	for _, t := range test.Tickets {
		mT[t.Id] = t.Owner
	}
	for _, w := range d.Winners {
		m[mT[w.T]]++
	}
	for k, v := range m {
		if v <= 1 {
			t.Errorf("Owner [%s] has only [%d] winning tickets", k, v)
		}
	}
}
func Test_10T_2B_9T_SAME_OWNER_FEATURE_MAX_2_PER_OWNER(t *testing.T) {
	testNum := 2
	multipleOwnerId := "o1"
	test := Data{
		Tickets: []Ticket{
			{Data: "a", Id: "1", Owner: multipleOwnerId},
			{Data: "a", Id: "2", Owner: "o2"},
			{Data: "a", Id: "3", Owner: multipleOwnerId},
			{Data: "a", Id: "4", Owner: multipleOwnerId},
			{Data: "a", Id: "5", Owner: multipleOwnerId},
			{Data: "a", Id: "6", Owner: multipleOwnerId},
			{Data: "a", Id: "7", Owner: multipleOwnerId},
			{Data: "a", Id: "8", Owner: multipleOwnerId},
			{Data: "a", Id: "9", Owner: multipleOwnerId},
			{Data: "a", Id: "10", Owner: multipleOwnerId},
		},
		Bunches: []Bunch{{Id: "toy1", Nb: 1}, {Id: "toy2", Nb: 1}},
		Mode:    "raffle",
		Features: []string{
			"max_2_per_owner",
		},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
	m := map[string]int{
		"o1": 0,
		"o2": 0,
	}
	mT := map[string]string{}

	for _, t := range test.Tickets {
		mT[t.Id] = t.Owner
	}
	for _, w := range d.Winners {
		m[mT[w.T]]++
	}
	for k, v := range m {
		if v > 2 {
			t.Errorf("Owner [%s] has [%d] winning tickets, not allowed", k, v)
		}
	}
}
func Test_6T_2B_CHOSEN_BUNCHES(t *testing.T) {
	test := Data{
		Tickets: []Ticket{
			{Data: "a", Id: "1", Owner: "o1", ChosenBunches: []string{"toy2"}},
			{Data: "a", Id: "2", Owner: "o1", ChosenBunches: []string{"toy2"}},
			{Data: "a", Id: "3", Owner: "o1", ChosenBunches: []string{"toy2"}},
			{Data: "a", Id: "4", Owner: "o1", ChosenBunches: []string{"toy2"}},
			{Data: "a", Id: "5", Owner: "o1", ChosenBunches: []string{"toy2"}},
			{Data: "a", Id: "6", Owner: "o2", ChosenBunches: []string{"toy1"}},
		},
		Bunches:  []Bunch{{Id: "toy1", Nb: 1}, {Id: "toy2", Nb: 1}},
		Mode:     "raffle",
		Features: []string{},
	}
	d, _ := CreateDraw(test)

	for _, w := range d.Winners {
		if w.To == "o1" && w.B != "toy2" {
			t.Errorf("Owner [%s] won a bunch it didn't choose [%s]", w.To, w.B)
		}
		if w.To == "o2" && w.B != "toy1" {
			t.Errorf("Owner [%s] won a bunch it didn't choose [%s]", w.To, w.B)
		}
	}
}
func Test_9T_3B_MAX_1_TAG(t *testing.T) {
	test := Data{
		Tickets: []Ticket{
			{Data: "a", Id: "1", Owner: "o1"},
			{Data: "a", Id: "2", Owner: "o1"},
			{Data: "a", Id: "3", Owner: "o1"},
			{Data: "a", Id: "4", Owner: "o1"},
			{Data: "a", Id: "5", Owner: "o1"},
			{Data: "a", Id: "6", Owner: "o1"},
			{Data: "a", Id: "7", Owner: "o1"},
			{Data: "a", Id: "8", Owner: "o1"},
			{Data: "a", Id: "9", Owner: "o2"},
		},
		Bunches: []Bunch{{Id: "toy1", Nb: 1, Tags: []string{"toy"}}, {Id: "toy2", Nb: 1, Tags: []string{"toy"}}, {Id: "travel1", Nb: 1, Tags: []string{"travel"}}},
		Mode:    "raffle",
		Features: []string{
			"max_1_per_tag_per_owner",
		},
	}
	d, _ := CreateDraw(test)

	m := map[string]map[string]int{
		"o1": {"toy": 0, "travel": 0},
		"o2": {"toy": 0, "travel": 0},
	}
	for _, w := range d.Winners {
		for _, tag := range w.Bt {
			m[w.To][tag]++
		}
	}
	for k, v := range m {
		for k2, v2 := range v {
			if v2 > 2 {
				t.Errorf("Owner [%s] has [%d] winning tickets with tag [%s], not allowed", k, v2, k2)
			}
		}
	}
}
