from threading import Thread, Lock

i = 0

def tellOpp(lock):
	global i
	for j in range(0, 1000000):
		lock.acquire(True)
		i += 1
		lock.release()
	
def tellNed(lock):
	global i
	for j in range(0, 1000000):
		lock.acquire(True)
		i -= 1
		lock.release()

def main():
	lock = Lock()
	ink = Thread(target = tellOpp, args = (lock,))
	dek = Thread(target = tellNed, args = (lock,))
	ink.start()
	dek.start()
	ink.join()
	dek.join()
	print(i)

main()
