#include <iostream>
#include <fstream>

void part2(std::string path)
{
  std::ifstream input;
  input.open(path);
  if (!input.is_open())
  {
    std::cerr << "failed to open file: " << path << std::endl;
    throw;
  }

  int floor = 0, pos = 0;
  char next;
  while (input >> next && floor >= 0)
  {
    if (next == '(')
    {
      ++floor;
    }
    else if (next == ')')
    {
      --floor;
    }
    else
    {
      std::cerr << "invalid character in input: " << next << std::endl;
      throw;
    }

    ++pos;
  }

  std::cout << pos << std::endl;
  input.close();
}
