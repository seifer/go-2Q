2Q LRU Algorithm

Examples #1 (by count)

    cache := lru.NewByCount(2)

    cache.Add(123)
    cache.Add(123)
    cache.Add(225)
    cache.Add(556)
    cache.Del(123)
    cache.Add(123)

    if res := cache.Reclaim(); len(res) > 0 {
        for _, v := range res {
            fmt.Println(v.(int)) 
        }
    }

Examples #2 (by size)

    cache := lru.NewBySize(2)

    cache.Add(123, 2)
    cache.Add(123, 3)
    cache.Add(225, 1)
    cache.Add(556, 1)
    cache.Del(123)
    cache.Add(123, 3)

    if res := cache.Reclaim(); len(res) > 0 {
        for _, v := range res {
            fmt.Println(v.(int))
        }
    }
