use std::collections::HashMap;

fn parse(output: &str) -> HashMap<String, String> {
    let mut ret = HashMap::new();
    if output.is_empty() {
        return  ret;
    }
    for row in output.lines() {
        if row.is_empty() {
            continue;
        }
        if let Some((file_name, mine_type)) = row.split_once(":") {
            ret.insert(file_name.trim().to_string(), mine_type.trim().to_string());
        }
    }
    return ret;
}
fn main() {
    let output_str = "go.mod: test/plain\ngo.sum: text/plain\nmain: application/x-match-binary".to_string();
    let parse_ary = parse(&output_str);
    for (k, v) in parse_ary.iter() {
        println!("{}:{}", k, v);
    }
}
