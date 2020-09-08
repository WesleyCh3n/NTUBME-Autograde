# NTU BME Autograde
This program generate the files which `autolab autograde` need for NTU BME Computer Programming Course. Simply as follow,
```
├───autograde-Makefile
├───autograde.tar
│   ├───score.py
│   └───gtest.cpp
├───answers.yaml
```

You can walk through `autograde-Makefile` to have the overview of whole concept.

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
### Generate Autograding Code
#### Using Yaml (Suggestion)
First, goto the homework folder, ex. `~/Autolab/courses/<course_name>/<hw_name>/`.

Create sample `answers.yaml` by typing
```
generate
```
In `answers.yml`, modify parameters to create the tests you want:
```yml
HW_NUM: <HW number>
TYPE_ANS: <Type of answers>
INPUTS: <User input of each test>
ANSWERS: <Right answers>
```

| Parameters |                     Info                    |                                                   Format                                                  |        Exmaple        |
|   :----:   |                     :-:                     |                                                    :-:                                                    |          :-:          |
| **HW_NUM** |                  HW number                  |                                             Zero padding first                                            |         `06c`         |
|**TYPE_ANS**|The tatal type of variables in this question.|                                         Use `,` seperate each type                                        |     `float,double`    |
| **INPUTS** |       The User input for Google Test.       |                    Use `,` seperate each input in that test. Use `;` seperate each test                   |     `12,23;31,42`     |
| **ANSWERS**|              The right answers.             |First charactor is which variable, 2nd is logical operator, 3rd is the value. Format is same as **INPUTS**. Fraction with double-type denominator is also acceptable. For example `2=11/3.0` |`1=30,2>89;1!=98,2<=40`|

Valid logical operators for **ANSWERS**:
- Binary operator: `=`, `>`, `<`, `!=`, `>=`, `<=`
- String comparasion: `&=`(String match), `&?`(String not match)

If there is no **INPUTS**, leave it blank.

Once you finished editing, type following cmd to generate code

```
generate -Y answers.yaml
```

---
#### Using cmdline (DEPRECATED)
```
generate -hw <hw number> \
         -t <type of answers> \
         -i "<user input of each test>"
         -ans "<right answer>"
```

- `-hw`: homework number with problem. Need zero padding at first charactor. ex: `-hw 03c` # means HW03C
- `-t`: type of each `answer` variables. ex: `-t "int,float,float"`
- `-i`: if this problem has user input. Need to be in "". Each test use `;` to seperate. Use space to seperate every input in each test. ex: `-i "33,12;93,22"`
- `-ans`: right answer of each test. Use space saparate each test. Use `;` to sparate each eq. ex: `-ans \"1=2,2!=23,3<=-12;1=12,2>21,3!=1212\"`
    - List of valid operator:
        - Binary: `=` `>` `<` `!=` `>=` `<=`
        - String: `&=`(string match) `&?`(string not match)

Example:

Create HW05D with 3 answers with type `int float float`, "12 2" & "9 6" as each google tests inputs. First google test answers are answer1==12 answer2>=6, second google test answer are answer1!=7 answer2<99.

```
generate -hw 05d \
         -t "int float float" \
         -i "12 2,9 6" \
         -ans "1=12;2>=6 1!=7;2<99"
```
