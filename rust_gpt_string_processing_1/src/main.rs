fn char_cnt(input: &str, search_char: char) -> usize {
    input.matches(search_char).count()
}

fn main() {
    let input = "Hello, world!".to_string();
    let cnt = char_cnt(&input, 'l');
    println!("{}", cnt);
}