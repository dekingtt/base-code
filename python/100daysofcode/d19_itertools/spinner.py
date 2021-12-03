import itertools
import sys
import time

def spinner(seconds):
    """Show anf animated spinner while sleep."""
    symbols = itertools.cycle('-\|/')
    tend = time.time() + seconds
    while time.time() < tend:
        sys.stdout.write('\rPlease wait...' + next(symbols)) 
        sys.stdout.flush()
        time.sleep(0.1)
    print()

if __name__ == "__main__":
    spinner(3)