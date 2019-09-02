class Person(object):
	def __init__(self, name, age):
		self.name = name
		self.age = age
	
	def say(self):
		print('python: ' + self.name + "'s age is: ", self.age)


