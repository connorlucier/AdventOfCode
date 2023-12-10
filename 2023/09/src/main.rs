use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let input = parse_input();
    println!("{}", part1(&input));
    println!("{}", part2(&input));
}

fn part1(input: &Vec<Vec<i64>>) -> i64 {
    let mut result = 0;

    for row in input.iter() {
        result += process(row, true);
    }

    return result;
}

fn part2(input: &Vec<Vec<i64>>) -> i64 {
    let mut result = 0;

    for row in input.iter() {
        result += process(row, false);
    }

    return result;
}

fn parse_input() -> Vec<Vec<i64>> {
    let file = File::open("input.txt").expect("Could not open file");
    let reader = BufReader::new(file);
    let mut result: Vec<Vec<i64>> = Vec::new();

    for line in reader.lines() {
        result.push(
            line.unwrap()
                .split_whitespace()
                .map(|v| v.parse::<i64>().unwrap())
                .collect::<Vec<_>>(),
        );
    }

    return result;
}

fn process(nums: &Vec<i64>, is_p1: bool) -> i64 {
    if nums.iter().all(|&v| v == 0) {
        return 0;
    }

    let mut next: Vec<i64> = Vec::new();
    for i in 1..nums.len() {
        next.push(nums[i] - nums[i - 1]);
    }

    if is_p1 {
        return nums.iter().last().unwrap() + process(&next, is_p1);
    }

    return nums[0] - process(&next, is_p1);
}
