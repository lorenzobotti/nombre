# nombre

nombre formatta numeri in italiano. immàginatelo come `itoa`, però in lettere.
```
$ nombre 0 1 100 -6 1984 8920942
zero
uno
cento
menosei
millenovecentoottantaquattro
ottomilioninovecentoventimilanovecentoquarantadue
```

## buildare come eseguibile
```
$ go build -o nombre ./cmd
```

## utilizzo come libreria
```
$ go get github.com/lorenzobotti/nombre
$ go mod tidy
```

La libreria esporta solo una funzione: `func Convert(int number) string`