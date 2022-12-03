#include <iostream>
#include <fstream>

void part1(std::string path)
{
  std::ifstream input;
  input.open(path);
  if (!input.is_open())
  {
    std::cerr << "failed to open file: " << path << std::endl;
    throw;
  }

  int floor = 0;
  char next;
  while (input >> next)
  {
    if (next == '(')
    {
      floor += 1;
    }
    else if (next == ')')
    {
      floor -= 1;
    }
    else
    {
      std::cerr << "invalid character in input: " << next << std::endl;
      throw;
    }
  }

  std::cout << floor << std::endl;
  input.close();
}
