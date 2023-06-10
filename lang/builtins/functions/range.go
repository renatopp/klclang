package functions

import (
	"klc/lang/builtins"
	"klc/lang/obj"
	"math"
)

func range_(args ...obj.Object) obj.Object {
	start := 0.0
	stop := 0.0
	step := 1.0
	direction := 1.0

	if len(args) > 1 {
		start = args[0].AsNumber()
		stop = args[1].AsNumber()
	} else {
		stop = args[0].AsNumber()
	}

	if len(args) > 2 {
		step = args[2].AsNumber()
	}

	if start > stop {
		direction = -1
	}
	step = math.Abs(step) * direction

	var result []float64
	for i := start; i*direction < stop*direction; i += step {
		result = append(result, i)
	}

	return builtins.NewNumberList(result...)
}

var Range = builtins.WithDoc(
	builtins.NewFunction(range_,
		builtins.NewParam("start", nil, false),
		builtins.NewParam("stop", builtins.NewNumber(0), false),
		builtins.NewParam("step", builtins.NewNumber(1), false),
	),
	`Generates a sequence of numbers from 'start' to 'stop' (exclusive) with an optional 'step' value.	`,
)
