#include <iostream>
#include <fstream>
#include "md5.h"

void part2(std::string path)
{
  std::ifstream input;
  input.open(path);
  if (!input.is_open())
  {
    std::cerr << "failed to open file: " << path << std::endl;
    throw;
  }

  std::string key;
  input >> key;

  std::string hash;
  int i = 0;
  do
  {
    hash = md5(key + std::to_string(++i));
  } while (hash.find("000000") != 0);

  std::cout << i << std::endl;
}