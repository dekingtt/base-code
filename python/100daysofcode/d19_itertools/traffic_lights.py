from time import sleep
import itertools
import random

colors  ='Red, Green, Amber'.split()
rotation = itertools.cycle(colors)

def rg_timer():
    return random.randint(3, 7)

def light_rotation(rotation):
    for color in rotation:
        if color == 'Amber':
            print(f'Caution! the light is {color}')
            sleep(3)
        if color == 'Red':
            print('STOP! the light is ' + color)
            sleep(rg_timer)
        else:
            print('Go! the light is %s' % color)
        sleep(rg_timer())

if __name__ == '__main__':
    light_rotation(rotation)