import time
from threading import Thread


def countdown(n):
    while n > 0:
        print("t-minus", n)
        n -= 1
        time.sleep(1)


if __name__ == "__main__":
    t = Thread(target=countdown, args=(10,))
    t.start()

    if t.is_alive():
        print("still running")
    else:
        print("completed")
