#include <string>

namespace log_line {
    std::string message(std::string line) {
        int start = line.find(":");
        std::string message = "";
        message.append(line,start+2,line.size());   
        return message;
    }

    std::string log_level(std::string line) {
        // return the log level
        int start = line.find("[");
        int end = line.find("]");
        std::string level = "";
        level.append(line,start+1,end-1);   

        return level;
    }

    std::string reformat(std::string line) {
        // return the reformatted message
         std::string msg = message(line);
         std::string level = log_level(line);
        
         return msg + " (" + level +")";
    }
}
