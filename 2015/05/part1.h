#include <iostream>
#include <fstream>
#include <algorithm>

bool is_naughty(std::string input)
{
  return input.find("ab") != std::string::npos ||
    input.find("cd") != std::string::npos ||
    input.find("pq") != std::string::npos ||
    input.find("xy") != std::string::npos;
}

bool has_double_letter(std::string input)
{
  for (int i = 0; i < input.length() - 1; ++i)
  {
    if (input[i] == input[i+1])
    {
      return true;
    }
  }

  return false;
}

int count_vowels(std::string input)
{
  return std::count(input.begin(), input.end(), 'a') +
    std::count(input.begin(), input.end(), 'e') +
    std::count(input.begin(), input.end(), 'i') +
    std::count(input.begin(), input.end(), 'o') +
    std::count(input.begin(), input.end(), 'u');
}

void part1(std::string path)
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
    if (!is_naughty(line) && count_vowels(line) >= 3 && has_double_letter(line))
    {
      ++nice_count;
    }
  }

  std::cout << nice_count << std::endl;
}
