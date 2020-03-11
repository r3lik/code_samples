#https://www.codevscolor.com/args-kwargs-python-difference/

def print_values(**values):
	for key, value in values.items():
		print("{} = {}".format(key,value))

print_values(one = 1,two = 2,three = 3,four = 4,five = 5)


def show_details(a,b,*args,**kwargs):
	print("a is ",a)
	print("b is ",b)
	print("args is ",args)
	print("kwargs is ",kwargs)


show_details(1,2,3,4,5,6,7,8,9)
print("-----------")
show_details(1,2,3,4,5,6,c= 7,d = 8,e = 9)
print("-----------")
