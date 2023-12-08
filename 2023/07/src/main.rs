use itertools::Itertools;
use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    println!("{}", process(false));
    println!("{}", process(true));
}

fn process(use_jokers: bool) -> usize {
    let mut hands = parse_input(use_jokers);
    let mut result = 0;
    hands.sort();
    for (i, h) in hands.iter().enumerate() {
        result += h.bid * (i + 1);
    }

    return result;
}

fn parse_input(use_jokers: bool) -> Vec<Hand> {
    let file = File::open("input.txt").expect("Could not open file");
    let reader = BufReader::new(file);
    let mut hands: Vec<Hand> = Vec::new();

    for l in reader.lines() {
        let line = l.unwrap();
        let (card_str, bid) = line.split_whitespace().collect_tuple().unwrap();
        let counts = card_str
            .chars()
            .unique()
            .map(|c| CharCount {
                val: c,
                count: card_str.chars().filter(|ch| *ch == c).count(),
            })
            .collect::<Vec<_>>();

        hands.push(Hand {
            cards: card_str.to_string(),
            kind: map_to_kind(counts, use_jokers),
            strength: card_str
                .chars()
                .map(|c| map_to_strength(c, use_jokers))
                .join("")
                .parse::<usize>()
                .unwrap(),
            bid: bid.parse::<usize>().unwrap(),
        });
    }

    return hands;
}

fn map_to_strength(val: char, use_jokers: bool) -> &'static str {
    match val {
        '2' => "02",
        '3' => "03",
        '4' => "04",
        '5' => "05",
        '6' => "06",
        '7' => "07",
        '8' => "08",
        '9' => "09",
        'T' => "10",
        'Q' => "12",
        'K' => "13",
        'A' => "14",
        'J' => {
            if use_jokers {
                return "01";
            }
            return "11";
        }
        _ => panic!("Invalid card value"),
    }
}

fn map_to_kind(orig_counts: Vec<CharCount>, use_jokers: bool) -> HandKind {
    let mut counts = orig_counts;

    if use_jokers {
        let joker_count = counts
            .iter()
            .find(|cc| cc.val == 'J')
            .unwrap_or(&&CharCount { val: 'J', count: 0 })
            .count;

        if joker_count > 0 && joker_count < 5 {
            counts = counts
                .iter()
                .cloned()
                .filter(|cc| cc.val != 'J')
                .collect::<Vec<_>>();

            let max_count: &mut CharCount = counts.iter_mut().max_by_key(|cc| cc.count).unwrap();
            max_count.count += joker_count;
        }
    }

    match counts.len() {
        1 => HandKind::FiveOfAKind,
        2 => {
            if counts.iter().any(|cc| cc.count == 4) {
                return HandKind::FourOfAKind;
            }

            return HandKind::FullHouse;
        }
        3 => {
            if counts.iter().any(|cc| cc.count == 3) {
                return HandKind::ThreeOfAKind;
            }

            return HandKind::TwoPair;
        }
        4 => HandKind::OnePair,
        _ => HandKind::HighCard,
    }
}

#[derive(Clone, Debug, Eq, PartialEq, PartialOrd, Ord)]
enum HandKind {
    HighCard,
    OnePair,
    TwoPair,
    ThreeOfAKind,
    FullHouse,
    FourOfAKind,
    FiveOfAKind,
}

#[derive(Clone, Debug, Eq, PartialEq, PartialOrd, Ord)]
struct Hand {
    kind: HandKind,
    strength: usize,
    cards: String,
    bid: usize,
}

#[derive(Clone, Copy, Debug)]
struct CharCount {
    val: char,
    count: usize,
}
