package constants

import "klc/lang/builtins"

var MG = builtins.WithDoc(builtins.NewNumber(0.000001), "Milligram, a unit of mass equal to 1/1000 grams.")
var CG = builtins.WithDoc(builtins.NewNumber(0.00001), "Centigram, a unit of mass equal to 1/100 grams.")
var DG = builtins.WithDoc(builtins.NewNumber(0.0001), "Decigram, a unit of mass equal to 1/10 grams.")
var G = builtins.WithDoc(builtins.NewNumber(0.001), "Gram, a unit of mass equal to 1 grams.")
var KG = builtins.WithDoc(builtins.NewNumber(1), "Kilogram, a unit of mass equal to 1000 grams.")
var T = builtins.WithDoc(builtins.NewNumber(1000), "Ton, a unit of mass equal to 1000 kilograms.")
var OZ = builtins.WithDoc(builtins.NewNumber(0.028349523125), "Ounce, a unit of mass equal to 1/16 pounds.")
var LB = builtins.WithDoc(builtins.NewNumber(0.45359237), "Pound, a unit of mass equal to 16 ounces.")
