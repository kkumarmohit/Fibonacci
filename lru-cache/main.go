package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type FibonacciResponse struct {
	Order int `json:"order"`
	Value int `json:"value"`
}

type CacheItem struct {
	Order int
	Value int
}

type LRUCache struct {
	capacity int
	items    map[int]*list.Element
	order    *list.List
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[int]*list.Element),
		order:    list.New(),
	}
}

func (c *LRUCache) Get(order int) (int, bool) {
	if elem, ok := c.items[order]; ok {
		c.order.MoveToFront(elem)
		return elem.Value.(CacheItem).Value, true
	}
	return 0, false
}

func (c *LRUCache) Put(order int, value int) {
	if elem, ok := c.items[order]; ok {
		c.order.MoveToFront(elem)
		elem.Value = CacheItem{Order: order, Value: value}
		return
	}
	if c.order.Len() >= c.capacity {
		oldest := c.order.Back()
		if oldest != nil {
			item := oldest.Value.(CacheItem)
			delete(c.items, item.Order)
			c.order.Remove(oldest)
		}
	}
	elem := c.order.PushFront(CacheItem{Order: order, Value: value})
	c.items[order] = elem
}

type CacheResponse struct {
	Fibonacci FibonacciResponse `json:"fibonacci"`
	Cached    bool              `json:"cached"`
}

func fetchFromDemoApp(order int, endpoint string) (FibonacciResponse, error) {
	url := fmt.Sprintf("http://fibonacci-api:8080%s?order=%d", endpoint, order)
	resp, err := http.Get(url)
	if err != nil {
		return FibonacciResponse{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return FibonacciResponse{}, err
	}
	var fibResp FibonacciResponse
	if err := json.Unmarshal(body, &fibResp); err != nil {
		return FibonacciResponse{}, err
	}
	return fibResp, nil
}

func fibonacciHandler(cache *LRUCache, endpoint string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		orderStr := req.URL.Query().Get("order")
		if orderStr == "" {
			http.Error(w, "Missing 'order' query parameter", 400)
			return
		}
		order, err := strconv.Atoi(orderStr)
		if err != nil || order < 1 {
			http.Error(w, "Invalid 'order' query parameter", 400)
			return
		}
		if value, ok := cache.Get(order); ok {
			resp := CacheResponse{
				Fibonacci: FibonacciResponse{Order: order, Value: value},
				Cached:    true,
			}
			json.NewEncoder(w).Encode(resp)
			return
		}
		fibResp, err := fetchFromDemoApp(order, endpoint)
		if err != nil {
			http.Error(w, "Failed to fetch from demo app", 500)
			return
		}
		cache.Put(order, fibResp.Value)
		resp := CacheResponse{
			Fibonacci: fibResp,
			Cached:    false,
		}
		json.NewEncoder(w).Encode(resp)
	}
}

func main() {
	cache := NewLRUCache(5)
	http.HandleFunc("/fibonacci", fibonacciHandler(cache, "/fibonacci"))
	http.HandleFunc("/recursive-fibonacci", fibonacciHandler(cache, "/recursive-fibonacci"))
	port := "8081"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	log.Printf("LRU Cache app listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
