use regex::Regex;
use std::cmp;
use std::fs::File;
use std::io::{BufRead, BufReader};

const TOTAL_RED: i32 = 12;
const TOTAL_GREEN: i32 = 13;
const TOTAL_BLUE: i32 = 14;

fn main() {
    println!("{}", part1());
    println!("{}", part2());
}

fn part1() -> i32 {
    let file = File::open("input.txt").expect("Could not open file");
    let reader = BufReader::new(file);
    let mut result = 0;

    let re1 = Regex::new("Game (?<game_id>[0-9]+)").unwrap();
    let re2 = Regex::new("(?<quantity>[0-9]+) (?<color>red|green|blue)").unwrap();

    for line in reader.lines() {
        let raw = line.expect("Could not read line");

        let caps = re1.captures(&raw).expect("No match for game_id");
        let game_id = &caps["game_id"].parse::<i32>().unwrap();

        let mut split = raw.split(": ");
        split.next();

        let all_sets = split.next().unwrap();
        let split_sets = all_sets.split("; ");

        let mut possible = true;
        for set in split_sets {
            if !possible {
                break;
            }

            let split_shows = set.split(", ");
            for show in split_shows {
                let caps = re2.captures(&show).expect("No match for quantity/color");
                let qty = &caps["quantity"].parse::<i32>().unwrap();
                let color = &caps["color"];
                if !is_possible(*qty, color) {
                    possible = false;
                    break;
                }
            }
        }

        if possible {
            result = result + game_id;
        }
    }

    return result;
}

fn part2() -> i32 {
    let file = File::open("input.txt").expect("Could not open file");
    let reader = BufReader::new(file);
    let mut result = 0;

    let re = Regex::new("(?<quantity>[0-9]+) (?<color>red|green|blue)").unwrap();

    for line in reader.lines() {
        let raw = line.expect("Could not read line");
        let mut max_red = 0;
        let mut max_green = 0;
        let mut max_blue = 0;

        let mut split = raw.split(": ");
        split.next();

        let all_sets = split.next().unwrap();
        let split_sets = all_sets.split("; ");
        for set in split_sets {
            let split_shows = set.split(", ");
            for show in split_shows {
                let caps = re.captures(&show).expect("No match for quantity/color");
                let qty = &caps["quantity"].parse::<i32>().unwrap();
                let color = &caps["color"];

                match color {
                    "red" => max_red = cmp::max(*qty, max_red),
                    "green" => max_green = cmp::max(*qty, max_green),
                    "blue" => max_blue = cmp::max(*qty, max_blue),
                    _ => continue,
                };
            }
        }

        result = result + max_red * max_green * max_blue;
    }

    return result;
}

fn is_possible(quantity: i32, color: &str) -> bool {
    return match color {
        "red" => quantity <= TOTAL_RED,
        "green" => quantity <= TOTAL_GREEN,
        "blue" => quantity <= TOTAL_BLUE,
        _ => false,
    };
}
