package protection

import (
	"encoding/json"
	"testing"
)

type testCase struct {
	Got  string
	Want Packet
}

func TestUnmarshal(t *testing.T) {
	templateTimestamp := Timestamp{0, 0, 0, 0, 0, 0, 0}
	templateDetails := Details{"test", "", nil}
	templatePacket := Packet{0, 0, templateTimestamp, templateDetails}

	tests := map[string]testCase{
		"out_of_bounds": {
			`{
	"boardId": 0,
	"timestamp": {
		"counter": 0,
		"second": 0,
		"minute": 0,
		"hour": 0,
		"day": 0,
		"month": 0,
		"year": 0
	} ,
	"protection": {
		"name": "test",
		"type": "OUT_OF_BOUNDS",
		"data": {
			"value": 0,
			"bounds": [-1, 1]
		}
	}
}`,
			withDetails(t, templatePacket, OutOfBounds{0, [2]float64{-1, 1}}),
		},
		"lower_bound": {
			`{
	"boardId": 0,
	"timestamp": {
		"counter": 0,
		"second": 0,
		"minute": 0,
		"hour": 0,
		"day": 0,
		"month": 0,
		"year": 0
	} ,
	"protection": {
		"name": "test",
		"type": "LOWER_BOUND",
		"data": {
			"value": 0,
			"bound": 1
		}
	}
}`,
			withDetails(t, templatePacket, LowerBound{0, 1}),
		},
		"upper_bound": {
			`{
	"boardId": 0,
	"timestamp": {
		"counter": 0,
		"second": 0,
		"minute": 0,
		"hour": 0,
		"day": 0,
		"month": 0,
		"year": 0
	} ,
	"protection": {
		"name": "test",
		"type": "UPPER_BOUND",
		"data": {
			"value": 0,
			"bound": -1
		}
	}
}`,
			withDetails(t, templatePacket, UpperBound{0, -1}),
		},
		"equals": {
			`{
	"boardId": 0,
	"timestamp": {
		"counter": 0,
		"second": 0,
		"minute": 0,
		"hour": 0,
		"day": 0,
		"month": 0,
		"year": 0
	} ,
	"protection": {
		"name": "test",
		"type": "EQUALS",
		"data": {
			"value": 0
		}
	}
}`,
			withDetails(t, templatePacket, Equals{0}),
		},
		"not_equals": {
			`{
	"boardId": 0,
	"timestamp": {
		"counter": 0,
		"second": 0,
		"minute": 0,
		"hour": 0,
		"day": 0,
		"month": 0,
		"year": 0
	} ,
	"protection": {
		"name": "test",
		"type": "NOT_EQUALS",
		"data": {
			"value": 0,
			"want": 1
		}
	}
}`,
			withDetails(t, templatePacket, NotEquals{0, 1}),
		},
		"time_accumulation": {
			`{
	"boardId": 0,
	"timestamp": {
		"counter": 0,
		"second": 0,
		"minute": 0,
		"hour": 0,
		"day": 0,
		"month": 0,
		"year": 0
	} ,
	"protection": {
		"name": "test",
		"type": "TIME_ACCUMULATION",
		"data": {
			"value": 0,
			"bound": 1,
			"timelimit": 1
		}
	}
}`,
			withDetails(t, templatePacket, TimeAccumulation{0, 1, 1}),
		},
		"error_handler": {
			`{
	"boardId": 0,
	"timestamp": {
		"counter": 0,
		"second": 0,
		"minute": 0,
		"hour": 0,
		"day": 0,
		"month": 0,
		"year": 0
	} ,
	"protection": {
		"name": "test",
		"type": "ERROR_HANDLER",
		"data": "error"
	}
}`,
			withDetails(t, templatePacket, ErrorHandler("error")),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var packet Packet
			err := json.Unmarshal([]byte(test.Got), &packet)
			if err != nil {
				t.Fatal(err)
			}

			if packet != test.Want {
				t.Fatalf("Expected:\n%#v\nGot:\n%#v", test.Want, packet)
			}
		})
	}
}

func withDetails(t *testing.T, packet Packet, data Data) Packet {
	t.Helper()

	packet.Protection.Type = data.Type()
	packet.Protection.Data = data

	return packet
}
