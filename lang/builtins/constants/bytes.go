package constants

import "klc/lang/builtins"

var B_ = builtins.WithDoc(builtins.NewNumber(0.125), "Bit, a unit of information equal to 1/8 bytes.")
var KB_ = builtins.WithDoc(builtins.NewNumber(128), "Kilobit, a unit of information equal to 1024 bits.")
var MB_ = builtins.WithDoc(builtins.NewNumber(131072), "Megabit, a unit of information equal to 1024 kilobits.")
var GB_ = builtins.WithDoc(builtins.NewNumber(134217728), "Gigabit, a unit of information equal to 1024 megabits.")
var TB_ = builtins.WithDoc(builtins.NewNumber(137438953472), "Terabit, a unit of information equal to 1024 gigabits.")
var PB_ = builtins.WithDoc(builtins.NewNumber(140737488355328), "Petabit, a unit of information equal to 1024 terabits.")
var B = builtins.WithDoc(builtins.NewNumber(1), "Byte, a unit of information equal to 8 bits.")
var KB = builtins.WithDoc(builtins.NewNumber(1024), "Kilobyte, a unit of information equal to 1024 bytes.")
var MB = builtins.WithDoc(builtins.NewNumber(1048576), "Megabyte, a unit of information equal to 1024 kilobytes.")
var GB = builtins.WithDoc(builtins.NewNumber(1073741824), "Gigabyte, a unit of information equal to 1024 megabytes.")
var TB = builtins.WithDoc(builtins.NewNumber(1099511627776), "Terabyte, a unit of information equal to 1024 gigabytes.")
var PB = builtins.WithDoc(builtins.NewNumber(1125899906842624), "Petabyte, a unit of information equal to 1024 terabytes.")
