from socket import socket
from json import loads
from base64 import b64decode

# character version
def main_():
    # 1. create socket
    client = socket()
    # 2. connect to server
    client.connect(('127.0.0.1', 6789))
    # 3. receive
    print(client.recv(1024).decode('utf-8'))
    client.close()

# mutiple media vesion
def main():
    client = socket()
    client.connect(('127.0.0.1', 6789))
    in_data = bytes()

    data = client.recv(1024)
    while data:
        in_data += data
        data = client.recv(1024)

    my_dict = loads(in_data.decode('utf-8'))
    filename = my_dict['filename']
    filedata = my_dict['filedata'].encode('utf-8')
    with open('_'+filename, 'wb') as f:
        f.write(b64decode(filedata))
    print("saved the data")

if __name__ == '__main__':
    main()