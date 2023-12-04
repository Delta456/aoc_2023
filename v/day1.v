import os

fn main() {
	mut sum := 0
	m := {
		'zero':  '0'
		'one':   '1'
		'two':   '2'
		'three': '3'
		'four':  '4'
		'five':  '5'
		'six':   '6'
		'seven': '7'
		'eight': '8'
		'nine':  '9'
	}
	for str in file_lines('../input/day1.txt')! {
		mut numbers := []string{}
		for i, s in str {
			if s.is_digit() {
				numbers << s.ascii_str()
			} else {
				for k, _ in m {
					if str[i..].len >= k.len && str[i..i + k.len] == k {
						numbers << m[k]
					}
				}
			}
		}
		num := '${numbers[0]}${numbers#[-1..][0]}'.int()
		sum += num
	}
	println(sum)
}

fn file_lines(file string) ![]string {
	return os.read_file(file)!.split_into_lines()
}
