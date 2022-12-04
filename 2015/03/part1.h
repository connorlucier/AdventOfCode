#include <iostream>
#include <fstream>
#include <map>
#include <tuple>

void part1(std::string path)
{
  std::ifstream input;
  input.open(path);
  if (!input.is_open())
  {
    std::cerr << "failed to open file: " << path << std::endl;
    throw;
  }

  char next;
  std::tuple<int,int> pos(0,0);
  std::map<std::tuple<int,int>,bool> houses;
  houses[pos] = true;

  while (input >> next)
  {
    switch (next)
    {
      case '^':
        ++std::get<1>(pos);
        break;
      case '>':
        ++std::get<0>(pos);
        break;
      case 'v':
        --std::get<1>(pos);
        break;
      case '<':
        --std::get<0>(pos);
        break;
      default:
        std::cerr << "invalid input: " << next << std::endl;
    }

    if (houses.find(pos) == houses.end())
    {
      houses[pos] = true;
    }
  }

  std::cout << houses.size() << std::endl;
}

