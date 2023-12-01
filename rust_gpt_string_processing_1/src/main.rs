fn char_cnt(input: &str, search_text: &str) -> i32 {
    let mut ret = 0;
    if input.is_empty() {
        return ret;
    }
    for c in input.chars() {
        if c.to_string()  == search_text {
            ret += 1;
        }
    }
    return  ret;
}

fn main() {
    let input = "Hello, world!".to_string();
    let search_text = "l".to_string();
    let cnt = char_cnt(&input, &search_text);
    println!("{}", cnt);
}