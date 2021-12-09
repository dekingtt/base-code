from random import randint
from threading import Thread
from time import time, sleep

def download(filename):
    print("start download %s ..." % filename)
    time_to_download = randint(5,10)
    sleep(time_to_download)
    print("%s finished, cost %s seconds" % (filename, time_to_download))

def main():
    start = time()
    t1 = Thread(target=download, args=("File_1.txt", ))
    t2 = Thread(target=download, args=('File_2.exe', ))
    t1.start()
    t2.start()
    t1.join()
    t2.join()
    end = time()
    print("totally cost %d seconds"% (end - start))


if __name__ == "__main__":
    main()