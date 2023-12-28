use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let (grid, row, col) = parse_input();
    let path = traverse(&grid, (row, col));

    println!("{}", part1(&path));
    println!("{}", part2(&path));
}

fn part1(path: &Vec<Node>) -> usize {
    return path.len() / 2;
}

fn part2(path: &Vec<Node>) -> usize {
    let mut area = 0;
    for i in 0..path.len() {
        let na = (i + 1) % path.len();
        area += path[i].row as isize * path[na].col as isize
            - path[na].row as isize * path[i].col as isize;
    }

    area /= 2;

    return area as usize + 1 - path.len() / 2;
}

fn parse_input() -> (Vec<Vec<Node>>, usize, usize) {
    let file = File::open("input.txt").expect("Could not open file");
    let reader = BufReader::new(file);
    let mut result: Vec<Vec<Node>> = Vec::new();
    let mut row = 0;
    let mut col = 0;

    for (r, l) in reader.lines().enumerate() {
        let line = l.expect("Could not read line");
        result.push(
            line.chars()
                .enumerate()
                .map(|(c, v)| Node {
                    row: r,
                    col: c,
                    val: v,
                })
                .collect::<Vec<_>>(),
        );

        if line.contains("S") {
            row = r;
            col = line.find("S").unwrap();
        }
    }

    return (result, row, col);
}

fn traverse(grid: &Vec<Vec<Node>>, start: (usize, usize)) -> Vec<Node> {
    let (mut r, mut c) = start;
    let mut dir = Dir::None;
    let mut path: Vec<Node> = Vec::new();

    loop {
        path.push(grid[r][c]);

        match (grid[r][c].val, dir) {
            ('S', _) => {
                if c < grid[r].len()
                    && (grid[r][c + 1].val == 'J'
                        || grid[r][c + 1].val == '7'
                        || grid[r][c + 1].val == '-')
                {
                    c += 1;
                    dir = Dir::Right;
                } else if r < grid.len()
                    && (grid[r + 1][c].val == 'J'
                        || grid[r + 1][c].val == 'L'
                        || grid[r + 1][c].val == '|')
                {
                    r += 1;
                    dir = Dir::Down;
                } else if c > 0
                    && (grid[r][c - 1].val == 'L'
                        || grid[r][c - 1].val == 'F'
                        || grid[r][c - 1].val == '-')
                {
                    c -= 1;
                    dir = Dir::Left;
                } else if r > 0
                    && (grid[r - 1][c].val == '7'
                        || grid[r - 1][c].val == 'F'
                        || grid[r - 1][c].val == '|')
                {
                    r -= 1;
                    dir = Dir::Up;
                }
            }
            ('F', Dir::Up) | ('L', Dir::Down) | ('-', Dir::Right) => {
                c += 1;
                dir = Dir::Right;
            }
            ('7', Dir::Right) | ('F', Dir::Left) | ('|', Dir::Down) => {
                r += 1;
                dir = Dir::Down;
            }
            ('7', Dir::Up) | ('J', Dir::Down) | ('-', Dir::Left) => {
                c -= 1;
                dir = Dir::Left;
            }
            ('J', Dir::Right) | ('L', Dir::Left) | ('|', Dir::Up) => {
                r -= 1;
                dir = Dir::Up;
            }
            (_, _) => break,
        }

        if grid[r][c].val == 'S' {
            break;
        }
    }

    return path;
}

#[derive(Clone, Copy, Eq, PartialEq)]
enum Dir {
    None,
    Right,
    Down,
    Left,
    Up,
}

#[derive(Clone, Copy, Debug, PartialEq, Eq, PartialOrd, Ord)]
struct Node {
    row: usize,
    col: usize,
    val: char,
}
