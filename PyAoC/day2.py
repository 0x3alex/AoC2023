import re

with open('input.txt') as doc:
    games = [line.rstrip('\n') for line in doc]


def day2_part1(input):
    id_sum = 0
    for x in range(0, len(input)):
        #remove leading "Game x:" text
        a_draws = input[x][input[x].index(":")+1:]
        #construct list with every draw as one entry
        s_draws = a_draws.split(";")
        valid_draw = True
        for i in range(0, len(s_draws)):
            s_color = s_draws[i].split(",")
            for j in s_color:
                if "green" in j:
                    j = re.sub('[abcdefghijklmnopqrstuvwxyz ]', "", j)
                    if int(j) > 13:
                        valid_draw = False
                if "red" in j:
                    j = re.sub('[abcdefghijklmnopqrstuvwxyz ]', "", j)
                    if int(j) > 12:
                        valid_draw = False
                if "blue" in j:
                    j = re.sub('[abcdefghijklmnopqrstuvwxyz ]', "", j)
                    if int(j) > 14:
                        valid_draw = False
        if valid_draw:
            id_sum = id_sum + (x+1)
    return id_sum


def day2_part2(input):
    power_sum = 0
    for x in range(0, len(input)):
        #remove leading "Game x:" text
        a_draws = input[x][input[x].index(":")+1:]
        #construct list with every draw as one entry
        s_draws = a_draws.split(";")
        max_green = 0
        max_red = 0
        max_blue = 0
        for i in range(0, len(s_draws)):
            s_color = s_draws[i].split(",")
            for j in s_color:
                if "green" in j:
                    j = re.sub('[abcdefghijklmnopqrstuvwxyz ]', "", j)
                    if int(j) > max_green:
                        max_green = int(j)
                if "red" in j:
                    j = re.sub('[abcdefghijklmnopqrstuvwxyz ]', "", j)
                    if int(j) > max_red:
                        max_red = int(j)
                if "blue" in j:
                    j = re.sub('[abcdefghijklmnopqrstuvwxyz ]', "", j)
                    if int(j) > max_blue:
                        max_blue = int(j)
        power_game = max_green * max_red * max_blue
        power_sum = power_sum + power_game
    return power_sum


print(day2_part1(games))
print(day2_part2(games))