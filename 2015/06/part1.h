#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>

void part1(std::string path)
{
  std::ifstream input;
  input.open(path);
  if (!input.is_open())
  {
    std::cerr << "failed to open file: " << path << std::endl;
    throw;
  }

  std::vector<std::vector<bool>> grid(1000, std::vector<bool>(1000));
  int count = 0;
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
          if (!grid[x][y])
          {
            ++count;
          }

          grid[x][y] = true;
        }
        else if (inst.name() == "turn off")
        {
          if (grid[x][y])
          {
            --count;
          }

          grid[x][y] = false;
        }
        else
        {
          count += grid[x][y] ? -1 : 1;
          grid[x][y] = !grid[x][y];
        }
      }
    }
  }

  std::cout << count << std::endl;
}
