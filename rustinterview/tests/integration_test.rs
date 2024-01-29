use rustinterview;
use rustinterview::windowquiz::min_sub_array_len;

#[test]
fn it_min_sub_array_len() {
    assert_eq!(0, min_sub_array_len(0, Vec::new()))
}