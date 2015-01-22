from threading import Thread

i = 0

def tellOpp():
	global i
	for j in range(0, 1000000):
		i += 1
	
def tellNed():
	global i
	for j in range(0, 1000000):
		i -= 1

def main():
	ink = Thread(target = tellOpp, args = (),)
	dek = Thread(target = tellNed, args = (),)
	ink.start()
	dek.start()
	ink.join()
	dek.join()
	print(i)

main()
