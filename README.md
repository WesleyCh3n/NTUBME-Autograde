# Autograde
This program generate the files which autolab's autograde need for BMECPL Class. Simply as follow,
```
├───autograde-Makefile
├───autograde.tar
│   ├───score.py
│   └───gtest.cpp
```

You can see through `autograde-Makefile` to have the overview of whole concept.

`score.py` is the main code to generate score.

`gtest.cpp` is the Google Test code.

## Installation
```
wget https://github.com/WesleyCh3n/NTUBME-Autograde/raw/master/generate
chmod +x ./generate
mv ./generate ~/.local/bin/
```
Remember to add `export PATH=$PATH:$HOME/.local/bin` to .bashrc.

## Usage
### Generate Homework
```
generate -hw <hw number> \
         -a <number of answer> \
         -t <number of test> \
         -s "<logical operator each test>" \
         -i "<user input of each test>"
```

- `-hw`: homework number with problem. Need zero padding at first charactor. ex: `-hw 03c` # means HW03C
- `-a`: the total number of answers which student declare, like answer1, answer2. ex: `-a 3` # means this problem have 3 answer variables.
- `-t`: number of google test. ex: `-t 3`
- `-s`: logical operator in each google test. Need to be in "" and with space each operator. Valid operator: `>` `<` `=` `>=` `<=` `!=`. ex: `-s "> = =!"`
- `-i`: if this problem has user input. Need to be in "". Each test use `,` to seperate. Use space to seperate every input in each test. ex: `-i "33 12,93 22"`

Example:
Create HW05D with 3 answers, 2 google tests, "= =" in each test, 2 google tests with "12 2" & "9 6" inputs.
```
generate -hw 05d \
         -a 3 \
         -t 2 \
         -s "= =" \
         -i "12 2,9 6"
```
![Attention](https://via.placeholder.com/15/f03c15/000000?text=+) `Attention`: After generating files, `gtest.cpp` still need to modify `{{TYPE}}` of answer and `{{NUM}}` of right answers in code.

### Generate tar
```
generate tar
```
