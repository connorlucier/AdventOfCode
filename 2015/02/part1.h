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

  int sq_ft = 0;
  std::string line;
  while (input >> line)
  {
    int l, w, h;
    sscanf(line.c_str(), "%dx%dx%d", &l, &w, &h);

    int box_area = 2 * l * w + 2 * l * h + 2 * w * h;
    int extra = l < w ? l * std::min(w, h) : w * std::min(l, h);
    sq_ft += box_area + extra;
  }

  std::cout << sq_ft << std::endl;
}
