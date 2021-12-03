import requests
import bs4

URL = "https://www.mattkaydiary.com/"

def pull_site():
    raw_site_page = requests.get(URL)
    raw_site_page.raise_for_status()
    return raw_site_page

def scraper(site):
    header_list = []
    soup = bs4.BeautifulSoup(site.text, 'html.parser')
    html_header_list = soup.select('.entry-title')

    for headers in html_header_list:
        header_list.append(headers.get_text())
    
    for header in header_list:
        print(header.strip())

if __name__ == "__main__":
    site = pull_site()
    scraper(site)