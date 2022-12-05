#include <iostream>
#include <fstream>
#include <algorithm>

bool has_double_substr(std::string input)
{
  for (int i = 0; i < input.size() - 1; ++i)
  {
    for (int j = i+2; j < input.size() - 1; ++j)
    {
      if (input.substr(i, 2) == input.substr(j, 2))
      {
        return true;
      }
    }
  }

  return false;
}

bool has_split_double_letter(std::string input)
{
  for (int i = 0; i < input.size() - 2; ++i)
  {
    if (input[i] == input[i+2])
    {
      return true;
    }
  }

  return false;
}

void part2(std::string path)
{
  std::ifstream input;
  input.open(path);
  if (!input.is_open())
  {
    std::cerr << "failed to open file: " << path << std::endl;
    throw;
  }

  std::string line;
  int nice_count = 0;
  while (input >> line)
  {
    if (has_double_substr(line) && has_split_double_letter(line))
    {
      ++nice_count;
    }
  }

  std::cout << nice_count << std::endl;
}
