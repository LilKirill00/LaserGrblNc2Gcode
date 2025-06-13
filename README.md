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
```
G90 (use absolute coordinates)
G0 X0 Y0 F1000
M3 S0
G1 X0.1 S3
X0.15 S15
G0 X26.35 Y0 S0
G1 X26.4 S11
X26.5 S3
S0
G0 X26.5 Y0.05 S0
```
### Выход
```
G90 (use absolute coordinates)
G0 X0 Y0 F1000
M3 S0
M3 S3
G1 X0.1
M3 S15
G1 X0.15
M3 S0
G0 X26.35 Y0
M3 S11
G1 X26.4
M3 S3
G1 X26.5
M3 S0
M3 S0
G0 X26.5 Y0.05
```
