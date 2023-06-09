package constants

import "klc/lang/builtins"

var KM = builtins.WithDoc(
	builtins.NewNumber(1000),
	"Kilometer, a unit of length equal to 1000 meters.",
)
