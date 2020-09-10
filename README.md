# NTU BME Autograde
This program generate the files which `autolab`'s `autograde` need for NTU BME Computer Programming Course. Simply as follow,
```
├───autograde-Makefile
├───autograde.tar
│   ├───score.py
│   └───gtest.cpp
├───answers.yml
```

You can walk through `autograde-Makefile` to have the overview of whole concept.

`score.py` is the main code to generate score.

`gtest.cpp` is the Google Test code.


## Prerequisites
- curl
- [jq](https://stedolan.github.io/jq/): command-line JSON processor.
- [yq](https://mikefarah.gitbook.io/yq/): command-line YAML processor.

## Installation
```bash
git clone https://github.com/WesleyCh3n/NTUBME-Autograde.git & cd NTUBME-Autograde
sudo make install
```

## Usage - Generate Autograding Code
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

### Variables
|   Parameters   |                     Info                    |         Format        |        Exmaple       |
|     :----:     |                     :--                     |          :--          |          :--         |
|  **Autograde** |                  Top level                  |                       |                      |
|  **Homework**  |                  HW number                  |     number+problem    |         `6c`         |
|**VariableType**|The tatal type of variables in this question.|use list to store types|`[float, double, int]`|

**Google Test field**(`Test:`): using yaml [list](https://docs.ansible.com/ansible/latest/reference_appendices/YAMLSyntax.html) syntax to store the lists of tests also lists of logical operation each test. Remember not to write unnecessary `-`, that may loss some of tests.

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

If there is no **input**, leave it blank. if there is no test, just delete `Test:` field.


## Uninstall
```bash
sudo make uninstall
```
