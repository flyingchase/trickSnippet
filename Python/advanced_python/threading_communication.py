import heapq
import threading
from queue import Queue
from threading import Thread


def producer(out_q):
    data = []
    while True:
        out_q.put(data)


def consumer(in_q):
    data = in_q.get()


class PriorityQueue:
    def __int__(self):
        self._queue = []
        self._count = 0
        self._cv = threading.Condition()

    def put(self, item, priority):
        with self._cv:
            heapq.heappush(self._queue, (-priority, self._count, item))
            self._count += 1
            self._cv.notify()

    def get(self):
        with self._cv:
            while len(self._queue) == 0:
                self._cv.wait()
            return heapq.heappop(self._queue)[-1]


if __name__ == "__main__":
    q = Queue()
    t1 = Thread(target=producer, args=(q,))
    t2 = Thread(target=consumer, args=(q,))
    t1.start()
    t2.start()
