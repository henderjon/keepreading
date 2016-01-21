#!/usr/bin/env awk -f
BEGIN{ FS="," }
NF==1{
	print("],\n")
	print("[\n")
}
NF>1{
	printf("{\"gospel\":\"" $1 "\",\"nt\":\"" $2 "\",\"wisdom\":\"" $3 "\",\"ot\":\"" $4 "\"}\n")
}


