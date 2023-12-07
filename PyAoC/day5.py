with open('input.txt') as doc:
    lines = [line.rstrip('\n') for line in doc]

seeds = lines[0][lines[0].find(":") + 1:].split()



def extract_maps(map_name):
    out_map = []
    for x in range(0, len(lines)):
        if map_name in lines[x]:
            y = x + 1
            while lines[y] != "":
                out_map.append(lines[y].split())
                y = y + 1
                if y > len(lines) - 1:
                    break
    return out_map


seed_to_soil = extract_maps("seed-to-soil")
soil_to_fert = extract_maps("soil-to-fertilizer")
fert_to_water = extract_maps("fertilizer-to-water")
water_to_light = extract_maps("water-to-light")
light_to_temp = extract_maps("light-to-temperature")
temp_to_humid = extract_maps("temperature-to-humidity")
humid_to_loc = extract_maps("humidity-to-location")
all_maps = [seed_to_soil, soil_to_fert, fert_to_water, water_to_light, light_to_temp, temp_to_humid, humid_to_loc]


def day5_part1():
    for seed in range(0, len(seeds)):
        for map_x in range(0, len(all_maps)):
            found = False
            for line in range(0, len(all_maps[map_x])):
                destination_min = int(all_maps[map_x][line][0])
                source_min = int(all_maps[map_x][line][1])
                source_max = int(all_maps[map_x][line][1]) + int(all_maps[map_x][line][2])
                if source_min <= int(seeds[seed]) <= source_max and found == False:
                    diff = int(seeds[seed]) - source_min
                    seeds[seed] = str(destination_min + diff)
                    found = True
    min_loc = seeds[0]
    for loc in range(0, len(seeds)):
        if int(seeds[loc]) < int(min_loc):
            min_loc = seeds[loc]
    return min_loc


print(day5_part1())