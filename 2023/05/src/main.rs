use itertools::Itertools;
use std::cmp;
use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    println!("{}", part1());
    println!("{}", part2());
}

fn part1() -> usize {
    let seeds = parse_seeds(true).iter().map(|s| s.min).collect::<Vec<_>>();
    let maps = parse_maps();

    let mut result = usize::MAX;
    for seed in seeds {
        let mut val = seed;
        for map in &maps {
            val = map_to_next(val, &map);
        }

        if val < result {
            result = val;
        }
    }

    return result;
}

fn part2() -> usize {
    let mut ranges = parse_seeds(false);
    let all_maps = parse_maps();

    for maps in &all_maps {
        let mut next_ranges: Vec<Range> = Vec::new();

        for range in &ranges {
            let mut filt_maps = maps
                .iter()
                .filter(|m| {
                    range.min <= m.src.max && m.src.max <= range.max
                        || range.min <= m.src.min && m.src.min <= range.max
                        || m.src.min <= range.min && range.max <= m.src.max
                })
                .collect::<Vec<_>>();
            filt_maps.sort_by(|a, b| a.src.min.cmp(&b.src.min));

            let mut cur = range.min;
            for m in filt_maps {
                if cur < m.src.min {
                    next_ranges.push(Range {
                        min: cur,
                        max: m.src.min - 1,
                    });

                    cur = m.src.min;
                }

                let offset = cur - m.src.min;
                let width = cmp::min(m.src.max, range.max) - cur;

                next_ranges.push(Range {
                    min: m.dest.min + offset,
                    max: m.dest.min + offset + width,
                });

                cur = m.src.max + 1;
            }

            if cur < range.max {
                next_ranges.push(Range {
                    min: cur,
                    max: range.max,
                })
            }
        }

        ranges = next_ranges;
    }

    return ranges.iter().map(|r| r.min).min().unwrap();
}

fn parse_seeds(is_part1: bool) -> Vec<Range> {
    let file = File::open("input.txt").expect("Could not open file");
    let reader = BufReader::new(file);
    let line = reader.lines().next().unwrap().unwrap();

    let mut split = line.split(":");
    split.next();

    if is_part1 {
        return split
            .next()
            .unwrap()
            .split_whitespace()
            .map(|v| {
                let val = v.parse::<usize>().unwrap();
                return Range { min: val, max: val };
            })
            .collect::<Vec<_>>();
    }

    let mut seeds: Vec<Range> = Vec::new();
    let nums = split
        .next()
        .unwrap()
        .split_whitespace()
        .map(|v| v.parse::<usize>().unwrap());

    let mut start = 0;
    for (i, num) in nums.enumerate() {
        if i % 2 == 0 {
            start = num;
        } else {
            seeds.push(Range {
                min: start,
                max: start + num - 1,
            })
        }
    }

    return seeds;
}

fn parse_maps() -> Vec<Vec<Map>> {
    let file = File::open("input.txt").expect("Could not open file");
    let reader = BufReader::new(file);
    let mut seeds_soil: Vec<Map> = Vec::new();
    let mut soil_fert: Vec<Map> = Vec::new();
    let mut fert_water: Vec<Map> = Vec::new();
    let mut water_light: Vec<Map> = Vec::new();
    let mut light_temp: Vec<Map> = Vec::new();
    let mut temp_hum: Vec<Map> = Vec::new();
    let mut hum_loc: Vec<Map> = Vec::new();

    let mut map_count = 0;
    for (i, l) in reader.lines().enumerate() {
        let line = l.expect("Could not read line");

        if i == 0 || line.trim().is_empty() {
            continue;
        } else if line.contains(":") {
            map_count += 1;
            continue;
        }

        // select the current map and push values
        let map: &mut Vec<Map> = match map_count {
            1 => &mut seeds_soil,
            2 => &mut soil_fert,
            3 => &mut fert_water,
            4 => &mut water_light,
            5 => &mut light_temp,
            6 => &mut temp_hum,
            _ => &mut hum_loc,
        };

        let (dest_start, src_start, width) = line
            .split_whitespace()
            .map(|v| v.parse::<usize>().unwrap())
            .collect_tuple()
            .unwrap();

        map.push(Map {
            src: Range {
                min: src_start,
                max: src_start + width - 1,
            },
            dest: Range {
                min: dest_start,
                max: dest_start + width - 1,
            },
        })
    }

    return vec![
        seeds_soil,
        soil_fert,
        fert_water,
        water_light,
        light_temp,
        temp_hum,
        hum_loc,
    ];
}

fn map_to_next(val: usize, next: &Vec<Map>) -> usize {
    return next
        .iter()
        .find(|v| v.src.min <= val && val <= v.src.max)
        .map_or_else(|| val, |v| val - v.src.min + v.dest.min);
}

#[derive(Debug)]
struct Range {
    min: usize,
    max: usize,
}

#[derive(Debug)]
struct Map {
    src: Range,
    dest: Range,
}
