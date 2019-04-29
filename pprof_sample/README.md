# pprof sample
introduction of pprof sample usage

## fib profiling

```bash
-> % ./fib
2019/04/29 17:02:10 profile: cpu profiling enabled, cpu.pprof
start
2019/04/29 17:02:16 profile: cpu profiling disabled, cpu.pprof
```

```bash
-> % go tool pprof fib cpu.pprof
File: fib
Type: cpu
Time: Apr 29, 2019 at 5:02pm (JST)
Duration: 5.23s, Total samples = 4.45s (85.16%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 4.44s, 99.78% of 4.45s total
Dropped 8 nodes (cum <= 0.02s)``
      flat  flat%   sum%        cum   cum%
     4.10s 92.13% 92.13%      4.29s 96.40%  main.fib
     0.19s  4.27% 96.40%      0.19s  4.27%  runtime.newstack
     0.12s  2.70% 99.10%      0.12s  2.70%  runtime.nanotime
     0.03s  0.67% 99.78%      0.03s  0.67%  runtime.usleep
         0     0% 99.78%      4.29s 96.40%  main.main
         0     0% 99.78%      4.29s 96.40%  runtime.main
         0     0% 99.78%      0.16s  3.60%  runtime.mstart
         0     0% 99.78%      0.15s  3.37%  runtime.mstart1
         0     0% 99.78%      0.15s  3.37%  runtime.sysmon
```

```bash
-> % go tool pprof http://localhost:6060/debug/pprof/profile
Fetching profile over HTTP from http://localhost:6060/debug/pprof/profile
Saved profile in /Users/kazukihigashiguchi/pprof/pprof.samples.cpu.001.pb.gz
Type: cpu
Time: Apr 29, 2019 at 5:09pm (JST)
Duration: 30.18s, Total samples = 26.03s (86.26%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 25900ms, 99.50% of 26030ms total
Dropped 4 nodes (cum <= 130.15ms)
      flat  flat%   sum%        cum   cum%
   24160ms 92.82% 92.82%    25450ms 97.77%  main.fib
    1290ms  4.96% 97.77%     1290ms  4.96%  runtime.newstack
     450ms  1.73% 99.50%      450ms  1.73%  runtime.nanotime
         0     0% 99.50%    25450ms 97.77%  main.main
         0     0% 99.50%    25450ms 97.77%  runtime.main
         0     0% 99.50%      570ms  2.19%  runtime.mstart
         0     0% 99.50%      570ms  2.19%  runtime.mstart1
         0     0% 99.50%      570ms  2.19%  runtime.sysmon
```

## refs
- https://godoc.org/net/http/pprof
- https://christina04.hatenablog.com/entry/golang-pprof-basic
- https://github.com/uber-archive/go-torch
- https://github.com/pkg/profile
