from itertools import permutations, combinations

friends = 'mike bob julian'.split()

print(list(combinations(friends, 2)))

print(list(permutations(friends,2)))