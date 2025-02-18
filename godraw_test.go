package godraw

import (
	"encoding/json"
	"fmt"
	"github.com/gotombola/godraw/types"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func Test_0T_1B_1Q_0W(t *testing.T) {
	testNum := 0
	test := Data{
		Tickets: []types.Ticket{},
		Bunches: []types.Bunch{{Data: "toy1", Quantity: 1}},
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
		Tickets: []types.Ticket{{Data: "a", Id: "1"}},
		Bunches: []types.Bunch{{Data: "toy1", Quantity: 1}},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_1T_0B_0Q_0W(t *testing.T) {
	testNum := 0
	test := Data{
		Tickets: []types.Ticket{{Data: "a", Id: "1"}},
		Bunches: []types.Bunch{},
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
		Tickets: []types.Ticket{{Data: "a", Id: "1"}, {Data: "a", Id: "1"}, {Data: "a", Id: "1"}},
		Bunches: []types.Bunch{{Data: "toy1", Quantity: 1}, {Data: "toy2", Quantity: 1}, {Data: "toy3", Quantity: 1}},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_3T_2B_2Q_2W(t *testing.T) {
	testNum := 2
	test := Data{
		Tickets: []types.Ticket{{Data: "a", Id: "1"}, {Data: "a", Id: "1"}, {Data: "a", Id: "1"}},
		Bunches: []types.Bunch{{Id: "toy1", Quantity: 1}, {Id: "toy2", Quantity: 1}},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_2T_3B_3Q_2W(t *testing.T) {
	testNum := 2
	test := Data{
		Tickets: []types.Ticket{{Data: "a", Id: "1"}, {Data: "a", Id: "1"}},
		Bunches: []types.Bunch{{Id: "toy1", Quantity: 1}, {Id: "toy2", Quantity: 1}, {Id: "toy3", Quantity: 1}},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_0T_3B_3Q_0W(t *testing.T) {
	testNum := 0
	test := Data{
		Tickets: []types.Ticket{},
		Bunches: []types.Bunch{{Id: "toy1", Quantity: 1}, {Id: "toy2", Quantity: 1}, {Id: "toy3", Quantity: 1}},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_3T_0B_0Q_0W(t *testing.T) {
	testNum := 0
	test := Data{
		Tickets: []types.Ticket{{Data: "a", Id: "1"}, {Data: "a", Id: "1"}, {Data: "a", Id: "1"}},
		Bunches: []types.Bunch{},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_3T_1B_4Q_3W(t *testing.T) {
	testNum := 3
	test := Data{
		Tickets: []types.Ticket{{Data: "a", Id: "1"}, {Data: "a", Id: "1"}, {Data: "a", Id: "1"}},
		Bunches: []types.Bunch{{Id: "toy1", Quantity: 4}},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_3T_2B_1Q_2Q_3W(t *testing.T) {
	testNum := 3
	test := Data{
		Tickets: []types.Ticket{{Data: "a", Id: "1"}, {Data: "a", Id: "1"}, {Data: "a", Id: "1"}},
		Bunches: []types.Bunch{{Id: "toy1", Quantity: 1}, {Id: "toy2", Quantity: 2}},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_3T_2B_1Q_10Q_3W(t *testing.T) {
	testNum := 3
	test := Data{
		Tickets: []types.Ticket{{Data: "a", Id: "1"}, {Data: "a", Id: "1"}, {Data: "a", Id: "1"}},
		Bunches: []types.Bunch{{Id: "toy1", Quantity: 1}, {Id: "toy2", Quantity: 10}},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != testNum {
		t.Errorf("Expected %d, got %d", testNum, len(d.Winners))
	}
}

func Test_3T_2B_1Q_10Q_2W(t *testing.T) {
	testNum := 2
	test := Data{
		Tickets: []types.Ticket{{Data: "a", Id: "1", Owner: ""}, {Data: "a", Id: "1", Owner: "test"}, {Data: "a", Id: "1", Owner: "test"}},
		Bunches: []types.Bunch{{Id: "toy1", Quantity: 10}, {Id: "toy2", Quantity: 10}},
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
		Tickets:           []types.Ticket{{Data: "a", Id: "1", Owner: ""}, {Data: "a", Id: "2", Owner: ""}, {Data: "a", Id: "3", Owner: ""}},
		Bunches:           []types.Bunch{{Id: "toy1", Quantity: 1}, {Id: "toy2", Quantity: 1}, {Id: "toy3", Quantity: 1}},
		Mode:              "lottery",
		IgnoredTickets:    []types.Ticket{{Data: "a", Id: "2", Owner: ""}},
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
		Tickets:           []types.Ticket{{Data: "a", Id: "1", Owner: ""}, {Data: "a", Id: "2", Owner: ""}, {Data: "a", Id: "3", Owner: ""}},
		Bunches:           []types.Bunch{{Id: "toy1", Quantity: 10}, {Id: "toy2", Quantity: 10}, {Id: "toy3", Quantity: 10}},
		Mode:              "lottery",
		IgnoredTickets:    []types.Ticket{},
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
		Tickets:        []types.Ticket{{Data: "a", Id: "1", Owner: ""}, {Data: "", Id: "2", Owner: ""}, {Data: "a", Id: "3", Owner: ""}},
		Bunches:        []types.Bunch{{Id: "toy1", Quantity: 1}, {Id: "toy2", Quantity: 1}, {Id: "toy3", Quantity: 1}},
		Mode:           "lottery",
		IgnoredBunches: []types.Bunch{{Id: "toy1", Quantity: 1}},
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
		Tickets: []types.Ticket{
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
		Bunches: []types.Bunch{{Id: "toy1", Quantity: 1}, {Id: "toy2", Quantity: 1}},
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
		m[mT[w.Ticket]]++
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
		Tickets: []types.Ticket{
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
		Bunches:  []types.Bunch{{Id: "toy1", Quantity: 1}, {Id: "toy2", Quantity: 1}},
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
		m[mT[w.Ticket]]++
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
		Tickets: []types.Ticket{
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
		Bunches: []types.Bunch{{Id: "toy1", Quantity: 1}, {Id: "toy2", Quantity: 1}},
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
		m[mT[w.Ticket]]++
	}
	for k, v := range m {
		if v > 2 {
			t.Errorf("Owner [%s] has [%d] winning tickets, not allowed", k, v)
		}
	}
}
func Test_6T_2B_CHOSEN_BUNCHES(t *testing.T) {
	test := Data{
		Tickets: []types.Ticket{
			{Data: "a", Id: "1", Owner: "o1", ChosenBunches: []string{"toy2"}},
			{Data: "a", Id: "2", Owner: "o1", ChosenBunches: []string{"toy2"}},
			{Data: "a", Id: "3", Owner: "o1", ChosenBunches: []string{"toy2"}},
			{Data: "a", Id: "4", Owner: "o1", ChosenBunches: []string{"toy2"}},
			{Data: "a", Id: "5", Owner: "o1", ChosenBunches: []string{"toy2"}},
			{Data: "a", Id: "6", Owner: "o2", ChosenBunches: []string{"toy1"}},
		},
		Bunches:  []types.Bunch{{Id: "toy1", Quantity: 1}, {Id: "toy2", Quantity: 1}},
		Mode:     "raffle",
		Features: []string{},
	}
	d, _ := CreateDraw(test)

	for _, w := range d.Winners {
		if w.TicketOwner == "o1" && w.Bunch != "toy2" {
			t.Errorf("Owner [%s] won a bunch it didn't choose [%s]", w.TicketOwner, w.Bunch)
		}
		if w.TicketOwner == "o2" && w.Bunch != "toy1" {
			t.Errorf("Owner [%s] won a bunch it didn't choose [%s]", w.TicketOwner, w.Bunch)
		}
	}
}
func Test_9T_3B_MAX_1_TAG(t *testing.T) {
	test := Data{
		Tickets: []types.Ticket{
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
		Bunches: []types.Bunch{{Id: "toy1", Quantity: 1, Tags: []string{"toy"}}, {Id: "toy2", Quantity: 1, Tags: []string{"toy"}}, {Id: "travel1", Quantity: 1, Tags: []string{"travel"}}},
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
		for _, tag := range w.BunchTags {
			m[w.TicketOwner][tag]++
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
func Test_Bunch_JSON_Marshalling(t *testing.T) {
	originalJSON := `{"id":"1","d":"data","n":5,"ro":2,"t":["tag1","tag2"]}`
	var bunch types.Bunch
	err := json.Unmarshal([]byte(originalJSON), &bunch)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	marshalledJSON, err := json.Marshal(bunch)
	if err != nil {
		t.Errorf("Failed to marshal struct: %v", err)
	}

	var resultMap, originalMap map[string]interface{}
	err = json.Unmarshal(marshalledJSON, &resultMap)
	if err != nil {
		t.Errorf("Failed to unmarshal marshalled JSON: %v", err)
	}

	err = json.Unmarshal([]byte(originalJSON), &originalMap)
	if err != nil {
		t.Errorf("Failed to unmarshal original JSON: %v", err)
	}

	if len(resultMap) != len(originalMap) {
		t.Errorf("Mismatch in JSON structure, expected %v but got %v", originalMap, resultMap)
	}

	if !reflect.DeepEqual(originalMap, resultMap) {
		t.Errorf("Mismatch in JSON structure:\nExpected: %v\nGot: %v", originalMap, resultMap)
	}
}
func Test_Ticket_JSON_Marshalling(t *testing.T) {
	originalJSON := `{"d":"data","id":"1","o":"o1","b":["bunch1","bunch2","bunch3"]}`

	var ticket types.Ticket
	err := json.Unmarshal([]byte(originalJSON), &ticket)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	marshalledJSON, err := json.Marshal(ticket)
	if err != nil {
		t.Errorf("Failed to marshal struct: %v", err)
	}

	var originalMap, resultMap map[string]interface{}
	err = json.Unmarshal([]byte(originalJSON), &originalMap)
	if err != nil {
		t.Errorf("Failed to unmarshal original JSON: %v", err)
	}

	err = json.Unmarshal(marshalledJSON, &resultMap)
	if err != nil {
		t.Errorf("Failed to unmarshal marshalled JSON: %v", err)
	}

	if !reflect.DeepEqual(originalMap, resultMap) {
		t.Errorf("Mismatch in JSON structure:\nExpected: %v\nGot: %v", originalMap, resultMap)
	}
}
func Test_9T_2B_7T_WITH_INVALID_TIMESTAMP(t *testing.T) {
	invalidTimestampOwner := "o1"
	test := Data{
		Tickets: []types.Ticket{
			{Data: "a", Id: "1", Owner: invalidTimestampOwner, Timestamp: 5},
			{Data: "a", Id: "3", Owner: invalidTimestampOwner, Timestamp: 5},
			{Data: "a", Id: "4", Owner: invalidTimestampOwner, Timestamp: 5},
			{Data: "a", Id: "5", Owner: invalidTimestampOwner, Timestamp: 5},
			{Data: "a", Id: "6", Owner: invalidTimestampOwner, Timestamp: 5},
			{Data: "a", Id: "7", Owner: invalidTimestampOwner, Timestamp: 5},
			{Data: "a", Id: "8", Owner: "o2", Timestamp: 15},
			{Data: "a", Id: "9", Owner: "o2", Timestamp: 15},
		},
		Bunches:        []types.Bunch{{Id: "toy1", Quantity: 1}, {Id: "toy2", Quantity: 1}},
		Mode:           "raffle",
		StartTimestamp: 10,
		EndTimestamp:   20,
		Features:       []string{},
	}
	d, _ := CreateDraw(test)
	for _, win := range d.Winners {
		if win.TicketOwner == invalidTimestampOwner {
			t.Errorf("Owner should not have won, invalid timestamp")
		}
	}
}

func Test_1MILLION_TICKETS_WITH_TIMESTAMPS(t *testing.T) {
	start := time.Now()
	nbTickets := 1000000
	nbBunches := 100

	tickets := make([]types.Ticket, nbTickets)
	for i := 0; i < nbTickets; i++ {
		tickets[i] = types.Ticket{
			Data:      "data",
			Id:        fmt.Sprintf("%d", i),
			Owner:     fmt.Sprintf("o%d", rand.Intn(1000)),
			Timestamp: rand.Intn(1000),
		}
	}

	bunches := make([]types.Bunch, nbBunches)
	for i := 0; i < nbBunches; i++ {
		bunches[i] = types.Bunch{
			Id:       fmt.Sprintf("bunch%d", i),
			Quantity: 1,
		}
	}

	test := Data{
		Tickets:        tickets,
		Bunches:        bunches,
		Mode:           "raffle",
		StartTimestamp: 500,
		EndTimestamp:   900,
		Features:       []string{},
	}
	_, _ = CreateDraw(test)
	duration := time.Since(start)
	if duration.Seconds() > 29.0 {
		t.Errorf("Computation took too long: %f seconds", duration.Seconds())
	}
}
func Test_9T_3B_DATA_TAGS(t *testing.T) {
	test := Data{
		Tickets: []types.Ticket{
			{Data: "a", Id: "1", Owner: "o1"},
			{Data: "a", Id: "2", Owner: "o1"},
			{Data: "a", Id: "3", Owner: "o1"},
			{Data: "a", Id: "4", Owner: "o1"},
			{Data: "a", Id: "5", Owner: "o2"},
			{Data: "a", Id: "6", Owner: "o3"},
			{Data: "a", Id: "7", Owner: "o4"},
			{Data: "a", Id: "8", Owner: "o5"},
			{Data: "a", Id: "9", Owner: "o5"},
		},
		Bunches:  []types.Bunch{{Id: "toy1", Quantity: 1, Tags: []string{"toy"}}, {Id: "toy2", Quantity: 1, Tags: []string{"toy"}}, {Id: "travel1", Quantity: 1, Tags: []string{"travel"}}},
		Mode:     "raffle",
		Features: []string{},
		Tags:     []string{"toy"},
	}
	d, _ := CreateDraw(test)

	if len(d.Winners) != 2 {
		t.Errorf("Only two bunches should have been won")
	}

}
