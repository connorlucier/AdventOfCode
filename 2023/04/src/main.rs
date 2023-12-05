use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    println!("{}", part1());
    println!("{}", part2());
}

fn part1() -> i32 {
    let file = File::open("input.txt").expect("Could not open file");
    let reader = BufReader::new(file);
    let mut result = 0;

    for line in reader.lines() {
        let (win_nums, hand_nums) = parse_line(line.unwrap());
        let mut score = 0;

        for num in hand_nums {
            if win_nums.contains(&num) {
                if score == 0 {
                    score = 1;
                } else {
                    score = score * 2;
                }
            }
        }

        result = result + score;
    }

    return result;
}

fn part2() -> usize {
    let file = File::open("input.txt").expect("Could not open file");
    let reader = BufReader::new(file);

    let mut cards: Vec<Card> = Vec::new();
    for line in reader.lines() {
        let (win_nums, hand_nums) = parse_line(line.unwrap());
        cards.push(Card {
            win_nums: win_nums,
            hand_nums: hand_nums,
        });
    }

    // initialize as 1 copy per card
    let mut copies: Vec<usize> = cards.iter().map(|_| 1).collect::<Vec<_>>();

    for (i, card) in cards.iter().enumerate() {
        let copies_to_add = card
            .hand_nums
            .iter()
            .filter(|num| card.win_nums.contains(num))
            .collect::<Vec<_>>()
            .len();

        for j in i + 1..i + copies_to_add + 1 {
            copies[j] += copies[i];
        }
    }

    return copies.iter().sum();
}

fn parse_line(line: String) -> (Vec<i32>, Vec<i32>) {
    let mut split1 = line.split("|");
    let mut split2 = split1.next().unwrap().split(":");
    split2.next(); // don't care about "Card #:"

    let win_nums = split2
        .next()
        .unwrap()
        .split_whitespace()
        .map(|v| v.parse::<i32>().unwrap())
        .collect::<Vec<_>>();

    let hand_nums = split1
        .next()
        .unwrap()
        .split_whitespace()
        .map(|v| v.parse::<i32>().unwrap())
        .collect::<Vec<_>>();

    return (win_nums, hand_nums);
}

struct Card {
    win_nums: Vec<i32>,
    hand_nums: Vec<i32>,
}
