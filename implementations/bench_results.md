# Результаты тестирования

# Как показаны результаты

Все решения приводятся с краткой аннотацией в порядке их появления.

В каждой группе бенчмарка указаны победитель (или победители, если значения очень близкие)

Выводы делайте сами

0. comdiv (fib[]+math.Pow) [ссылка](./comdiv/initial_solution.go)
1. gubkin (fib+BinaryPow) [ссылка](./gubkin/initial_binary_pow.go)
2. gubkin (fib[]+BinaryPow) [ссылка](./gubkin/optimized_with_fib_array.go)
3. comdiv  (fib[]+BinaryPow(optimal)) [ссылка](./comdiv/with_gubkin_binary_pow.go)
4. tetyuev (linear) [ссылка](./tetyuev/initial_version.go)
5. pirozhkov  (fib+math.Pow) [ссылка](./pirozhkov/initial_version.go)
6. tetyuev (fib[]+math.Pow) [ссылка](./tetyuev/fib_version.go)
7. harisov (fib+SqrtPower) [ссылка](./harisov/main_version.go)
8. gavrilov (linear) [ссылка](./gavrilov/main_impl.go)
9. metelkin (linear) [ссылка](./metelkin/main_impl.go)
10. abzaev (fib+math.Pow)  [ссылка](./abzaev/main_version.go)
11. gubkin (fib[]+CachedPow)  [ссылка](./gubkin/cached_pow1.go)
12. gubkin (CachedPow)  [ссылка](./gubkin/plain.go)
13. comdiv (fib[]+CachedPow+BinaryPowDual) [ссылка](./comdiv/try_to_make_best_function.go)

# Бенчмарк на предподготовленных кейсах (который основной использовался)

| n      | author    | ver | method                            | time  |
|--------|-----------|-----|-----------------------------------|-------|
| 0      | comdiv    | v1  | `(fib[]+math.Pow)`                | 192.3 |
| 1      | gubkin    | v1  | `(fib+BinaryPow)`                 | 67.73 |
| 2      | gubkin    | v2  | `(fib[]+BinaryPow)`               | 54.82 |
| 3      | comdiv    | v2  | `(fib[]+BinaryPow(optimal))`      | 51.33 |
| 4      | tetyuev   | v1  | `(linear)`                        | 35.60 |
| 5      | pirozhkov | v1  | `(fib+math.Pow)`                  | 229.6 |
| 6      | tetyuev   | v2  | `(fib[]+math.Pow)`                | 209.6 |
| 7      | harisov   | v1  | `(fib+SqrtPower)`                 | 67.84 |
| 8      | gavrilov  | v1  | `(linear)`                        | 36.21 |
| 9      | metelkin  | v1  | `(linear)`                        | 36.21 |
| 10     | abzaev    | v1  | `(fib+math.Pow)`                  | 213.3 |
| **11** | gubkin    | v3  | `(fib[]+CachedPow)`               | 29.43 |
| **12** | gubkin    | v4  | `(CachedPow)`                     | 25.76 |
| **13** | comdiv    | v3  | `(fib[]+CachedPow+BinaryPowDual)` | 27.93 |




# Бенчмарк на предподготовленных данных на большой степени 77

| n      | author    | ver | method                            | time  |
|--------|-----------|-----|-----------------------------------|-------|
| 0      | comdiv    | v1  | `(fib[]+math.Pow)`                | 185.0 |
| 1      | gubkin    | v1  | `(fib+BinaryPow)`                 | 156.0 |
| 2      | gubkin    | v2  | `(fib[]+BinaryPow)`               | 117.3 |
| 3      | comdiv    | v2  | `(fib[]+BinaryPow(optimal))`      | 108.7 |
| 4      | tetyuev   | v1  | `(linear)`                        | 83.80 |
| 5      | pirozhkov | v1  | `(fib+math.Pow)`                  | 238.3 |
| 6      | tetyuev   | v2  | `(fib[]+math.Pow)`                | 188.0 |
| 7      | harisov   | v1  | `(fib+SqrtPower)`                 | 129.3 |
| 8      | gavrilov  | v1  | `(linear)`                        | 83.34 |
| 9      | metelkin  | v1  | `(linear)`                        | 83.79 |
| 10     | abzaev    | v1  | `(fib+math.Pow)`                  | 238.1 |
| 11     | gubkin    | v3  | `(fib[]+CachedPow)`               | 112.8 |
| 12     | gubkin    | v4  | `(CachedPow)`                     | 111.4 |
| **13** | comdiv    | v3  | `(fib[]+CachedPow+BinaryPowDual)` | 74.34 |





# Рандомный тест на малых степенях 2 .. 20

| n     | author    | ver | method                            | time  |
|-------|-----------|-----|-----------------------------------|-------|
| 0     | comdiv    | v1  | `(fib[]+math.Pow)`                | 150.6 |
| 1     | gubkin    | v1  | `(fib+BinaryPow)`                 | 65.18 |
| 2     | gubkin    | v2  | `(fib[]+BinaryPow)`               | 50.75 |
| 3     | comdiv    | v2  | `(fib[]+BinaryPow(optimal))`      | 48.15 |
| **4** | tetyuev   | v1  | `(linear)`                        | 36.57 |
| 5     | pirozhkov | v1  | `(fib+math.Pow)`                  | 170.3 |
| 6     | tetyuev   | v2  | `(fib[]+math.Pow)`                | 162.8 |
| 7     | harisov   | v1  | `(fib+SqrtPower)`                 | 66.02 |
| **8** | gavrilov  | v1  | `(linear)`                        | 36.42 |
| **9** | metelkin  | v1  | `(linear)`                        | 36.65 |
| 10    | abzaev    | v1  | `(fib+math.Pow)`                  | 160.4 |
| 11    | gubkin    | v3  | `(fib[]+CachedPow)`               | 45.56 |
| 12    | gubkin    | v4  | `(CachedPow)`                     | 43.91 |
| 13    | comdiv    | v3  | `(fib[]+CachedPow+BinaryPowDual)` | 41.81 |





# Рандомный тест на средних степенях 21 .. 40

| n     | author    | ver | method                            | time  |
|-------|-----------|-----|-----------------------------------|-------|
| 0     | comdiv    | v1  | `(fib[]+math.Pow)`                | 235.3 |
| 1     | gubkin    | v1  | `(fib+BinaryPow)`                 | 103.7 |
| 2     | gubkin    | v2  | `(fib[]+BinaryPow)`               | 87.85 |
| 3     | comdiv    | v2  | `(fib[]+BinaryPow(optimal))`      | 84.20 |
| **4** | tetyuev   | v1  | `(linear)`                        | 54.83 |
| 5     | pirozhkov | v1  | `(fib+math.Pow)`                  | 288.1 |
| 6     | tetyuev   | v2  | `(fib[]+math.Pow)`                | 268.5 |
| 7     | harisov   | v1  | `(fib+SqrtPower)`                 | 121.8 |
| **8** | gavrilov  | v1  | `(linear)`                        | 57.23 |
| **9** | metelkin  | v1  | `(linear)`                        | 56.37 |
| 10    | abzaev    | v1  | `(fib+math.Pow)`                  | 267.1 |
| 11    | gubkin    | v3  | `(fib[]+CachedPow)`               | 60.01 |
| 12    | gubkin    | v4  | `(CachedPow)`                     | 64.88 |
| 13    | comdiv    | v3  | `(fib[]+CachedPow+BinaryPowDual)` | 72.65 |

# Рандомный тест на средних степенях 70 .. 77

| n     | author    | ver | method                            | time  |
|-------|-----------|-----|-----------------------------------|-------|
| 0     | comdiv    | v1  | `(fib[]+math.Pow)`                | 223.8 |
| 1     | gubkin    | v1  | `(fib+BinaryPow)`                 | 180.8 |
| 2     | gubkin    | v2  | `(fib[]+BinaryPow)`               | 157.8 |
| 3     | comdiv    | v2  | `(fib[]+BinaryPow(optimal))`      | 146.3 |
| **4** | tetyuev   | v1  | `(linear)`                        | 89.69 |
| 5     | pirozhkov | v1  | `(fib+math.Pow)`                  | 277.2 |
| 6     | tetyuev   | v2  | `(fib[]+math.Pow)`                | 232.4 |
| 7     | harisov   | v1  | `(fib+SqrtPower)`                 | 165.5 |
| **8** | gavrilov  | v1  | `(linear)`                        | 92.61 |
| **9** | metelkin  | v1  | `(linear)`                        | 91.72 |
| 10    | abzaev    | v1  | `(fib+math.Pow)`                  | 271.2 |
| 11    | gubkin    | v3  | `(fib[]+CachedPow)`               | 120.1 |
| 12    | gubkin    | v4  | `(CachedPow)`                     | 118.0 |
| 13    | comdiv    | v3  | `(fib[]+CachedPow+BinaryPowDual)` | 119.5 |

# Выводы

Пока побеждает простота!
Простое элементарное линейное решение показывает себя лучше на всех режимах кроме
кэшированного варианта на самых больших степенях.
Скорее всего это связано с более эффективной работой кэша в вариантах 11,12,13 (я имею в виду кэша процессора)
и именно если прогоняются по кругу одни и те же входные параметры.

Пока победила линейка!
Но с другой стороны посмотрите на варианты 11, 12, 13 и вообще на то какие изыскания провел Андрей, начиная 
со своего первого варианта (номер 1)

# Сырой ответ

```text
goos: linux
goarch: amd64
pkg: github.com/comdiv/task_func_optimize_base_go/implementations
cpu: AMD Ryzen 5 2600 Six-Core Processor            
BenchmarkDefault/0._comdiv____v1_(fib[]+math.Pow)                6215539               192.3 ns/op
BenchmarkDefault/1._gubkin____v1_(fib+BinaryPow)                17725952                67.73 ns/op
BenchmarkDefault/2._gubkin____v2_(fib[]+BinaryPow)              21919173                54.82 ns/op
BenchmarkDefault/3._comdiv____v2_(fib[]+BinaryPow(optimal))             23354980                51.33 ns/op
BenchmarkDefault/4._tetyuev___v1_(linear)                               33726814                35.60 ns/op
BenchmarkDefault/5._pirozhkov_v1_(fib+math.Pow)                          5319892               229.6 ns/op
BenchmarkDefault/6._tetyuev___v2_(fib[]+math.Pow)                        5744156               209.6 ns/op
BenchmarkDefault/7._harisov___v1_(fib+SqrtPower)                        17661408                67.84 ns/op
BenchmarkDefault/8._gavrilov__v1_(linear)                               33320715                36.21 ns/op
BenchmarkDefault/9._metelkin__v1_(linear)                               33789726                36.21 ns/op
BenchmarkDefault/10._abzaev____v1_(fib+math.Pow)                         5497677               213.3 ns/op
BenchmarkDefault/11._gubkin____v3_(fib[]+CachedPow)                     41216056                29.43 ns/op
BenchmarkDefault/12._gubkin____v4_(CachedPow)                           46439714                25.76 ns/op
BenchmarkDefault/13._comdiv____v3_(fib[]+CachedPow+BinaryPowDual)       42669422                27.93 ns/op
BenchmarkRandom_Low_2_20/0._comdiv____v1_(fib[]+math.Pow)                7940508               150.6 ns/op
BenchmarkRandom_Low_2_20/1._gubkin____v1_(fib+BinaryPow)                18447506                65.18 ns/op
BenchmarkRandom_Low_2_20/2._gubkin____v2_(fib[]+BinaryPow)              23539304                50.75 ns/op
BenchmarkRandom_Low_2_20/3._comdiv____v2_(fib[]+BinaryPow(optimal))     24682069                48.15 ns/op
BenchmarkRandom_Low_2_20/4._tetyuev___v1_(linear)                       33422211                36.57 ns/op
BenchmarkRandom_Low_2_20/5._pirozhkov_v1_(fib+math.Pow)                  7053662               170.3 ns/op
BenchmarkRandom_Low_2_20/6._tetyuev___v2_(fib[]+math.Pow)                7378737               162.8 ns/op
BenchmarkRandom_Low_2_20/7._harisov___v1_(fib+SqrtPower)                18406586                66.02 ns/op
BenchmarkRandom_Low_2_20/8._gavrilov__v1_(linear)                       34028223                36.42 ns/op
BenchmarkRandom_Low_2_20/9._metelkin__v1_(linear)                       33191913                36.65 ns/op
BenchmarkRandom_Low_2_20/10._abzaev____v1_(fib+math.Pow)                 7472377               160.4 ns/op
BenchmarkRandom_Low_2_20/11._gubkin____v3_(fib[]+CachedPow)             26427739                45.56 ns/op
BenchmarkRandom_Low_2_20/12._gubkin____v4_(CachedPow)                   28579252                43.91 ns/op
BenchmarkRandom_Low_2_20/13._comdiv____v3_(fib[]+CachedPow+BinaryPowDual)               29219487                41.81 ns/op
BenchmarkRandom_High_21_40/0._comdiv____v1_(fib[]+math.Pow)                              5104814               235.3 ns/op
BenchmarkRandom_High_21_40/1._gubkin____v1_(fib+BinaryPow)                              11306476               103.7 ns/op
BenchmarkRandom_High_21_40/2._gubkin____v2_(fib[]+BinaryPow)                            13642075                87.85 ns/op
BenchmarkRandom_High_21_40/3._comdiv____v2_(fib[]+BinaryPow(optimal))                   14282408                84.20 ns/op
BenchmarkRandom_High_21_40/4._tetyuev___v1_(linear)                                     21979848                54.83 ns/op
BenchmarkRandom_High_21_40/5._pirozhkov_v1_(fib+math.Pow)                                4534417               288.1 ns/op
BenchmarkRandom_High_21_40/6._tetyuev___v2_(fib[]+math.Pow)                              4615033               268.5 ns/op
BenchmarkRandom_High_21_40/7._harisov___v1_(fib+SqrtPower)                              10793326               121.8 ns/op
BenchmarkRandom_High_21_40/8._gavrilov__v1_(linear)                                     21126110                57.23 ns/op
BenchmarkRandom_High_21_40/9._metelkin__v1_(linear)                                     21853443                56.37 ns/op
BenchmarkRandom_High_21_40/10._abzaev____v1_(fib+math.Pow)                               4682384               267.1 ns/op
BenchmarkRandom_High_21_40/11._gubkin____v3_(fib[]+CachedPow)                           25312566                60.01 ns/op
BenchmarkRandom_High_21_40/12._gubkin____v4_(CachedPow)                                 19608717                64.88 ns/op
BenchmarkRandom_High_21_40/13._comdiv____v3_(fib[]+CachedPow+BinaryPowDual)             16906556                72.65 ns/op
BenchmarkRandom_Large_70_77/0._comdiv____v1_(fib[]+math.Pow)                             5213182               223.8 ns/op
BenchmarkRandom_Large_70_77/1._gubkin____v1_(fib+BinaryPow)                              6711955               180.8 ns/op
BenchmarkRandom_Large_70_77/2._gubkin____v2_(fib[]+BinaryPow)                            8013178               157.8 ns/op
BenchmarkRandom_Large_70_77/3._comdiv____v2_(fib[]+BinaryPow(optimal))                   7725942               146.3 ns/op
BenchmarkRandom_Large_70_77/4._tetyuev___v1_(linear)                                    13471202                89.69 ns/op
BenchmarkRandom_Large_70_77/5._pirozhkov_v1_(fib+math.Pow)                               4329211               277.2 ns/op
BenchmarkRandom_Large_70_77/6._tetyuev___v2_(fib[]+math.Pow)                             5165058               232.4 ns/op
BenchmarkRandom_Large_70_77/7._harisov___v1_(fib+SqrtPower)                              6922506               165.5 ns/op
BenchmarkRandom_Large_70_77/8._gavrilov__v1_(linear)                                    13201156                92.61 ns/op
BenchmarkRandom_Large_70_77/9._metelkin__v1_(linear)                                    13254768                91.72 ns/op
BenchmarkRandom_Large_70_77/10._abzaev____v1_(fib+math.Pow)                              4423490               271.2 ns/op
BenchmarkRandom_Large_70_77/11._gubkin____v3_(fib[]+CachedPow)                          10067815               120.1 ns/op
BenchmarkRandom_Large_70_77/12._gubkin____v4_(CachedPow)                                10170728               118.0 ns/op
BenchmarkRandom_Large_70_77/13._comdiv____v3_(fib[]+CachedPow+BinaryPowDual)            10014868               119.5 ns/op
BenchmarkFixed_77/0._comdiv____v1_(fib[]+math.Pow)                                       6597138               185.0 ns/op
BenchmarkFixed_77/1._gubkin____v1_(fib+BinaryPow)                                        7689319               156.0 ns/op
BenchmarkFixed_77/2._gubkin____v2_(fib[]+BinaryPow)                                     10458312               117.3 ns/op
BenchmarkFixed_77/3._comdiv____v2_(fib[]+BinaryPow(optimal))                            10971444               108.7 ns/op
BenchmarkFixed_77/4._tetyuev___v1_(linear)                                              14335090                83.80 ns/op
BenchmarkFixed_77/5._pirozhkov_v1_(fib+math.Pow)                                         5037984               238.3 ns/op
BenchmarkFixed_77/6._tetyuev___v2_(fib[]+math.Pow)                                       6383760               188.0 ns/op
BenchmarkFixed_77/7._harisov___v1_(fib+SqrtPower)                                        9288810               129.3 ns/op
BenchmarkFixed_77/8._gavrilov__v1_(linear)                                              14404894                83.34 ns/op
BenchmarkFixed_77/9._metelkin__v1_(linear)                                              14319126                83.79 ns/op
BenchmarkFixed_77/10._abzaev____v1_(fib+math.Pow)                                        5086422               238.1 ns/op
BenchmarkFixed_77/11._gubkin____v3_(fib[]+CachedPow)                                    10659434               112.8 ns/op
BenchmarkFixed_77/12._gubkin____v4_(CachedPow)                                          10783416               111.4 ns/op
BenchmarkFixed_77/13._comdiv____v3_(fib[]+CachedPow+BinaryPowDual)                      15502162                74.34 ns/op

```