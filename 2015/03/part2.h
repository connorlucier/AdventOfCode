#include <iostream>
#include <fstream>
#include <map>
#include <tuple>

void part2(std::string path)
{
  std::ifstream input;
  input.open(path);
  if (!input.is_open())
  {
    std::cerr << "failed to open file: " << path << std::endl;
    throw;
  }

  bool robo = false;
  char next;
  std::tuple<int,int> santa_pos(0,0);
  std::tuple<int,int> robo_santa_pos(0,0);
  std::map<std::tuple<int,int>,bool> houses;
  houses[santa_pos] = true;

  while (input >> next)
  {
    std::tuple<int,int> *pos = robo ? &robo_santa_pos : &santa_pos; 
    switch (next)
    {
      case '^':
        ++std::get<1>(*pos);
        break;
      case '>':
        ++std::get<0>(*pos);
        break;
      case 'v':
        --std::get<1>(*pos);
        break;
      case '<':
        --std::get<0>(*pos);
        break;
      default:
        std::cerr << "invalid input: " << next << std::endl;
    }

    if (houses.find(*pos) == houses.end())
    {
      houses[*pos] = true;
    }

    robo = !robo;
  }

  std::cout << houses.size() << std::endl;
}

