#include <stdio.h>
#include <stdlib.h>

/**
* display_factors - factorizes a number into
*                   a product of two smaller numbers.
* @n: the number to factorize
*
* Return: (void) - always returns nothing
*/
void display_factors(int n)
{
	int i, flag = 0;

	for (i = 2; i <= n / 2; ++i)
	{
		if (n % i == 0)
		{
			printf("%d=%d*%d\n", n, i, n / i);
			flag = 1;
			break;
		}
	}

}

/**
* main - the main function for the monty code interpreter
* @argc: argument count
* @argv: argument value
*
* Return: 0 on success
*/
int main(int argc, char *argv[])
{
	char *content;
	FILE *file;
	size_t size = 0;
	ssize_t read_line = 1;

	if (argc != 2)
	{
		fprintf(stderr, "USAGE: factors <file>\n");
		exit(EXIT_FAILURE);
	}

	file = fopen(argv[1], "r");

	if (!file)
	{
		fprintf(stderr, "Error: Can't open file %s\n", argv[1]);
		exit(EXIT_FAILURE);
	}

	while (read_line > 0)
	{
		content = NULL;
		read_line = getline(&content, &size, file);

		if (read_line > 0)
		{
			display_factors(atoi(content));
		}

	}

	fclose(file);

	return (EXIT_SUCCESS);
}
