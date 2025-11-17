package CircularQueue

type Data struct {
	key   string
	value string
}
type CircularQueue struct {
	data    []Data //
	head    int    // 出队指针
	tail    int    // 入队指针
	maxSize int    // 总容量
	size    int    // 当前大小
}

// NewCircularQueue 创建一个环形队列（固定容量）
func NewCircularQueue(cap int) *CircularQueue {
	return &CircularQueue{
		data:    make([]Data, cap),
		maxSize: cap,
		head:    0,
		tail:    0,
		size:    0,
	}
}

// Empty 是否空
func (q *CircularQueue) Empty() bool {
	return q.size == 0
}

// Full 是否满
func (q *CircularQueue) Full() bool {
	return q.size == q.maxSize
}

// Enqueue 入队
func (q *CircularQueue) Enqueue(key, value string) bool {
	if q.Full() { // 当前大小等于 最大容量说明队列已满
		return false
	}

	q.data[q.tail] = Data{key: key, value: value}
	q.tail = (q.tail + 1) % q.maxSize // 队首和队尾都进行取模运算
	q.size++
	return true
}

// Dequeue 出队 (队列一般是 先进先出 FIFO 所以 出队 一般出队首元素)
func (q *CircularQueue) Dequeue() (Data, bool) {
	if q.Empty() { // 当前大小为 0 说明队列为空
		return Data{}, false
	}

	val := q.data[q.head]             // 队首
	q.head = (q.head + 1) % q.maxSize // 队首和队尾都进行取模运算
	q.size--
	return val, true
}

// Front 返回队头元素
func (q *CircularQueue) Front() (Data, bool) {
	if q.Empty() {
		return Data{}, false
	}
	return q.data[q.head], true
}

func Do() {
	q := NewCircularQueue(2)

	q.Enqueue("msg", "1")
	q.Enqueue("msg", "2")
	q.Enqueue("msg", "3")
	q.Dequeue()
	q.Dequeue()
	q.Enqueue("msg", "4")
	q.Dequeue()
}
