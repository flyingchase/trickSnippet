def foo(self):
    return "foo"


name = "yourname"

if __name__ == "__main__":
    B = type("MyClass", (object,), {"name": name, "foo": foo})
    print(B.name)
    print(B().foo())
