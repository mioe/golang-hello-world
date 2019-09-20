package main

import (
  "fmt"
  "math"
  "math/cmplx"
  "math/rand"
)

func basicTypes() {
  var b bool = false
/*
  Оператор var объявляет список переменных; как и в случае со
  списком аргументов функции, тип указывается последним.
 */
  var s string = `hello`
  var i int = 1
  // int  int8  int16  int32  int64
  // uint uint8 uint16 uint32 uint64 uintptr

  // byte псевдоним для uint8
  // rune псевдоним для int32
  // | представляет Unicode код

  var f32 float32 = 3.14
  var f64 float64 = float64(f32)
  var r rune = 'は'
  fmt.Println(`basicTypes:`)
  fmt.Println("\t", b, s, i, f32, f64, r)
/*
  Типы int, uint, и uintptr обычно размером 32 бита на 32-битных
  системах и 64 бита на 64-битных системах. Когда вам надо
  целочисленное значение, то следует использовать int, если
  у вас нет особой причины использовать целочисленное значение
  конкретного размера или беззнаковое значение.
 */
}

func shortVars() {
  b := false
  s := `hello`
  i := 1
  f64 := 3.14
  r := 'は'
  fmt.Println(`shortVars:`)
  fmt.Println("\t", b, s, i, f64, r)
/*
  Внутри функции, краткий оператор присваивания := с неявным
  типом может быть использован вместо объявления с помощью var.
  Вне функции, каждый оператор начинается с ключевого слова (var,
  func, и так далее), и поэтому конструкция := недоступна.
 */
}

func complexVar() {
  // package - math/cmplx
  var z complex128 = cmplx.Sqrt(-5 + 12i)
  fmt.Println(`complexVar:`)
  fmt.Println("\t", z)
}

func piVar() {
  // package - math
  pi := math.Pi
  fmt.Println(`piVar:`)
  fmt.Println("\t", pi)
/*
  Когда вы импортируете пакет, то можете ссылаться только на его
  экспортируемые имена. Любые не экспортированные имена недоступны
  за пределами пакета.
 */
}

func randVar() {
  // package - math/rand
  rand := rand.Intn(100)
  fmt.Println(`randVar:`)
  fmt.Println("\t", rand)
/*
  Замечание: окружение, в котором выполняются эти программы,
  заранее предопределено, так что каждый раз,когда вы
  запускаете пример rand.Intn будет возвращать одно и то же
  число. (Чтобы получить другое число, инициализируйте
  генератор случайных чисел; подробнее rand.Seed.)
 */
}

func main() {
  basicTypes()
  shortVars()
  complexVar()
  piVar()
  randVar()
}
