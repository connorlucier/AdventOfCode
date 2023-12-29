use std::cmp;
use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    println!("{}", part1());
    println!("{}", part2());
}

fn part1() -> usize {
    let input = parse_input(true);
    return sum_distances(&input);
}

fn part2() -> usize {
    let input = parse_input(false);
    return sum_distances(&input);
}

fn parse_input(is_p1: bool) -> Vec<(usize, usize)> {
    let file = File::open("input.txt").expect("Could not open file");
    let reader = BufReader::new(file);
    let mut result: Vec<(usize, usize)> = Vec::new();
    let mut row_offsets: Vec<usize> = Vec::new();
    let mut col_offsets: Vec<usize> = Vec::new();
    let lines = reader
        .lines()
        .map(|l| l.expect("Could not read line"))
        .collect::<Vec<_>>();

    let mut off = 0;
    for r in 0..lines.len() {
        row_offsets.push(off);

        if !lines[r].contains("#") {
            if is_p1 {
                off += 1;
            } else {
                off += 999999;
            }
        }
    }

    off = 0;
    for c in 0..lines.first().unwrap().len() {
        col_offsets.push(off);
        if lines.iter().all(|l| l.chars().nth(c).unwrap() == '.') {
            if is_p1 {
                off += 1;
            } else {
                off += 999999;
            }
        }
    }

    for (r, l) in lines.iter().enumerate() {
        for (c, ch) in l.chars().enumerate() {
            if ch == '#' {
                result.push((r + row_offsets[r], c + col_offsets[c]));
            }
        }
    }

    return result;
}

fn sum_distances(input: &Vec<(usize, usize)>) -> usize {
    let mut result = 0;

    for i in 0..input.len() {
        for j in i + 1..input.len() {
            result += cmp::max(input[i].0, input[j].0) - cmp::min(input[i].0, input[j].0)
                + cmp::max(input[i].1, input[j].1)
                - cmp::min(input[i].1, input[j].1);
        }
    }

    return result;
}
