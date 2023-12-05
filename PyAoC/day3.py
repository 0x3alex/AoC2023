
with open('input.txt') as doc:
    lines = [line.rstrip('\n') for line in doc]


# function to check neighbouring "cells" for symbols
def check_adjacent(line_index, start_index, end_index):
    exist_symbol = False
    for x in range(line_index - 1, line_index + 2):
        if x < 0 or x > len(lines) - 1:
            continue
        for y in range(start_index - 1, end_index + 2):
            if y < 0 or y > len(lines[x]) - 1:
                continue
            if lines[x][y].isdigit():
                continue
            if lines[x][y] != ".":
                exist_symbol = True
    return exist_symbol


def day3_part1(input_arr):
    start_number = 0
    end_number = 0
    all_sum = 0
    for x in range(0, len(input_arr)):
        cast_number = 0
        number = ""
        for y in range(0, len(input_arr[x])):
            if input_arr[x][y].isdigit():
                number = number + input_arr[x][y]
                if y - 1 < 0:
                    start_number = y
                if not input_arr[x][y - 1].isdigit():
                    start_number = y
                if y == len(input_arr) - 1:
                    end_number = y
                    cast_number = int(number)
                    if check_adjacent(x, start_number, end_number):
                        all_sum = all_sum + cast_number
                    cast_number = 0
                    number = ""
            else:
                if y - 1 < 0:
                    continue
                if input_arr[x][y - 1].isdigit():
                    cast_number = int(number)
                    end_number = y - 1
                    if check_adjacent(x, start_number, end_number):
                        all_sum = all_sum + cast_number
                cast_number = 0
                number = ""
    return all_sum


print(day3_part1(lines))