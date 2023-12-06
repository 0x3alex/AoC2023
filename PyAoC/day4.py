import math

with open('input.txt') as doc:
    lines = [line.rstrip('\n') for line in doc]


def day4_part1(input_arr):
    point_sum = 0
    for x in range(0, len(input_arr)):
        match_count = 0
        split_numbers = input_arr[x].split("|")
        winning = split_numbers[0][split_numbers[0].find(":")+1:].split()
        on_card = split_numbers[1].split()
        for y in range(0, len(on_card)):
            if on_card[y] in winning:
                match_count = match_count + 1
        if match_count == 1:
            point_sum = point_sum + 1
        if match_count > 1:
            point_sum = point_sum + int(math.pow(2, (match_count-1)))
        print(match_count)
    return point_sum





print(day4_part1(lines))