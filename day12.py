import ujson

def add_numbers(incoming):
    if type(incoming) == int:
        return incoming
    elif type(incoming) == dict:
        if "red" in incoming.values():
            return 0
        return sum(map(add_numbers, incoming.values()))
    elif type(incoming) == list:
        return sum(map(add_numbers, incoming))
    else:
        return 0

data = ujson.decode(open('day12.input', 'r').read())
print add_numbers(data)