fn main() {
    let data = include_str!("../input.txt").trim_end();
    println!("Part 1: {}", data.len() - 2 * (data.matches(')').count()),);

    let mut floor = 0;
    for (i, c) in data.chars().enumerate() {
        match c {
            '(' => floor += 1,
            ')' => floor -= 1,
            _ => {}
        }

        if floor < 0 {
            println!("Part 2: {}", i + 1);
            break;
        }
    }
}
