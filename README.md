<h1 align="center" style="text-shadow: 2px 2px 2px #878787;"> NTU BME Autograde </h1>

![docker action](https://github.com/WesleyCh3n/NTUBME-Autograde/actions/workflows/docker.yml/badge.svg)
![Bash - 5.0.17](https://img.shields.io/badge/Bash-5.0.17-informational?logo=gnu-bash)
![Made with - Docker](https://img.shields.io/badge/With-Docker-informational?logo=Docker)

<a href="http://autolabproject.com">
  <img src="https://autolabproject.com/images/autolab_red.svg" width="380px" height="100px">
</a>

This program ([ga](https://github.com/WesleyCh3n/NTUBME-Autograde/blob/main/src/ga)) create the files which [Autolab](https://autolabproject.com/)'s `autograde` needed for [NTU BME Computer Programming Course](https://nol.ntu.edu.tw/nol/coursesearch/print_table.php?course_id=611%2018300&class=&dpt_code=6110&ser_no=12565&semester=110-1&lang=CH). File tree is as follow,
```
├───autograde-Makefile
├───autograde.tar
│   ├───score.py
│   └───gtest.cpp
├───answers.yml
```

You can walk through [`autograde-Makefile`](https://github.com/WesleyCh3n/NTUBME-Autograde/blob/main/src/Makefile) to have the overview of whole concept.

[`score.py`](https://github.com/WesleyCh3n/NTUBME-Autograde/blob/main/src/score.py) is the main code to grade student score.

[`gtest.cpp`](https://github.com/WesleyCh3n/NTUBME-Autograde/blob/main/src/gtest.cpp) is the [Google Test](https://github.com/google/googletest) code.


## Prerequisites
- curl: get source code template.
- [jq](https://stedolan.github.io/jq/) (v1.5): command-line JSON processor.
- [yq](https://mikefarah.gitbook.io/yq/) (v3.4.0): command-line YAML processor.

## Installation
```bash
git clone https://github.com/WesleyCh3n/NTUBME-Autograde.git & cd NTUBME-Autograde
sudo make install
```

## Usage - Generate Autograding Code

### Basic usage:
```
ga [-n <HW number>] [-L] -Y <answer.yml>
```
`-n`: create blank `answer.yml` with `HW` number

`-Y`: parsing YAML file

`-L`: for scoring lab purpose

In detail:
1. First, goto the homework folder, ex. `~/Autolab/courses/<course_name>/<hw_name>/`.

2. Create sample `answers.yml` by typing

    ```
    ga
    ```

3. In `answers.yml`, modify parameters to create the tests you want:

    ```yml
    ---
    Autograde:
      Homework: 4c
      AdditionalTar: ["./Document.txt", "./foo.txt"]
      VariableType: [int, float, double]
      Test:
        - input: [1, 0.2, -32]
          answer:
            - L: abs(ans1)
              op: '<='
              R: 11/3
    ```

4. Once you finished editing, type following cmd to generate autograding code

    ```
    ga -Y answers.yml
    ```

### YAML Parameters
| Parameters        | Info                                                 | Format                  | Exmaple                |
| :----:            | :--                                                  | :--                     | :--                    |
| **Autograde**     | Top level                                            |                         |                        |
| **Homework**      | HW number                                            | number+problem          | `6c`                   |
| **AdditionalTar** | Additional files                                     | file1, file2            | `[file1, file2]`       |
| **VariableType**  | The tatal type of answer variables in this question. | use list to store types | `[float, double, int]` |

**Google Test field** (`Test:`): using yaml [list](https://docs.ansible.com/ansible/latest/reference_appendices/YAMLSyntax.html) syntax to store the lists of tests also lists of logical operation each test. Remember not to write unnecessary `-`, that may loss some of tests.

This is an example of two test with one and two logical operations.
```yml
...
  Test:
    - input: []
      answer:
        - L:
          op: ''
          R:
    - input: []
      answer:
        - L:
          op: ''
          R:
        - L:
          op: ''
          R:
```

|Parameters|Info                                                                              |Format                          |Exmaple   |
|:----:    |:--                                                                               |:--                             |:--       |
|**input** |The User input for Google Test.                                                   |use list to store the test input|`[12, 13]`|
|**answer**|The right answers. For example, answer1 should be 32, answer2 less than 0.02, etc.|see `L`, `op`, `R` below        |          |

<!-- |**answer**|       The right answers.      |First charactor is which variable, 2nd is logical operator, 3rd is the value. Format is same as **INPUTS**. Fraction with double-type denominator is also acceptable. For example `2=11/3.0`|`1=30,2>89;1!=98,2<=40`| -->

|Parameters|Info                  |Format                                                                                                                                             |Exmaple      |
|:----:    |:--                   |:--                                                                                                                                                |:--          |
|**L**     |Left side of operator |Use `ans<num>` describe which variable you want to test. Can implement with mathematical operation like: `abs(ans1-ans2)`                          |`sqrt(ans1)` |
|**op**    |logical operator      |Use quote.<ul><li>Binary operator: `=`, `>`, `<`, `!=`, `>=`, `<=`</li><li>String comparasion: `&=`(String match), `&?`(String not match)</li></ul>|`'<='`       |
|**R**     |Right side of operator|Value. Fraction with double-type denominator is also acceptable.                                                                                   |`(11/3.0)-12`|

If there is no **input**, leave it empty. If there is no test, just delete `Test:` field.

## Run with Docker

- Put following code to shellrc(e.g. `~/.bashrc`) to use `ga` normally as cli.
```bash
ga() {
    docker run --rm -e HOST_USER_ID=$(id -u) -e HOST_USER_GID=$(id -g) -v `pwd`:/workdir/ ghcr.io/wesleych3n/autograde:latest ga $@
}
  
# After re-login
# ga -n {HW number}
# ga -Y {filename}.yml
```

## Uninstall
```bash
sudo make uninstall
```
