use regex::Regex;
use std::collections::HashMap;
use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let (instr, nodes) = parse_input();
    println!("{}", part1(&instr, &nodes));
    println!("{}", part2(&instr, &nodes));
}

fn part1(instr: &Vec<char>, nodes: &HashMap<String, Node>) -> usize {
    return process("AAA", |v| v == "ZZZ", instr, nodes);
}

fn part2(instr: &Vec<char>, nodes: &HashMap<String, Node>) -> usize {
    let instr_counts = nodes
        .iter()
        .filter(|(k, _)| k.ends_with("A"))
        .map(|(k, _)| process(k, |v| v.ends_with("Z"), instr, nodes))
        .collect::<Vec<_>>();

    return lcm(instr_counts);
}

fn process(
    start: &str,
    end_check: impl Fn(&str) -> bool,
    instr: &Vec<char>,
    nodes: &HashMap<String, Node>,
) -> usize {
    let mut cur = start;
    let mut result = 0;

    while !end_check(cur) {
        for i in instr.iter() {
            if end_check(cur) {
                break;
            }

            cur = match i {
                'L' => &nodes[cur].left,
                'R' => &nodes[cur].right,
                _ => panic!("Invalid instruction"),
            };

            result += 1;
        }
    }

    return result;
}

fn parse_input() -> (Vec<char>, HashMap<String, Node>) {
    let file = File::open("input.txt").expect("Could not open file");
    let reader = BufReader::new(file);
    let mut instr: Vec<char> = Vec::new();
    let mut nodes: HashMap<String, Node> = HashMap::new();
    let re = Regex::new(r"(?<src>[A-Z]+)\s=\s\((?<left>[A-Z]+),\s(?<right>[A-Z]+)\)").unwrap();

    for (i, l) in reader.lines().enumerate() {
        let line = l.expect("Could not read line");
        if i == 0 {
            instr = line.chars().collect::<Vec<_>>();
            continue;
        } else if line.trim().is_empty() {
            continue;
        }

        assert!(re.is_match(&line));
        let caps = re.captures(&line).unwrap();

        nodes.insert(
            caps["src"].to_string(),
            Node {
                left: caps["left"].to_string(),
                right: caps["right"].to_string(),
            },
        );
    }

    return (instr, nodes);
}

#[derive(Debug)]
struct Node {
    left: String,
    right: String,
}

fn lcm(nums: Vec<usize>) -> usize {
    let mut factors = HashMap::new();

    for mut n in &mut nums.into_iter() {
        let mut primes = HashMap::new();

        while n % 2 == 0 {
            let prime = primes.entry(2).or_insert(1);
            *prime *= 2;
            n /= 2;
        }

        let sqrt = (n as f64).sqrt() as usize + 1;
        for i in (3..sqrt).step_by(2) {
            while n % i == 0 {
                let prime = primes.entry(i).or_insert(1);
                *prime *= i;
                n /= i;
            }
        }

        if n > 2 {
            *primes.entry(n).or_insert(n) = n;
        }

        for (k, v) in primes {
            let fact = factors.entry(k).or_insert(0);

            if v > *fact {
                *fact = v;
            }
        }
    }

    let mut result = 1;
    for (_, v) in factors.iter() {
        result *= v;
    }

    return result;
}
