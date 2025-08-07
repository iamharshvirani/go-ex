package algos

import "fmt"

// Node represents a single element in the doubly linked list.
type Node struct {
	key   string
	value int
	prev  *Node
	next  *Node
}

// Cache is the main struct for our LRU cache. It holds a map
// for O(1) lookups and a doubly linked list for order.
type Cache struct {
	capacity int
	length   int
	head     *Node
	tail     *Node
	data     map[string]*Node
}

// NewCache creates and returns a new Cache instance with a given capacity.
func NewCache(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		length:   0,
		head:     nil,
		tail:     nil,
		data:     make(map[string]*Node),
	}
}

// addNode adds a new node to the head of the doubly linked list.
// This is an internal helper method.
func (c *Cache) addNode(node *Node) {
	node.prev = nil
	node.next = c.head

	if c.head != nil {
		c.head.prev = node
	}
	c.head = node

	if c.tail == nil {
		c.tail = node
	}
	c.length++
}

// removeNode removes a node from the doubly linked list.
// This is an internal helper method.
func (c *Cache) removeNode(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		// Node is the head
		c.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		// Node is the tail
		c.tail = node.prev
	}
	c.length--
}

// moveToHead moves an existing node to the head of the linked list.
// This signifies that it has been "recently used."
// This is an internal helper method.
func (c *Cache) moveToHead(node *Node) {
	c.removeNode(node)
	c.addNode(node)
}

// Get retrieves a value from the cache. If the key exists, it moves the node
// to the front of the list to mark it as most recently used.
func (c *Cache) Get(key string) (int, bool) {
	if node, ok := c.data[key]; ok {
		c.moveToHead(node)
		return node.value, true
	}
	return 0, false
}

// Set adds a new key-value pair to the cache or updates an existing one.
func (c *Cache) Set(key string, value int) {
	if node, ok := c.data[key]; ok {
		// Key exists, update the value and move to head.
		node.value = value
		c.moveToHead(node)
		return
	}

	// Key is new
	node := &Node{key: key, value: value}
	c.data[key] = node
	c.addNode(node)

	// Check if we've exceeded capacity, and if so, remove the tail.
	if c.length > c.capacity {
		tailKey := c.tail.key
		c.removeNode(c.tail)
		delete(c.data, tailKey)
	}
}

func RunLRUCache() {
	// Create a new LRU cache with a capacity of 3.
	lru := NewCache(3)

	fmt.Println("Setting key-value pairs...")
	lru.Set("a", 10) // Cache: {a: 10}
	lru.Set("b", 20) // Cache: {b: 20, a: 10}
	lru.Set("c", 30) // Cache: {c: 30, b: 20, a: 10}

	fmt.Println("\nAccessing 'b' to make it most recently used.")
	val, ok := lru.Get("b") // Accessing 'b'
	if ok {
		fmt.Printf("Get('b'): %d\n", val) // Cache: {b: 20, c: 30, a: 10}
	}

	fmt.Println("\nAdding 'd', which will evict 'a' (least recently used).")
	lru.Set("d", 40) // Cache: {d: 40, b: 20, c: 30}
	val, ok = lru.Get("a")
	fmt.Printf("Get('a') exists? %t\n", ok) // Should be false

	fmt.Println("\nAdding 'e', which will evict 'c'.")
	lru.Set("e", 50) // Cache: {e: 50, d: 40, b: 20}
	val, ok = lru.Get("c")
	fmt.Printf("Get('c') exists? %t\n", ok) // Should be false
}
