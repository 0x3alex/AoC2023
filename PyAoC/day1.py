import string
import re

with open('input.txt') as doc:
    lines = [line.rstrip('\n') for line in doc]


##Part 1
def part_1(input):
    out_sum = 0

    for x in range(0, len(input)):
        all_num = re.sub('[abcdefghijklmnopqrstuvwxyz]', "",input[x])
        if(len(all_num) == 0):
            continue
        num_x = int(all_num[0] + all_num[len(all_num)-1])
        out_sum = out_sum + num_x
    return out_sum


##Part 2
def part_2(input):
    sum_2 = 0
    number_words = {
        "one": "1",
        "two": "2",
        "three": "3",
        "four": "4",
        "five": "5",
        "six": "6",
        "seven": "7",
        "eight": "8",
        "nine": "9"
    }

    for x in range(0, len(input)):
        rep_words = input[x]
        result = ""
        acc = ""
        for y in rep_words:
            acc = acc + y
            for i,j in number_words.items():
                if i in acc:
                    result = result + j
                    acc = y
                    break
            if y.isdigit():
                result = result + y
                acc = ""
        num_rep = int(result[0] + result[len(result) - 1])
        sum_2 = sum_2 + num_rep

    return sum_2



print(part_1(lines))
print(part_2(lines))