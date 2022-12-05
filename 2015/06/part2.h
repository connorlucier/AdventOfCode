#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>

void part2(std::string path)
{
  std::ifstream input;
  input.open(path);
  if (!input.is_open())
  {
    std::cerr << "failed to open file: " << path << std::endl;
    throw;
  }

  std::vector<std::vector<int>> grid(1000, std::vector<int>(1000));
  int total_brightness = 0;
  std::string line;
  while (std::getline(input, line))
  {
    instruction inst = parse_line(line);
    for (int x = inst.x1; x <= inst.x2; ++x)
    {
      for (int y = inst.y1; y <= inst.y2; ++y)
      {
        if (inst.name() == "turn on")
        {
          ++total_brightness;
          ++grid[x][y];
        }
        else if (inst.name() == "turn off")
        {
          if (grid[x][y] > 0)
          {
            --total_brightness;
            --grid[x][y];
          }
        }
        else
        {
          total_brightness += 2;
          grid[x][y] += 2;
        }
      }
    }
  }

  std::cout << total_brightness << std::endl;
}
