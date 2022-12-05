#include <iostream>
#include <sstream>

class instruction {
  private:
    std::string instr;
  public:
    int x1;
    int y1;
    int x2;
    int y2;

    instruction(std::string instr, int x1, int y1, int x2, int y2)
    {
      this->instr = instr;
      this->x1 = x1;
      this->y1 = y1;
      this->x2 = x2;
      this->y2 = y2;
    }

    std::string name() const
    {
      return instr;
    }
};

std::ostream& operator<<(std::ostream &strm, const instruction &instr) {
  return strm << instr.name() << " " <<
    instr.x1 << "," << instr.x2 << " to " <<
    instr.y1 << "," << instr.y2;
}

instruction parse_line(std::string input)
{
  std::string instr;
  int x1, y1, x2, y2;
  bool is_start = true;

  std::istringstream ss(input);
  std::string token;
  while (ss >> token)
  {
    if (token == "through")
    {
      continue;
    }
    if (token == "turn" || token == "toggle")
    {
      instr = token;
    }
    else if (token == "on" || token == "off")
    {
      instr += " " + token;
    }
    else
    {
      sscanf(token.c_str(), "%d,%d", is_start ? &x1 : &x2, is_start ? &y1 : &y2);
      is_start = !is_start;
    }
  }

  return instruction(instr, x1, y1, x2, y2);
}