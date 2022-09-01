from functools import lru_cache


def decorate(func):
    def wrapped():
        print("running wrapped()")
    return wrapped


@decorate
def target():
    print("running target")


registry = []


def register(func):
    print("running register {}".format(func))
    registry.append(func)
    return func


@register
def f1():
    print("running f1()")


@register
def f2():
    print("running f2()")


@lru_cache()
def func(n):
    print(n, "called")
    return n


if __name__ == "__main__":
    target()
    print("registery ->", registry)

    print()

    print(func(1))
    print(func(1))
    print(func(2))
