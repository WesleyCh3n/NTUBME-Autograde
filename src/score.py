import re
import os
import sys
import json
import argparse
from subprocess import run,PIPE


def info(str):
    print(str.center(90, "="))

parser = argparse.ArgumentParser()
parser.add_argument('--hw', help="Homework name", default='')
parser.add_argument('--N', help="Number of google testing", default=0, type=int)
parser.add_argument('--inputs', help="input of testing", default='')
args = parser.parse_args()


## Preprocessing arguments
# list of Testing assessment's inputs
# inputs = args.inputs.replace(',',' ')
inputs = [bytes(i, encoding='utf8') for i in args.inputs.split(';') if i != ""]
# list of Testing commands
if args.N != 0:
    gTest = [f"./out --gtest_filter=GoogleTest.test{i+1}" for i in range(args.N)]
else:
    gTest = ["./out"]

if __name__ == "__main__":
    # Compiling
    info(" Autograde Start ")
    compile_out = run(f"g++ {args.hw} -o out_origin -std=c++17 -Wall".split(), stdout = PIPE, stderr = PIPE)

    # Print compiling outputs
    info(" Warning Output ")
    compile_output = re.sub('‘|’', '"', compile_out.stderr.decode()) # Autograde somehow can't deal with these two characters
    print(compile_output)

    # Count numbers of warning
    warn_list = re.findall(fr'{args.hw}:\d+:\d+: warning',compile_out.stderr.decode())
    warn = 10 if len(warn_list) >= 10 else len(warn_list) # Numbers of warning equal 10 at most

    # Clean compile binary file
    if os.path.isfile('./out_origin'):
        run(f"rm ./out_origin".split())

    # If compile successfull, do google test
    if compile_out.returncode == 0:
        info(" Compiling Complete ")

        # Store testing result
        results = []
        # Compile with google test
        compile_out = run(f"g++ gtest.cpp -o out -std=c++17 -Wall -lgtest -lpthread".split(), stdout = PIPE, stderr = PIPE)

        # Find do compile output have the "not declared..." or not, to test if student declare the right variable name
        info(" Finding Answer ")
        find_ans = re.findall(r'‘answer\d’ was not declared in this scope',compile_out.stderr.decode())

        if len(find_ans) > 0: # if the list > 0 means student didn't declare right name
            print('Unable to find the answer.\nPlease check your answer.\n(Which should be declared in Global.)')
            results.append(0)
            warn = 0
        else:
            info(" Testing Output ")
            # if this test need to input somthing
            if len(args.inputs) > 0: # Need input
                for i in range(len(inputs)):
                    # Run google test
                    test_out = run(gTest[i].split(), input=inputs[i],stdout = PIPE, stderr = PIPE)

                    # if pass test result append 1, else 0
                    results.append(1 if test_out.returncode == 0 else 0)
                    print(f"Pass Tesing Num.{i}".center(90, " ") if test_out.returncode == 0 else f"Fail Testing Num.{i}".center(90, " "))

            elif len(args.inputs) == 0: # Don't need input
                if args.N == 0: # if the test need google test
                    # Run with no google test
                    test_out = run(gTest[0].split(), stdout = PIPE, stderr = PIPE)

                    # if pass test result append 1, else 0
                    results.append(1 if test_out.returncode == 0 else 0)
                    print(f"Pass Tesing Num.{0}".center(90, " ") if test_out.returncode == 0 else f"Fail Testing Num.{0}".center(90, " "))

                else:
                    for i in range(args.N):
                        # Run google test
                        test_out = run(gTest[i].split(), stdout = PIPE, stderr = PIPE)

                        # if pass test result append 1, else 0
                        results.append(1 if test_out.returncode == 0 else 0)
                        print(f"Pass Tesing Num.{i}".center(90, " ") if test_out.returncode == 0 else f"Fail Testing Num.{i}".center(90, " "))

        # Calculate final score
        final = int( 30 + 40*(sum(results)/len(results)) + (-3)*warn)
        info(" Testing Finish ")
        info(" Autograde Finish ")
        print("\n")
        info(" Final Score ")
        print(json.dumps({"scores":{"Autograde":final}}).center(90, " "))

    else:
        # Compile fail
        print(compile_output)
        info(" Compiling Fail ")
        info(" Autograde Finish ")
        info(" Final Score ")
        print(json.dumps({"scores":{"Autograde":0}}).center(90, " "))