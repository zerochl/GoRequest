package threadpool

type ThreadPool struct {
	queue        chan func() error
	corePoolSize chan int
	maxIdle      int
	keepAlive    int
}

func NewThreadPool(maxIdle, keepAlive int) *ThreadPool {
	threadPool := &ThreadPool{
		queue:        make(chan func() error, maxIdle),
		corePoolSize: make(chan int, keepAlive),
		maxIdle:      maxIdle,
		keepAlive:    keepAlive,
	}
	// New完立即开启线程池
	threadPool.startPool()
	return threadPool
}

func (threadPool *ThreadPool) startPool() {
	// 开启一个线程执行loop，防止外部调用者被阻塞
	go func() {
		threadPool.loopQueue()
	}()
}

// 阻塞调用，如果queue已满则会被阻塞
func (threadPool *ThreadPool) AddTaskBlock(task func() error) {
	threadPool.queue <- task
}

// 无阻塞添加任务，方法调用用并不会因为queue已满而被阻塞
// 无阻塞调用有个风险，同一时间的调用，并不能确保先后顺序
func (threadPool *ThreadPool) AddTask(task func() error) {
	go func() {
		threadPool.queue <- task
	}()
}

// 从任务队列获取任务,并转移到核心队列
func (threadPool *ThreadPool) loopQueue() {
	for true {
		// 如果queue为空，此处代码会阻塞，防止无限制for循环
		newTask := <-threadPool.queue
		// 如果corePoolSize已满，此时再写入，此处代码会阻塞
		threadPool.corePoolSize <- 1
		go threadPool.startTask(newTask)
	}
}

func (threadPool *ThreadPool) startTask(task func() error) {
	task()
	// 线程执行完毕，执行corePoolSize读取，此操作会释放一个core位置
	<-threadPool.corePoolSize
}
