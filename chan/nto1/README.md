## yang gua pelajarin

### 1. synced work 1toN, ready ke workers

```
goroutine itu concurrent, meaning kalo lu sleep main go routine, mereka bakal tetep lanjut. nah gmn cara delay work start biar synced, antara 1 goroutine dgn  yg lain?
```

#### a. block worker pake \<-ready

```
dengan begini, worker ga bakal lanjutin kodenya, sampai, either
```

##### 1) channel di close

```
making use fiturnya closed channel in which lu bisa receive infinitely
```

##### 2) ready di pass value

```
sequentially pass value struct{}{} ke ready, biar <-ready di worker bisa di eksekusi
```

### 2. wait group lite Nto1, workers ke done

```
workers bakal inform ke done kalo work nya selesai, dengan pass struct{}{} ke done
<-done di main bakal blocking sampai go routine semua selesai (in case ada pass value ke done lagi)
```
