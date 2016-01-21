#!/usr/bin/env awk -f
BEGIN{
	FS=","
	N = 0
	for(i = 1; i < 10; i += 1){
		NUMBERS[i] = i
	}
}
NF==1{
	# print("\n#-------------------------------------------------------------------------------#\n")
	printf("\n\n%s\n\n", $1)
	N = 0
}
NF>1{
	char = substr($1, 1, 1)
	if($1 && $2 && $3 && $4 && NUMBERS[char]){
		N += 1
		printf("%2s) %-16s%-16s%-16s%-16s\n", N, $1, $2, $3, $4)
	}else{
		printf("    %-16s%-16s%-16s%-16s\n", $1, $2, $3, $4)
	}

}


