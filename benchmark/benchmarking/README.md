# Test de performance

## 1 POST /transaction

php script.php 0.000000000001 60

### noasync:

 - ./API.exe: pas d'option -async

```
Test 1:

2018201820182018-1111-0404 11:27:48
Benchmark: 157 requests in 60.190945863724seconds.
Equivalent to 2.6083657225699 requests/sec

Test 2:

2018201820182018-1111-0404 11:39:45 B
enchmark: 160 requests in 60.244254112244seconds.
Equivalent to 2.6558549418156 requests/sec

Test 3:

2018201820182018-1111-0404 11:43:52
Benchmark: 160 requests in 60.222141027451seconds.
Equivalent to 2.656830150344 requests/sec
```

### async gorout 500:

 - ./API.exe -async -gorout 500

```
Test 1:

2018201820182018-1111-0404 11:52:38
Benchmark: 670 requests in 60.039159059525seconds.
Equivalent to 11.159383483965 requests/sec

Test 2:

2018201820182018-1111-0505 12:00:20
Benchmark: 657 requests in 60.185122013092seconds.
Equivalent to 10.916319150389 requests/sec

Test 3:

2018201820182018-1111-0505 12:08:20
Benchmark: 658 requests in 60.270732879639seconds.
Equivalent to 10.917404991806 requests/sec
```

### async gorout 20000:

 - ./API.exe -async -gorout 20000

Avec cette configuration, on monte à un maximum de 4660 goroutines sur 60 seconde.
Dans le cas du test 1, avec 4778 transactions ajoutées en base en 1 min, il faudra
environ 300ms x 4660 / 1000 = 1398 secondes pour tout indexer, soit un peu plus de 23 minutes.
A noté que l'API reste disponible pour toute autre opération durant ce laps de temps.

```
Test 1:

2018201820182018-1111-0505 12:16:51
Benchmark: 4778 requests in 60.006369113922seconds.
Equivalent to 79.624881001031request/sec

Test 2:

2018201820182018-1111-0505 12:57:12
Benchmark: 3824 requests in 60.001321077347seconds.
Equivalent to 63.731930086515request/sec

Test 3:

2018201820182018-1111-0505 01:02:38 Benchmark: 4360 requests in 60.000090837479seconds.
Equivalent to 72.666556652554request/sec
```