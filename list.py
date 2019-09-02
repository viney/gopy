# list.py
from person import Person



p1 = Person('golang', 10)
p2 = Person('influx', 19)

pList = [p1, p2]
	
p = Person('python', 11)
pList.append(p)

for p in pList[:]:
	p.say()


def list(i):
	return pList[:i]
