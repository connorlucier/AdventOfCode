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

  int ribbon_ft = 0;
  std::string line;
  while (input >> line)
  {
    int l, w, h;
    sscanf(line.c_str(), "%dx%dx%d", &l, &w, &h);

    int ribbon = l < w ? 2 * l + 2 * std::min(w, h) : 2 * w + 2 * std::min(l, h);
    int extra = l * w * h;
    ribbon_ft += ribbon + extra;
  }

  std::cout << ribbon_ft << std::endl;
}
