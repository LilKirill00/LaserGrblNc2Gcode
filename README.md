# LaserGrblNc2Gcode

Конвертер для преобразования кодов из LaserGRBL в GCode коды для Marlin

## Пример преобразования
### Пример команды
```sh
go run main.go -i=src/sample.nc -o=dst/sample.gcode
```
Где:
- -i - входной файл полученный из LaserGRBL
- -o - файл который получаем

### Вход
![[src/sample.nc]]
### Выход
![[dst/sample.gcode]]
