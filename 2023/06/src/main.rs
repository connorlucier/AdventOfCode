use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    println!("{}", part1());
    println!("{}", part2());
}

fn part1() -> usize {
    let races = parse_input(true);
    let mut result = 0;

    for race in races {
        let (low, high) = quad(race);
        let ways = high - low + 1;

        if result == 0 {
            result = ways;
        } else {
            result *= ways;
        }
    }
    return result;
}

fn part2() -> usize {
    let mut races = parse_input(false);
    let (low, high) = quad(races.pop().unwrap());
    return high - low + 1;
}

fn parse_input(is_p1: bool) -> Vec<Race> {
    let file = File::open("input.txt").expect("Could not open file");
    let mut lines = BufReader::new(file).lines();
    let mut races: Vec<Race> = Vec::new();

    if is_p1 {
        let times = p1_parse_line(lines.next().unwrap().unwrap());
        let dists = p1_parse_line(lines.next().unwrap().unwrap());

        assert_eq!(times.len(), dists.len());

        for i in 0..times.len() {
            races.push(Race {
                time: times[i],
                dist: dists[i],
            });
        }

        return races;
    }

    return vec![Race {
        time: p2_parse_input(lines.next().unwrap().unwrap()),
        dist: p2_parse_input(lines.next().unwrap().unwrap()),
    }];
}

fn p1_parse_line(line: String) -> Vec<f64> {
    let mut split = line.split(":");
    split.next();

    return split
        .next()
        .unwrap()
        .split_whitespace()
        .map(|v| v.parse::<f64>().unwrap())
        .collect::<Vec<_>>();
}

fn p2_parse_input(line: String) -> f64 {
    let mut split = line.split(":");
    split.next();

    let mut result_str: String = String::new();
    for s in split.next().unwrap().split_whitespace() {
        result_str += s;
    }

    return result_str.parse::<f64>().unwrap();
}

fn quad(r: Race) -> (usize, usize) {
    let low = (r.time - (r.time * r.time - 4.0 * r.dist).sqrt()) / 2.0;
    let high = (r.time + (r.time * r.time - 4.0 * r.dist).sqrt()) / 2.0;
    return (low.ceil() as usize, high.floor() as usize);
}

#[derive(Debug, Copy, Clone)]
struct Race {
    time: f64,
    dist: f64,
}
