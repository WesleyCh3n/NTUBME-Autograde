HW = HW{{HW_NUM}}.cpp

all:
	@tar xf autograde.tar
	@# Delete stdafx.h pch.h
	@sed -i -E "s/#( *)include( *)\"*<*stdafx.h\"*>*//" ${HW}
	@sed -i -E "s/#( *)include( *)\"*<*pch.h\"*>*//" ${HW}
	@# Copy and backup
	@cp ${HW} "your.cpp"
	@cp ${HW} "${HW}_"
	@# Replace main() to your_main_in_testin()
	@sed -i -E "s/main( *)\(([^)]*+)\)/your_main_in_testing()/" your.cpp
	@python3 score.py --hw ${HW} --N {{N_TEST}} --inputs {{INPUTS}}

clean:
	@mv "${HW}_" ${HW}
	@rm -f out out_origin stdafx.h pch.h score.py gtest.cpp your.cpp
