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

## utilizzo come eseguibile
### build
```
$ go build -o nombre ./cmd
```
### utilizzo
```
  -spazi
        spazi tra gli ordini di grandezza (esempio: "trentamila duecento quarantre")
  -un
        accorcia l'uno (esempio: "ventunmila")
```

## utilizzo come libreria
```
$ go get github.com/lorenzobotti/nombre
$ go mod tidy
```

la libreria esporta le funzioni: 
 - `func Convert(int number) string`
 - `func ConvertWithOptions(int number, Options options) string`
 - `func DefaultOptions() Options`

## note di utilizzo
nombre segue le [regole dell'Academia della Crusca](https://accademiadellacrusca.it/it/consulenza/quarantaquattro-gatti-in-fila-per-sei-col-resto-di-due-o-di-quando-e-come-scrivere-i-numeri-in-lettere/1077) per la scrittura dei numeri in lettere.