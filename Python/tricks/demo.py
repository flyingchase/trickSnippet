
from operator import itemgetter
from fnmatch import fnmatchcase
from collections import ChainMap


# 移出队列中的重复元素并保持元素间顺序不变
def removeDuplicateKeep(items, key=None):
    seen = set()
    for item in items:
        val = item if key is None else key(item)
        if val not in seen:
            yield item
            seen.add(val)

    return seen


if __name__ == "__main__":
    # 公共键对字典列表排序
    rows = [
        {'frame': "Brain"},
        {"frame": "David"},
        {"frame": "John"},
        {"frame": "Big"}
    ]
    rows_by_name = sorted(rows, key=itemgetter("frame"))
    print(rows_by_name)

    # 自定义排序
    class User:
        def __init__(self, user_id) -> None:
            self.user_id = user_id

        def __repr__(self) -> str:
            return 'User({})'.format(self.user_id)

    users = (User(1), User(100), User(99), User(24))
    sorted(users, key=lambda u: u.user_id)
    print(users)

    a = {1: 1, 2: "2", 3: 3, 4: '4'}
    b = {1: "1", 2: 2, 3: 3}
    # 两字典交集
    print(a.keys() & b.keys())
    print(a.items() & b.items())
    # 两字典差集
    print(a.keys()-b.keys())
    print(a.items()-b.items())
    # 两字典映射为同一个字典，相同key 取前者
    c = ChainMap(a, b)
    print(dict(c))

    res = {key: value for key, value in a.items() if key > 2}
    print(res)

    print('+++++++++{}++++++++++'.format("END"))

    # 字符串
    data = ['s', 100, 1, 'a', {'k': 'v'}]
    print(','.join(str(d) for d in data))

    # 字符串匹配
    addresses = [
        '5412 N CLARK ST',
        '1060 W ADDISON ST',
        '1039 W GRANVILLE AVE',
        '2122 N CLARK ST',
        '4802 N BROADWAY',
    ]
    print([addr for addr in addresses if fnmatchcase(addr, '* ST')])

    # 字符串格式化
    s = '{name} has {n} message'.format(name='wo', n=1)
    print(s)

    # 序列中移出重复项且稳定
    items = [0, 1, 9, 2, 8, 3, 7, 4, 6, 5, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0]
    print(list(removeDuplicateKeep(items, key=None)))
