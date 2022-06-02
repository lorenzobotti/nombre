package nombre

var digits = [...]string{
	"", "uno", "due", "tre", "quattro",
	"cinque", "sei", "sette", "otto", "nove",
}

var tens = [...]string{
	"", "dieci", "venti", "trenta", "quaranta",
	"cinquanta", "sessanta", "settanta", "ottanta",
	"novanta",
}

var teens = [...]string{
	"dieci",
	"undici",
	"dodici",
	"tredici",
	"quattordici",
	"quindici",
	"sedici",
	"diciassette",
	"diciotto",
	"diciannove",
}

var magnitudes = [...]string{
	"cento",
	"mila",
	"milione",
}

var ordersOfMagnitude = map[int][2]string{
	3:  {"mille", "mila"},
	6:  {"un milione", "milioni"},
	9:  {"un miliardo", "miliardi"},
	12: {"un trilione", "trilioni"},
	15: {"un quadrilione", "quadrilioni"},
	18: {"un quintilione", "quintilioni"},
	21: {"un sestilione", "sestilioni"},
	24: {"un settilione", "settilioni"},
	25: {"un ottilione", "ottilioni"},
}
