from random import randint
from threading import Thread, Lock
from time import time, sleep

class DownloadTask(Thread):
    def __init__(self, filename) -> None:
        super().__init__()
        self._filename = filename

    def run(self):
        print("start to download %s ..."%self._filename)
        time_to_download = randint(5, 10)
        sleep(time_to_download)
        print('%s finished, cost %d seconds'%(self._filename, time_to_download))

# Lock
class Account(object):
    def __init__(self) -> None:
        super().__init__()
        self._balance = 0
        self._lock = Lock()

    def deposit(self, money):
        self._lock.acquire()
        try:
            new_balance = self._balance + money
            sleep(0.01)
            self._balance = new_balance
        finally:
            self._lock.release()

    @property
    def balance(self):
        return self._balance

class AddMondeyThread(Thread):
    def __init__(self, account, money) -> None:
        super().__init__()
        self._account = account
        self._money = money

    def run(self):
        self._account.deposit(self._money)

def main():
    start = time()
    t1 = DownloadTask('File_1.pdf')
    t2 = DownloadTask('File_2.mp4')

    t1.start()
    t2.start()
    t1.join()
    t2.join()

    end = time()

    print("totlaly cost %d seconds"%(end - start))

    print()
    account = Account()
    threads = []
    for _ in range(100):
        t = AddMondeyThread(account=account, money=1)
        threads.append(t)
        t.start()

    for t in threads:
        t.join()
    
    print("Account balance is %d"% account.balance)

if __name__=='__main__':
    main()