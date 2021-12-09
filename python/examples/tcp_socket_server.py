import json
from socket import socket, SOCK_STREAM, AF_INET, timeout
from base64 import b64decode, b64encode
from json import dumps
from datetime import datetime
from sys import setrecursionlimit
from threading import Thread

# single thread
def main_():
    # 1.创建套接字对象并指定使用哪种传输服务
    # family=AF_INET - IPv4地址
    # family=AF_INET6 - IPv6地址
    # type=SOCK_STREAM - TCP套接字
    # type=SOCK_DGRAM - UDP套接字
    # type=SOCK_RAW - 原始套接字
    server = socket(family=AF_INET, type=SOCK_STREAM)
    # 2.绑定IP地址和端口(端口用于区分不同的服务)
    # 同一时间在同一个端口上只能绑定一个服务否则报错
    server.bind(('127.0.0.1', 6789))
    # 3.开启监听 - 监听客户端连接到服务器
    # 参数512可以理解为连接队列的大小
    server.listen(512)
    print('Listening...')
    while True:
        # 4.通过循环接收客户端的连接并作出相应的处理(提供服务)
        # accept方法是一个阻塞方法如果没有客户端连接到服务器代码不会向下执行
        # accept方法返回一个元组其中的第一个元素是客户端对象
        # 第二个元素是连接到服务器的客户端的地址(由IP和端口两部分构成)
        client, addr = server.accept()
        print(str(addr) + "connected to server.")
        # 5.发送数据
        client.send(str(datetime.now()).encode('utf-8'))
        client.send('\n'.encode('utf-8'))
        # 6.断开连接
        client.close()
        exit(0)

# mutiple thread version
def main():
    class FileTransferHandler(Thread):
        def __init__(self, cclient) -> None:
            super().__init__()
            self.cclient = cclient

        def run(self) -> None:
            my_dict = {}
            my_dict['filename'] = "guido.png"
            my_dict['filedata'] = data
            json_str = dumps(my_dict)
            self.cclient.send(json_str.encode('utf-8'))
            self.cclient.close()

    # 1 create socket
    server = socket()
    server.settimeout(1)
    # 2 bind
    server.bind(('127.0.0.1', 6789))
    # 3. listening
    server.listen(512)

    with open('../100daysofcode/pybites.png', 'rb') as f:
        data = b64encode(f.read()).decode('utf-8')

    while True:
        try: 
            client, addr = server.accept()
        except timeout as x:
            pass
        else:
            FileTransferHandler(client).start()

if __name__ == '__main__':
    main()