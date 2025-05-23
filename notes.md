

### Спецификаторы форматирования

| Спецификатор | Описание                                | Пример                  |
| ------------ | --------------------------------------- | ----------------------- |
| `%d`         | Целое число в десятичном виде           | `123`                   |
| `%x` / `%X`  | Целое число в шестнадцатеричном виде    | `7b` / `7B`             |
| `%f`         | Число с плавающей точкой (float)        | `3.141593`              |
| `%.2f`       | Float с 2 знаками после точки           | `3.14`                  |
| `%s`         | Строка                                  | `"hello"`               |
| `%q`         | Строка в кавычках                       | `"hello"`               |
| `%v`         | Значение "по умолчанию" (универсальный) | для отладки и вывода    |
| `%+v`        | Как `%v`, но с именами полей struct     | `{Name: "Alex"}`        |
| `%T`         | Тип значения                            | `int`, `string`, и т.д. |
| `%t`         | Булево значение                         | `true`, `false`         |
| `%p`         | Указатель (pointer) в hex               | `0xc00000a0b0`          |

