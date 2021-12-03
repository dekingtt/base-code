import requests

URL = "http://store.steampowered.com/feeds/newreleases.xml"

def main():
    r = requests.get(URL)
    with open('newreleases.xml', 'wb') as f:
        f.write(r.content)

if __name__ == "__main__":
    main()