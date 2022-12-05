#include <iostream>
#include "part1.h"
#include "part2.h"

int main(int argc, char *argv[])
{
  if (argc < 2)
  {
    std::cerr << "missing required argument for day" << std::endl;
    return 1;
  }

  std::string path = "/input.txt";
  path = path.insert(0, argv[1]);
  part1(path);
  part2(path);
}
