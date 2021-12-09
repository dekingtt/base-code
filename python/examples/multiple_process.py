from multiprocessing import Process, process
from os import getppid
from random import randint
from time import time, sleep


def download_task(filename):
    print('Download started: pid is %d'.format(getppid()))
    print('Downloading %s...' % filename)
    time_to_end = randint(5, 10)
    sleep(time_to_end)
    print("%s Finished, cost %d second" % (filename, time_to_end))


def main():
    start = time()
    p1 = Process(target=download_task, args=('File_1.pdf',))
    p1.start()
    p2 = Process(target=download_task, args=('File_2.exe',))
    p2.start()
    p1.join()
    p2.join()
    end = time()
    print("totally cost %d seconds" % (end-start))

if __name__=='__main__':
    main()
