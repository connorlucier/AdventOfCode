use fancy_regex::{Match, Regex};
use std::collections::VecDeque;
use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    println!("{}", part_1());
    println!("{}", part_2());
}

fn part_1() -> i32 {
    let file = File::open("input.txt").expect("Could not open file");
    let reader = BufReader::new(file);

    let re = Regex::new(r"[^0-9]").unwrap();
    let mut result = 0;

    for line in reader.lines() {
        let raw = line.expect("Could not read line");
        let nums = re.replace_all(&raw, "");
        let mut val = String::from("");

        val.push(nums.chars().next().unwrap());
        val.push(nums.chars().next_back().unwrap());
        result = result + val.parse::<i32>().unwrap();
    }

    return result;
}

fn part_2() -> i32 {
    let file = File::open("input.txt").expect("Could not open file");
    let reader = BufReader::new(file);

    let re = Regex::new(r"(?=(one|two|three|four|five|six|seven|eight|nine|[0-9]))").unwrap();
    let mut result = 0;

    for line in reader.lines() {
        let raw = line.expect("Could not read line");
        let mut start = 0;
        let mut matches: VecDeque<Match<'_>> = VecDeque::new();
        while let Some(m) = re.captures_from_pos(&raw, start).unwrap() {
            matches.push_back(m.get(1).unwrap());
            start = m.get(0).unwrap().start() + 1;
        }

        let match1 = matches.pop_front().expect("No match");
        let match2: Match<'_>;
        let match2_opt = &matches.pop_back();

        if match2_opt.is_none() {
            match2 = match1;
        } else {
            match2 = match2_opt.unwrap();
        }

        let val1 = get_value(match1.as_str());
        let val2 = get_value(match2.as_str());
        result = result + (val1 * 10 + val2);
    }

    return result;
}

fn get_value(val: &str) -> i32 {
    return match val {
        "one" => 1,
        "two" => 2,
        "three" => 3,
        "four" => 4,
        "five" => 5,
        "six" => 6,
        "seven" => 7,
        "eight" => 8,
        "nine" => 9,
        v => v.parse::<i32>().unwrap(),
    };
}
