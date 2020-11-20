### baseline

```
100 bytes:

./wrk -c1 -t1 http://localhost:31952/byte/100
Running 10s test @ http://localhost:31952/byte/100
  1 threads and 1 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   660.57us  182.68us   6.55ms   98.10%
    Req/Sec     1.53k    41.45     1.58k    80.00%
  15237 requests in 10.00s, 3.15MB read
Requests/sec:   1523.65
Transfer/sec:    322.88KB

./wrk -c16 -t16 http://localhost:31952/byte/100
Running 10s test @ http://localhost:31952/byte/100
  16 threads and 16 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     9.69ms  770.89us  29.34ms   95.19%
    Req/Sec   103.63      5.27   121.00     63.13%
  16512 requests in 10.01s, 3.42MB read
Requests/sec:   1649.81
Transfer/sec:    349.62KB

./wrk -c64 -t64 http://localhost:31952/byte/100
Running 10s test @ http://localhost:31952/byte/100
  64 threads and 64 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    38.98ms    3.64ms  84.95ms   94.76%
    Req/Sec    25.63      5.25    50.00     54.64%
  16464 requests in 10.10s, 3.41MB read
Requests/sec:   1630.18
Transfer/sec:    345.46KB

./wrk -c128 -t128 http://localhost:31952/byte/100
Running 10s test @ http://localhost:31952/byte/100
  128 threads and 128 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    77.77ms   10.66ms 207.52ms   94.46%
    Req/Sec    12.84      4.56    40.00     71.80%
  16585 requests in 10.09s, 3.43MB read
Requests/sec:   1642.94
Transfer/sec:    348.16KB

1MB:

/wrk -c1 -t1 http://localhost:31952/byte/1000000
Running 10s test @ http://localhost:31952/byte/1000000
  1 threads and 1 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    95.86ms    3.18ms 123.53ms   92.31%
    Req/Sec    10.40      1.97    20.00     96.00%
  104 requests in 10.01s, 10.17MB read
Requests/sec:     10.39
Transfer/sec:      1.02MB

./wrk -c16 -t16 http://localhost:31952/byte/1000000
Running 10s test @ http://localhost:31952/byte/1000000
  16 threads and 16 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.36s   374.43ms   1.91s    84.00%
    Req/Sec     0.34      1.26    10.00     93.33%
  105 requests in 10.01s, 10.27MB read
  Socket errors: connect 0, read 0, write 0, timeout 5
Requests/sec:     10.48
Transfer/sec:      1.03MB

./wrk -c64 -t64 http://localhost:31952/byte/1000000
Running 10s test @ http://localhost:31952/byte/1000000
  64 threads and 64 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.01s   569.31ms   1.91s    56.67%
    Req/Sec     0.34      1.06     5.00     93.33%
  105 requests in 10.10s, 10.27MB read
  Socket errors: connect 0, read 0, write 0, timeout 75
Requests/sec:     10.40
Transfer/sec:      1.02MB

./wrk -c128 -t128 http://localhost:31952/byte/1000000
Running 10s test @ http://localhost:31952/byte/1000000
  128 threads and 128 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.03s   557.28ms   1.92s    60.00%
    Req/Sec     0.26      0.84     5.00     93.33%
  105 requests in 10.10s, 10.27MB read
  Socket errors: connect 0, read 0, write 0, timeout 75
Requests/sec:     10.40
Transfer/sec:      1.02MB
```
### istio

```
100 bytes:

./wrk -c1 -t1 http://localhost:31952/byte/100
Running 10s test @ http://localhost:31952/byte/100
  1 threads and 1 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.41ms  293.69us   9.76ms   97.14%
    Req/Sec   714.42     42.30   767.00     65.00%
  7113 requests in 10.00s, 2.33MB read
Requests/sec:    711.12
Transfer/sec:    238.89KB

./wrk -c16 -t16 http://localhost:31952/byte/100
Running 10s test @ http://localhost:31952/byte/100
  16 threads and 16 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    14.22ms    4.38ms  61.46ms   78.42%
    Req/Sec    70.61     10.72   101.00     85.38%
  11299 requests in 10.01s, 3.72MB read
Requests/sec:   1128.80
Transfer/sec:    380.18KB

./wrk -c64 -t64 http://localhost:31952/byte/100
Running 10s test @ http://localhost:31952/byte/100
  64 threads and 64 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    57.84ms   11.10ms 105.98ms   79.10%
    Req/Sec    17.26      5.15    40.00     66.47%
  11094 requests in 10.10s, 3.65MB read
Requests/sec:   1098.50
Transfer/sec:    370.09KB

./wrk -c128 -t128 http://localhost:31952/byte/100
Running 10s test @ http://localhost:31952/byte/100
  128 threads and 128 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   114.67ms   19.76ms 235.56ms   85.79%
    Req/Sec     9.29      2.25    30.00     82.03%
  11191 requests in 10.10s, 3.69MB read
Requests/sec:   1108.11
Transfer/sec:    374.31KB

1MB:

./wrk -c1 -t1 http://localhost:31952/byte/1000000
Running 10s test @ http://localhost:31952/byte/1000000
  1 threads and 1 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    98.46ms    2.92ms 115.54ms   92.08%
    Req/Sec    10.15      1.51    20.00     96.97%
  101 requests in 10.01s, 9.89MB read
Requests/sec:     10.09
Transfer/sec:      0.99MB

./wrk -c16 -t16 http://localhost:31952/byte/1000000
Running 10s test @ http://localhost:31952/byte/1000000
  16 threads and 16 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.38s   382.62ms   1.94s    83.67%
    Req/Sec     0.29      0.96     5.00     93.20%
  103 requests in 10.01s, 10.08MB read
  Socket errors: connect 0, read 0, write 0, timeout 5
Requests/sec:     10.28
Transfer/sec:      1.01MB

./wrk -c64 -t64 http://localhost:31952/byte/1000000
Running 10s test @ http://localhost:31952/byte/1000000
  64 threads and 64 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.04s   577.72ms   1.96s    60.00%
    Req/Sec     0.35      1.28    10.00     93.14%
  102 requests in 10.02s, 9.99MB read
  Socket errors: connect 0, read 0, write 0, timeout 72
Requests/sec:     10.18
Transfer/sec:      1.00MB

./wrk -c128 -t128 http://localhost:31952/byte/1000000
Running 10s test @ http://localhost:31952/byte/1000000
  128 threads and 128 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.04s   572.93ms   1.95s    60.00%
    Req/Sec     0.28      0.93     5.00     95.19%
  104 requests in 10.10s, 10.18MB read
  Socket errors: connect 0, read 0, write 0, timeout 74
Requests/sec:     10.30
Transfer/sec:      1.01MB
```
### appmesh

```
100 bytes:

./wrk -c1 -t1 http://localhost:31952/byte/100
Running 10s test @ http://localhost:31952/byte/100
  1 threads and 1 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.26ms  393.13us  11.03ms   96.12%
    Req/Sec   807.70     99.53     1.03k    67.00%
  8044 requests in 10.01s, 2.04MB read
Requests/sec:    803.75
Transfer/sec:    208.79KB

./wrk -c16 -t16 http://localhost:31952/byte/100
Running 10s test @ http://localhost:31952/byte/100
  16 threads and 16 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    11.85ms    2.25ms  39.60ms   88.30%
    Req/Sec    84.53     13.88   262.00     69.35%
  13559 requests in 10.10s, 3.45MB read
Requests/sec:   1342.50
Transfer/sec:    349.89KB

./wrk -c64 -t64 http://localhost:31952/byte/100
Running 10s test @ http://localhost:31952/byte/100
  64 threads and 64 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    48.97ms    9.65ms 127.02ms   85.95%
    Req/Sec    20.41      4.79    40.00     77.57%
  13135 requests in 10.10s, 3.34MB read
Requests/sec:   1300.39
Transfer/sec:    339.06KB

./wrk -c128 -t128 http://localhost:31952/byte/100
Running 10s test @ http://localhost:31952/byte/100
  128 threads and 128 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    98.15ms   19.05ms 226.94ms   85.00%
    Req/Sec    10.37      2.69    30.00     89.59%
  13097 requests in 10.10s, 3.34MB read
Requests/sec:   1297.09
Transfer/sec:    338.73KB

1MB:

/wrk -c1 -t1 http://localhost:31952/byte/1000000
Running 10s test @ http://localhost:31952/byte/1000000
  1 threads and 1 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    97.04ms    3.18ms 114.81ms   89.32%
    Req/Sec    10.41      2.35    20.00     92.86%
  103 requests in 10.00s, 10.08MB read
Requests/sec:     10.30
Transfer/sec:      1.01MB

./wrk -c16 -t16 http://localhost:31952/byte/1000000
Running 10s test @ http://localhost:31952/byte/1000000
  16 threads and 16 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.37s   382.14ms   1.92s    83.67%
    Req/Sec     0.29      0.96     5.00     93.20%
  103 requests in 10.02s, 10.08MB read
  Socket errors: connect 0, read 0, write 0, timeout 5
Requests/sec:     10.28
Transfer/sec:      1.01MB

./wrk -c64 -t64 http://localhost:31952/byte/1000000
Running 10s test @ http://localhost:31952/byte/1000000
  64 threads and 64 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.03s   570.19ms   1.94s    60.00%
    Req/Sec     0.32      0.99     5.00     92.31%
  104 requests in 10.10s, 10.18MB read
  Socket errors: connect 0, read 0, write 0, timeout 74
Requests/sec:     10.30
Transfer/sec:      1.01MB

./wrk -c128 -t128 http://localhost:31952/byte/1000000
Running 10s test @ http://localhost:31952/byte/1000000
  128 threads and 128 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.03s   571.57ms   1.94s    60.00%
    Req/Sec     0.29      0.95     5.00     93.27%
  104 requests in 10.10s, 10.18MB read
  Socket errors: connect 0, read 0, write 0, timeout 74
Requests/sec:     10.30
Transfer/sec:      1.01MB
```