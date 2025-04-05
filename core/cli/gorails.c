#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "command_handler.h"

int main(int argc, char *argv[]) {
    if (argc < 2) {
        printf("Usage: %s <command>\n", argv[0]);
        return 1;
    }

    if (strcmp(argv[1], "new") == 0) {
        if (argc < 3) {
            printf("Usage: %s new <project_name>\n", argv[0]);
            return 1;
        }
        create_project(argv[2]);
    }
    else if (strcmp(argv[1], "start") == 0) {
        run_server();
    } 
    else if(strcmp(argv[1], "demo") == 0) {
        if (argc < 3) {
            printf("Usage: %s demo <project_name>\n", argv[0]);
            return 1;
        }
        generate_demo(argv[2]);
    }
    else if (strcmp(argv[1], "controller") == 0) {
        printf("Not implemented yet\n");
    } else {
        printf("GoRails: Invalid command\n");
        return 1;
    }
}
