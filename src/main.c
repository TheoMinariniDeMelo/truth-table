#include <stdio.h>
#include <errno.h>
#include <limits.h>
#include <stdlib.h>
#include <stdbool.h>
#include <time.h>

long input = 0;
long number_of_rows_now;
bool* create_row(){
	bool* row = malloc(input*sizeof(bool));

	for(int i = 0; i < input; i++){
		row[i] = rand() % 2;
	}
	return row;
}
bool DoesItAlreadyExist(bool* row, bool** arr){
	for(int i = 0; i < number_of_rows_now ; i++){
		for(int j = 0; j < input; j++){
			if(arr[i][j] != row[j]) break;
			else if(j == input - 1) return true; 
		}
	}
	return false;
}
void print_table(bool** truth_table){
	printf("|");
	for(long i = 0; i < input; i++) printf(" i%ld |", i + 1);
	printf("\n");	
	long number_of_rows = 1L << input;

	for(int i = 0; i < number_of_rows; i++){
		printf("|");
		for(int j = 0; j < input; j++){
			printf(" %d |", truth_table[i][j] ? 1 : 0);
		}
		printf("\n");
	}
}
int main(int argc, char **argv){
	if(argc != 2) {
		perror("Erro: Enter with only one input!");
		exit(1);
	}
	errno = 0;
	char* input_rest;
	input = strtol(argv[1], &input_rest, 10);
	if (errno == ERANGE) {
		perror("Erro: valor fora do intervalo de long.\n");
		exit(1);
	}
	if (*input_rest != '\0'){
		perror("Erro: nenhum número foi convertido.\n");
		exit(1);
	}
	long number_of_rows = 1L << input;

	bool** truth_table = malloc(sizeof(bool*) * number_of_rows);

	srand(time(NULL));

	for(int i = 0; i < number_of_rows; i++){
		number_of_rows_now = i;
		bool* row = create_row();
		if(DoesItAlreadyExist(row, truth_table)){
			free(row);
			i--;
			continue;
		}
		truth_table[i] = row;
	}
	print_table(truth_table);
	for (int i = 0; i < number_of_rows; i++) {
		free(truth_table[i]);
	}
	free(truth_table);

	return 0;
}
