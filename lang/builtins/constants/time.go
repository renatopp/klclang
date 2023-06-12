package constants

import "klc/lang/builtins"

var MS = builtins.WithDoc(builtins.NewNumber(0.001), "Millisecond, a unit of time equal to 1/1000 seconds.")
var MICROS = builtins.WithDoc(builtins.NewNumber(0.000001), "Microsecond, a unit of time equal to 1/1000000 seconds.")
var NS = builtins.WithDoc(builtins.NewNumber(0.000000001), "Nanosecond, a unit of time equal to 1/1000000000 seconds.")
var PS = builtins.WithDoc(builtins.NewNumber(0.000000000001), "Picosecond, a unit of time equal to 1/1000000000000 seconds.")
var S = builtins.WithDoc(builtins.NewNumber(1), "Second, a unit of time equal to 1 seconds.")
var MIN = builtins.WithDoc(builtins.NewNumber(60), "Minute, a unit of time equal to 60 seconds.")
var H = builtins.WithDoc(builtins.NewNumber(3600), "Hour, a unit of time equal to 3600 seconds.")
var DAY = builtins.WithDoc(builtins.NewNumber(86400), "Day, a unit of time equal to 86400 seconds.")
var WEEK = builtins.WithDoc(builtins.NewNumber(604800), "Week, a unit of time equal to 604800 seconds.")
var MONTH = builtins.WithDoc(builtins.NewNumber(2629800), "Month, a unit of time equal to 2629800 seconds.")
var YEAR = builtins.WithDoc(builtins.NewNumber(31557600), "Year, a unit of time equal to 31557600 seconds.")
var DECADE = builtins.WithDoc(builtins.NewNumber(315576000), "Decade, a unit of time equal to 315576000 seconds.")
var CENTURY = builtins.WithDoc(builtins.NewNumber(3155760000), "Century, a unit of time equal to 3155760000 seconds.")
var MILLENNIUM = builtins.WithDoc(builtins.NewNumber(31557600000), "Millennium, a unit of time equal to 31557600000 seconds.")
