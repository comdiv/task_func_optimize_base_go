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
9. abzaev (fib+math.Pow)  [ссылка](./abzaev/main_version.go)
10. gubkin (fib[]+CachedPow)  [ссылка](./gubkin/cached_pow1.go)
11. gubkin (CachedPow)  [ссылка](./gubkin/plain.go)
12. comdiv (fib[]+CachedPow+BinaryPowDual) [ссылка](./comdiv/try_to_make_best_function.go)

# Бенчмарк на предподготовленных кейсах (который основной использовался)

| n        | author     | ver    | method                          | time        |
|----------|------------|--------|---------------------------------|-------------|
| 0        | comdiv     | v1     | (fib[]+math.Pow)                | 192.5       |
| 1        | gubkin     | v1     | (fib+BinaryPow)                 | 67.68       |
| 2        | gubkin     | v2     | (fib[]+BinaryPow)               | 55.28       |
| 3        | comdiv     | v2     | (fib[]+BinaryPow(optimal))      | 52.28       |
| 4        | tetyuev    | v1     | (linear)                        | 35.80       |
| 5        | pirozhkov  | v1     | (fib+math.Pow)                  | 227.1       |
| 6        | tetyuev    | v2     | (fib[]+math.Pow)                | 201.5       |
| 7        | harisov    | v1     | (fib+SqrtPower)                 | 68.22       |
| 8        | gavrilov   | v1     | (linear)                        | 35.54       |
| 9        | abzaev     | v1     | (fib+math.Pow)                  | 215.0       |
| ***10*** | gubkin     | v3     | (fib[]+CachedPow)               | ***29.80*** |
| **11**   | **gubkin** | **v4** | **(CachedPow)**                 | **25.90**   |
| ***12*** | comdiv     | v3     | (fib[]+CachedPow+BinaryPowDual) | ***28.33*** |

# Бенчмарк на предподготовленных данных на большой степени 77

| n      | author     | ver    | method                                | time      |
|--------|------------|--------|---------------------------------------|-----------|
| 0      | comdiv     | v1     | `(fib[]+math.Pow)`                    | 184.5     |
| 1      | gubkin     | v1     | `(fib+BinaryPow)`                     | 156.8     |
| 2      | gubkin     | v2     | `(fib[]+BinaryPow)`                   | 115.4     |
| 3      | comdiv     | v2     | `(fib[]+BinaryPow(optimal))`          | 109.2     |
| 4      | tetyuev    | v1     | `(linear)`                            | 84.50     |
| 5      | pirozhkov  | v1     | `(fib+math.Pow)`                      | 240.5     |
| 6      | tetyuev    | v2     | `(fib[]+math.Pow)`                    | 190.3     |
| 7      | harisov    | v1     | `(fib+SqrtPower)`                     | 130.0     |
| 8      | gavrilov   | v1     | `(linear)`                            | 83.77     |
| 9      | abzaev     | v1     | `(fib+math.Pow)`                      | 235.9     |
| 10     | gubkin     | v3     | `(fib[]+CachedPow)`                   | 112.8     |
| 11     | gubkin     | v4     | `(CachedPow)`                         | 111.7     |
| **12** | **comdiv** | **v3** | **`(fib[]+CachedPow+BinaryPowDual)`** | **73.68** |

# Рандомный тест на малых степенях 2 .. 20

| n     | author       | ver    | method                            | time      |
|-------|--------------|--------|-----------------------------------|-----------|
| 0     | comdiv       | v1     | `(fib[]+math.Pow)`                | 150.5     |
| 1     | gubkin       | v1     | `(fib+BinaryPow)`                 | 65.10     |
| 2     | gubkin       | v2     | `(fib[]+BinaryPow)`               | 50.75     |
| 3     | comdiv       | v2     | `(fib[]+BinaryPow(optimal))`      | 48.19     |
| **4** | **tetyuev**  | **v1** | **`(linear)`**                    | **36.34** |
| 5     | pirozhkov    | v1     | `(fib+math.Pow)`                  | 169.3     |
| 6     | tetyuev      | v2     | `(fib[]+math.Pow)`                | 161.7     |
| 7     | harisov      | v1     | `(fib+SqrtPower)`                 | 65.86     |
| **8** | **gavrilov** | v1     | **`(linear)`**                    | **36.34** |
| 9     | abzaev       | v1     | `(fib+math.Pow)`                  | 161.7     |
| 10    | gubkin       | v3     | `(fib[]+CachedPow)`               | 45.66     |
| 11    | gubkin       | v4     | `(CachedPow)`                     | 43.97     |
| 12    | comdiv       | v3     | `(fib[]+CachedPow+BinaryPowDual)` | 41.76     |

# Рандомный тест на средних степенях 21 .. 40

| n        | author       | ver      | method                            | time        |
|----------|--------------|----------|-----------------------------------|-------------|
| 0        | comdiv       | v1       | `(fib[]+math.Pow)`                | 237.1       |
| 1        | gubkin       | v1       | `(fib+BinaryPow)`                 | 103.8       |
| 2        | gubkin       | v2       | `(fib[]+BinaryPow)`               | 88.68       |
| 3        | comdiv       | v2       | `(fib[]+BinaryPow(optimal))`      | 84.91       |
| **4**    | **tetyuev**  | **v1**   | **`(linear)`**                    | **54.90**   |
| 5        | pirozhkov    | v1       | `(fib+math.Pow)`                  | 261.6       |
| 6        | tetyuev      | v2       | `(fib[]+math.Pow)`                | 244.4       |
| 7        | harisov      | v1       | `(fib+SqrtPower)`                 | 99.79       |
| **8**    | **gavrilov** | **v1**   | **`(linear)`**                    | **56.75**   |
| 9        | abzaev       | v1       | `(fib+math.Pow)`                  | 259.4       |
| ***10*** | ***gubkin*** | ***v3*** | `(fib[]+CachedPow)`               | ***61.92*** |
| 11       | gubkin       | v4       | `(CachedPow)`                     | 63.02       |
| 12       | comdiv       | v3       | `(fib[]+CachedPow+BinaryPowDual)` | 71.57       |

# Рандомный тест на средних степенях 70 .. 77

| n     | author       | ver    | method                            | time      |
|-------|--------------|--------|-----------------------------------|-----------|
| 0     | comdiv       | v1     | `(fib[]+math.Pow)`                | 223.1     |
| 1     | gubkin       | v1     | `(fib+BinaryPow)`                 | 178.7     |
| 2     | gubkin       | v2     | `(fib[]+BinaryPow)`               | 147.9     |
| 3     | comdiv       | v2     | `(fib[]+BinaryPow(optimal))`      | 140.3     |
| **4** | **tetyuev**  | **v1** | **`(linear)`**                    | **90.04** |
| 5     | pirozhkov    | v1     | `(fib+math.Pow)`                  | 276.7     |
| 6     | tetyuev      | v2     | `(fib[]+math.Pow)`                | 234.0     |
| 7     | harisov      | v1     | `(fib+SqrtPower)`                 | 168.2     |
| **8** | **gavrilov** | **v1** | `(linear)`                        | **90.06** |
| 9     | abzaev       | v1     | `(fib+math.Pow)`                  | 273.5     |
| 10    | gubkin       | v3     | `(fib[]+CachedPow)`               | 121.9     |
| 11    | gubkin       | v4     | `(CachedPow)`                     | 118.8     |
| 12    | comdiv       | v3     | `(fib[]+CachedPow+BinaryPowDual)` | 120.1     |

# Выводы

Пока побеждает простота!
Простое элементарное линейное решение показывает себя лучше на всех режимах кроме
кэшированного варианта на самых больших степенях.
Скорее всего это связано с более эффективной работой кэша в вариантах 10,11,12 (я имею в виду кэша процессора)
и именно если прогоняются по кругу одни и те же входные параметры.

Пока победила линейка!
Но с другой стороны посмотрите на варианты 10, 11, 12 и вообще на то какие изыскания провел Андрей, начиная 
со своего первого варианта (номер 1)

# Сырой ответ

```text
goos: linux
goarch: amd64
pkg: github.com/comdiv/task_func_optimize_base_go/basis
cpu: AMD Ryzen 5 2600 Six-Core Processor            
BenchmarkBasicSuperFuncImpl         1051           1129977 |
PASS
ok      github.com/comdiv/task_func_optimize_base_go/basis      1.318s
goos: linux
goarch: amd64
pkg: github.com/comdiv/task_func_optimize_base_go/implementations
cpu: AMD Ryzen 5 2600 Six-Core Processor            
BenchmarkDefault/0._comdiv____v1_(fib[]+math.Pow)                6220890               192.5 ns/op
BenchmarkDefault/1._gubkin____v1_(fib+BinaryPow)                17734603                67.68 ns/op
BenchmarkDefault/2._gubkin____v2_(fib[]+BinaryPow)              21867381                55.28 ns/op
BenchmarkDefault/3._comdiv____v2_(fib[]+BinaryPow(optimal))             23110713                52.28 ns/op
BenchmarkDefault/4._tetyuev___v1_(linear)                               43461694                35.80 ns/op
BenchmarkDefault/5._pirozhkov_v1_(fib+math.Pow)                          5327708               227.1 ns/op
BenchmarkDefault/6._tetyuev___v2_(fib[]+math.Pow)                        5959140               201.5 ns/op
BenchmarkDefault/7._harisov___v1_(fib+SqrtPower)                        17758226                68.22 ns/op
BenchmarkDefault/8._gavrilov__v1_(linear)                               33845356                35.54 ns/op
BenchmarkDefault/9._abzaev____v1_(fib+math.Pow)                          5629910               215.0 ns/op
BenchmarkDefault/10._gubkin____v3_(fib[]+CachedPow)                     41262303                29.80 ns/op
BenchmarkDefault/11._gubkin____v4_(CachedPow)                           45566968                25.90 ns/op
BenchmarkDefault/12._comdiv____v3_(fib[]+CachedPow+BinaryPowDual)       40437264                28.33 ns/op
BenchmarkRandom_Low_2_20/0._comdiv____v1_(fib[]+math.Pow)                7957298               150.5 ns/op
BenchmarkRandom_Low_2_20/1._gubkin____v1_(fib+BinaryPow)                18371292                65.10 ns/op
BenchmarkRandom_Low_2_20/2._gubkin____v2_(fib[]+BinaryPow)              23557147                50.75 ns/op
BenchmarkRandom_Low_2_20/3._comdiv____v2_(fib[]+BinaryPow(optimal))     24743973                48.19 ns/op
BenchmarkRandom_Low_2_20/4._tetyuev___v1_(linear)                       33426841                36.34 ns/op
BenchmarkRandom_Low_2_20/5._pirozhkov_v1_(fib+math.Pow)                  7085553               169.3 ns/op
BenchmarkRandom_Low_2_20/6._tetyuev___v2_(fib[]+math.Pow)                7409898               161.7 ns/op
BenchmarkRandom_Low_2_20/7._harisov___v1_(fib+SqrtPower)                18168246                65.86 ns/op
BenchmarkRandom_Low_2_20/8._gavrilov__v1_(linear)                       33509650                36.34 ns/op
BenchmarkRandom_Low_2_20/9._abzaev____v1_(fib+math.Pow)                  7401091               161.7 ns/op
BenchmarkRandom_Low_2_20/10._gubkin____v3_(fib[]+CachedPow)             28599098                45.66 ns/op
BenchmarkRandom_Low_2_20/11._gubkin____v4_(CachedPow)                   27089617                43.97 ns/op
BenchmarkRandom_Low_2_20/12._comdiv____v3_(fib[]+CachedPow+BinaryPowDual)               28570917                41.76 ns/op
BenchmarkRandom_High_21_40/0._comdiv____v1_(fib[]+math.Pow)                              5041204               237.1 ns/op
BenchmarkRandom_High_21_40/1._gubkin____v1_(fib+BinaryPow)                              11595182               103.8 ns/op
BenchmarkRandom_High_21_40/2._gubkin____v2_(fib[]+BinaryPow)                            13539183                88.68 ns/op
BenchmarkRandom_High_21_40/3._comdiv____v2_(fib[]+BinaryPow(optimal))                   13696000                84.91 ns/op
BenchmarkRandom_High_21_40/4._tetyuev___v1_(linear)                                     21896962                54.90 ns/op
BenchmarkRandom_High_21_40/5._pirozhkov_v1_(fib+math.Pow)                                4592438               261.6 ns/op
BenchmarkRandom_High_21_40/6._tetyuev___v2_(fib[]+math.Pow)                              4910998               244.4 ns/op
BenchmarkRandom_High_21_40/7._harisov___v1_(fib+SqrtPower)                              11968050                99.79 ns/op
BenchmarkRandom_High_21_40/8._gavrilov__v1_(linear)                                     21090129                56.75 ns/op
BenchmarkRandom_High_21_40/9._abzaev____v1_(fib+math.Pow)                                4505362               259.4 ns/op
BenchmarkRandom_High_21_40/10._gubkin____v3_(fib[]+CachedPow)                           20040206                61.92 ns/op
BenchmarkRandom_High_21_40/11._gubkin____v4_(CachedPow)                                 19508112                63.02 ns/op
BenchmarkRandom_High_21_40/12._comdiv____v3_(fib[]+CachedPow+BinaryPowDual)             16969706                71.57 ns/op
BenchmarkRandom_Large_70_77/0._comdiv____v1_(fib[]+math.Pow)                             5382337               223.1 ns/op
BenchmarkRandom_Large_70_77/1._gubkin____v1_(fib+BinaryPow)                              6716575               178.7 ns/op
BenchmarkRandom_Large_70_77/2._gubkin____v2_(fib[]+BinaryPow)                            8063720               147.9 ns/op
BenchmarkRandom_Large_70_77/3._comdiv____v2_(fib[]+BinaryPow(optimal))                   8466577               140.3 ns/op
BenchmarkRandom_Large_70_77/4._tetyuev___v1_(linear)                                    13134283                90.04 ns/op
BenchmarkRandom_Large_70_77/5._pirozhkov_v1_(fib+math.Pow)                               4317555               276.7 ns/op
BenchmarkRandom_Large_70_77/6._tetyuev___v2_(fib[]+math.Pow)                             5104638               234.0 ns/op
BenchmarkRandom_Large_70_77/7._harisov___v1_(fib+SqrtPower)                              7190355               168.2 ns/op
BenchmarkRandom_Large_70_77/8._gavrilov__v1_(linear)                                    12872006                90.06 ns/op
BenchmarkRandom_Large_70_77/9._abzaev____v1_(fib+math.Pow)                               4418335               273.5 ns/op
BenchmarkRandom_Large_70_77/10._gubkin____v3_(fib[]+CachedPow)                           9628729               121.9 ns/op
BenchmarkRandom_Large_70_77/11._gubkin____v4_(CachedPow)                                10021816               118.8 ns/op
BenchmarkRandom_Large_70_77/12._comdiv____v3_(fib[]+CachedPow+BinaryPowDual)            10004498               120.1 ns/op
BenchmarkFixed_77/0._comdiv____v1_(fib[]+math.Pow)                                       6544897               184.5 ns/op
BenchmarkFixed_77/1._gubkin____v1_(fib+BinaryPow)                                        7722958               156.8 ns/op
BenchmarkFixed_77/2._gubkin____v2_(fib[]+BinaryPow)                                     10283184               115.4 ns/op
BenchmarkFixed_77/3._comdiv____v2_(fib[]+BinaryPow(optimal))                            10843225               109.2 ns/op
BenchmarkFixed_77/4._tetyuev___v1_(linear)                                              14233248                84.50 ns/op
BenchmarkFixed_77/5._pirozhkov_v1_(fib+math.Pow)                                         4991678               240.5 ns/op
BenchmarkFixed_77/6._tetyuev___v2_(fib[]+math.Pow)                                       6373161               190.3 ns/op
BenchmarkFixed_77/7._harisov___v1_(fib+SqrtPower)                                       10712284               130.0 ns/op
BenchmarkFixed_77/8._gavrilov__v1_(linear)                                              14152489                83.77 ns/op
BenchmarkFixed_77/9._abzaev____v1_(fib+math.Pow)                                         5083718               235.9 ns/op
BenchmarkFixed_77/10._gubkin____v3_(fib[]+CachedPow)                                    10673214               112.8 ns/op
BenchmarkFixed_77/11._gubkin____v4_(CachedPow)                                          10783995               111.7 ns/op
BenchmarkFixed_77/12._comdiv____v3_(fib[]+CachedPow+BinaryPowDual)                      16065031                73.68 ns/op

```