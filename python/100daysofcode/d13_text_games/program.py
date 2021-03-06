import math
from typing import ChainMap, Match
from actors import Creature, Dragon, Wizard
import random

def print_header():
    print('-----------------------------')
    print('         WIZARD Game')
    print('-----------------------------')
    print()

def game_loop():
    creature = [
        Creature('Bat', 5),
        Creature('Toad', 1),
        Creature('Tiger', 12),
        Dragon('Black Dragon', 50, scaliness=2, breaths_fire=False),
        Wizard('Evil wizard', 1000),
    ]
    hero = Wizard('Gandoif', 75)
    while True:
        active_creature = random.choice(creature)

        print(f'{active_creature.name} of level {active_creature.level} has appear from a dark and foggy forest...')
        print()

        cmd = input('Do you [a]ttck, [r]unaway or [l]ook around?')

        if cmd == 'a':
            if hero.attack(active_creature):
                creature.remove(active_creature)
                print(f"The wizard defeated {active_creature.name}")
            else:
                print(f'The wizard has been defeat by powerful {active_creature.name}')

        elif cmd == 'r':
            print('The wizard has become unsure of his power and flees!!!')
        elif cmd =='l':
            print(f'The wizard {hero.name} takes in the surroundings and sees:')
            for c in creature:
                print(f'{c.name} of level {c.level}')
        else:
            print('OK, exiting game... bye!')
            break
        if not creature:
            print("You've defeated all cretures, well done!")
            break

        print()

def main():
    print_header()
    game_loop()

if __name__ == '__main__':
    main()

