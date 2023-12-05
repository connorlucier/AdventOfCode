use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let (num_coords, sym_coords) = read_coords();
    println!("{}", part1(&num_coords, &sym_coords));
    println!("{}", part2(&num_coords, &sym_coords));
}

fn part1(num_coords: &Vec<Num>, sym_coords: &Vec<Sym>) -> usize {
    let mut result: usize = 0;

    for num in num_coords {
        let adj_symbols = sym_coords
            .iter()
            .filter(|pt| {
                pt.pos.row as isize >= num.start.row as isize - 1
                    && pt.pos.row <= num.end.row + 1
                    && pt.pos.col as isize >= num.start.col as isize - 1
                    && pt.pos.col <= num.end.col + 1
            })
            .collect::<Vec<_>>();

        if adj_symbols.len() > 0 {
            result = result + num.value;
        }
    }

    return result;
}

fn part2(num_coords: &Vec<Num>, sym_coords: &Vec<Sym>) -> usize {
    let mut result: usize = 0;
    for ast in sym_coords
        .iter()
        .filter(|pt| pt.value == '*')
        .collect::<Vec<_>>()
    {
        // only safe because number lengths are short. if numbers
        // were longer (5+ digits), the middle digit(s) could be
        // adjacent to a symbol as well as start/end digits and
        // would have to check those.
        let adj_numbers = num_coords
            .iter()
            .filter(|num| {
                num.start.row as isize >= ast.pos.row as isize - 1
                    && num.start.row <= ast.pos.row + 1
                    && ((num.start.col as isize >= ast.pos.col as isize - 1
                        && num.start.col <= ast.pos.col + 1)
                        || (num.end.col as isize >= ast.pos.col as isize - 1
                            && num.end.col <= ast.pos.col + 1))
            })
            .collect::<Vec<_>>();

        if adj_numbers.len() == 2 {
            result = result + adj_numbers[0].value * adj_numbers[1].value;
        }
    }

    return result;
}

fn read_coords() -> (Vec<Num>, Vec<Sym>) {
    let file = File::open("input.txt").expect("Could not open file");
    let reader = BufReader::new(file);
    let mut start: isize = -1;
    let mut num_str: String = String::new();
    let mut sym_coords: Vec<Sym> = Vec::new();
    let mut num_coords: Vec<Num> = Vec::new();

    for (i, l) in reader.lines().enumerate() {
        let line = l.unwrap();

        // handle special case for a number at the end of a line
        // also handles end of input file, since last line is an empty line
        if start >= 0 {
            num_coords.push(Num {
                start: Point {
                    row: i - 1,
                    col: start as usize,
                },
                end: Point {
                    row: i - 1,
                    col: line.len() - 1,
                },
                value: num_str.parse::<usize>().unwrap(),
            });

            start = -1;
            num_str = String::new();
        }

        for (j, c) in line.chars().enumerate() {
            if c.is_numeric() {
                if start < 0 {
                    start = j as isize;
                }
                num_str.push(c);
            } else {
                if start >= 0 {
                    num_coords.push(Num {
                        start: Point {
                            row: i,
                            col: start as usize,
                        },
                        end: Point { row: i, col: j - 1 },
                        value: num_str.parse::<usize>().unwrap(),
                    });
                }

                if c != '.' {
                    sym_coords.push(Sym {
                        pos: Point { row: i, col: j },
                        value: c,
                    });
                }

                start = -1;
                num_str = String::new();
            }
        }
    }

    return (num_coords, sym_coords);
}

struct Point {
    row: usize,
    col: usize,
}

struct Sym {
    pos: Point,
    value: char,
}

struct Num {
    start: Point,
    end: Point,
    value: usize,
}
